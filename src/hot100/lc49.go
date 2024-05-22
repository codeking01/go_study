package main

import (
	"sort"
	"strings"
)

/**
 * @author xiongjl
 * @since 2023/12/25  20:17
 */

func groupAnagrams(strs []string) [][]string {
	// 只需要全部排序即可，然后排序后只要判断两个字符串是否相等就行
	m := make(map[string][]string)
	//var dictKey string
	for _, str := range strs {
		dictKey := sortString(str)
		if _, ok := m[dictKey]; ok {
			m[dictKey] = append(m[dictKey], str)
		} else {
			m[dictKey] = make([]string, 0)
			m[dictKey] = append(m[dictKey], str)
		}
	}
	result := make([][]string, 0)
	// 取出来放数组里面
	for _, value := range m {
		result = append(result, value)
	}
	return result
}

func sortString(str string) string {
	strChar := strings.Split(str, "")
	sort.Strings(strChar)
	return strings.Join(strChar, "")
}

/*func main() {
	temp := "bca"
	fmt.Println(temp)
	runes := []rune(temp)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	temp = string(runes)
	fmt.Print(temp)
	m := make(map[string][]string)
	m["test"] = make([]string, 0)
	fmt.Println("m", m)
	stringsToSort := []string{"banana", "apple", "orange", "grape"}
	// 使用sort.Strings进行排序
	sort.Strings(stringsToSort)
	fmt.Println("Sorted Strings:", stringsToSort)
	// 使用内部里面的排序
	for i, str := range stringsToSort {
		stringsToSort[i] = sortString(str)
	}
	fmt.Println("内部排序的 Strings:", stringsToSort)
	values := []int{1, 2, 3, 4, 5}
	// values... 将切片中的元素展开传递给函数
	sum := calculateSum(values...)
	fmt.Println("Sum:", sum)
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
func calculateSum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
*/
