package wordcounter

import (
	"bytes"
	"math/rand"
	"regexp"
)

// WordCount will store word bytes and their occurances
type WordCount struct {
	Word  []byte
	Count int
}

//CleanUpText removes all special character(,:,* etc) from doc
func CleanUpText(words []byte) [][]byte {
	reg := regexp.MustCompile(`\p{Ll}+`)
	matches := reg.FindAll(bytes.ToLower(words), -1)
	return matches
}

//GetWordFrequencies return slice of WordCount
func GetWordFrequencies(matches [][]byte) []*WordCount {
	var wordCounts []*WordCount
	for _, match := range matches {
		isMatch := false
		for _, v := range wordCounts {
			if bytes.Equal(v.Word, match) {
				v.Count++
				isMatch = true
				break
			}
		}
		if !isMatch {
			wordCounts = append(wordCounts, &WordCount{match, 1})
		}
	}

	return wordCounts
}

// SortWordsCountDesc sorts the WordsCount Slice in desc order using MergeSort
func SortWordsCountDesc(a []*WordCount) []*WordCount {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i := range a {
		if a[i].Count > a[right].Count {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	SortWordsCountDesc(a[:left])
	SortWordsCountDesc(a[left+1:])

	return a
}
