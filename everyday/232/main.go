package main

// 一个是加入的结构向正确结构转移, 一个是正确结构向加入结构转移
type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{stack1: make([]int, 0), stack2: make([]int, 0)}
}

func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	x := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[:len(this.stack2)-1]
	return x

}

func (this *MyQueue) Peek() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) != 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	x := this.stack2[len(this.stack2)-1]
	return x
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
