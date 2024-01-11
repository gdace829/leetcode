package main

import "math"

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := len(word) - 1; i >= 0; i-- {
		num := word[i] - 'a'
		if node.children[num] != nil {
			node = node.children[num]
		} else {
			node.children[num] = &Trie{}
			node = node.children[num]
		}
	}
	node.isEnd = true
}
func (this *Trie) SearchPre(word string) *Trie {
	node := this
	for _, v := range word {
		num := v - 'a'
		if node.children[num] != nil {
			node = node.children[num]
		} else {
			return nil
		}
	}
	return node
}
func (this *Trie) Search(word string) bool {
	return this.SearchPre(word) != nil && this.SearchPre(word).isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPre(prefix) != nil
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minExtraChar1(s string, dictionary []string) int {
	dp := make([]int, len(s)+1)
	map1 := make(map[string]int)
	for _, v := range dictionary {
		map1[v]++
	}
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if _, ok := map1[s[j:i]]; ok {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[len(s)]
}

// 判断连续递增的字符串是否存在可以用字典树
func minExtraChar(s string, dictionary []string) int {
	dp := make([]int, len(s)+1)
	t := Trie{}
	for i := 0; i < len(dictionary); i++ {
		t.Insert(dictionary[i])
	}
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + 1
		node := t
		for j := i - 1; j >= 0; j-- {
			v := node.SearchPre(s[j : j+1])
			if v == nil {
				break
			}
			node = *v
			if node.isEnd == true {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[len(s)]
}
