package main

import (
	"fmt"
	"strings"
)

func capitalizeTitle(title string) string {
	strs := strings.Split(title, " ")
	var res []string
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) > 2 {
			res = append(res, strings.ToUpper(string(strs[i][0]))+strings.ToLower(strs[i][1:]))
		} else {
			res = append(res, strings.ToLower(strs[i]))
		}
	}
	return strings.Join(res, " ")
}
func main() {
	fmt.Println('z' - 'a')
}
