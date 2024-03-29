package check

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// getconfigCmd 基础验证命令
/*
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
*/

var GetProjectLabelCmd = &cobra.Command{
	Use:   "getprojectlabel",
	Short: "获取腾讯云项目的所有标签",
	Long:  `通过命令行获取所有项目的标签`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		// 判断是否有传入标志
		if flags.NFlag() == 0 {
			fmt.Println("Please specify a flag")
			os.Exit(1)
		}
		// region,_ := cmd.Flags().GetString("region")

	},
}


var GetProjectcheckCmd = &cobra.Command{
	Use:   "getprojectcheck",
	Short: "获取腾讯云项目指定标签项目的账单账单",
	Long:  `通过命令行获取项目所有项目账单`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		// 判断是否有传入标志
		if flags.NFlag() == 0 {
			fmt.Println("Please specify a flag")
			os.Exit(1)
		}
	},
}