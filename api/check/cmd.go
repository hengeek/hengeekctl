package check

import (
	"fmt"

	"github.com/hengeek/hengeekctl/public"
	"github.com/spf13/cobra"
)

var GetConfigCmd = &cobra.Command{
	Use:   "getconfig",
	Short: "获取配置信息，这是一个示例",
	Long:  `通过命令行获取配置信息`,
	Run: func(cmd *cobra.Command, args []string) {
		word, _ := cmd.Flags().GetString("word")
		fmt.Println("通过配置文件获取到的用户名:", public.Config.TcSecretID)
		fmt.Println("通过配置文件获取到的密码:", public.Config.TcSecretKey)
		fmt.Println("通过命令行获取到的内容是:", word)
	},
}
