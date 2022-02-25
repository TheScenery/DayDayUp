package yonglianggezhanshixianduilielcof

import "container/list"

type CQueue struct {
	s1 *list.List
	s2 *list.List
}

func Constructor() CQueue {
	cQueue := CQueue{
		s1: list.New(),
		s2: list.New(),
	}
	return cQueue
}

func (this *CQueue) AppendTail(value int) {
	this.s1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s2.Len() > 0 {
		r := this.s2.Back()
		this.s2.Remove(r)
		return r.Value.(int)
	}

	if this.s1.Len() > 0 {
		for this.s1.Len() > 0 {
			v := this.s1.Back()
			this.s1.Remove(v)
			this.s2.PushBack(v.Value)
		}
		return this.DeleteHead()
	}

	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
