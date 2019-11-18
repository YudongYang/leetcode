/*
 * @lc app=leetcode id=55 lang=java
 *
 * [55] Jump Game
 */

// @lc code=start
class Solution {
  public boolean canJump(int[] nums) {
    // int[] r = new int[nums.length];
    // r[0] = 1;
    // for (int i = 0; i < (nums.length - 1); i++) {
    //   if (r[i] != 1) {
    //     continue;
    //   }
    //   int n = nums[i];
    //   for (int j = 1; j <= n && (i + j) < nums.length; j++) {
    //     r[(i + j)] = 1;
    //   }
    // }
    // return r[nums.length - 1] == 1;
    int p = nums.length - 1;
    for (int i = nums.length - 1; i >= 0; i--) {
      if ((i + nums[i]) >= p) {
        p = i;
      }
    }
    return p == 0;
  }
}
// @lc code=end

