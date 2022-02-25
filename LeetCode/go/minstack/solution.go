package minstack

import "container/list"

type MinStack struct {
	s    *list.List
	sMin *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		s:    list.New(),
		sMin: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.s.PushBack(x)
	minE := this.sMin.Back()
	if minE == nil {
		this.sMin.PushBack(x)
	} else {
		min := minE.Value.(int)
		if x < min {
			min = x
		}
		this.sMin.PushBack(min)
	}
}

func (this *MinStack) Pop() {
	if this.s.Len() > 0 {
		this.s.Remove(this.s.Back())
		this.sMin.Remove(this.sMin.Back())
	}
}

func (this *MinStack) Top() int {
	if this.s.Len() > 0 {
		return this.s.Back().Value.(int)
	}
	return -1
}

func (this *MinStack) GetMin() int {
	e := this.sMin.Back()
	if e != nil {
		return e.Value.(int)
	}
	return -1
}
