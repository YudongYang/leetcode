/*
 * @lc app=leetcode id=136 lang=java
 *
 * [136] Single Number
 */

// @lc code=start
class Solution {
  public int singleNumber(int[] nums) {
    int r = 0;
    for (int i:nums) {
      r = r ^ i;
    }
    return r;
  }
}
// @lc code=end

