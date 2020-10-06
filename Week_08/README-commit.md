学习笔记

# 总结
- 常用位运算，背吧
  - golang 取反， ^x
  - 除以2，乘以2，不用背了吧
  - 奇偶 if x&1 == 1; then odd ; else even
  - 右侧(低位)n位清零  x & (^0 << n)
  - 左侧(高位)清零至n位 x & ((1<<n) - 1 )
  - 得到最低位1(的二进制表示) x & (-x)  ???
  - 干掉最低位1 x & ( x - 1)
  - 干掉自己 x &^ x 或 x & (^x)
  
- 模板

|模板|笔记|理解|历史|
|---|---|---|---|
|双向BFS|https://shimo.im/docs/JpK8hPCgqjxxTX3k/||week07|
|并查集|https://shimo.im/docs/xtKGQWhYwJrCRtJP/||week07|
|Trie树|https://shimo.im/docs/RTr33Rxtgc6tcXxv/ ||week07|
|DFS递归|https://shimo.im/docs/9CYPpdcPGwXT93QV/|实际上回溯是带数据处理的DFS，在递归前后调用数据的处理和revert就成回溯了，另外回溯更侧重考虑剪枝|week04|
|DFS迭代|https://shimo.im/docs/JqXDvhW9jt6Y9hQV/ |这就是树的先序遍历模板 |week04|
|BFS|https://shimo.im/docs/VkVGpccqqxwvqtgR/|用队列|week04|
|二分查找|https://shimo.im/docs/vkPwvRcktgHdWWdW/|细节需要研究，都是边界问题|week04|
|回溯|https://shimo.im/docs/3c3VYtCW9kyCcGYc|模板很好用|week03|
|单调栈单调队列|https://shimo.im/docs/hyWwqQ39xVcwk3xG/| 理解的还凑合，记得不牢|week02|

- 某些场景可以显著优化性能的点
  - golang遍历字符串时，直接for string比for []byte(string)要快，注意range string得到的是rune，要得到byte需要用string\[i]
  - 给map，slice设置capacity，减少扩容次数
  - 剪枝、双向BFS什么的大套路
  
- 还没练习过的类型
  - 二分查找不等于的系列
  - 
  
# 随笔

# 提交作业

```
#学号:G20200343110147
#姓名:benben
#班级:14期1班3组
#语言:go
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/blob/golang/Week_08
```

