/*
 * Copyright (c) 2019. The ango Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 */
package cmd

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	AnsibleBin          = "/usr/bin/ansible-playbook "
	NoTag               = "penglai-release,ypl-back,course-job,course-web,daka-web"
	ErrorPkgUrlNotExist = "Your yml url is incorrect."
	ErrorFileNotExist   = "your yml file is not exist."
)

var (
	DeployType = "deploy"
	FileConfig string
	exampleDeploy = `
	# use file to deploy, file name must in your AngoBaseDir
	ango deploy -f api.yml -t v1.2.0

	# support read stdin of -f
	cat api.yml | ango deploy -f - -t v1.2.0
`
	projCmd = &cobra.Command{
		Use:     "deploy [flags]",
		Short:   "to deploy project",
		Long:    "use ango to deploy project with webhook to dingding",
		Example: exampleDeploy,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// support ReadStdin
			if Config == "-" {
				err := ReadStdin()
				if err != nil {
					log.Fatal(err)
				}
			}
			var flag bool
			FileConfig , _ = downloadFile(Config)
			yml, baseYml, baseProject := GetProjectName(FileConfig)
			for _, v := range strings.Split(NoTag, ",") {
				if strings.Split(baseYml, ".")[0] == v {
					flag = true
					break
				}
			}
			err := Deploy(yml, baseProject, flag)
			if err != nil {
				fmt.Println(err)
				return
			}
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			// 配置文件为空， 配置文件不存在。
			if ExitCase(Config) {
				log.Info("yml config is required, Exit now")
				os.Exit(-1)
			}
		},
	}
)


func Deploy(yml, baseProject string, flag bool) error {
	var cmdStr string
	if flag {
		cmdStr = fmt.Sprintf("%s %s  -f 1", AnsibleBin, yml)
	} else {
		cmdStr = fmt.Sprintf("%s %s -e version=%s -f 1", AnsibleBin, yml, Tag)
	}

	if Verbose {
		cmdStr += " -v"
	}
	return Exec(cmdStr, DeployType, baseProject)
}

func ReadStdin() error {
	var b bytes.Buffer
	_, err := b.ReadFrom(os.Stdin)
	if err != nil {
		return err
	}
	Config = PathName + "/tmp.yml"
	return ioutil.WriteFile(Config, b.Bytes(), 0660)
}

func ExitCase(ymlUrl string) bool {
	if ymlUrl == "-" {
		return false
	}
	return ymlUrlCheck(ymlUrl)
}

func ymlUrlCheck(ymlUrl string) bool {
	if !strings.HasPrefix(ymlUrl, "http") && !FileExist(ymlUrl) {
		message := ErrorFileNotExist
		log.Error(message + "please check where your PkgUrl is right?")
		return true
	}
	// 判断PkgUrl, 有http前缀时, 下载的文件如果不是text/plain,则报错.
	return strings.HasPrefix(ymlUrl, "http") && !downloadFileCheck(ymlUrl)
}

func downloadFileCheck(ymlUrl string) bool {
	u, err := url.Parse(ymlUrl)
	if err != nil {
		return false
	}
	if u != nil {
		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			log.Error(ErrorPkgUrlNotExist, "please check where your PkgUrl is right?")
			return false
		}
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
		_, err = client.Do(req)
		if err != nil {
			log.Error(err)
			return false
		}
		//if tp := resp.Header.Get("Content-Type"); tp != "application/octet-streamfile" {
		//	log.Error("your pkg url is  a ", tp, "file, please check your PkgUrl is right?")
		//	return false
		//}
	}
	return true
}

// 根据location判断，url则下载并返回os的路径及md5。
func downloadFile(location string) (filePATH, md5 string) {
	if _, isUrl := isUrl(location); isUrl {
		absPATH := PathName + "/" + path.Base(location)
		if !FileExist(absPATH) {
			//generator download cmd
			dwnCmd := downloadCmd(location)
			//os exec download command
			exCmd := fmt.Sprintf("mkdir -p %s && cd %s && %s ", PathName, PathName, dwnCmd)
			execCmd("/bin/sh", "-c", exCmd)
		}
		location = absPATH
	}
	//file md5
	md5 = fromLocal(location)
	return location, md5
}

//根据url 获取command
func downloadCmd(url string) string {
	//only http
	u, isHttp := isUrl(url)
	var c = ""
	if isHttp {
		param := ""
		if u.Scheme == "https" {
			param = "--no-check-certificate"
		}
		c = fmt.Sprintf(" wget -c %s %s", param, url)
	}
	return c
}

// 判断是不是url
func isUrl(u string) (url.URL, bool) {
	if uu, err := url.Parse(u); err == nil && uu != nil && uu.Host != "" {
		return *uu, true
	}
	return url.URL{}, false
}

// 根据本地文件判断 md5
func fromLocal(localPath string) string {
	cmd := fmt.Sprintf("md5sum %s | cut -d\" \" -f1", localPath)
	c := exec.Command("sh", "-c", cmd)
	out, err := c.CombinedOutput()
	if err != nil {
		log.Error(err)
	}
	md5 := string(out)
	md5 = strings.ReplaceAll(md5, "\n", "")
	md5 = strings.ReplaceAll(md5, "\r", "")

	return md5
}