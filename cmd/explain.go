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
	"github.com/spf13/cobra"
	"github.com/aroq/uniconf/uniconf"
)

var explainJsonPath string

var explainKey string

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explains given config parameter",
	Long: `Explains given config parameter. For example:
unipipe explain -jsonpath ".params.jobs"
    .`,
	Run: func(cmd *cobra.Command, args []string) {
		uniconf.Explain(explainJsonPath, explainKey)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)
	explainCmd.Flags().StringVarP(&explainJsonPath,"jsonpath", "j", ".", "Jsonpath expression to get collect params from ('.' by default)")
	explainCmd.Flags().StringVarP(&explainKey, "key", "k", "params", "Element name to collect params from ('params' by default)")
}
