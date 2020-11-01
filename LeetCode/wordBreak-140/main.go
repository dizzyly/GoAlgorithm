package main

import (
	"fmt"
	"strings"
)

/*
	给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，
	在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。
	返回所有这些可能的句子。
*/

func main() {
	wordDict := []string{
		"apple",
		"pen",
		"pine",
	}
	for _, v := range wordDict {
		fmt.Println(v[0])
	}
}

func wordBreak(s string, wordDict []string) []string {
	var sentences []string
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordsList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return sentences
}