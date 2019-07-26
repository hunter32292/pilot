/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a managed Kubernetes cluster",
	Long:  `Create a managed Kubernetes cluster according to the config provided`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
		configPath, _ := cmd.Flags().GetString("config")
		config := parseConfig(configPath)
		fmt.Printf("Config is: %+v", config)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringP("config", "c", "config.json", "path to the cluster config.json file")
}

func parseConfig(configPath string) Config {
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println("error!")
	}
	var config Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		fmt.Println("error!")
	}
	return config
}

// Config represent the data in a cluster configuration file
type Config struct {
	CloudProvider string `json:"cloudProvider"`
	WorkerNodes   int    `json:"workerNodes"`
}
