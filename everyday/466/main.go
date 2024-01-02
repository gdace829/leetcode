package main

/*
https://leetcode.cn/problems/count-the-repetitions/description/
466. 统计重复个数
困难
233
相关企业
定义 str = [s, n] 表示 str 由 n 个字符串 s 连接构成。

例如，str == ["abc", 3] =="abcabcabc" 。
如果可以从 s2 中删除某些字符使其变为 s1，则称字符串 s1 可以从字符串 s2 获得。

例如，根据定义，s1 = "abc" 可以从 s2 = "abdbec" 获得，仅需要删除加粗且用斜体标识的字符。
现在给你两个字符串 s1 和 s2 和两个整数 n1 和 n2 。由此构造得到两个字符串，其中 str1 = [s1, n1]、str2 = [s2, n2] 。

请你找出一个最大整数 m ，以满足 str = [str2, m] 可以从 str1 获得。
*/
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	map1 := make(map[int]int) //take r1 index
	map2 := make(map[int]int) //take r2 index
	index1 := 0               // r1 index
	index2 := 0               //r2 index
	ans := 0                  // num of s2
	// bian li r1
	for index1/len(s1) < n1 {
		//when index at the last s1
		if index1%len(s1) == len(s1)-1 {
			if val, ok := map1[index2%len(s2)]; ok {
				//how much s1 one circle
				numS1 := index1/len(s1) - val/len(s1)
				//how much s2 one circle
				numS2 := index2/len(s2) - map2[index2%len(s2)]/len(s2)
				//left how much circle
				numCircle := (n1 - index1/len(s1) - 1) / numS1
				ans += numCircle * numS2
				//index of R1
				index1 += numCircle * numS1 * len(s1)

			} else {
				map1[index2%len(s2)] = index1
				map2[index2%len(s2)] = index2
			}
		}

		//when s1 equal s2
		if s1[index1%len(s1)] == s2[index2%len(s2)] {
			if index2%len(s2) == len(s2)-1 {
				ans++
			}
			index2++
		}
		index1++
	}
	return ans / n2
}

// func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
// 	len1, len2 := len(s1), len(s2)
// 	index1, index2 := 0, 0 // 注意此处直接使用 Ra Rb 的下标，不取模

// 	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
// 		return 0
// 	}

// 	map1, map2 := make(map[int]int), make(map[int]int)
// 	ans := 0 // 注意，此处存储的是 Ra 中 Sb 的个数，而非 Ra 中 Rb 的个数

// 	for index1/len1 < n1 { // 遍历整个 Ra

// 		if index1%len1 == len1-1 { //在 Sa 末尾
// 			if val, ok := map1[index2%len2]; ok { // 出现了循环结
// 				cycleLen := index1/len1 - val/len1                 // 每个循环占多少个 Sa
// 				cycleNum := (n1 - 1 - index1/len1) / cycleLen      // 还有多少个循环
// 				cycleS2Num := index2/len2 - map2[index2%len2]/len2 // 当前数量减去上次已经匹配的数量, 每个循环含有多少个 Sb
// 				index1 += cycleNum * cycleLen * len1               // 将 Ra 快进到相应的位置
// 				ans += cycleNum * cycleS2Num                       // 把快进部分的答案数量加上
// 			} else { // 第一次，注意存储的是未取模的, 这个相对位置对应index1 index2
// 				map1[index2%len2] = index1
// 				map2[index2%len2] = index2
// 			}

// 		}

// 		// 出现过多少次s2
// 		if s1[index1%len1] == s2[index2%len2] {
// 			if index2%len2 == len2-1 {
// 				ans += 1
// 			}
// 			index2 += 1
// 		}
// 		index1 += 1
// 	}
// 	return ans / n2
// }

func main() {

}
