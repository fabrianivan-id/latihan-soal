package main

import (
	"fmt"
	"strings"
)

func GetDemolitionScore(arr []int, k int) int64 {
	memo := map[string]int64{}
	var key = func(a []int, k int) string {
		sb := strings.Builder{}
		for i, v := range a {
			if i > 0 {
				sb.WriteByte(',')
			}
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte('|')
		sb.WriteString(fmt.Sprint(k))
		return sb.String()
	}
	var max = func(a, b int64) int64 {
		if a > b {
			return a
		}
		return b
	}
	var weaken = func(a []int, sub int) {
		for i := range a {
			a[i] -= sub
			if a[i] < 0 {
				a[i] = 0
			}
		}
	}
	var best func([]int, int) int64
	best = func(a []int, k int) int64 {
		if k == 0 || len(a) == 0 {
			return 0
		}
		sk := key(a, k)
		if v, ok := memo[sk]; ok {
			return v
		}
		var ans int64
		for i := range a {
			score := int64(a[i])
			// decide remaining partition
			left := append([]int(nil), a[:i]...)
			right := append([]int(nil), a[i+1:]...)
			var rem []int
			if len(left) == len(right) {
				rem = left
			} else if len(left) > len(right) {
				rem = left
			} else {
				rem = right
			}
			weaken(rem, a[i])
			cur := score + best(rem, k-1)
			ans = max(ans, cur)
		}
		memo[sk] = ans
		return ans
	}
	cp := append([]int(nil), arr...)
	return best(cp, k)
}

func main() {
	arr := []int{3, 1, 5, 6, 2}
	k := 2
	fmt.Println(GetDemolitionScore(arr, k))
}
