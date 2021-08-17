# 
## golang的heap

golang的container包下的heap是golang实现堆的一种方式
heap主要代码如下
```
// 堆这个接口的定义如下
// 通过实现这个接口来定义一个堆
type Interface interface {
	sort.Interface //这是堆的主要数据类型，sort库中的一个接口。实现这个接口的结构都可以作为堆的存储结构。
	Push(x interface{}) // add x as element Len() 定义堆要是实现一个push方法，这个方法根据要定义的堆怎么弹出数据的逻辑来定义。
	Pop() interface{}   // remove and return element Len() - 1. 义堆要是实现一个push方法，这个方法根据要定义的堆怎么弹出数据的逻辑来定义。
}


```