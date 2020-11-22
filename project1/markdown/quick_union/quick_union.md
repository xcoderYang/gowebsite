线性表中，栈和队列是非常重要的两种数据结构，本文将就这两种数据结构进行 golang语言实现

## 一.栈的实现

我们需要实现如下几个方法

1. push()&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;向栈中压入一个元素
2. pop()&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;从栈顶取出一个元素
3. isEmpty()&nbsp;&nbsp;&nbsp;&nbsp;判断栈是否为空
4. length()&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;获取栈中元素的数目
5. peer()&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;查询栈顶元素

我们需要注意 peer() 方法并不会将栈顶元素删除

数组实现如下：

```
type stack struct {
	cache []int
}

func (sk *stack) push(n int) {
	sk.cache = append(sk.cache, n)
}

func (sk stack) length() int {
	return len(sk.cache)
}
func (sk *stack) pop() int {
	if sk.length() == 0 {
		return 0
	}
	item := sk.cache[sk.length()-1]
	sk.cache = sk.cache[:len(sk.cache)-1]
	return item
}
func (sk stack) isEmpty() bool {
	return len(sk.cache) == 0
}
func (sk stack) peer() int {
	return sk.cache[sk.length()-1]
}
```

接下来，我们将用链表实现以下项目，并使用 interface{} 来代替 int实现多种类型的兼容

```

type stackLink struct {
	Top    *node
	Length int
}

type node struct {
	Val  interface{}
	Prev *node
}

func (sl *stackLink) push(value interface{}) {
	newNode := &node{
		Val:  value,
		Prev: sl.Top}
	sl.Top = newNode
	sl.Length++
}

func (sl *stackLink) pop() interface{} {
	topNodeVal := sl.Top.Val
	sl.Top = sl.Top.Prev
	sl.Length--
	return topNodeVal
}

func (sl stackLink) length() int {
	return sl.Length
}

func (sl stackLink) isEmpty() bool {
	return sl.Length == 0
}

func (sl stackLink) peer() interface{} {
	return sl.Top.Val
}
```
由于任何的变量都实现了空接口，所以我们可以通过传递空接口来实现在栈中压入不同元素的目的

## 二.队列实现
同样，我们对于队列，实现了如下方法：

 1. enqueue()&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;入队列
 2. dequeue()&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;出队列
 3. isEmpty()&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;判断队列是否为空
 4. getLength()&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;获队列的长度

链表实现方式如下：

```
type queue struct {
	First *node
	Last  *node
	Len   int
}

type node struct {
	Val  interface{}
	Next *node
	Pre  *node
}
func (qu *queue) enqueue(data interface{}) {
	nNode := &node{
		Val:  data,
		Pre:  qu.First,
		Next: nil}
	if qu.First == nil {
		qu.First = nNode
	} else {
		qu.First.Next = nNode
		qu.First = nNode
	}
	if qu.Last == nil {
		qu.Last = nNode
	}
	qu.Len++
}

func (qu *queue) dequeue() interface{} {
	if qu.Len > 0 {
		nNode := qu.Last.Val
		if qu.Last.Next != nil {
			qu.Last.Next.Pre = nil
		}
		qu.Last = qu.Last.Next
		qu.Len--
		return nNode
	}
	return errors.New("error")
}

func (qu queue) isEmpty() bool {
	return qu.Len <= 0
}

func (qu queue) getLength() int {
	return qu.Len
}

```
## 三.栈和队列的应用
在这一部分，我们通过栈来实现表达式的计算
例如：我们需要计算 (1+((2+3)\*(4\*5)))

我们维护两个栈，一个是值栈，一个是操作栈，我们在读取表达式的时候采取如下的策略：

 1. 如果遇到 '('，我们将忽略它
 2. 如果遇到数字，将其压入值栈
 3. 如果遇到操作符，将其压入操作栈
 4. 如果遇到 ')'，我们从值栈中取出两个值 n1和 n2，在操作栈中，我们取出一个操作符 op
 5. 我们进行 n2 op n1的操作（例如 n1 = 3，n2 = 9，op = '/'，我们将执行 9/3 ）
 6. 将所得的结果压入值栈

```
func stackCompute(str string) int {
	var (
		vsk  stackLink
		opsk stackLink
	)
	for _, v := range str {
		if v <= '9' && v >= '0' {
			vsk.push(int(v) - '0')
		} else if v == '+' || v == '-' || v == '*' || v == '/' {
			opsk.push(int(v))
		} else if v == ')' {
			n1 := vsk.pop().(int)
			n2 := vsk.pop().(int)
			op := opsk.pop().(int)
			var ans int
			switch op {
			case '+':
				ans = n2 + n1
			case '-':
				ans = n2 - n1
			case '*':
				ans = n2 * n1
			case '/':
				ans = n2 / n1
			}
			vsk.push(int(ans))
		}
	}
	for !opsk.isEmpty() {
		n1 := vsk.pop().(int)
		n2 := vsk.pop().(int)
		op := opsk.pop().(int)
		var ans int
		switch op {
		case '+':
			ans = n2 + n1
		case '-':
			ans = n2 - n1
		case '*':
			ans = n2 * n1
		case '/':
			ans = n2 / n1
		}
		vsk.push(int(ans))
	}
	char := vsk.pop().(int)
	return int(char)
}
```

我们如下调用 
```
stackCompute("(1+((2+3)\*(4\*5)))")
将会得到结果 101
```



------
参考文献：
[Dynamic Connectivity - 普林斯顿大学 | Coursera](https://www.coursera.org/learn/algorithms-part1/lecture/fjxHC/dynamic-connectivity)

