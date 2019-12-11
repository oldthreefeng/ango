package cmd

import (
	"fmt"
	"github.com/oldthreefeng/ango/play"
	"github.com/spf13/cobra"
	"os/exec"
)

var (
	projCmd = &cobra.Command{
		Use:     "deploy [ some project to deploy ]",
		Short:   "to deploy project ",
		Long:    "to deploy project ",
		Example: "api, yj-mall, yj-adamll",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				err := PlayBook(arg)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		},
	}
)

const (
	AnsibleBin  = "/usr/bin/ansible-playbook "
	MallApiUrl  = "https://mall.youpenglai.com/apis/version"
	AdMallUrl   = "https://admall.youpenglai.com"
	AdComApiUrl = "https://ad.youpenglai.com/Public/version"
	CardApiUrl  = "https://card.youpenglai.com/card/nologin/version"
	WWWUrl      = "https://www.youpenglai.com/"
	PlMall      = "https://plmall.youpenglai.com/"
)

func PlayBook(args string) error {
	//运行完了才打印. 不方便查看
	//cmd := AnsibleBin + args + ".yml" + " -e version=" + Tag
	//output, err := exec.Command("sh", "-c", cmd).Output()
	//if err != nil {
	//	return err
	//}
	//fmt.Printf("%s", output)
	cmdStr := AnsibleBin + args + ".yml" + " -e version=" + Tag
	fmt.Println(cmdStr)
	cmd := exec.Command("sh","-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	if err = cmd.Start(); err != nil {
		return err
	}

	for  {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	if err = cmd.Wait() ; err != nil {
		return err
	}

	link := play.Linking{}
	link.Msgtype = "link"
	link.Link.Title = args
	link.Link.Text = args +":"+ Tag + "部署成功"
	link.Link.PicUrl = "http://icons.iconarchive.com/icons/paomedia/small-n-flat/1024/sign-check-icon.png"
	switch args {
	case "api", "yj-mall", "yj-h5":
		link.Link.MessageUrl = MallApiUrl
	case "card":
		link.Link.MessageUrl = CardApiUrl
	case "adcom":
		link.Link.MessageUrl = AdComApiUrl
	case "www-ypl":
		link.Link.MessageUrl = WWWUrl
	case "yj-admall":
		link.Link.MessageUrl = AdMallUrl
	case "plmall":
		link.Link.MessageUrl = PlMall
	default:
		link.Link.MessageUrl = MallApiUrl
	}
	//fmt.Println(link)
	err = link.Dingding(DingDingToken)
	if err != nil {
		return err
	}
	return nil
}
