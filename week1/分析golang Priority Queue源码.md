# 
## golang的heap
golang 通过 heap 实现 Priority Queue。
golang的container包下的heap是golang实现堆的一种方式
heap主要代码如下
```
// 堆这个接口的定义如下
// 通过实现这个接口来定义一个堆
type Interface interface {
	sort.Interface //这是堆的主要数据类型，sort库中的一个接口。实现这个接口的结构都可以作为堆的存储结构。
	Push(x interface{}) // add x as element Len() 定义堆要是实现一个push方法，这个方法根据要定义的堆怎么添加数据的逻辑来定义。
	Pop() interface{}   // remove and return element Len() - 1. 定义堆要是实现一个pop方法，这个方法定义怎么推出一个数据的逻辑。
}

```
这个sort.interface 定义如下,根据需要可以重写下面的方法
```
type Interface interface {
	Len() int // 判断长度
	Less(i, j int) bool // 堆排序方式，比如，子节点比父节点小的大根堆，或者子节点比父节点大的小根堆 
	Swap(i, j int) // 堆中元素交换
}	
```

heap.init (h interface) 初始化堆，堆化过程
```
func Init(h Interface) {
	// heapify  堆化
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {  //从左边倒数第一个根节点开始维护堆
		down(h, i, n) // 保证堆按照定义的规则进行维护，从上往下进行调整
	}
}
```
heap.Push (h interface, x interface{}) 将一个堆元素压入堆中

```
func Push(h Interface, x interface{}) {
	h.Push(x) // 根据定义的Push规则添加一个元素到堆中，此时堆得数据结构中是非堆化，也就是非整理过的堆
	up(h, h.Len()-1) // 调整新加入元素根据堆的规则应该在什么位置，从下往上进行调整
}
```

heap.Pop (h Interface) interface{}
```
func Pop(h Interface) interface{} {
	n := h.Len() - 1 
	h.Swap(0, n) //最后一个元素 与第一个元素进行调换，此时堆是不符合定义的堆得规则的。
	down(h, 0, n)  // 从上往下进行调整，将第一个元素调整到合适的位置
	return h.Pop() // 根据定义的数据结构，将最后一位拿出，这不要根据定义堆的数据结构自己实现
}
``` 
heap.up(h Interface, j int), 自下而上调整堆
```
func up(h Interface, j int) { // h 是 堆， j是要调整的堆元素位置
	for {
		i := (j - 1) / 2 // parent // 找到当前元素的父节点
		if i == j || !h.Less(j, i) { // 如果与父节点值相等或者复核定义这堆得要求，比如 比父节点小的大根堆，或者比父节点大的小根堆 
			break // 退出调整
		}
		h.Swap(i, j) // 因为不符合定义堆的要求，与父节点交换。 
		j = i // 元素的新位置再次进行调整
	}
}
```

heap.down(h Interface, i0, n int) bool, 自上而下维护一个堆。
```
func down(h Interface, i0, n int) bool { // h是堆， i0是  要调整的元素， n是堆中元素的总数
	i := i0
	for {
		j1 := 2*i + 1 // 找到要调整元素的左节点
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow // 如果 左节点不存在 退出
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) { // 如果右节点存在，并且 与左节点相比， 选出最适合的堆定义的一个
			j = j2 // = 2*i + 2  // right child 右节点适合
		}
		if !h.Less(j, i) {  // 左右节点适合的，与根节点进行比较, 根节点合适退出
			break
		}
		h.Swap(i, j) // 左右节点合适，与根节点交换
		i = j // 调整元素变为根节点进行下一次调整，直到调整完成。
	}
	return i > i0
}
```