// Copyright © 2018 Alexander Tolstikov <tolstikov@gmail.com>
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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/aroq/uniconf/uniconf"
	"path"
)

var cfgFile string

var cfgEnvVar string

var outputFormat string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "unipipe",
	Short: "A tool for managing CICD pipelines",
	Long: `A tool for managing CICD pipelines.`,
	//	Run: func(cmd *cobra.Command, args []string) { },
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

	// Global persistent flags.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config file", "c", path.Join("config.yaml"), "config file ('.unipipe/config.yaml' by default)")
	rootCmd.PersistentFlags().StringVarP(&cfgEnvVar, "config env var", "e", "UNICONF", "config ENV VAR name ('UNICONF' by default)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "yaml", "output format, e.g. 'yaml' or 'json' ('yaml' by default)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	uniconf.AddConfigProvider(defaultUniconfConfig)
}

// defaultUniconfConfig provides default Uniconf configuration.
func defaultUniconfConfig() interface{} {
	return map[string]interface{}{
		"sources": map[string]interface{}{
			"env": map[string]interface{}{
				"type": "env",
			},
			"project": map[string]interface{}{
				"type": "file",
				"path": ".unipipe",
			},
		},
		"from": []interface{}{
			"env:UNICONF",
			"project:/" + cfgFile,
			"env:" + cfgEnvVar,
		},
	}
}
