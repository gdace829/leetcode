package main

type MyStack struct {
	que1 []int
	que2 []int
}

func Constructor() MyStack {
	return MyStack{que1: make([]int, 0), que2: make([]int, 0)}
}

func (this *MyStack) Push(x int) {
	this.que2 = append(this.que2, x)
	for len(this.que1) > 0 {
		this.que2 = append(this.que2, this.que1[0])
		this.que1 = this.que1[1:]
	}
	this.que1, this.que2 = this.que2, this.que1

}

func (this *MyStack) Pop() int {
	x := this.que1[0]
	this.que1 = this.que1[1:]
	return x
}

func (this *MyStack) Top() int {
	return this.que1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.que1) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
