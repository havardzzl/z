package leetcode

import (
	"math/rand"
)

type RandomizedSet struct {
	r    map[int]int
	data []int
}

func ConstructorRS() RandomizedSet {
	return RandomizedSet{r: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.r[val]; ok {
		return false
	}
	l := len(this.data)
	this.r[val] = l
	this.data = append(this.data, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.r[val]; !ok {
		return false
	}
	idx := this.r[val]
	lastIdx := len(this.data) - 1
	this.data[idx] = this.data[lastIdx]
	this.r[this.data[idx]] = idx
	this.data = this.data[:lastIdx]
	delete(this.r, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.r))]
}
