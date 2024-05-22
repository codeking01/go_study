package main

/**
 * @author xiongjl
 * @since 2023/12/27  13:19
 */
/*
func findAnagrams(s string, p string) []int {
	// 双指针
	left, right := 0, len(p)
	target := sortString(p)
	result := make([]int, 0)
	for right <= len(s) {
		temp := sortString(s[left:right])
		if temp == target {
			result = append(result, left)
		}
		left++
		right++
	}
	return result
}*/

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	// 排序会超时
	sFreq, pFreq := [26]int{}, [26]int{}
	result := make([]int, 0)
	for _, item := range p {
		pFreq[item-'a']++
	}
	// 初始化的频次
	for index := range p {
		sFreq[s[index]-'a']++
	}
	// 最开始的位置判断一下
	if isSameSize(sFreq, pFreq) {
		result = append(result, 0)
	}
	pLen := len(p)
	sLen := len(s)
	for i := pLen; i < sLen; i++ {
		sFreq[s[i]-'a']++
		sFreq[s[i-pLen]-'a']--
		if isSameSize(sFreq, pFreq) {
			result = append(result, i-pLen+1)
		}
	}
	return result
}

func isSameSize(a [26]int, b [26]int) bool {
	for index, item := range a {
		if item != b[index] {
			return false
		}
	}
	return true
}

//func main() {
//	s := "bac"
//	fmt.Println(sortString(s[0:2]))
//}
