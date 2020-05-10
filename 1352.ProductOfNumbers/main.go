package main

type ProductOfNumbers struct {
	Nums []int
	Len  int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		Nums: make([]int, 0, 10),
		Len:  0,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	this.Nums = append(this.Nums, num)
	this.Len = len(this.Nums)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	var num = 1
	for i := 0; i < k; i++ {
		num = num * this.Nums[this.Len-1-i]
	}
	return num
}

func main() {
	//obj := Constructor()
}
