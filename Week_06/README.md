学习笔记

# 总结

# 随笔

## 大概的一点感觉

- 子问题就是从0开始到n-1, n-2等问题的某种求和或求最值，这种dp状态数组和转移方程虽然也不好想，但最容易理解，dp初始化一般也都是一个二维数组（？）其dp\[i]\[0], dp\[0]\[j] 都是0，这些也可能是额外补出来辅助的，时间复杂度一般都是O(n),只要遍历一次就行了，dp一般也可以压缩成一个一维数组，运气好的还能压缩成几个变量
  - 偷房子问题
  - 最大矩形
  
- 子问题是一个数组里i, j区间之内的最优解的某种求和方式，是从中间向两边扩散的形态，这种一般会以ij作为维度制作二维dp数组，而数组初始化是以i=j,i+1=j这种形态开始，一般是一个上三角矩阵，从对角线开始递推，这种问题是要穷举ij的，所以时间复杂度至少是O(n^2)，然后因为递推的过程是斜着并且反复遍历的，所以dp数组也基本上不能压缩
  - 戳气球


  

# 提交作业
```
#学号:G20200343110147
#姓名:benben
#班级:14期1班3组
#语言:go
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/tree/golang/Week_06

```

# 实战

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [62 unique-paths](https://leetcode.com/problems/unique-paths/discuss/?currentPage=1&orderBy=most_votes&query=) | [不同路径](https://leetcode-cn.com/problems/unique-paths/)| 🟡 中等 | 数组、动态规划 | - |9.8✅|9.9|9.10|9.17||
| [63 unique-paths-ii](https://leetcode.com/problems/unique-paths-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [不同路径II](https://leetcode-cn.com/problems/unique-paths-ii/)| 🟡 中等 | 数组、动态规划 | - |9.8✅|9.9|9.10|9.17||
| [1143 longest-common-subsequence](https://leetcode.com/problems/longest-common-subsequence/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)| 🟡 中等 | 动态规划 | - |9.8✅|9.9|9.10|9.17||
| [120 triangle](https://leetcode.com/problems/triangle/discuss/?currentPage=1&orderBy=most_votes&query=) | [三角形最小路径和](https://leetcode-cn.com/problems/triangle/)| 🟡 中等 | 数组、动态规划 | - |9.9✅|9.10|9.11|9.18||
| [53 maximum-subarray](https://leetcode.com/problems/maximum-subarray/discuss/?currentPage=1&orderBy=most_votes&query=) | [最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)| 简单 | 数组、动态规划 | - |9.9✅|9.10|9.11|9.18||
| [152 maximum-product-subarray](https://leetcode.com/problems/maximum-product-subarray/discuss/?currentPage=1&orderBy=most_votes&query=) | [乘积最大子数组](https://leetcode-cn.com/problems/maximum-product-subarray/)| 🟡 中等 | 数组、动态规划 | - |9.9✅|9.10|9.11|9.18||
| [322 coin-change](https://leetcode.com/problems/coin-change/discuss/?currentPage=1&orderBy=most_votes&query=) | [零钱兑换](https://leetcode-cn.com/problems/coin-change/)| 🟡 中等 | 动态规划 | - |9.10✅|9.1|9.4|9.11||
| [198 house-robber](https://leetcode.com/problems/house-robber/discuss/?currentPage=1&orderBy=most_votes&query=) | [打家劫舍](https://leetcode-cn.com/problems/house-robber/)| 简单 | 动态规划 | - |9.10✅|9.1|9.4|9.11||
| [213 house-robber-ii](https://leetcode.com/problems/house-robber-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [打家劫舍II](https://leetcode-cn.com/problems/house-robber-ii/)| 🟡 中等 | 动态规划 | - |9.10✅|9.1|9.4|9.11||
| [121 best-time-to-buy-and-sell-stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)| 简单 | 数组、动态规划 | - |9.10✅|9.1|9.4|9.11||
| [122 best-time-to-buy-and-sell-stock-ii](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)| 简单 | 贪心算法、动态规划 | - |9.7|9.1|9.4|9.11||
| [123 best-time-to-buy-and-sell-stock-iii](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)| 🔴 困难 | 数组、动态规划 | - |9.7|9.1|9.4|9.11||
| [309 best-time-to-buy-and-sell-stock-with-cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/discuss/?currentPage=1&orderBy=most_votes&query=) | [最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)| 🟡 中等 | 数组、动态规划 | - |9.7|9.1|9.4|9.11||
| [188 best-time-to-buy-and-sell-stock-iv](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)| 🔴 困难 | 数组、动态规划 | - |9.7|9.1|9.4|9.11||
| [714 best-time-to-buy-and-sell-stock-with-transaction-fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)| 🟡 中等 | 数组、动态规划 | - |9.7|9.1|9.4|9.11||

# HomeWork

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [64 minimum-path-sum](https://leetcode.com/problems/minimum-path-sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)| 🟡 中等 | 数组、动态规划 | - |9.7|9.1|9.4|9.11||
| [91 decode-ways](https://leetcode.com/problems/decode-ways/discuss/?currentPage=1&orderBy=most_votes&query=) | [解码方法](https://leetcode-cn.com/problems/decode-ways/)| 🟡 中等 | 字符串、动态规划 | - |9.7|9.1|9.4|9.11||
| [221 maximal-square](https://leetcode.com/problems/maximal-square/discuss/?currentPage=1&orderBy=most_votes&query=) | [最大正方形](https://leetcode-cn.com/problems/maximal-square/)| 🟡 中等 | 动态规划 | - |9.7|9.1|9.4|9.11||
| [621 task-scheduler](https://leetcode.com/problems/task-scheduler/discuss/?currentPage=1&orderBy=most_votes&query=) | [任务调度器](https://leetcode-cn.com/problems/task-scheduler/)| 🟡 中等 | 贪心算法、队列、数组 | - |9.7|9.1|9.4|9.11||
| [647 palindromic-substrings](https://leetcode.com/problems/palindromic-substrings/discuss/?currentPage=1&orderBy=most_votes&query=) | [回文字串](https://leetcode-cn.com/problems/palindromic-substrings/)| 🟡 中等 | 字符串、动态规划 | - |9.7|9.1|9.4|9.11||
| [32 longest-valid-parentheses](https://leetcode.com/problems/longest-valid-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/)| 🔴 困难  | 字符串、动态规划 | - |9.7|9.1|9.4|9.11||
| [72 edit-distance](https://leetcode.com/problems/edit-distance/discuss/?currentPage=1&orderBy=most_votes&query=) | [编辑距离](https://leetcode-cn.com/problems/edit-distance/)| 🔴 困难  | 字符串、动态规划 | - |9.7|9.1|9.4|9.11||
| [363 max-sum-of-rectangle-no-larger-than-k](https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/discuss/?currentPage=1&orderBy=most_votes&query=) | [矩形区域不超过 K 的最大数值和](https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/)| 🔴 困难  | 队列、二分查找、动态规划 | - |9.7|9.1|9.4|9.11||
| [403 frog-jump](https://leetcode.com/problems/frog-jump/discuss/?currentPage=1&orderBy=most_votes&query=) | [青蛙过河](https://leetcode-cn.com/problems/frog-jump/)| 🔴 困难  | 动态规划 | - |9.7|9.1|9.4|9.11||
| [410 split-array-largest-sum](https://leetcode.com/problems/split-array-largest-sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [分割数组的最大值](https://leetcode-cn.com/problems/split-array-largest-sum/)| 🔴 困难  | 二分查找、动态规划 | - |9.7|9.1|9.4|9.11||
| [552 student-attendance-record-ii](https://leetcode.com/problems/student-attendance-record-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [学生出勤记录II](https://leetcode-cn.com/problems/student-attendance-record-ii/)| 🔴 困难  | 动态规划 | - |9.7|9.1|9.4|9.11||
| [76 minimum-window-substring](https://leetcode.com/problems/minimum-window-substring/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/)| 🔴 困难  | 双指针、滑动窗口 | - |9.7|9.1|9.4|9.11||
| [312 burst-balloons](https://leetcode.com/problems/burst-balloons/discuss/?currentPage=1&orderBy=most_votes&query=) | [戳气球](https://leetcode-cn.com/problems/burst-balloons/)| 🔴 困难 | 分治算法、动态规划 | - |9.7|9.1|9.4|9.11||
| [62 unique-paths](https://leetcode.com/problems/unique-paths/discuss/?currentPage=1&orderBy=most_votes&query=) | [不同路径](https://leetcode-cn.com/problems/unique-paths/)| 🔴 困难  | 数组、动态规划 | - |9.7|9.1|9.4|9.11||
