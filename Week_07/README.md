学习笔记

# 总结

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



# 随笔

- 双向BFS
  - 发现BFS完全可以不用queue来实现了。。
  - 就直接整一个数组从头遍历，结果丢到另一个数组里，另一个数组就是下一层，然后再遍历那个数组
  - 而且如果同层不需要保序的话，用个hash表存都可以
  - 模板是没变的，只要这个数组长度大于0就没完

# 提交作业

```
#学号:G20200343110147
#姓名:benben
#班级:14期1班3组
#语言:go
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/tree/golang/Week_07
```

# 实战

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [102 binary-tree-level-order-traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)| 🟡 中等 | 树、广度优先搜索 | - |9.20✅|9.21|9.22|9.30||
| [208 implement-trie-prefix-tree](https://leetcode.com/problems/implement-trie-prefix-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)| 🟡 中等 | 设计、字典树 | - |9.20✅|9.21|9.22|9.30||
| [212 word-search-ii](https://leetcode.com/problems/word-search-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [单词搜索 II](https://leetcode-cn.com/problems/word-search-ii/)| 🔴 困难   | 字典树、回溯算法 | - |9.20✅|9.21|9.22|9.30||
| [547 friend-circles](https://leetcode.com/problems/friend-circles/discuss/?currentPage=1&orderBy=most_votes&query=) | [朋友圈](https://leetcode-cn.com/problems/friend-circles/)| 🟡 中等  | 深度优先搜索、并查集 | - |9.20✅|9.21|9.22|9.30||
| [200 number-of-islands](https://leetcode.com/problems/number-of-islands/discuss/?currentPage=1&orderBy=most_votes&query=) | [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)| 🟡 中等  | 深、广度优先搜索、并查集 | - |9.20✅|9.21|9.22|9.30||
| [130 surrounded-regions](https://leetcode.com/problems/surrounded-regions/discuss/?currentPage=1&orderBy=most_votes&query=) | [被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions/)| 🟡 中等 | 深、广度优先搜索、并查集 | - |9.20✅|9.21|9.22|9.30||
| [70 climbing-stairs](https://leetcode.com/problems/climbing-stairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)| 简单 | 动态规划 | - |9.26✅|9.27|9.29|10.4||
| [22 generate-parentheses](https://leetcode.com/problems/generate-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [括号生成](https://leetcode-cn.com/problems/generate-parentheses/)| 🟡 中等  | 字符串、回溯算法 | - |9.26✅|9.27|9.29|10.4||
| [51 n-queens](https://leetcode.com/problems/n-queens/discuss/?currentPage=1&orderBy=most_votes&query=) | [n皇后](https://leetcode-cn.com/problems/n-queens/)|  🔴 困难 | 回溯算法 | - |9.26✅|9.27|9.29|10.4||
| [36 valid-sudoku](https://leetcode.com/problems/valid-sudoku/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的数独](https://leetcode-cn.com/problems/valid-sudoku/)| 🟡 中等 | 哈希表？ | - |9.22✅|9.25|9.26|10.4||
| [37 sudoku-solver](https://leetcode.com/problems/sudoku-solver/discuss/?currentPage=1&orderBy=most_votes&query=) | [解数独](https://leetcode-cn.com/problems/sudoku-solver/)| 🔴 困难 | 哈希表、回溯算法 | - |9.22✅|9.25|9.26|10.4||
| [127 word-ladder](https://leetcode.com/problems/word-ladder/discuss/?currentPage=1&orderBy=most_votes&query=) | [单词接龙](https://leetcode-cn.com/problems/word-ladder/)| 🟡 中等 | 广度优先搜索 | - |9.20✅|9.21|9.22|9.30||
| [433 minimum-genetic-mutation](https://leetcode.com/problems/minimum-genetic-mutation/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)| 🟡 中等 | 广度优先搜索 | - |9.20✅|9.21|9.22|9.30||
| [1091 shortest-path-in-binary-matrix](https://leetcode.com/problems/shortest-path-in-binary-matrix/discuss/?currentPage=1&orderBy=most_votes&query=) | [二进制矩阵中的最短路径](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)| 🟡 中等 | 广度优先搜索 | - |9.24✅|9.26|9.27|10.3||
| [773 sliding-puzzle](https://leetcode.com/problems/sliding-puzzle/discuss/?currentPage=1&orderBy=most_votes&query=) | [滑动谜题](https://leetcode-cn.com/problems/sliding-puzzle/)| 🔴 困难 | 广度优先搜索 | - |9.26✅|9.27|9.29|10.4||


# HomeWork

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |