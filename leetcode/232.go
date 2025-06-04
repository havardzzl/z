package leetcode

type MyQueue struct {
	inStack  []int
	outStack []int
}

func ConstructorDQ() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) reallocate() {
	if len(this.outStack) == 0 && len(this.inStack) > 0 {
		for i := len(this.inStack) - 1; i >= 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i])
		}
		this.inStack = this.inStack[:0]
	}
}

func (this *MyQueue) Pop() int {
	this.reallocate()
	var ans int
	if len(this.outStack) > 0 {
		ans = this.outStack[len(this.outStack)-1]
		this.outStack = this.outStack[:len(this.outStack)-1]
	}
	return ans
}

func (this *MyQueue) Peek() int {
	this.reallocate()
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	this.reallocate()
	return len(this.outStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
