package main

import "fmt"

type MinStack struct {
	Nums []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Nums: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.Nums = append(this.Nums, x)
}

func (this *MinStack) Pop() {
	this.Nums = append([]int{}, this.Nums[0:len(this.Nums)-1]...)
}

func (this *MinStack) Top() int {
	return this.Nums[len(this.Nums)-1]
}

func (this *MinStack) GetMin() int {
	fmt.Println(this.Nums)
	if len(this.Nums) == 0 {
		return 0
	}
	var min = this.Nums[0]
	for _, v := range this.Nums {
		if min >= v {
			min = v
		}
	}
	return min
}

func main() {
	c := Constructor()
	c.Push(2)
	c.Push(0)
	c.Push(3)
	c.Push(0)
	a := c.GetMin()
	fmt.Println(a)
	c.Pop()
	b := c.GetMin()
	fmt.Println(b)
	c.Pop()
	d := c.GetMin()
	fmt.Println(d)
	c.Pop()
	e := c.GetMin()
	fmt.Println(e)
}
