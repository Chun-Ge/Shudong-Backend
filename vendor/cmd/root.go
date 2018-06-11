// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"args"
	"fmt"
	"os"
	"server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shudong server",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: runner,
}

// runner for the CLI command
func runner(cmd *cobra.Command, cmdArgs []string) {
	port := viper.GetString("Port")
	mysqlURL := viper.GetString("MySQLURL")
	mysqlPort := viper.GetString("MySQLPort")
	mysqlUser := viper.GetString("MySQLUser")
	mysqlPassword := viper.GetString("MYSQLPassword")

	// update the var ars in args/args.go
	args.UpdateVarArgs(port, mysqlURL, mysqlPort, mysqlUser, mysqlPassword)

	// update secret key
	args.UpdateSecretKey()

	// start service
	server.StartWithConfiguration("./config/sample.yml")
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "",
		"config file (default is yourApp/config/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("port", "p", "",
		"the port that the server will be attached to")
	rootCmd.PersistentFlags().String("mysql-port", "",
		"the port that the mysql-server-side is running on")
	rootCmd.PersistentFlags().String("mysql-user", "",
		"the username for mysql")
	rootCmd.PersistentFlags().String("mysql-password", "",
		"the password for mysql")
	rootCmd.PersistentFlags().String("mysql-url", "",
		"the url for mysql")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// bind the flags into config
	// viper.BindPFlag("Config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("Port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("MySQLPort", rootCmd.PersistentFlags().Lookup("mysql-port"))
	viper.BindPFlag("MySQLUser", rootCmd.PersistentFlags().Lookup("mysql-user"))
	viper.BindPFlag("MySQLPassword", rootCmd.PersistentFlags().Lookup("mysql-password"))
	viper.BindPFlag("MySQLURL", rootCmd.PersistentFlags().Lookup("mysql-url"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".demo" (without extension).
		viper.AddConfigPath("./config")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
