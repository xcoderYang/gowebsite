堆排序与优先队列

## 一.优先队列

之前我们讲过队列这种数据结构，队列的特点是先进先出，那什么是优先队列呢？一般来说，优先队列分为两种，一种是最大优先队列，一种是最小优先队列

我们以最大优先队列为例来讲解今天的内容，最大优先队列的进出队列顺序并不是按照先进先出的准则，而是最大值先出

譬如队列 a = [3，1，5，4，2，0，6，9，7]，按照最大优先队列的出队列顺序，下一个值应该是 9，然后是 7

使用数组解决最大优先队列一般有两种方法：
1. 将数组排序，每次取出第一个值即可
2. 每次出队列都遍历数组，取出最大值

有序数组的入队时间复杂度为 O(n)，出队时间复杂度为 O(1)

无序数组的入队时间复杂度为 O(1)，出队时间复杂度为 O(n)

无论有序还是无序，每次出入队列的操作的时间复杂度都是 O(n)

pic1：

接下来我们介绍一种数据结构，它能够在优先队列的操作中达到 O(logn)的时间复杂度

### 堆
堆的本质是完全二叉树（也有多叉的堆，不在今天讨论范围内）。对比于普通的完全二叉树，堆（大顶堆）有如下特性：
1. 根节点为整个堆的最大值
2. 每个节点的值都不小于它的子节点

由于完全二叉树的特性，我们可以将堆使用数组存储起来，而不用担心浪费空间

我们假设使用 a数组表示一个大顶堆，则
 1. a[1]元素为整个堆的最大值（a[0]元素暂被省略）
 2. k节点的父节点为 k/2
 3. k节点的孩子节点为 2k和 2k+1

pic2:

在建堆的过程中，我们会遇到一些违反以上规则的情况，我们将通过以下两个函数来维护堆
 
 ```
func swim(k int, datas *[]int){
	for k>1 && less((*datas)[k/2],(*datas)[k]){
		exch(&(*datas)[k/2], &(*datas)[k])
		k = k/2
	}
}
```

```
func sink(k int, datas []int){
	N:=len(datas)-2
	for 2*k<=N{
		j:=2*k
		// 选取 j和 j+1中的最大值,假设最大值为 j，如果 j小于 j+1，则 j++变成 j+1
		// j这个变量永远指向子元素中的最大值
		if j<N && less(datas[j], datas[j+1]){
			j++
		}
		// 如果父节点比子节点要大，说明不用下沉了
		if !less(datas[k], datas[j]){
			break
		}
		// 下沉
		exch(&datas[k], &datas[j])
		k = j
	}
}
```

第一个函数是 swim主要用于元素入队，我们通过 swim方法找到新入队元素应该在的位置

第二个函数是 sink主要用于元素出队，我们在提取最大值时，会将队列末尾的元素与最大值（也就是队首）交换，然后下沉该元素，找到他应该在的位置

整个建堆的过程如下：

1. 在一个空堆中插入第一个元素，此时该堆是符合堆的性质的
2. 我们插入一个新元素，然后使用 swim 函数调整该新元素的位置
3. 循环进行第二步，直到所有元素均已插入，此时整个堆建立完毕

整个建堆的过程可以看成不断执行入队列的过程，由于每个元素只与其父元素比较，所以整个建堆过程的时间复杂度为 O(logn)

然后我们考虑出队列，出队的过程如下：

1. 由于队列的第一个值为整个队列的最大值，所以将队列的最后一个元素与队列的第一个元素互换，然后取出该最大值
2. 此时由于队列最后一个元素在队首，肯定是不符合堆的性质的，所以我们使用 sink下沉该元素，通过与子节点比较，找到该元素合适的位置

此时我们的最大元素已经被取出，而整个队列还维持这堆的性质。由于出队列下沉的过程中，每个元素只与该元素的子节点比较，所以整个出队列过程的时间复杂度也是 O(logn)

golang代码实现如下：

```
func main(){
	datas:=[]int{
		4,1,5,7,8,1,2,12,63,64,23,46,76,7,123,12535,3,
	}
	buildHead(datas)
}

func less(a int,b int) bool {
	return a<b
}

func exch(pa *int, pb *int){
	temp:=*pa
	*pa = *pb
	*pb = temp
}

func swim(k int, datas *[]int){
	for k>1 && less((*datas)[k/2],(*datas)[k]){
		exch(&(*datas)[k/2], &(*datas)[k])
		k = k/2
	}
}

func insert(x int, datas []int) []int{
	datas = append(datas,x)
	//fmt.Println(datas)
	swim(len(datas)-1, &datas)
	return datas
}

func sink(k int, datas []int){
	N:=len(datas)-2
	for 2*k<=N{
		j:=2*k
		// 选取 j和 j+1中的最大值,假设最大值为 j，如果 j小于 j+1，则 j++变成 j+1
		// j这个变量永远指向子元素中的最大值
		if j<N && less(datas[j], datas[j+1]){
			j++
		}
		// 如果父节点比子节点要大，说明不用下沉了
		if !less(datas[k], datas[j]){
			break
		}
		// 下沉
		exch(&datas[k], &datas[j])
		k = j
	}
}

func delMax(datas *[]int)int{
	max:=(*datas)[1]
	exch(&(*datas)[1], &(*datas)[len((*datas))-1])
	sink(1, (*datas))
	*datas = (*datas)[:len((*datas))-1]
	return max
}

func buildHead(datas []int){
	ans:=make([]int,1)
	for i:=0; i<len(datas); i++{
		ans = insert(datas[i], ans)
	}
	cache:=len(ans)
	for i:=0; i<cache-1; i++{
		fmt.Println(delMax(&ans))
	}
	return ans
}
```

除了大顶堆，还有小顶堆，也就是队首元素为整个队列的最小值，整体逻辑基本不变，只是在元素比较时需要转变符号

## 堆排序

有了以上优先队列和堆的基础，堆排序的概念也很容易理解了。

我们可以将整个待排序的数组维护成一个堆，然后不断的使用删除最大值的操作，此时整个堆的输出即为降序排列。如果维护小顶堆，则输出为升序队列

如果我们仔细观察，会发现其实我们不需要对队列中的每一个元素执行 sink操作，如果此时整个数组已经确定，我们只需要对前 n/2个元素执行 sink操作即可（因为每个 sink操作会使该元素与子节点比较，所以对第n/2个元素执行 sink操作时，第 n个元素也被加入比较中）

pic3:



------
参考文献：
[Dynamic Connectivity - 普林斯顿大学 | Coursera](https://www.coursera.org/learn/algorithms-part1/lecture/fjxHC/dynamic-connectivity)

