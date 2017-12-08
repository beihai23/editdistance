package editdistance

import (
	"fmt"
	"math"
)

// 动态规划实现
func EditDistanceDP(a string, b string) int {
	// 初始化表格, a 为用横  b 为竖
	x := len([]rune(a)) + 1
	y := len([]rune(b)) + 1
	history := make([][]int, y)
	for r := 0; r < y; r++ {
		history[r] = make([]int, x)
		history[r][0] = r
	}
	for c := 0; c < x; c++ {
		history[0][c] = c
	}

	for r := 1; r <= len(b); r++ {
		for c := 1; c <= len(a); c++ {
			if a[c-1] == b[r-1] {
				history[r][c] = history[r-1][c-1]
			} else {
				history[r][c] = _MinOfAll(history[r-1][c]+1, history[r][c-1]+1, history[r-1][c-1]+1)
			}
		}
	}

	return history[len(b)][len(a)]
}

// 动态规划实现, 对空间进行优化
func EditDistance(a string, b string) int {
	// 保存中间结果的表格可以优化为只有一行。
	// 为了进一步节省空间，把较短的字符调整为原本表格的横轴
	x := len([]rune(a))
	y := len([]rune(b))
	shorter := x
	if y < x {
		shorter = y
		a, b = b, a
	}

	shorter += 1
	history := make([]int, shorter)
	for i := 0; i < shorter; i++ {
		history[i] = i
	}

	up := 0
	leftUp := 0
	for r := 1; r <= len(b); r++ {
		history[0] = r
		leftUp = r - 1
		for c := 1; c <= len(a); c++ {
			up = history[c]

			if a[c-1] == b[r-1] {
				history[c] = leftUp
			} else {
				history[c] = _MinOfAll(leftUp+1, history[c]+1, history[c-1]+1)
			}

			leftUp = up
		}
	}

	return history[len(history)-1]
}

// 递归实现
func __EditDistanceRecursion(a []rune, i int, b []rune, j int) int {
	if i == -1 {
		return j + 1 // j是下标，要返回的是长度，所以+1
	}

	if j == -1 {
		return i + 1 // i是下标，要返回的是长度，所以+1
	}

	if a[i] == b[j] {
		return __EditDistanceRecursion(a, i-1, b, j-1)
	}

	return _MinOfAll(__EditDistanceRecursion(a, i-1, b, j)+1, __EditDistanceRecursion(a, i, b, j-1)+1, __EditDistanceRecursion(a, i-1, b, j-1)+1)
}

func _MinOfAll(numbers ...int) int {
	var min int = math.MaxInt64
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}

	return min
}

func _PrintTab(history [][]int) {
	for i := 0; i < len(history); i++ {
		for j := 0; j < len(history[i]); j++ {
			if history[i][j] != -1 {
				fmt.Printf("%2d ", history[i][j])
			}
		}
		fmt.Println()
	}
}
