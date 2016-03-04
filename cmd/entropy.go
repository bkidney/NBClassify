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
	"errors"
	"fmt"
	"math"

	"github.com/bkidney/NBClassify/utils"
	"github.com/spf13/cobra"
)

// classifyCmd represents the classify command
var entropyCmd = &cobra.Command{
	Use:   "entropy",
	Short: "Calculates the entropy of a training set.",
	Long: `Calculates the entropy of a data set based on the probability
	of the words within it.`,
	RunE: Entropy,
}

func Entropy(cmd *cobra.Command, args []string) error {

	var training map[string]utils.TallyCount

	if len(args) < 1 {
		return errors.New("Missing Data Set")
	}

	// Load databases foor calculating entropy
	training = make(map[string]utils.TallyCount)
	for i := 0; i < len(args); i++ {
		training[args[i]] = utils.LoadTrainingData(args[i])
	}

	entropies := make(map[string]float64)
	for trainingSetName, trainingSet := range training {
		var entropy float64
		total := trainingSet.Total
		for _, count := range trainingSet.WordCounts {
			prob := float64(count) / float64(total)
			entropy += (prob * math.Log2(prob))
		}

		entropies[trainingSetName] = entropy
		entropies[trainingSetName] = entropies[trainingSetName] * -1

		//fmt.Printf("%s = %f\n", trainingSetName, score[trainingSetName])
	}

	for name, entropy := range entropies {
		fmt.Printf("%s -> %f \n", name, entropy)
	}

	return nil

}

func init() {
	RootCmd.AddCommand(entropyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// classifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// classifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
