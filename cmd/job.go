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

	"github.com/spf13/cobra"
	"github.com/aroq/unipipe/unipipe"
)

var jobName string

// jobCmd represents the job command
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "job operations",
	Long: `Job operations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(unipipe.Job(jobName))
	},
}

func init() {
	rootCmd.AddCommand(jobCmd)
	jobCmd.Flags().StringVarP(&jobName,"name", "n", "", "Job name")
}
