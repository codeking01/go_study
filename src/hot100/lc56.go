package main

/**
 * @author xiongjl
 * @since 2024/1/3  11:02
 */

func merge(intervals [][]int) [][]int {
	//先排序
	quickSortTwoDim(intervals, 0, len(intervals)-1)
	left, right := intervals[0][0], intervals[0][1]
	var res [][]int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			// 更新有边界
			right = maxValue(intervals[i][1], right)
		} else {
			temp := []int{left, right}
			res = append(res, temp)
			left, right = intervals[i][0], intervals[i][1]
		}
	}
	temp := []int{left, right}
	res = append(res, temp)
	left, right = intervals[len(intervals)-1][0], intervals[len(intervals)-1][1]
	return res
}

func maxValue(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 写一个快排，关于二维数组的快排，先写一维的
func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pivot := nums[left]
	l, r := left, right
	for l < r {
		for (l < r) && (nums[r] >= pivot) {
			r--
		}
		if (l < r) && (nums[r] < pivot) {
			nums[l] = nums[r]
		}
		for (l < r) && (nums[l] <= pivot) {
			l++
		}
		if (l < r) && (nums[l] > pivot) {
			nums[r] = nums[l]
		}
	}
	// 更新中间的坐标
	nums[l] = pivot
	quickSort(nums, left, l)
	quickSort(nums, l+1, right)
}

// 二维数组的快排
func quickSortTwoDim(nums [][]int, left int, right int) {
	if left >= right {
		return
	}
	pivot := nums[left][0]
	temp := append([]int{}, nums[left]...)
	l, r := left, right
	for l < r {
		for l < r && nums[r][0] >= pivot {
			r--
		}
		if l < r {
			// 数组是引用类型，所以不可以这样
			//nums[l] = nums[r]
			nums[l] = append([]int{}, nums[r]...)
		}
		for l < r && nums[l][0] <= pivot {
			l++
		}
		if l < r {
			//nums[r] = nums[l]
			nums[r] = append([]int{}, nums[l]...)
		}
	}
	//nums[l] = nums[left]
	nums[l] = append([]int{}, temp...)
	quickSortTwoDim(nums, left, l)
	quickSortTwoDim(nums, l+1, right)
}
