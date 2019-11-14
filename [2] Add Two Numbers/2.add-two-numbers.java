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
    int a, b, x = 0, y, r;
    while (l1 != null || l2 != null) {
      if (l1 != null) {
        a = l1.val; l1 = l1.next;
      } else {
        a = 0; l1 = null;
      }
      if (l2 != null) {
        b = l2.val; l2 = l2.next;
      } else {
        b = 0; l2 = null;
      }
      r = x + a + b;
      y = r % 10;
      if (l3 != null) {
        l3.next = new ListNode(y);
        l3 = l3.next;
      } else {
        l3 = new ListNode(y);
        result = l3;
      }
      x = r / 10;
    }
    if (x != 0) {
      l3.next = new ListNode(x);
    }
    return result;
  }
}
// @lc code=end

