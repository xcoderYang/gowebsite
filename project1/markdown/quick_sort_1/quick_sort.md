快速排序被誉为20世纪科学和工程领域的十大算法之一。听名字就能了解，快速排序的特点，就是快

## 快速排序

快速排序采用了二分递归的思想，通过一趟排序将整个数组划分为两个部分，低位部分的值全部小于高位部分的值，然后对低位和高位部分分别排序

快速排序算法的具体步骤如下：

1. 对数组 arr 进行随机洗牌
2. 划分数组，我们通过交换操作，找到合适的 j，保证
   
    * j 左边的元素全部小于 arr[j]
    * j 右边的元素全部大于 arr[j]
  
3. 对每一块被划分的切片进行排序

![图片1]()

我们可以通过双指针在O(n)的时间复杂度内获取合适的 j

我们设立两个指针 i 和 j，同时设置一个标志值 arr[low]，一般来说，标志值取数组第一个元素
* 当 arr[i] < arr[low]时，指针 i 从左至右一直扫描
* 当 arr[j] > arr[low]时，指针 j 从右至左一直扫描
* 交换 arr[i]和 a[j]的元素
* 当 i < j 时重复以上步骤

上述算法结束之后，j 所在的位置即为我们寻找的 j

golang代码实现如下：

```
func partition(arr []int, low int, high int) int {
	i, j := low+1, high
	for true {
		for arr[i] < arr[low] {
			i++
			if i == high {
				break
			}
		}
		for arr[low] < arr[j] {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		exch(arr, i, j)
	}
	exch(arr, low, j)
	return j
}
```

对于整个快排算法，我们的 golang实现如下:
```
func quickSort(arr []int) []int {
    arr = shuffling(arr)
	sort(arr, 0, len(arr)-1)
	return arr
}

func sort(arr []int, low int, high int) {
	if high <= low {
		return
	}
	j := partition(arr, low, high)
	sort(arr, low, j-1)
	sort(arr, j+1, high)
}
func partition(arr []int, low int, high int) int {
	i, j := low+1, high
	for true {
		for arr[i] < arr[low] {
			i++
			if i == high {
				break
			}
		}
		for arr[low] < arr[j] {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		exch(arr, i, j)
	}
	exch(arr, low, j)
	return j
}
func exch(arr []int, a int, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}
```

时间复杂度：

最好情况 O(nlogn)
最坏情况 O(n^2)

由于我们添加了洗牌方法（见基本排序实现），所以整个数组趋向于无序，快速排序往往能取得比其他排序好的多的性能

------
参考文献：
[Dynamic Connectivity - 普林斯顿大学 | Coursera](https://www.coursera.org/learn/algorithms-part1/lecture/fjxHC/dynamic-connectivity)

