这次我们介绍一下归并排序

## 一.归并排序

归并排序很好的体现了分治法的应用，排序的大致思路如下：

1. 将数组切片为相同长度的两部分，一个是 nums[0...LEN/2] 另一个是 nums[LEN/2+1...LEN]
2. 递归的对两个部分进行相同的切片操作，直到数组长度为1
3. 对已经排好序的两个切片进行合并操作

简单解释一下上述思路，由于 1 的操作，我们已经将原数组切片为两个子数组，通过递归使用归并排序方法，子数组均已排序完成，我们接下来只需要将两个已经排好序的子数组进行合并即可

```
func merge_sort(start int, end int, nums string) string {
    // 如何 start == end，也就是 nums 长度为1，直接返回该元素
    // 我们可以认为长度为1的数组是已经排序好了的
	if start == end {
		return string(nums[start])
	}
    // 划分两个子数组
	mid := (start + end) / 2
    // 将低位数组排序
	str1 := merge_sort(start, mid, nums)
    // 将高位数组排序
	str2 := merge_sort(mid+1, end, nums)
    // 将两个有序数组合并
	return merge(str1, str2)
}
func merge(str1 string, str2 string) string {
	newString := ""
	i := 0
	j := 0
	for i < len(str1) && j < len(str2) {
		if str1[i] < str2[j] {
			newString += string(str1[i])
			i++
		} else {
			newString += string(str2[j])
			j++
		}
	}
	if i == len(str1) {
		newString += string(str2[j:])
	}
	if j == len(str2) {
		newString += string(str1[i:])
	}
	return newString
}
```

初次使用递归，会对递归的过程感到疑惑

我们可以追溯一下整个排序的过程

整个 nums长度为16，划分为两个数组 0~7和 8~15。

我们继续划分，将 0~7划分为 0~3和 4~7。 将 0~3继续划分为 0~1 和 2~3

如果我们继续划分，则最后的数组被划分为 0和 1

根据我们的代码和思路，此时直接返回单个元素

接着，我们合并两个单元素，也就是排序了的 M和 E，将其合并为 E和 M

同理，我们可以排序 R和 G，将其合并为 G和 R

接着，我们将 EM和GR进行合并，合并为 EGMR。此时 0~3的过程已经排序完成

回顾整个排序过程，我们将整个数组细分为一个一个的数字，然后两两之间进行合并，然后继续往上合并，最后将整个数组合并。合并过程就是我们的排序过程

接下来我们探讨一下归并排序的时间复杂度和稳定性

归并排序的过程主要消耗在将两个有序数组合并的过程上，整个归并排序的时间复杂度为 O(nlgn)，归并排序本身也是稳定的。

本次我们使用的是二路归并，几路归并取决于我们划分数组的数量，在本次介绍中，我们每次都将数组划分为两个子数组，所以叫做二路归并

------
参考文献：
[Dynamic Connectivity - 普林斯顿大学 | Coursera](https://www.coursera.org/learn/algorithms-part1/lecture/fjxHC/dynamic-connectivity)

