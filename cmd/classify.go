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
	"fmt"
	"log"
	"math"
	"os"

	"github.com/bkidney/NBClassify/utils"
	"github.com/spf13/cobra"
)

// classifyCmd represents the classify command
var classifyCmd = &cobra.Command{
	Use:   "classify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: Classify,
}

func Classify(cmd *cobra.Command, args []string) error {

	var training map[string]utils.TallyCount
	var scores map[string]float64

	if len(args) < 1 {
		return errors.New("Missing Filename")
	}

	classifyFile := args[0]
	results := utils.Tally(classifyFile)

	// Load databases against which classification is done.
	training = make(map[string]utils.TallyCount)
	for i := 1; i < len(args); i++ {
		training[args[i]] = loadTrainingData(args[i])
	}

	var totalTrainingWords int
	for _, trainingData := range training {
		totalTrainingWords += trainingData.Total
	}

	scores = make(map[string]float64)
	for trainingSetName, trainingSet := range training {
		for word, _ := range results.WordCounts {
			if count, exists := trainingSet.WordCounts[word]; exists {
				scores[trainingSetName] += math.Log(float64(count) / float64(trainingSet.Total))
			} else {
				scores[trainingSetName] += math.Log(0.01 / float64(trainingSet.Total))
			}
		}

		scores[trainingSetName] += math.Log(float64(trainingSet.Total) / float64(totalTrainingWords))

		//fmt.Printf("%s = %f\n", trainingSetName, score[trainingSetName])
	}

	var mostLikelyName string
	var mostLikelyScore float64
	for name, score := range scores {
		if len(mostLikelyName) == 0 {
			mostLikelyName = name
			mostLikelyScore = score
		} else {
			if score > mostLikelyScore {
				mostLikelyName = name
				mostLikelyScore = score
			}
		}
	}

	fmt.Printf("%s\n", mostLikelyName)

	return nil

}

func loadTrainingData(datafile string) (trainingData utils.TallyCount) {

	file, err := os.Open(datafile + ".gob")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dataDec := gob.NewDecoder(file)
	err = dataDec.Decode(&trainingData)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	//	trainingData.Print()

	return trainingData
}

func init() {
	RootCmd.AddCommand(classifyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// classifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// classifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
