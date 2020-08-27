学习笔记

## 总结
- 回溯模板实际上就是继承了递归的模板
  - 所谓回溯，就是在遍历同层兄弟节点时，互相不要干扰，每个节点用完了path变量要把其对path的修改revert掉，以至于当前的深度遍历回来时，还是从前那个少年
    - 因此，回溯问题最主要的就是怎么定义如何定义层，如何定义每层的元素
    - 进而获取如何在每一层处理数据，在每一层进行剪枝，如何向下递进一层
    - 搞清楚每层数据处理逻辑，就搞清楚了path怎么定义，怎么递进和回退
    - 搞清楚了剪枝，最主要是结果去重，同时也简化了复杂度，这不是一个优化，是必要的环节
    - 搞清楚了如何向下递进一层就搞清楚了递归函数怎么写，怎么入参
    - 上面三个都搞清楚了，套用模板即可
  - 不一定非要通过revert的方式进行兄弟间往复
    - 也可以在每个兄弟调用前复制一份递归入参
    - 但是通过直接修改path和revert的方式是高效的
  - 组合问题和回溯的关系
    - 组合问题是按照长度遍历一个集合，目标节点是每个元素i，将结果列表在第j位的所有可能作为递归树的同一层， 树的层数一般是一个小于n的值
    - 树的第j层节点是当结果已经获取到0~j-1层数据的情况下，还可以使用的节点列表，每层的可用节点列表和已经排出来的层是有关的， 这一点和排列是相同的
    - 组合问题与排列问题不同的是，每层可以选择的元素不仅和其自身path(即根节点过来的路径)有关，还和其兄弟有关，其兄弟也是每个节点需要排除的
      - 因此组合问题剪枝，不但要剪父子关系，还要剪兄弟关系
      - 一个模板方法就是，每层递归向下传递当前i所在的位置，而下层递归在遍历集合时，从i开始，就可以简单的过滤掉所有兄弟了
      - 举个例子，遍历1，2，3，4，5， 1向下递归时传2， 下层递归只会看2，3，4，5； 2向下递归传3， 下层递归只会看3，4，5； 3向下递归传4，下层递归只会看5
      - 因此，组合还可以从后边剪枝，就是剩下的元素不够结果所需的剩下的个数时，可以直接剪掉了， 比如 3向下传递4，只会有个5， 那么如果剩下的需求大于1个的话，这只有一个5可选的路径肯定就不能用了
    - path里放置处理结果，也要将当前元素的索引向下传，以通知下层元素不必再选择该索引之前的值
    - 结果就是集合中元素本身的列表综合，这里是元素本身

  - 排列问题和回溯的关系
    - 排列问题是按长度遍历一个集合，目标节点是每个元素i，将结果列表在第j位的所有可能作为递归树的同一层，树的层数一般是集合元素的个数n
    - 树的第j层节点是当结果已经获取到0~j-1层数据的情况下，还可以使用的节点列表，每层的可用节点列表和已经排出来的层是有关的
      - 全排列问题，已经在层路径列表中的元素不能作为下一层的元素使用，剪枝
      - n皇后问题，已经在层路径上出现过的同列或对角线上的点都不能作为下一层元素使用，剪枝
    - path里放置处理结果，同时，path本身的长度刚好和level相同，因此只向下面的递归函数传递path即可
    - 结果是集合中元素本身的列表的总和，这里是元素本身

  - 子集问题和回溯的关系
    - 子集问题是按长度遍历一个集合，目标节点是每个元素i对应的若干种状态，将同一个元素i作为递归树的同一层，树的层数是集合元素的个数n或者是小于n的一个数
    - 树的第i层节点就是子集合中第i个元素可以选择的内容
      - 子集就是这些内容无关，比如纯子集问题，就是当前元素选和不选，电话号码问题，就是每个数字对应的几个字母选项
      - 无关的子集问题就没有机会剪枝, 复杂度就是 x的n次方，x是每层的可选择内容树
      - 有关的子集问题一般要具体分析，比如()()()问题，实际上隐含的条件是任何时候)不能比(多，由此达到剪枝的目的
    - 可以直接使用 level+1 向下层推进， path里放置处理结果，而level+1用于表明是否已经递进到n，可以返回了。 返回的时候把path的结果丢到总结果里，path长度于level+1无关，所以需要传递这两个值到下一次递归
    - 结果是集合中元素对应的状态的列表的总和， 这里是元素对应的状态。

## 思考

- 组合问题
  - 组合问题一般不是回溯法解决的问题，但是可以看做回溯方式的一个特例，就是不回溯的回溯法。。。  
    - 套用回溯模板就是，不要调用撤销，同时每次递归在从选择列表中只拿一个备选项向下递归即可，不能遍历所有剩下的
    - 上面理解的有问题，为什么要撤销，如果递归前的操作对下次递归有影响才需要撤销，如果递归前的操作是复制了数据，就不用再撤销了，但是这样会额外浪费很多空间，所以撤销的目的更是为了降低空间开销？
    - 所以组合问题是不是要撤销也不是关键，关键的是 在每层选择的时候，一次性把该层次所有的可能都递归出去时，各个递归之间究竟有没有影响，如果有影响，在每次递归之后都要撤销操作，如果没有影响，就不用管了。
    - 所以所谓撤销实际的含义是：  撤销同层兄弟节点之间的互相影响。
  - 上面的结论有待进一步证实
  - 已经整理成总结

- 对剪枝的理解
  - 不仅仅是优化
  - 同时也是保证结果正确性的最佳办法，即在遍历过程中筛选正确结果
  
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
| [70 climbing-stairs](https://leetcode.com/problems/climbing-stairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)| 🟢 简单 | 泛型递归、树的递归 | - |8.24✅ |8.25|8.26|9.2||
| [22 generate-parentheses](https://leetcode.com/problems/generate-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [括号生成](https://leetcode-cn.com/problems/generate-parentheses/)| 🟡 中等 | 泛型递归、树的递归 | - |8.24✅  |||||
| [50 powx-n](https://leetcode.com/problems/powx-n/discuss/?currentPage=1&orderBy=most_votes&query=) | [Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)| 🟡 中等 | 分治、回溯 | - |8.24✅  |||||
| [78 subsets](https://leetcode.com/problems/subsets/discuss/?currentPage=1&orderBy=most_votes&query=) | [子集](https://leetcode-cn.com/problems/subsets/)| 🟡 中等 | 分治、回溯 | - | 8.24✅ |||||
| [17 letter-combinations-of-a-phone-number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/?currentPage=1&orderBy=most_votes&query=) | [电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/)| 🟡 中等 | 分治、回溯 | - | 8.24✅ |||||
| [51 n-queens](https://leetcode.com/problems/n-queens/discuss/?currentPage=1&orderBy=most_votes&query=) | [N皇后](https://leetcode-cn.com/problems/n-queens/)| 🔴 困难 | 分治、回溯 | - |8.24✅  |||||
| [169 majority-element](https://leetcode-cn.com/problems/majority-element) | [多数元素](https://leetcode-cn.com/problems/majority-element)|  🟢 简单 | 分治、回溯 | - | 8.24✅ |||||

### 课后作业
| 题号 | 名称 | 难度 | 分类 | 备注 |#1 | #2 | #3 | #4 | #? |
| --- | --- | --- | --- | --- |--- | --- | --- | --- | --- |
| [236 lowest-common-ancestor-of-a-binary-tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||
| [105 construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||
| [77 combinations](https://leetcode.com/problems/combinations/discuss/?currentPage=1&orderBy=most_votes&query=) | [组合](https://leetcode-cn.com/problems/combinations/)| 🟡 中等 | 泛型递归、树的递归 | - | |||||
| [46 permutations](https://leetcode.com/problems/permutations
