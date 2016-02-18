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

package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type TallyCount struct {
	Total      int
	WordCounts map[string]int
}

func countWords(f *fileScanner) (counts map[string]int, total int) {

	counts = make(map[string]int)

	scanner := f.GetScanner()
	for scanner.Scan() {
		for _, w := range strings.Fields(scanner.Text()) {
			w = strings.Trim(strings.ToLower(w), ",.!?:;()\"' ")
			if len(w) > 3 {
				counts[w]++
				total++
			}
		}
	}

	return counts, total
}

func Tally(path string) (results TallyCount) {

	fscanner := NewFileScanner()
	err := fscanner.Open(path)
	if err == nil {
		defer fscanner.Close()
		results.WordCounts, results.Total = countWords(fscanner)
	}
	return results
}

func (tc *TallyCount) Print() {
	fmt.Println("Total Words = " + strconv.Itoa(tc.Total))
	fmt.Println("===================================================")
	for k, v := range tc.WordCounts {
		fmt.Printf("%s => %d\n", k, v)
	}
}
