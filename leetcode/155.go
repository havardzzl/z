package leetcode

type MinStack struct {
	miniRecord [][2]int
	stack      []int
}

func StackConstructor() MinStack {
	return MinStack{miniRecord: [][2]int{}, stack: make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.miniRecord) == 0 || val < this.miniRecord[len(this.miniRecord)-1][1] {
		this.miniRecord = append(this.miniRecord, [2]int{len(this.stack) - 1, val})
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	if this.miniRecord[len(this.miniRecord)-1][0] > len(this.stack)-1 {
		this.miniRecord = this.miniRecord[:len(this.miniRecord)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.miniRecord[len(this.miniRecord)-1][1]
}
