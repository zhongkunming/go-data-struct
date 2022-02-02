package main

import (
	"fmt"
)

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	// BubbleSort(list)
	// BubbleSort2(list)
	// SelectSort(list)
	// SelectSort2(list)
	// InsertSort(list)
	// ShellSort(list)
	// MergeSort1(list, 0, len(list))
	MergeSort2(list, 0, len(list))
	fmt.Println(list)
}

// 冒泡排序
func BubbleSort(list []int) {
	num := len(list)
	// 循环n-1次
	for i := num - 1; i > 0; i-- {
		for j := 0; j < num-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// 优化的冒泡
func BubbleSort2(list []int) {
	num := len(list)
	didSwap := false
	// 循环n-1次
	for i := num - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		if !didSwap {
			return
		}
	}
}

func SelectSort(list []int) {
	num := len(list)
	// 进行n-1次轮训
	for i := 0; i < num-1; i++ {
		minIndex := i  // 最小数索引
		min := list[i] // 最小数
		for j := i + 1; j < num; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		if i != minIndex {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

// 改进的选择排序
func SelectSort2(list []int) {
	num := len(list)
	// 进行n-1次轮训
	for i := 0; i < num/2; i++ {
		minIndex := i // 最小数索引
		maxIndex := i // 最大值索引
		for j := i + 1; j < num-i; j++ {
			// 找到最大值下标
			if list[j] > list[maxIndex] {
				maxIndex = j // 这一轮这个是大的，直接 continue
				continue
			}
			// 找到最小值下标
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if maxIndex == i && minIndex != num-i-1 {
			// 最大的元素在开头，最小值不在最后
			// 大值与尾值交换 num-i-1尾值， maxIndex大值
			list[num-i-1], list[maxIndex] = list[maxIndex], list[num-i-1]
			// 小值放在最开头
			list[i], list[minIndex] = list[minIndex], list[i]
		} else if maxIndex == i && minIndex == num-i-1 {
			list[maxIndex], list[minIndex] = list[minIndex], list[maxIndex]
		} else {
			// 小值放开头
			list[i], list[minIndex] = list[minIndex], list[i]
			list[num-i-1], list[maxIndex] = list[maxIndex], list[num-i-1]
		}

	}
}

func InsertSort(list []int) {
	n := len(list)

	for i := 1; i < n; i++ {
		currentNumber := list[i]
		j := i - 1 //前面数字坐标
		if currentNumber < list[j] {
			for ; j >= 0 && currentNumber < list[j]; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = currentNumber
		}
	}
}

func ShellSort(list []int) {
	n := len(list)

	// 设置步长每次减半，直到1
	for step := n / 2; step >= 1; step /= 2 {
		// 开始插入排序
		// 取 d = 6 对 [5 x x x x x 6 x x x x x] 进行直接插入排序，没有变化。
		// 取 d = 3 对 [5 x x 6 x x 6 x x 4 x x] 进行直接插入排序，排完序后：[4 x x 5 x x 6 x x 6 x x]
		// 取 d = 1 对 [4 9 1 5 8 14 6 49 25 6 6 3] 进行直接插入排序，因为 d=1 完全就是直接插入排序了
		for i := step; i < n; i += step {
			// 固定步长
			// 比较
			for j := i - step; j >= 0; j -= step {
				if list[j+step] < list[j] {
					list[j], list[j+step] = list[j+step], list[j]
					continue
				}
				break
			}

		}
	}
}

// 归并排序 自顶向下排序
func MergeSort1(list []int, begin int, end int) {
	if end-begin > 1 {
		mid := begin + (end-begin+1)/2
		MergeSort1(list, begin, mid)
		MergeSort1(list, mid, end)
		merge1(list, begin, mid, end)
	}
}

// 归并排序，自底向上排序
func MergeSort2(list []int, begin int, end int) {
	// 步数为1开始，step长度的数组表示一个有序的数组
	step := 1
	// 范围大于 step 的数组才可以进入归并
	for end-begin > step {
		// 从头到尾对数组进行归并操作
		// step << 1 = 2 * step 表示偏移到后两个有序数组将它们进行归并
		for i := begin; i < end; i += step << 1 {
			lo := i                // 第一个有序数组的上界
			mid := lo + step       // 第一个有序数组的下界，第二个有序数组的上界
			hi := lo + (step << 1) // 第二个有序数组的下界
			// 不存在第二个数组，直接返回
			if mid > end {
				return
			}
			// 第二个数组长度不够
			if hi > end {
				hi = end
			}
			// 两个有序数组进行合并
			merge1(list, lo, mid, hi)
		}
		// 上面的 step 长度的两个数组都归并成一个数组了，现在步长翻倍
		step <<= 1
	}
}

func merge1(list []int, begin, mid, end int) {
	leftSize := mid - begin
	rightSize := end - mid
	newSize := leftSize + rightSize
	result := make([]int, 0, newSize)
	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := list[begin+l]
		rValue := list[mid+r]
		if lValue < rValue {
			result = append(result, lValue)
			l++
		} else {
			result = append(result, rValue)
			r++
		}
	}
	// 将剩下的元素追加到辅助数组后面
	result = append(result, list[begin+l:mid]...)
	result = append(result, list[mid+r:end]...)
	// 将辅助数组的元素复制回原数组，释放辅助空间
	for i := 0; i < newSize; i++ {
		list[begin+i] = result[i]
	}
	return
}
