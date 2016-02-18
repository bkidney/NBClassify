// Copyright © 2016 Brian Kidney <bkidney@briankidney.ca>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/gob"
	"errors"
	"os"

	"github.com/bkidney/NBClassify/utils"
	"github.com/spf13/cobra"
)

var classname string

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train [path]",
	Short: "Create a new class to which text can be matched.",
	Long: `Read in a file and create a classification set from it. Sets 
	created can be specified when attempting to classify a piece of text.`,
	RunE: Train,
}

func Train(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		return errors.New("Missing Filename")
	}

	trainingFile := args[0]
	results := utils.Tally(trainingFile)

	datafile, err := os.Create(args[1] + ".gob")
	if err != nil {
		panic(err)
	}
	defer datafile.Close()

	dataEnc := gob.NewEncoder(datafile)
	err = dataEnc.Encode(results)
	if err != nil {
		panic(err)
	}

	return nil

}

func init() {
	trainCmd.Flags().StringVarP(&classname, "classname", "c", "", "Name for new class")
	RootCmd.AddCommand(trainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trainCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
