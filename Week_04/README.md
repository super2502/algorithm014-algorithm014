学习笔记

# 总结 

- 套路
  - 深度优先搜索，DFS， 递归或迭代， 用栈。
  - 广度优先搜索，BFS， 迭代， 用队列
- DFS实际上是树的先序遍历 ？？
  - 应该是，因为在做树的染色法迭代遍历时，曾经发现过，只有先序遍历不需要把当前节点压栈两次，其简化版就和现在的dfs模板一样
- 模板
  - 非常惊诧的发现了这样一个模板，和标准二分模板一模一样，而且能获取比要搜索的值小的最大的那个。。。, 完全搞不懂为什么
  ```
    left := lo
    right := hi
    for left <= right {
        mid := left + (left-right) / 2 
        if target == nums[mid] {
            return true
        } else if target < nums[mid] {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    if left == lo {
        return false
    }
    left--
  ```
|模板|笔记|理解|历史|
|---|---|---|---|
|DFS递归|https://shimo.im/docs/9CYPpdcPGwXT93QV/|这玩意和回溯模板怎么区分，都是dfs为什么这儿可以不回溯|week04|
|DFS迭代|https://shimo.im/docs/JqXDvhW9jt6Y9hQV/ |这就是树的先序遍历模板 |week04|
|BFS|https://shimo.im/docs/VkVGpccqqxwvqtgR/|用队列|week04|
|二分查找|https://shimo.im/docs/vkPwvRcktgHdWWdW/|细节需要研究，都是边界问题|week04|
|回溯|https://shimo.im/docs/3c3VYtCW9kyCcGYc|模板很好用|week03|
|单调栈单调队列|https://shimo.im/docs/hyWwqQ39xVcwk3xG/| 理解的还凑合，记得不牢|week02|

- 想不出来的问题整理，只能背了
  - 二叉树公共祖先节点
  - 二叉树展开成右单链表
  - 丑数
  


# 思考

- DFS和回溯
  - 回溯也是dfs，这两个玩意模板还不一样，具体区别在哪？还是没有区别
- BFS和回溯
  - 广度优先不需要回溯，一个节点的哥哥节点被探索过，之后，visited里永远都有了，后面的弟弟节点也不需要再走哥哥节点，因此跟回溯不同，不会撤销visited
  

# 随笔

- 第四周的问题明显的需要背诵的更多了，根本想不出来
- 二分查找，x平方根， 需要取整，这里已经懵了， mid小的时候有可能确是解，所以是不是要cache一下这个mid

# 提交作业
```
#学号:G20200343110147
#姓名:benben
#班级:14期1班3组
#语言:go
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/tree/golang/Week_04

```

# 实战

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [102 binary-tree-level-order-traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)| 🟡 中等 | 深度优先、广度优先 | - | 8.30✅|8.31|9.1|9.8||
| [433 minimum-genetic-mutation](https://leetcode.com/problems/minimum-genetic-mutation/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)| 🟡 中等 | 深度优先、广度优先 | - | 8.30✅|8.31|9.1|9.8||
| [22 generate-parentheses](https://leetcode.com/problems/generate-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [括号生成](https://leetcode-cn.com/problems/generate-parentheses/)| 🟡 中等 | 深度优先、广度优先 | - | 8.30✅|8.31|9.1|9.8||
| [515 find-largest-value-in-each-tree-row](https://leetcode.com/problems/find-largest-value-in-each-tree-row/discuss/?currentPage=1&orderBy=most_votes&query=) | [在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)| 🟡 中等 | 深度优先、广度优先 | - | 8.30✅|8.31|9.1|9.8||
| [322 coin-change](https://leetcode.com/problems/coin-change/discuss/?currentPage=1&orderBy=most_votes&query=) | [零钱兑换](https://leetcode-cn.com/problems/coin-change/)| 🟡 中等 | 动态规划 | - |8.30✅|8.31|9.1|9.8||
| [69 sqrtx](https://leetcode.com/problems/sqrtx/discuss/?currentPage=1&orderBy=most_votes&query=) | [x 的平方根](https://leetcode-cn.com/problems/sqrtx/)| 🟢 简单 | 二分查找 | - | 8.30✅|8.31|9.1|9.8||
| [367 valid-perfect-square](https://leetcode.com/problems/valid-perfect-square/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/)| 🟢 简单 | 二分查找 | - |8.30✅|8.31|9.1|9.8||

### 课后作业

| 题号 | 名称 | 难度 | 分类 | 备注 |
| --- | --- | --- | --- | --- |
| [860 lemonade-change](https://leetcode.com/problems/lemonade-change/discuss/?currentPage=1&orderBy=most_votes&query=) | [柠檬水找零](https://leetcode-cn.com/problems/lemonade-change/)| 🟢 简单 | 贪心算法 | - | |||||
| [122 best-time-to-buy-and-sell-stock-ii](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)| 🟢 简单 | 贪心算法 | - | |||||
| [455 assign-cookies](https://leetcode.com/problems/assign-cookies/discuss/?currentPage=1&orderBy=most_votes&query=) | [分发饼干](https://leetcode-cn.com/problems/assign-cookies/)| 🟢 简单 | 贪心算法 | - | |||||
| [874 walking-robot-simulation](https://leetcode.com/problems/walking-robot-simulation/discuss/?currentPage=1&orderBy=most_votes&query=) | [模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/)| 🟢 简单 | 贪心算法 | - | |||||
| [127 word-ladder](https://leetcode.com/problems/word-ladder/discuss/?currentPage=1&orderBy=most_votes&query=) | [单词接龙](https://leetcode-cn.com/problems/word-ladder/)| 🟡 中等 | 深度优先、广度优先 | - | |||||
| [200 number-of-islands](https://leetcode.com/problems/number-of-islands/discuss/?currentPage=1&orderBy=most_votes&query=) | [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)| 🟡 中等 | 深度优先、广度优先 | - | |||||
| [529 minesweeper](https://leetcode.com/problems/minesweeper/discuss/?currentPage=1&orderBy=most_votes&query=) | [扫雷游戏](https://leetcode-cn.com/problems/minesweeper/)| 🟡 中等 | 深度优先、广度优先 | - | |||||
| [55 jump-game](https://leetcode.com/problems/jump-game/discuss/?currentPage=1&orderBy=most_votes&query=) | [跳跃游戏](https://leetcode-cn.com/problems/jump-game/)| 🟡 中等 | 贪心算法 | - | |||||
| [33 search-in-rotated-sorted-array](https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)| 🟡 中等 | 二分查找 | - | |||||
| [74 search-a-2d-matrix](https://leetcode.com/problems/search-a-2d-matrix/discuss/?currentPage=1&orderBy=most_votes&query=) | [搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)| 🟡 中等 | 二分查找 | - | |||||
| [153 find-minimum-in-rotated-sorted-array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)| 🟡 中等 | 二分查找 | - | |||||
| [126 word-ladder-ii](https://leetcode.com/problems/word-ladder-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [单词接龙 II](https://leetcode-cn.com/problems/word-ladder-ii/)| 🔴 困难 | 深度优先、广度优先 | - | |||||
| [45 jump-game-ii](https://leetcode.com/problems/jump-game-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/)| 🔴 困难 | 贪心算法 | - | |||||
