学习笔记

## 总结
- 组合问题
  - 组合问题一般不是回溯法解决的问题，但是可以看做回溯方式的一个特例，就是不回溯的回溯法。。。  
    - 套用回溯模板就是，不要调用撤销，同时每次递归在从选择列表中只拿一个备选项向下递归即可，不能遍历所有剩下的
    - 上面理解的有问题，为什么要撤销，如果递归前的操作对下次递归有影响才需要撤销，如果递归前的操作是复制了数据，就不用再撤销了，但是这样会额外浪费很多空间，所以撤销的目的更是为了降低空间开销？
    - 所以组合问题是不是要撤销也不是关键，关键的是 在每层选择的时候，一次性把该层次所有的可能都递归出去时，各个递归之间究竟有没有影响，如果有影响，在每次递归之后都要撤销操作，如果没有影响，就不用管了。
    - 所以所谓撤销实际的含义是：  撤销同层兄弟节点之间的互相影响。
  - 上面的结论有待进一步证实

## 随笔

### 关于模板

- 回溯模板
  - 关键还是在于找到什么是路径，以及在这个路径制约下的选择列表是什么，以及如何过滤出这个选择列表
  - 注意在path或path值处理时，golang需要copy一下数组，否则append到最终结果中是个数组指针，当下一次修改append的内容时会影响最终结果里上一次的记录
  - 在处理组合问题时，因为结果没有顺序，所以顺序反而成为约束条件，即可以按照顺序遍历路径， 于是在模板中不必再回退，同时每次递归的路径只要递一个步长即可，退化的模板如下
```
    if 满足结束条件:
        result.add(路径)
        return
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
        
变成了
    if 满足结束条件:
        result.add(路径)
        return
    选择下一个
    backtrack(路径, 选择列表)
```

### 提交作业

```
#学号:G20200343110147
#姓名:benben
#班级:14期1班3组
#语言:go
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/tree/golang/Week_03
```
### 实战
| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #? |
| --- | --- | --- | --- | --- |  --- | --- | --- | --- | --- |
| [70 climbing-stairs](https://leetcode.com/problems/climbing-stairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)| 🟢 简单 | 泛型递归、树的递归 | - | ✅ | ||||
| [22 generate-parentheses](https://leetcode.com/problems/generate-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [括号生成](https://leetcode-cn.com/problems/generate-parentheses/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||
| [50 powx-n](https://leetcode.com/problems/powx-n/discuss/?currentPage=1&orderBy=most_votes&query=) | [Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)| 🟡 中等 | 分治、回溯 | - | |||||
| [78 subsets](https://leetcode.com/problems/subsets/discuss/?currentPage=1&orderBy=most_votes&query=) | [子集](https://leetcode-cn.com/problems/subsets/)| 🟡 中等 | 分治、回溯 | - | |||||
| [17 letter-combinations-of-a-phone-number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/?currentPage=1&orderBy=most_votes&query=) | [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)| 🟡 中等 | 分治、回溯 | - | |||||
| [51 n-queens](https://leetcode.com/problems/n-queens/discuss/?currentPage=1&orderBy=most_votes&query=) | [N皇后](https://leetcode-cn.com/problems/n-queens/)| 🔴 困难 | 分治、回溯 | - | |||||
| [169 majority-element](https://leetcode-cn.com/problems/majority-element) | [多数元素](https://leetcode-cn.com/problems/majority-element)|  🟢 简单 | 分治、回溯 | - | |||||

### 课后作业
| 题号 | 名称 | 难度 | 分类 | 备注 |#1 | #2 | #3 | #4 | #? |
| --- | --- | --- | --- | --- |--- | --- | --- | --- | --- |
| [236 lowest-common-ancestor-of-a-binary-tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||
| [105 construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||
| [77 combinations](https://leetcode.com/problems/combinations/discuss/?currentPage=1&orderBy=most_votes&query=) | [组合](https://leetcode-cn.com/problems/combinations/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||
| [46 permutations](https://leetcode.com/problems/permutations/discuss/?currentPage=1&orderBy=most_votes&query=) | [全排列](https://leetcode-cn.com/problems/permutations/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||
| [47 permutations-ii](https://leetcode.com/problems/permutations-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [全排列 II](https://leetcode-cn.com/problems/permutations-ii/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||

### 下周预习
| 题号 | 名称 | 难度 | 分类 | 备注 |#1 | #2 | #3 | #4 | #? |
| --- | --- | --- | --- | --- |--- | --- | --- | --- | --- |
| [102](https://leetcode.com/problems/binary-tree-level-order-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)| 🟡 中等 | 深度优先、广度优先 | - ||||||
| [322](https://leetcode.com/problems/coin-change/discuss/?currentPage=1&orderBy=most_votes&query=) | [零钱兑换](https://leetcode-cn.com/problems/coin-change/)| 🟡 中等 | 动态规划 | - ||||||
| [69](https://leetcode.com/problems/sqrtx/discuss/?currentPage=1&orderBy=most_votes&query=) | [x 的平方根](https://leetcode-cn.com/problems/sqrtx/)| 🟢 简单 | 二分查找 | - ||||||
| [367](https://leetcode.com/problems/valid-perfect-square/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/)| 🟢 简单 | 二分查找 | - ||||||
| [122](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)| 🟢 简单 | 贪心算法 | - ||||||
| [455](https://leetcode.com/problems/assign-cookies/discuss/?currentPage=1&orderBy=most_votes&query=) | [分发饼干](https://leetcode-cn.com/problems/assign-cookies/)| 🟢 简单 | 贪心算法 | - ||||||
| [55](https://leetcode.com/problems/jump-game/discuss/?currentPage=1&orderBy=most_votes&query=) | [跳跃游戏](https://leetcode-cn.com/problems/jump-game/)| 🟡 中等 | 贪心算法 | - ||||||

