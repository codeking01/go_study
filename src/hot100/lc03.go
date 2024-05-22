package main

/**
 * @author xiongjl
 * @since 2023/12/27  10:13
 */

func lengthOfLongestSubstring(s string) int {
	// 字典存储
	dict := make(map[rune]int)
	max := 0
	left := 0
	for index, item := range s {
		// 判断是否在字典里面，并且left是要和 前一个item的索引比较
		if value, ok := dict[item]; ok && (value >= left) {
			left = dict[item] + 1
		}
		dict[item] = index
		if index-left+1 > max {
			max = index - left + 1
		}
	}
	return max
}

/*func main() {
	s := "abcabc"
	str := make([]string, 0)
	for _, item := range s {
		str = append(str, string(item))
	}
	for index, item := range str {
		if item == "a" {
			fmt.Println(index, item)
		}
	}
	str1 := "abcabcbb"
	lengthOfLongestSubstring(str1)

}*/
