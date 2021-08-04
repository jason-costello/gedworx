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
	"fmt"

	"github.com/spf13/cobra"
)

// matchesCmd represents the matches command
var matchesCmd = &cobra.Command{
	Use:   "matches",
	Short: "Returns list of autosomal dna matches from various providers",
	Long: `Returns list of autosomal dna matches from various providers such as gedmatch.com.
	
	Usage:
	  *  using defaults:  goworx matches -s gm
	  *  explicitly defining constraints: 	goworx matches -s gm -n 50 -o 0 -c 20 -l 45000 -t none
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("matches called")
	},
}

func init() {
	rootCmd.AddCommand(matchesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// matchesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	matchesCmd.Flags().StringP("source", "s", "gm", "source of matches.  Currently only gm (gedmatch) is a valid option and is default option.")
	matchesCmd.Flags().StringP("id", "i", "", "ID of person to find matches for.  This is a source specific id.")
	matchesCmd.Flags().IntP("num", "n", 50, "Number of matches returned.  Default is 50.")
	matchesCmd.Flags().IntP("offset", "o", 0, "Number of items to offset by.  Used in combination with number or returned for paging.  Default is 0")
	matchesCmd.Flags().IntP("cm", "c", 7, "Only return results with a shared cm value greater than this. Default is 7")
	matchesCmd.Flags().IntP("overlap-cutoff", "l", 45000, "Cutoff value for overlap.  Default is 45000.")
	matchesCmd.Flags().StringP("tag-groups", "t", "none", "Return tag-groups. Acceptable values are all, none, one  Default is none. ")
}
