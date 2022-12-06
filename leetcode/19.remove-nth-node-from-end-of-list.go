/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */package main

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
// if head == nil {
// 	return nil
// }
// if n <= 0 {
// 	return nil
// }
// current := head
// len := 0
// for current != nil {
// 	len++
// 	current = current.Next
// }
// if n > len {
// 	return head
// }
// if n == len {
// 	current = head
// 	head = head.Next
// 	current.Next = nil
// 	return head
// }
// current = head
// i := 0
// for current != nil {
// 	if i == len-n-1 {
// 		deleteNode := current.Next
// 		current.Next = current.Next.Next
// 		deleteNode.Next = nil
// 		break
// 	}
// 	i++
// 	current = current.Next
// }
// return head
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}

// @lc code=end
