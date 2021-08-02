package main

import (
	"container/list"
	"reflect"
)

type MinStack struct {
	MinEleStack
	l *list.List
}

type MinEleStack struct {
	li *list.List
}


func Constructor() MinStack {
	return MinStack{
		MinEleStack: MinEleStack{li: list.New()},
		l:           list.New(),
	}
}


func (this *MinStack) Push(val int)  {
	this.l.PushBack(val)
	if this.MinEleStack.li.Len() == 0 {
		this.MinEleStack.li.PushBack(val)
	} else {
		min := this.MinEleStack.li.Back().Value.(int)
		if min > val {
			this.MinEleStack.li.PushBack(val)
		}
	}
}


func (this *MinStack) Pop()  {
	ele := this.l.Back()
	if min := this.MinEleStack.li.Back(); reflect.DeepEqual(ele.Value, min.Value) {
		this.MinEleStack.li.Remove(min)
	}
	this.l.Remove(ele)
}


func (this *MinStack) Top() int {
	return this.l.Back().Value.(int)
}


func (this *MinStack) GetMin() int {
	if this.MinEleStack.li.Len() == 0 {
		return 0
	}
	return this.MinEleStack.li.Back().Value.(int)
}
