package gee

import "strings"

type node struct{
	pattern string
	part string
	children []*node
	isWild bool
}

func (n *node) matchChild(part string) *node {
	for _,child := range n.children{
		if child.part == part || child.isWild{
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node{
	nodes:=make([]*node, 0)
	for _,child := range n.children{
		if child.part == part || child.isWild{
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/**
addRoute中的调用：
insert(pattern, parts, 0)

pattern: /p/:lang/doc
parts: ["p", ":lang", "doc"]
height: 0->1->2->3
 */
/**
一层层的进行匹配与插入
先从 根节点开始，找到根节点孩子中匹配的节点，找不到则插入，找到则从该孩子中执行插入方法
 */
func (n *node) insert(pattern string, parts[]string, height int){
	// parts的数目等于 height时，其实已经到最底层了，此时在当前节点上加上 pattern
	if len(parts) == height{
		n.pattern = pattern
		return
	}
	// 取出当前层的 part
	part := parts[height]
	// 从 n的孩子中找出第一个匹配 part的
	child := n.matchChild(part)

	// 如果没有找到，则追加到 n的孩子集合中（insert）
	if child == nil{
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// 此层插入已完成，从找出的第一个孩子中插入下一层
	child.insert(pattern, parts, height + 1)
}

/**
getRoute中调用:
search(searchParts, 0)

parts: ["p",":lang","docs"]
height
 */
func (n *node) search(parts []string, height int) *node{
	// 如果匹配的最深层，判断 pattern是否有值，如果有值，说明有到当前节点的路由（insert中，路由的中间节点是不会添加 pattern的）
	// 或者判断当前节点是否为全匹配节点 "*"
	// 如果为是，则当前节点是匹配节点
	// 如果找到最深层，没找到 pattern，说明不存在到此节点的路由（当前路由只是中间节点）

	// 这里有一点，就是*号是最短匹配，也就是说遇到*号就直接返回，
	if len(parts) == height || strings.HasPrefix(n.part, "*"){
		if n.pattern == ""{
			return nil
		}
		return n
	}

	// 获取路由的当前部分
	part := parts[height]
	// 从当前节点中找到所有匹配 part的子节点集合
	children := n.matchChildren(part)

	// 遍历子节点集合，递归调用 search
	for _, child := range children{
		// 从当前子节点中寻找
		result := child.search(parts, height + 1)
		// 如果找到，则返回，否则继续寻找
		if result != nil{
			return result
		}
	}
	// 全部没有找到，则返回 nil
	return nil
}