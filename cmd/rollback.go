package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	rollbackCmd = &cobra.Command{
		Use:     "rollback  the project",
		Short:   "rollback the project",
		Example: "api, mall, admall",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("rollback :" + strings.Join(args, " "))
		},
	}
)
