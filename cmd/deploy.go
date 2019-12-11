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
					continue
				}
			}
		},
	}
)

const (
	AnsibleBin = "/usr/bin/ansible-playbook "
	MallUrl    = "https://mall.youpenglai.com/apis/version"
	AdMallUrl  = "https://admall.youpenglai.com"
	AdComUrl   = "https://ad.youpenglai.com/Public/version"
	CardUrl    = "https://card.youpenglai.com/card/nologin/version"
	WWWUrl     = "https://www.youpenglai.com/"
	PlMall     = "https://plmall.youpenglai.com/"
)

func PlayBook(args string) error {
	cmd := AnsibleBin + args + ".yml" + " -e version=" + Tag
	fmt.Println(cmd)
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	link := play.Linking{}
	link.Msgtype = "link"
	link.Link.Title = args
	link.Link.Text = args + Tag + "部署成功"
	link.Link.PicUrl = "http://icons.iconarchive.com/icons/paomedia/small-n-flat/1024/sign-check-icon.png"
	switch args {
	case "api", "yj-mall", "yj-h5":
		link.Link.MessageUrl = MallUrl
	case "card":
		link.Link.MessageUrl = CardUrl
	case "adcom":
		link.Link.MessageUrl = AdComUrl
	case "www-ypl":
		link.Link.MessageUrl = WWWUrl
	case "yj-admall":
		link.Link.MessageUrl = AdMallUrl
	case "plmall":
		link.Link.MessageUrl = PlMall
	default:
		link.Link.MessageUrl = MallUrl
	}
	err = link.Dingding(DingDingToken)
	if err != nil {
		return err
	}
	return nil
}
