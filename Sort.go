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
	// MergeSort2(list, 0, len(list))
	MergeSort3(list, len(list))
	fmt.Println(list)
}

// 冒泡排序
func BubbleSort(list []int) {
	n := len(list)
	// n-1次轮训
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// 优化的冒泡
func BubbleSort2(list []int) {
	n := len(list)
	didSwap := false
	// 循环n-1次
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
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
	n := len(list)
	for i := 0; i < n-1; i++ {
		minIndex := i  // 最小值索引
		min := list[i] // 最小值
		for j := minIndex; j < n; j++ {
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

// 改进的选择排序(最大最小一起)
func SelectSort2(list []int) {
	n := len(list)
	// n-1次轮训
	for i := 0; i < n/2; i++ {
		minIndex := i
		maxIndex := i
		for j := i + 1; j < n-i; j++ {
			// 最大值
			if list[j] > list[maxIndex] {
				maxIndex = j // 这一轮这个是大的，直接 continue
				continue
			}
			// 最小值
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if maxIndex == i && minIndex == n-i-1 {
			// 最大值在开头，最小值在最后
			// 直接交换
			list[maxIndex], list[minIndex] = list[minIndex], list[maxIndex]
		} else if maxIndex == 1 && minIndex != n-i-1 {
			// 最大值在开头，最小值不在最后
			// 先吧最大值放在最后
			list[maxIndex], list[n-i-1] = list[n-i-1], list[maxIndex]
			// 把最小值放前面
			list[i], list[minIndex] = list[minIndex], list[i]
		} else {
			// 最小值放前面
			list[i], list[minIndex] = list[minIndex], list[i]
			// 交换最大值
			list[n-i-1], list[maxIndex] = list[maxIndex], list[n-i-1]
		}

	}
}

func InsertSort(list []int) {
	n := len(list)
	for i := 1; i < n; i++ {
		currentNumber := list[i] // 当前值
		j := i - 1               //前面一个数字的坐标
		// 当前值比前面的值小才比较
		if currentNumber < list[j] {
			for ; j >= 0 && list[j] > currentNumber; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = currentNumber
		}
	}
}

// 把不那么有序的列表变得相对有序， 不同步长试用插入排序
func ShellSort(list []int) {
	n := len(list)
	// 每次步长减半
	for step := n >> 1; step >= 1; step >>= 1 {
		// 根据步长，对分组后的数据进行插入排序
		for i := step; i < n; i += step {
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

// 归并排序 自顶向下
func MergeSort1(list []int, begin int, end int) {
	if end-begin > 1 {
		mid := begin + (end-begin+1)/2
		MergeSort1(list, begin, mid)
		MergeSort1(list, mid, end)
		merge1(list, begin, mid, end)
	}
}

// 归并排序 自底向上
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

// 归并排序 自底向上优化
// 小规模的数组，使用直接插入排序
// 原地排序
func MergeSort3(list []int, n int) {
	// 按照3个一组进行排序，使用直接插入
	blockSize := 3
	a, b := 0, blockSize
	for b <= n {
		InsertSort(list[a:b])
		a = b
		b += blockSize
	}
	InsertSort(list[a:n])

	// 把小数组进行归并操作
	for blockSize < n {
		a, b := 0, blockSize<<1
		for b < n {
			merge2(list, a, a+blockSize, b)
			a = b
			b += blockSize << 1
		}
		if m := a + blockSize; m < n {
			merge2(list, a, m, n)
		}
		blockSize <<= 1
	}
}

// 原地归并操作
func merge2(list []int, begin, mid, end int) {
	i, j, k := begin, mid, end-1 // 因为数组下标从0开始，所以 k = end-1
	for j-i > 0 && k-j >= 0 {
		step := 0
		// 从 i 向右移动，找到第一个 list[i]>list[j]的索引
		for j-i > 0 && list[i] <= list[j] {
			i++
		}
		// 从 j 向右移动，找到第一个 list[j]>list[i]的索引
		for k-j >= 0 && list[j] <= list[i] {
			j++
			step++
		}
		// 进行手摇翻转，将 list[i,mid] 和 [mid,j-1] 进行位置互换
		// mid 是从 j 开始向右出发的，所以 mid = j-step
		rotation(list, i, j-step, j-1)
		i = i + step
	}
}

func rotation(array []int, l, mid, r int) {
	reverse(array, l, mid-1)
	reverse(array, mid, r)
	reverse(array, l, r)
}

func reverse(array []int, l, r int) {
	for l < r {
		// 左右互相交换
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}
