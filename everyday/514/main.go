package main

func findRotateSteps(s, t string) int {
	n, m := len(s), len(t)

	// 先算出每个字母的最后一次出现的下标
	// 由于 s 是环形的，循环结束后的 pos 就刚好是 left[0]
	pos := [26]int{} // 初始值不重要
	for i, b := range s {
		pos[b-'a'] = i
	}
	// 计算每个 s[i] 左边 a-z 的最近下标（左边没有就从 n-1 往左找）
	left := make([][26]int, n)
	for i, b := range s {
		left[i] = pos
		pos[b-'a'] = i // 更新下标
	}

	// 先算出每个字母的首次出现的下标
	// 由于 s 是环形的，循环结束后的 pos 就刚好是 right[n-1]
	for i := n - 1; i >= 0; i-- {
		pos[s[i]-'a'] = i
	}
	// 计算每个 s[i] 右边 a-z 的最近下标（左边没有就从 0 往右找）
	right := make([][26]int, n)
	for i := n - 1; i >= 0; i-- {
		right[i] = pos
		pos[s[i]-'a'] = i // 更新下标
	}

	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1 // -1 表示没有计算过
		}
	}
	var dfs func(int, int) int
	dfs = func(j, i int) (res int) {
		if j == m {
			return 0
		}
		p := &memo[j][i]
		if *p != -1 { // 之前计算过
			return *p
		}
		defer func() { *p = res }() // 记忆化
		if s[i] == t[j] {           // 无需旋转
			return dfs(j+1, i)
		}
		// 左边最近 or 右边最近，取最小值
		l := left[i][t[j]-'a']
		res1 := dfs(j+1, l)
		if l > i {
			res1 += n - l + i
		} else {
			res1 += i - l
		}
		r := right[i][t[j]-'a']
		res2 := dfs(j+1, r)
		if r < i {
			res2 += n - i + r
		} else {
			res2 += r - i
		}
		return min(res1, res2)
	}
	return dfs(0, 0) + m
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
