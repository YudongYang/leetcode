/*
 * @lc app=leetcode id=2 lang=java
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
  public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
    ListNode l3 = null;
    ListNode result = null;
    int a, b, x = 0, y;
    while (l1 != null || l2 != null) {
      a = l1 != null ? l1.val : 0;
      b = l2 != null ? l2.val : 0;
      y = (x + a + b) % 10;
      if (l3 == null) {
        l3 = new ListNode(y);
        result = l3;
      } else {
        l3.next = new ListNode(y);
        l3 = l3.next;
      }
      x = (x + a + b) / 10;
      l1 = l1 != null ? l1.next : null;
      l2 = l2 != null ? l2.next : null;
    }
    if (x != 0) {
      l3.next = new ListNode(x);
    }
    return result;
  }
}
// @lc code=end

