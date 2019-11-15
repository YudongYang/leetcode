/*
 * @lc app=leetcode id=19 lang=java
 *
 * [19] Remove Nth Node From End of List
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
  public ListNode removeNthFromEnd(ListNode head, int n) {
    ListNode cur = head, cut;
    int i = 0;
    while (i <= n && cur != null) {
      cur = cur.next;
      i++;
    }
    if (cur == null && i == n) {
      return head.next;
    }
    cut = head;
    while (cur != null) {
      cur = cur.next;
      cut = cut.next;
    }
    if (cut.next != null) {
      cut.next = cut.next.next;
    }
    return head;
  }
}
// @lc code=end