# 实战

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [773 sliding](https://leetcode.com/problems/sliding/discuss/?currentPage=1&orderBy=most_votes&query=) | [题目模板](https://leetcode-cn.com/problems/sliding/)| 🔴 困难 🟡 中等 简单| 广度优先搜索 | - |9.26✅|9.27|9.29|10.4||


# HomeWork

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [191 number-of-1-bits](https://leetcode.com/problems/number-of-1-bits/discuss/?currentPage=1&orderBy=most_votes&query=) | [位1的个数](https://leetcode-cn.com/problems/number-of-1-bits/)| 简单| 位运算 | - |9.27✅|10.2✅|10.4|10.11||
| [231 power-of-two](https://leetcode.com/problems/power-of-two/discuss/?currentPage=1&orderBy=most_votes&query=) | [2的幂](https://leetcode-cn.com/problems/power-of-two/)| 简单| 位运算 | - |9.27✅|10.2✅|10.4|10.11||
| [190 reverse-bits](https://leetcode.com/problems/reverse-bits/discuss/?currentPage=1&orderBy=most_votes&query=) | [颠倒二进制位](https://leetcode-cn.com/problems/reverse-bits/)| 简单| 广度优先搜索 | - |9.27✅|10.2✅|10.4|10.11||
| [1122 relative-sort-array](https://leetcode.com/problems/relative-sort-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [数组的相对排序](https://leetcode-cn.com/problems/relative-sort-array/)| 简单| 排序、数组 | - |9.27✅|10.2✅|10.4|10.11||
| [242 valid-anagram](https://leetcode.com/problems/valid-anagram/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)| 简单| 排序、哈希表 | - |9.27✅|10.2✅|10.4|10.11||
| [146 lru-cache](https://leetcode.com/problems/lru-cache/discuss/?currentPage=1&orderBy=most_votes&query=) | [LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/)| 🟡 中等 | 设计 | - |9.27✅|10.2✅|10.4|10.11||
| [1244 design-a-leaderboard](https://leetcode.com/problems/design-a-leaderboard/discuss/?currentPage=1&orderBy=most_votes&query=) | [力扣排行榜](https://leetcode-cn.com/problems/design-a-leaderboard/)| 🟡 中等 | 排序、设计 | - |10.1✅|10.2✅|10.4|10.11||
| [56 merge-intervals](https://leetcode.com/problems/merge-intervals/discuss/?currentPage=1&orderBy=most_votes&query=) | [合并区间](https://leetcode-cn.com/problems/merge-intervals/)| 🟡 中等 | 排序、数组 | - |9.29✅|10.2✅|10.4|10.11||
| [51 n-queens](https://leetcode.com/problems/n-queens/discuss/?currentPage=1&orderBy=most_votes&query=) | [N 皇后](https://leetcode-cn.com/problems/n-queens/)| 🔴 困难 | 回溯算法、位运算 | - |10.1✅|10.2✅|10.4|10.11||
| [52 n-queens-ii](https://leetcode.com/problems/n-queens-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [N皇后 II](https://leetcode-cn.com/problems/n-queens-ii/)| 🔴 困难| 回溯算法、位运算 | - |10.1✅|10.2✅|10.4|10.11||
| [493 reverse-pairs](https://leetcode.com/problems/reverse-pairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [翻转对](https://leetcode-cn.com/problems/reverse-pairs/)| 🔴 困难 | 排序、分治算法 | - |9.30✅|10.2✅|10.4|10.11||
| [338 counting-bits](https://leetcode.com/problems/counting-bits/discuss/?currentPage=1&orderBy=most_votes&query=) | [比特位计数](https://leetcode-cn.com/problems/counting-bits/)| 🟡 中等 | 位运算、动态规划 | - |10.1✅|10.2✅|10.4|10.11||

# 常错和容易忘记的题

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [17 letter-combinations-of-a-phone-number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/?currentPage=1&orderBy=most_votes&query=) | [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)| 🟡 中等 | 分治、回溯 | 有个边界问题 |  | | | ||
| [84 largest-rectangle-in-histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/discuss/?currentPage=1&orderBy=most_votes&query=) | [柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)| 🔴️ 困难 | 栈、队列 | -| | | | ||
| [239 sliding-window-maximum](https://leetcode.com/problems/sliding-window-maximum/discuss/?currentPage=1&orderBy=most_votes&query=) | [滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)| 🔴️ 困难 | 栈、队列 | -| | | | ||
| [42 trapping-rain-water](https://leetcode.com/problems/trapping-rain-water/discuss/?currentPage=1&orderBy=most_votes&query=) | [接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)| 🔴️ 困难 | 栈、队列 |-| | | | ||
| [297 serialize-and-deserialize-binary-tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)| 🔴️ 困难 | 泛型递归、树的递归 | - | | | | ||
| [264 ugly-number-ii](https://leetcode.com/problems/ugly-number-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [丑数](https://leetcode-cn.com/problems/ugly-number-ii/)| 🟡 中等 | 泛型递归、树的递归 | 背背背，理解不了 | | | | ||
| [347 top-k-frequent-elements](https://leetcode.com/problems/top-k-frequent-elements/discuss/?currentPage=1&orderBy=most_votes&query=) | [前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)| 🟡 中等 | 泛型递归、树的递归 | - | | | | ||
| [77 combinations](https://leetcode.com/problems/combinations/discuss/?currentPage=1&orderBy=most_votes&query=) | [组合](https://leetcode-cn.com/problems/combinations/)| 🟡 中等 | 泛型递归、树的递归 | 组合问题dfs里只有一种情况| | | | ||
| [47 permutations-ii](https://leetcode.com/problems/permutations-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [全排列 II](https://leetcode-cn.com/problems/permutations-ii/)| 🟡 中等 | 回溯算法 | 去重问题先排序！！！ | | | | ||
| [40 combination-sum-ii](https://leetcode.com/problems/combination-sum-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/)| 🟡 中等 | 数组、回溯算法 | 向下drilldown时传的是i+1 || | | ||
| [5 longest-palindromic-substring](https://leetcode.com/problems/longest-palindromic-substring/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)| 🟡 中等 | 字符串、动态规划 | 有个加间隔符的套路 || | | ||

# 可进一步练习的题

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [148 sort-list](https://leetcode.com/problems/sort-list/discuss/?currentPage=1&orderBy=most_votes&query=) | [排序链表](https://leetcode-cn.com/problems/sort-list/)| 🟡 中等 | 排序、链表 | 把递归改为迭代以降低空间复杂度 |  | | | ||


