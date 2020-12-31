/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var file string
var subnet string
var singleip string
var timeout int
var routinepool int
var output string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-fping",
	Short: "批量IP网络探测CLI工具",
	Long: `批量IP网络探测工具，支持TCP,UDP,ICMP，支持高并发以及文件读取。`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-fping.yaml)")

	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "基于文件内IP探测")
	rootCmd.PersistentFlags().IntVarP(&routinepool, "routinepool", "r", 300, "批量探测的并发池, 默认300个goroutine")	
	rootCmd.PersistentFlags().StringVarP(&singleip, "singleip", "i", "", "探测单个IP，EP: 192.168.1.1")
	rootCmd.PersistentFlags().StringVarP(&subnet, "subnet", "g", "", "探测整个网段, EP: 192.168.1.1/16")
	rootCmd.PersistentFlags().IntVarP(&timeout, "timeout", "T", 1000, "探测超时时间, 单位MS，默认1000ms")
	rootCmd.PersistentFlags().IntVarP(&count, "count", "c", 2, "ICMP 探测IP的Packet 数量")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "探测结果输出位置")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-fping" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-fping")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
