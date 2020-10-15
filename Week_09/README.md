学习笔记

# 总结

- 高级动态规划

- 字符串算法

- 实用库方法，既是单题也是其他题目的关键步骤

|名称|单题|适用题目||
|---|---|---|---|
|合并两个有序链表/数组|21，88|归并排序，23合并k个有序链表||
|两三数之和|1，15|16最接近的三数之和，18四数之和，259较小的三数之和||
|移动零|283|快速排序，魔术索引等||
|二叉树遍历系列||||
|反转链表|206|反转链表II,25k个一组翻转链表||
|前缀和，前缀min||560和为k的子数组，1-3-2数||
  
  
- 模板

|模板|笔记|理解|历史|
|---|---|---|---|
|双向BFS|https://shimo.im/docs/JpK8hPCgqjxxTX3k/|单词接龙II，就用这个练|week07|
|并查集|https://shimo.im/docs/xtKGQWhYwJrCRtJP/||week07|
|Trie树|https://shimo.im/docs/RTr33Rxtgc6tcXxv/ ||week07|
|DFS递归|https://shimo.im/docs/9CYPpdcPGwXT93QV/|实际上回溯是带数据处理的DFS，在递归前后调用数据的处理和revert就成回溯了，另外回溯更侧重考虑剪枝|week04|
|DFS迭代|https://shimo.im/docs/JqXDvhW9jt6Y9hQV/ |这就是树的先序遍历模板 |week04|
|BFS|https://shimo.im/docs/VkVGpccqqxwvqtgR/|用队列|week04|
|二分查找|https://shimo.im/docs/vkPwvRcktgHdWWdW/|细节需要研究，都是边界问题|week04|
|回溯|https://shimo.im/docs/3c3VYtCW9kyCcGYc|模板很好用|week03|
|单调栈单调队列|https://shimo.im/docs/hyWwqQ39xVcwk3xG/| 理解的还凑合，记得不牢|week02|

# 随笔

- 学习专题
  - 背包九讲 http://www2.lssh.tp.edu.tw/~hlf/class-1/lang-c/DP.pdf
  - 算法系列，有时间阅读
    - Deer系列: https://leetcode-cn.com/problems/longest-consecutive-sequence/solution/tao-lu-jie-jue-zui-chang-zi-xu-lie-deng-yi-lei-wen/
  - 正则表达式、通配符匹配算法
  - KMP算法

- 还没练习过的类型以及需要找时间仔细分析的东西
  - 二分查找不等于的系列
    - 二分查找最后一个小于等于目标的位置， done
    - 二分查找第一个大于等于目标的位置， done
  - 归并排序的复杂度计算方式
  - 快排的复杂度计算方式
  - 布隆过滤器需要练习一下吗？
  - dp的递归加备忘录写法

# 提交作业

```
#学号:G20200343110147
#姓名:benben
#班级:14期1班3组
#语言:go
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/blob/golang/Week_09/README-commit.md
```

# 实战

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [1143 longest-common-subsequence](https://leetcode.com/problems/longest-common-subsequence/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)| 🟡 中等| 动态规划 | - |10.11✅|10.12|10.17|10.24||
| [72 edit-distance](https://leetcode.com/problems/edit-distance/discuss/?currentPage=1&orderBy=most_votes&query=) | [编辑距离](https://leetcode-cn.com/problems/edit-distance/)| 🔴 困难 | 字符串、动态规划 | - |10.11✅|10.12|10.17|10.24||
| [5 longest-palindromic-substring](https://leetcode.com/problems/longest-palindromic-substring/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)|🟡 中等| 字符串、动态规划 | - |10.11✅|10.12|10.17|10.24||
| [773 sliding](https://leetcode.com/problems/sliding/discuss/?currentPage=1&orderBy=most_votes&query=) | [题目模板](https://leetcode-cn.com/problems/sliding/)| 🔴 困难 🟡 中等 简单| 广度优先搜索 | - |10.11✅|10.12|10.17|10.24||

# HomeWork

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [387 first-unique-character-in-a-string](https://leetcode.com/problems/first-unique-character-in-a-string/discuss/?currentPage=1&orderBy=most_votes&query=) | [字符串中的第一个唯一字符](https://leetcode-cn.com/problems/first-unique-character-in-a-string/)| 简单| 哈希表、字符串 | - |10.11✅|10.12|10.17|10.24||
| [541 reverse-string-ii](https://leetcode.com/problems/reverse-string-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [反转字符串 II](https://leetcode-cn.com/problems/reverse-string-ii/)|简单| 字符串 | - |10.11✅|10.12|10.17|10.24||
| [151 reverse-words-in-a-string](https://leetcode.com/problems/reverse-words-in-a-string/discuss/?currentPage=1&orderBy=most_votes&query=) | [翻转字符串里的单词](https://leetcode-cn.com/problems/reverse-words-in-a-string/)|简单| 字符串 | - |10.11✅|10.12|10.17|10.24||
| [557 reverse-words-in-a-string-iii](https://leetcode.com/problems/reverse-words-in-a-string-iii/discuss/?currentPage=1&orderBy=most_votes&query=) | [反转字符串中的单词 III](https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/)|简单| 字符串 | - |10.11✅|10.12|10.17|10.24||
| [917 reverse-only-letters](https://leetcode.com/problems/reverse-only-letters/discuss/?currentPage=1&orderBy=most_votes&query=) | [仅仅反转字母](https://leetcode-cn.com/problems/reverse-only-letters/)|简单| 字符串 | - |10.11✅|10.12|10.17|10.24||
| [205 isomorphic-strings](https://leetcode.com/problems/isomorphic-strings/discuss/?currentPage=1&orderBy=most_votes&query=) | [同构字符串](https://leetcode-cn.com/problems/isomorphic-strings/)|简单| 字符串 | - |10.11✅|10.12|10.17|10.24||
| [680 valid-palindrome-ii](https://leetcode.com/problems/valid-palindrome-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [验证回文字符串 Ⅱ](https://leetcode-cn.com/problems/valid-palindrome-ii/)|简单| 字符串 | - |10.11✅|10.12|10.17|10.24||
| [63 unique-paths-ii](https://leetcode.com/problems/unique-paths-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/)| 🟡 中等 | 数组、动态规划 | - |10.11|10.12|10.17|10.24||
| [300 longest-increasing-subsequence](https://leetcode.com/problems/longest-increasing-subsequence/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)| 🟡 中等 | 二分查找、动态规划 | - |10.11|10.12|10.17|10.24||
| [91 decode-ways](https://leetcode.com/problems/decode-ways/discuss/?currentPage=1&orderBy=most_votes&query=) | [解码方法](https://leetcode-cn.com/problems/decode-ways/)| 🟡 中等 | 字符串、动态规划 | - |10.11|10.12|10.17|10.24||
| [8 string-to-integer-atoi](https://leetcode.com/problems/string-to-integer-atoi/discuss/?currentPage=1&orderBy=most_votes&query=) | [字符串转换整数 (atoi)](https://leetcode-cn.com/problems/string-to-integer-atoi/)| 🟡 中等 | 字符串 | - |10.11✅|10.12|10.17|10.24||
| [438 find-all-anagrams-in-a-string](https://leetcode.com/problems/find-all-anagrams-in-a-string/discuss/?currentPage=1&orderBy=most_votes&query=) | [找到字符串中所有字母异位词](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)| 🟡 中等 | 字符串 | - |10.11✅|10.12|10.17|10.24||
| [5 longest-palindromic-substring](https://leetcode.com/problems/longest-palindromic-substring/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)| 🟡 中等 | 字符串 | - |10.11✅|10.12|10.17|10.24||
| [32 longest-valid-parentheses](https://leetcode.com/problems/longest-valid-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/)| 🔴 困难 | 字符串、动态规划 | - |10.11✅|10.12|10.17|10.24||
| [818 race-car](https://leetcode.com/problems/race-car/discuss/?currentPage=1&orderBy=most_votes&query=) | [赛车](https://leetcode-cn.com/problems/race-car/)| 🔴 困难| 动态规划 | - |10.15✅|10.18|10.23|10.30||
| [10 regular-expression-matching](https://leetcode.com/problems/regular-expression-matching/discuss/?currentPage=1&orderBy=most_votes&query=) | [正则表达式匹配](https://leetcode-cn.com/problems/regular-expression-matching/)| 🔴 困难| 字符串、动态规划、回溯算法 | - |10.13✅|10.14|10.17|10.24||
| [44 wildcard-matching](https://leetcode.com/problems/wildcard-matching/discuss/?currentPage=1&orderBy=most_votes&query=) | [通配符匹配](https://leetcode-cn.com/problems/wildcard-matching/)| 🔴 困难|  贪心算法、字符串、动态规划、回溯算法  | - |10.13✅|10.14|10.17|10.24||
| [115 distinct-subsequences](https://leetcode.com/problems/distinct-subsequences/discuss/?currentPage=1&orderBy=most_votes&query=) | [不同的子序列](https://leetcode-cn.com/problems/distinct-subsequences/)| 🔴 困难| 字符串、动态规划 | - |10.13✅|10.14|10.17|10.24||


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
| [23 merge-k-sorted-lists](https://leetcode.com/problems/merge-k-sorted-lists/discuss/?currentPage=1&orderBy=most_votes&query=) | [合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/)| 🔴️ 困难 | 排序、链表 | 堆、链表、分治算法 | 试试用堆做一遍 | | | ||
| [300 longest-increasing-subsequence](https://leetcode.com/problems/longest-increasing-subsequence/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长上升子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/)| 🟡 中等 | 二分查找、动态规划 | 数组套路，背背 | | | | ||
| [16.16 sub-sort-lcci](https://leetcode.com/problems/sub-sort-lcci/discuss/?currentPage=1&orderBy=most_votes&query=) | [部分排序](https://leetcode-cn.com/problems/sub-sort-lcci/)| 🟡 中等 | 排序、数组 | 数组套路，背背 | | | | ||
| [456 132-pattern](https://leetcode.com/problems/132-pattern/discuss/?currentPage=1&orderBy=most_votes&query=) | [132模式](https://leetcode-cn.com/problems/132-pattern/)| 🟡 中等 | 栈 | 单调栈练习 | | | | ||
| [394 decode-string](https://leetcode.com/problems/decode-string/discuss/?currentPage=1&orderBy=most_votes&query=) | [字符串解码](https://leetcode-cn.com/problems/decode-string/)| 🟡 中等 | 栈、递归 |  | | | | ||
