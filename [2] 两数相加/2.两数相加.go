/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := l1
	cur := l1
	mov := 0
	for l1 != nil && !(l2 == nil && mov == 0) {
		val1, val2 := l1.Val, 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		l1.Val = val1 + val2 + mov
		mov = l1.Val / 10
		l1.Val %= 10
		if l1.Next == nil && l2 != nil {
			l1.Next = l2
			l2 = nil
		}
		cur = l1
		l1 = l1.Next
	}
	if mov > 0 {
		cur.Next = &ListNode{
			Val:  mov,
			Next: nil,
		}
	}
	return result
}

// @lc code=end

