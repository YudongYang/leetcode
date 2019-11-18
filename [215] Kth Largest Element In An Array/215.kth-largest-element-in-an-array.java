/*
 * @lc app=leetcode id=215 lang=java
 *
 * [215] Kth Largest Element in an Array
 */

// @lc code=start
class Solution {
  public int findKthLargest(int[] nums, int k) {
    // int[] q = new int[k];
    // int i = 0;
    // for (i = 0; i < k; i++) {
    //   q[i] = nums[i];
    // }
    // Arrays.sort(q);
    // for (; i < nums.length; i++) {
    //   if (nums[i] < q[0]) {
    //     continue;
    //   } else {
    //     q[0] = nums[i];
    //     Arrays.sort(q);
    //   }
    // }
    // return q[0];
    Queue<Integer> q = new PriorityQueue<Integer>(k);
    for (int n:nums) {
      q.offer(n);
      if (q.size() > k) {
        q.poll();
      }
    }
    return q.peek();
  }
}
// @lc code=end

