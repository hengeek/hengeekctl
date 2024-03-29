/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/hengeek/hengeekctl/api/check"
	"github.com/spf13/cobra"
)

// exCmd represents the jenkins command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: `与腾讯云账单费用的一些交互`,
}

func init() {
	rootCmd.AddCommand(checkCmd)
	// 获取配置文件
	checkCmd.AddCommand(check.GetProjectLabelCmd)
	cset := check.GetProjectLabelCmd.Flags()
	cset.StringP("list", "l", "", "项目费用的标签")


	checkCmd.AddCommand(check.GetProjectcheckCmd) 
	bset := check.GetProjectcheckCmd.Flags()
	bset.StringP("region", "r", "ap-guangzhou", "传入所在的区域，默认为广州")

}



