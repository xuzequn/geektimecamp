package main

type MyCircularDeque struct {
	queue []int // 队列
	front int   // 头
	rear  int   // 尾
}

// 初始化
/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, k+1),
		front: 0,
		rear:  0,
	}

}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.front = (this.front + cap(this.queue) - 1) % cap(this.queue) //从队头插入先调整指针，
		this.queue[this.front] = value                                    //后赋值，
		return true
	}
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.queue[this.rear] = value                 //从队尾插入，先复制
		this.rear = (this.rear + 1) % cap(this.queue) //后调整指针

		return true
	}

}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.front = (this.front + 1) % cap(this.queue)
		return true
	}

}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.rear = (this.rear + cap(this.queue) - 1) % cap(this.queue)
		return true
	}
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.front]

}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[(this.rear+cap(this.queue)-1)%cap(this.queue)] // 数组模拟循环队列，指针取模处理

}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.front == this.rear {
		return true
	} else {
		return false
	}

}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if (this.rear+1)%cap(this.queue) == this.front {
		return true
	} else {
		return false
	}

}
