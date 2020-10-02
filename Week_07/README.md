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

- AVL & 红黑树
  - 以理解概念和特点为主
    - 严格平衡二叉搜索树
    - 近似平衡二叉树
  - AVL： 平衡二叉搜索树
    - 每个节点保存平衡因子{-1, 0 , 1}, 用于标识右子树比左子树高度差，并且整个树要一直保持该特点，当平衡被打破时要立即恢复平衡
    - 恢复平衡通过4种旋转操作来实现
      - 插入导致的失衡，自下向上，每遇到一个失衡的点(-2,+2)，其子节点必有一个同符号1的，用其判断是否同侧失衡
      - 同侧失衡LL或RR，用对其子节点一次左旋或右旋解决，相当于将子节点升一级，失衡节点下降一级
      - 异侧失衡LR或RL，对其子节点一次左旋或者右旋将孙节点旋上来变成子节点，变成LL或者RR形态，再执行上一个步骤
      - 删除导致的失衡，看是如何实现的删除
        - 删除的节点直接旋转下去变成叶子节点，细节不太清楚
        - 交换删除的节点和最近的叶子节点，再由失去叶子节点的节点根据是否失衡自下而上调平，行为和插入是一样的
    - 不足： 
      - 平衡因子有三个状态，不能用bit保存，需要的存储空间会大一些，数据规模大时需要考虑空间成本
      - 任意插入删除都需要从叶子节点到根节点进行调平，写操作频繁时性能会下降，树的维护成本高
  - 红黑树
    - 特点：近似平衡的二叉搜索树，左右子树高度差小于两倍。其特点由以下条件来保证：
      - 节点非红即黑
      - 叶子节点是空节点，根节点和叶子节点都是黑色
      - 不能有两个相连接的红色节点
      - 从任意节点到其每个叶子节点的所有路径包含相等的黑色节点
  - AVL VS 红黑树
    - 查找性能AVL更好，因为其更严格平衡
    - 插入和删除红黑树性能更好，其需要旋转调平的情况更少
    - AVL要存储平衡因子和树高，红黑树只要一个bit存状态
    - 读多写少用AVL，写多用红黑树
          
- B树
  - 未系统学习，把了解到的内容先总结
  - B树一般是用来应对类似多次磁盘IO查找的场景，尽量减少查找次数，因此会尽量将树的高度降低，这样在每个节点上就会保存更多的数据或地址块
  - B树的节点上用于划分子节点的元素本身也保存了该元素对应的数据，这样就会导致每个节点能存放的总信息量下降
  - 遍历不是很稳定，数据不一定在哪一层的节点上
- B+树
  - 非叶子节点的划分元素本身不保存数据，自身也会一路向下同步到叶子节点上，成为叶子节点的最大或者最小元素
  - 叶子节点保存所有元素以及元素对应的数据记录指针
  - 叶子节点的全部元素还有一个按顺序链接的指针
  - 查找稳定，所有数据都在叶子节点，由于叶子节点本身还是一个链表，遍历数据也更容易
- 2-3树
  - 看起来就是一个三阶B树 ？  
    

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
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/blob/golang/Week_07/README-commit.md
```

# 实战

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [102 binary-tree-level-order-traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)| 🟡 中等 | 树、广度优先搜索 | - |9.20✅|9.26✅|10.2✅|10.9||
| [208 implement-trie-prefix-tree](https://leetcode.com/problems/implement-trie-prefix-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [实现 Trie (前缀树)](https://leetcode-cn.com/problems/implement-trie-prefix-tree/)| 🟡 中等 | 设计、字典树 | - |9.20✅|9.26✅|10.2✅|10.9||
| [212 word-search-ii](https://leetcode.com/problems/word-search-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [单词搜索 II](https://leetcode-cn.com/problems/word-search-ii/)| 🔴 困难   | 字典树、回溯算法 | - |9.20✅|9.26✅|10.2✅|10.9||
| [547 friend-circles](https://leetcode.com/problems/friend-circles/discuss/?currentPage=1&orderBy=most_votes&query=) | [朋友圈](https://leetcode-cn.com/problems/friend-circles/)| 🟡 中等  | 深度优先搜索、并查集 | - |9.20✅|10.2✅|10.5|10.12||
| [200 number-of-islands](https://leetcode.com/problems/number-of-islands/discuss/?currentPage=1&orderBy=most_votes&query=) | [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)| 🟡 中等  | 深、广度优先搜索、并查集 | - |9.20✅|10.2✅|10.5|10.12||
| [130 surrounded-regions](https://leetcode.com/problems/surrounded-regions/discuss/?currentPage=1&orderBy=most_votes&query=) | [被围绕的区域](https://leetcode-cn.com/problems/surrounded-regions/)| 🟡 中等 | 深、广度优先搜索、并查集 | - |9.20✅|10.2✅|10.5|10.12||
| [70 climbing-stairs](https://leetcode.com/problems/climbing-stairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)| 简单 | 动态规划 | - |9.26✅|10.2✅|10.5|10.12||
| [22 generate-parentheses](https://leetcode.com/problems/generate-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [括号生成](https://leetcode-cn.com/problems/generate-parentheses/)| 🟡 中等  | 字符串、回溯算法 | - |9.26✅|10.2✅|10.5|10.12||
| [51 n-queens](https://leetcode.com/problems/n-queens/discuss/?currentPage=1&orderBy=most_votes&query=) | [n皇后](https://leetcode-cn.com/problems/n-queens/)|  🔴 困难 | 回溯算法 | - |9.26✅|10.2✅|10.5|10.12||
| [36 valid-sudoku](https://leetcode.com/problems/valid-sudoku/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的数独](https://leetcode-cn.com/problems/valid-sudoku/)| 🟡 中等 | 哈希表？ | - |9.22✅|10.2✅|10.5|10.12||
| [37 sudoku-solver](https://leetcode.com/problems/sudoku-solver/discuss/?currentPage=1&orderBy=most_votes&query=) | [解数独](https://leetcode-cn.com/problems/sudoku-solver/)| 🔴 困难 | 哈希表、回溯算法 | - |9.22✅|10.2✅|10.5|10.12||
| [127 word-ladder](https://leetcode.com/problems/word-ladder/discuss/?currentPage=1&orderBy=most_votes&query=) | [单词接龙](https://leetcode-cn.com/problems/word-ladder/)| 🟡 中等 | 广度优先搜索 | - |9.20✅|10.2✅|10.5|10.12||
| [433 minimum-genetic-mutation](https://leetcode.com/problems/minimum-genetic-mutation/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)| 🟡 中等 | 广度优先搜索 | - |9.20✅|10.2✅|10.5|10.12||
| [1091 shortest-path-in-binary-matrix](https://leetcode.com/problems/shortest-path-in-binary-matrix/discuss/?currentPage=1&orderBy=most_votes&query=) | [二进制矩阵中的最短路径](https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/)| 🟡 中等 | 广度优先搜索 | - |9.24✅|10.2✅|10.5|10.12||
| [773 sliding-puzzle](https://leetcode.com/problems/sliding-puzzle/discuss/?currentPage=1&orderBy=most_votes&query=) | [滑动谜题](https://leetcode-cn.com/problems/sliding-puzzle/)| 🔴 困难 | 广度优先搜索 | - |9.26✅|10.2✅|10.5|10.12||


# HomeWork

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |