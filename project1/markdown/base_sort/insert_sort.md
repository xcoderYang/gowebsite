基本排序包括简单选择排序和插入排序，本文将就这两种排序进行 golang语言实现，并引出希尔排序

## 一.简单选择排序

简单排序将数组分为两个部分，从左到当前索引的前一个元素为已排序部分，从当前索引到数组的末尾为未排序部分

简单选择排序算法思路如下：

1. 从未排序部分中选取最小的一个元素 A
2. 将 A元素与当前索引所在元素交换
3. 重复 1，2步骤直到未排序部分为空

golang代码如下：

```
func selection_sort(nums []node) []node {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if (nums[j] < nums[minIndex]) {
				minIndex = j
			}
		}
		temp := nums[minIndex]
		nums[minIndex] = nums[i]
		nums[i] = temp
	}
	return nums
}
```
平均情况下：

时间复杂度 O(n^2)

空间复杂度 O(1)


## 二.插入排序

插入排序维护一个已排序数组，每次从未排序数组中选取第一个元素，然后找到该元素应处的位置，进行插入即可

插入排序算法步骤如下:

1. 从未排序数组中选取第一个元素
2. 从排序数组末尾往前找到该元素的插入位置
3. 在对应位置进行插入

```
func insert_sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				temp := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
			} else {
				break
			}
		}
	}
	return nums
}
```

平均情况下：

时间复杂度 O(n^2)

空间复杂度 O(1)

## 三.shell 排序

我们知道插入排序主要耗时用在数组元素的比较上，接下来，我们介绍一种分组插入排序，又称shell排序

shell排序将数组按照 h间隔分成交错的子数组，子数组内部使用插入排序进行排序。我们先选择大一点的 h间隔进行排序，然后逐渐缩小 h，直到 h=1进行最后一次排序

关于 h的选取非常重要，合适的 h会很好的降低的时间复杂度，比较常见的 h选取如下：

1. 2的幂:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1,2,4,8,16,32,...
2. 2的幂-1:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1,3,7,15,31,63,...
3. 3x+1:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1,4,13,40,121,364,...
4. Sedgewick:&nbsp;&nbsp;1,5,19,41,109,209,505,929,2161,3905,...

关于第四种 Sedgewick是一种实际使用中非常优质的 h间隔

Sedgewick的生成方法是两个多项式 (9 ⨉ 4i) – (9 ⨉ 2i) + 1和 4i – (3 ⨉ 2i) + 1 交叉序列

我们在这里选择 3x+1的序列
```
func shell_sort(nums []int) []int {
	N := len(nums)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
				temp := nums[j]
				nums[j] = nums[j-h]
				nums[j-h] = temp
			}
		}
		h /= 3
	}
	return nums
}
```

对比插入排序，如果选择合适的 h间隔，shell排序的时间复杂度会降至 O(nlogn)

## 四.shuffling

接下来，我们介绍一种线性时间内的随机洗牌算法，该洗牌算法又叫做 Knuth shuffle

该洗牌算法从后往前洗牌，算法大致逻辑如下：

算法从后往前洗牌，对于其中的每一张牌，从该牌前面的牌中随机选择一张牌与此牌交换，当前位置的牌就已经洗牌完毕。

算法实现如下：

```
func knuth_shuffle(nums []int) []int {
	Len := len(nums)
	for i := Len - 1; i >= 0; i-- {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(i + 1)
		temp := nums[i]
		nums[i] = nums[index]
		nums[index] = temp
	}
	return nums
}

```
------
参考文献：
[Dynamic Connectivity - 普林斯顿大学 | Coursera](https://www.coursera.org/learn/algorithms-part1/lecture/fjxHC/dynamic-connectivity)

