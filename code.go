package main

import (
	"fmt"
	"sort"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	a := []int{}
	result := []int{}
	b := 0

	var tmp map[int]int
	tmp = make(map[int]int)
	for i, value := range arr2 {
		tmp[value] = i
	}
	for _, value := range arr1 {
		if _, ok := tmp[value]; ok {
		} else {
			a = append(a, value)
			b++
		}
	}
	sort.Ints(a)
	for _, value := range arr2 {
		for _, value2 := range arr1 {
			if value == value2 {
				result = append(result, value)
			}
		}
	}
	for _, value := range a {
		result = append(result, value)
	}
	return result
}
func reverseWords1(s string) string {
	split := strings.Split(s, " ")
	for i, value := range split {
		if i == 0 {
			split[i] = reverseString(value)
		} else {
			split[i] = " " + reverseString(value)
		}
	}
	return strings.Join(split, "")
}

func reverseWords2(s string) string {
	split := strings.Split(reverseString(s), " ")
	result := make([]string, len(split))
	for i := 0; i < len(split); i++ {
		if i == len(split)-1 {
			result[i] = split[len(split)-i-1]
		} else {
			result[i] = split[len(split)-i-1] + " "
		}
	}
	return strings.Join(result, "")
}

func reverseString(s string) string {
	a := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		a[i] = s[len(s)-i-1]
	}
	return string(a)
}

func test() []int {
	a := []int{1, 2, 3}
	for i, value := range a {
		a[i] = plusOne(value)
	}
	fmt.Print(a)
	return a
}

func plusOne(n int) int {
	return n + 1
}

func productExceptSelf(nums []int) []int {
	leftMulti := make([]int, len(nums))
	rightMulti := make([]int, len(nums))
	result := make([]int, len(nums))
	//左边的乘积
	leftMulti[0] = 1
	for i := 1; i < len(nums); i++ {
		leftMulti[i] = nums[i-1] * leftMulti[i-1]
	}
	//右边的乘积
	rightMulti[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightMulti[i] = nums[i+1] * rightMulti[i+1]
	}
	for i := 0; i < len(nums); i++ {
		result[i] = rightMulti[i] * leftMulti[i]
	}
	return result
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func removeElements(head *ListNode, val int) *ListNode {
	//删除开头的
	//for {
	//	if head != nil && head.Val == val {
	//		head = head.Next
	//	} else {
	//		break
	//	}
	//}
	for head != nil && head.Val ==val{
		head = head.Next
	}
	//记录原始head
	head1 := head
	//删除指定节点
	if head1!=nil {
		for head1.Next != nil {
			if head1.Next.Val == val {
				head1.Next = head1.Next.Next
			} else {
				head1 = head1.Next
			}
		}
	}
	return head
}

func tribonacci1(n int) int {
	list := make([]int, 38)
	list[0] = 0
	list[1] = 1
	list[2] = 1
	if n < 3 {
		return list[n]
	}else {
		for i := 3; i < n+1; i++ {
			list[i] = list[i-3] + list[i-2] + list[i-1]
		}
		return list[n]
	}
}
var a [38]int
func tribonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n==1 {
		return 1
	}
	if n==2 {
		return 1
	}else {
		a[n] = tribonacci2(n-1) + tribonacci2(n-2) + tribonacci2(n-3)
	}
	return a[n]
}
func main() {
	//arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	//arr2 := []int{2,1,4,3,9,6}
	//s := "ni hao"
	//arr := []int{1, 2, 3, 4}
	//array := relativeSortArray(arr1, arr2)
	fmt.Print(tribonacci2(35))
}
