接下来，我们会讨论快速排序的更多细节

## 标志位的选取

在上篇博文中，我们讲到了标志位的选取一般是取数组第一个元素，但是由于快排的本质是通过O(n)的时间排序好一个数（标志位）以及划分两个相对大小的数组（小于标志位的数组和大于标志位的数组），如果标志位是整个数组的最大值或者最小值，会导致本轮排序在划分数组时不能获得优质的结果

什么样的标志位才是效率最高的呢？

1. 最好的标志位是数组的中位数，本轮排序会将数组划分为相同大小的两个子数组，由此每轮划分的元素最多，排序效果最好
2. 中位数的选取最快也需要 O(n) 的时间，每轮排序取额外选取中位数得不偿失，学过统计学，我们知道，可以用样本的中位数去估计真实的中位数

基于以上两点，我们可以试试三分取中法

### 三分取中法

三分取中法的大致思路如下：

首先，我们找出数组中位置在第一个的值，最后一个的值和最中间的值

然后，我们选取上述三个数中的中位数作为我们本轮排序的标志位

这样选择的标志位虽然不一定是整个数组的中位数，但是比起选择第一个元素，要科学不少，而且在数组基本有序时，也能有效的避免最差的时间复杂度。

如果对于标志位的选取更加严格，我们可以从 2*(n-1)位中选取中位数，n的数值越大，中位数越有效，但是选取中位数所消耗的时间也越多

## 重复键值

我们在使用快速排序时，有时候会遇到许多元素相等的项，譬如

1. 按照年龄排序
2. 邮件列表中的重复邮件
3. 根据地名排序

而这类情况往往有几个特征

1. 数组元素很多
2. 相对于数组元素的个数，元素的种类很少（很大一部分元素相等）

特别，当数组中的元素全相同时，如果我们在元素项与标志位相等时继续排序，时间复杂度会骤升至 O(n^2)，如果我们在元素项与标志位相等时停止，也会产生 O(nlgn) 的时间复杂度。

接下来我们介绍快排的一种优化变种，在数组元素基本相等时，可以取得近似 O(n)的排序效果

### 三路快排





