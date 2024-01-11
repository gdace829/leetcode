package main

import "fmt"

// 字符串题目可以考虑用数字表示 xx-‘a’
func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for i := 0; i < len(magazine); i++ {
		cnt[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		cnt[ransomNote[i]-'a']--
		if cnt[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
