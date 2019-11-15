/*
 * @lc app=leetcode id=21 lang=java
 *
 * [21] Merge Two Sorted Lists
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
  public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
    ListNode r, t;
    if (l1 == null && l2 == null) {
      return null;
    }
    if (l1 == null) {
      return l2;
    }
    if (l2 == null) {
      return l1;
    }
    if (l1.val > l2.val) {
      t = l1; l1 = l2; l2 = t;
    }
    r = l1;
    while (l1 != null && l2 != null) {
      while (l1.next != null && l1.next.val <= l2.val) {
        l1 = l1.next;
      }
      if (l1.next == null) {
        l1.next = l2;
        break;
      }
      t = l1.next;
      l1.next = l2; l1 = l2; l2 = t;
    }
    return r;
  }
}
// @lc code=end

