package main

/**
 * @author xiongjl
 * @since 2023/12/26  19:50
 */

func maxArea(height []int) int {
	// 思路就是，双指针，两边不断往中间靠，l和r相遇就结束
	var w int
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			w = height[l]
			if w*(r-l) > max {
				max = w * (r - l)
			}
			l++
		} else {
			w = height[r]
			if w*(r-l) > max {
				max = w * (r - l)
			}
			r--
		}
	}
	return max
}
