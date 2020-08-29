学习笔记

## 总结
- 发现了几个特别有用的基本功练习题
  - twoSum
  - moveZeros
  - golang sort.Interface接口实现
  - golang heap.Interface接口实现
  - 双指针revert数组
- 随着题目的增多，开始觉得有些东西混淆和模糊了，尤其是看着熟悉的题目基本不再审题，这可能是个问题。
- 我的脑图
  - 学的东西还不多，我的脑里还没有图[:发愁]
  
|模板|笔记|理解|#|
|---|---|---|---|
|单调栈单调队列|https://shimo.im/docs/hyWwqQ39xVcwk3xG/| 理解的还凑合，记得不牢||
|二叉树展开成右单链表|https://shimo.im/docs/6qJrdVCxvJyKXy6j/ |理解的不怎么样||
|回溯|https://shimo.im/docs/3c3VYtCW9kyCcGYc|刚看完理解了，等5遍之后再看看还记不记得||
- 关于复杂度的推算做的还不好，经常忘了做这件事
  - 树的遍历一般的时间复杂度都是O(N)，因为不论递归还是迭代，每个节点只会访问一次（实际上染色法迭代每个节点丢进去取出来又丢进去又取出来，大约每个节点搞了4次）
  - 空间复杂度看要做什么，比如处理完更新某个全局数字，那就是O(1)，如果处理完得到某种数组类型的东西，还要传给下一次递归或者迭代，还要看这个数组具体是怎么操作的，数组本身还有O(N)的操作复杂度，这时候丢进去一个指针用来指示下次处理时候数组从哪儿开始，能显著的降低复杂度。 例如前序遍历中序遍历恢复二叉树的问题，直接丢数组给下一次就不如丢索引，但是索引处理起来还是太复杂，很容易错
- 本周背诵模板大于理解含义
  - 二叉树的两套模板
    - 递归模板，非常简单，也逐渐开始理解稍复杂一点的本层处理逻辑了
    - 迭代模板，染色法模板还是比较万用的， 基本上套哪儿都好使，注意队列的使用，pop(pool)是从前面取
  - 顺便背诵了一下回溯的模板
    - 竟然一次做完了N皇后，还是很神奇的东西，模板还是要继续背
- hash表 golang的底层实现也挺复杂的，还没看明白
  - 一些理论的东西，装载因子，动态扩容，hash冲突的解决办法能简易的描述
- heap看golang的实现好几次了
  - 接口文档仔细阅读，全是细节， push的时候从哪儿推，pop的时候从哪儿取
- 什么时候能把二叉树遍历运用的和数组遍历一个感觉，可能才是真的弄明白了。

## 其他

### 还不能理解

- 丑数， 还是不太能理解，三指针怎么想出来的
- 为什么用前序遍历？
- 为什么用中序遍历？
  - 看到的一般都是二叉搜索树，因为中序遍历出来的是按顺序排列的
- 为什么用后序遍历？
  - [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)  为什么？？？
  - 目前这个只能按照一个范式背诵下来

### 随笔
- golang的heap理解
  - 虽然heap已被实现，但使用时候还是要套用模板来实现接口
  - 首先自己要定义一个数组
  - 严格按照模板来实行，push就是append数组， pop就是弹出最后一个元素，top就是数组[0]，这几个顺序不能变
  - 小顶堆的swap按照 i < j 返回
  - 另外golang的数组在使用的时候作为指针，取其索引值的时候要这样用 (*a)[i]
  
- 关于树的递归遍历
  - 能理解遍历模板和遍历过程
  - 但是， 只要在遍历过程中有数据比较的过程，就完全无法理解了，什么是上一个数据？？？从哪儿归来的？？
    - 显式的指定返回值可能会好一点。
- 关于树的迭代遍历，不管是几叉树都是一样的
  - 套用模板，使用栈和涂色node
  - 切记，栈要反着用，比如前序遍历，入栈的顺序是  right，left，root，  中序遍历  right，root，left
  - n叉树的children一样要倒着遍历入栈
  
- 使用queue
  - 应该用跟java一样的api， 入队列 add(node)/offer(node), 出队列 remove()/poll(), 看队首 element()/peek()
  
### 提交作业
```
#学号:G20200343110147
#姓名:benben
#班级:14期1班3组
#语言:go
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/tree/golang/Week_02
```
## 练习记录

### 实战
| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #5 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [226 invert-binary-tree](https://leetcode.com/problems/invert-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)| 🟢 简单 | 泛型递归、树的递归 | - |8.16✅  | 8.17✅ | 8.18✅| 8.25✅| |
| [98 validate-binary-search-tree](https://leetcode.com/problems/validate-binary-search-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)| 🟡 中等 | 泛型递归、树的递归 | - |8.16✅  | 8.17✅ |8.18✅ |8.25✅ | |
| [104 maximum-depth-of-binary-tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)| 🟢 简单 | 泛型递归、树的递归 | - |8.17✅  | 8.18✅  |8.19✅ | 8.26✅||
| [111 minimum-depth-of-binary-tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)| 🟢 简单 | 泛型递归、树的递归 | - | 8.17✅ |8.18✅  |8.19✅ |8.26✅ | |
| [297 serialize-and-deserialize-binary-tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)| 🔴️ 困难 | 泛型递归、树的递归 | - |8.17✅ |8.18✅ |8.20✅ | 8.27| |

### 课后作业
| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #5 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [242 valid-anagram](https://leetcode.com/problems/valid-anagram/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)| 🟢 简单 | 哈希表、映射、集合 | ⭐️ 也是实战题目 | 8.16✅  |8.17✅ |8.18✅ | 8.25✅| |
| [49 group-anagrams](https://leetcode.com/problems/group-anagrams/discuss/?currentPage=1&orderBy=most_votes&query=) | [字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)| 🟡 中等 | 哈希表、映射、集合 | ⭐️ 也是实战题目 | 8.16✅ |8.17✅ | 8.18✅|8.25✅ | |
| [1 two-sum](https://leetcode.com/problems/two-sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [两数之和](https://leetcode-cn.com/problems/two-sum/)| 🟢 简单 | 哈希表、映射、集合 | ⭐️ 也是实战题目 |8.16✅  | 8.17| 8.18|8.25 | |
| [94 binary-tree-inorder-traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)| 🟡 中等 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 | 8.16✅|8.17✅ | 8.18✅| 8.25✅| |
| [144 binary-tree-preorder-traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)| 🟡 中等 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 | 8.16✅|8.17✅ |8.18✅ |8.25✅ | |
| [590 n-ary-tree-postorder-traversal](https://leetcode.com/problems/n-ary-tree-postorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [N叉树的后序遍历](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)| 🟢 简单 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 |8.16✅|8.16✅ |8.17✅|8.24✅| |
| [589 n-ary-tree-preorder-traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [N叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)| 🟢 简单 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 | 8.16✅|8.16✅| 8.17✅|8.24✅| |
| [429 n-ary-tree-level-order-traversal](https://leetcode.com/problems/n-ary-tree-level-order-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [N叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)| 🟡 中等 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 | 8.16✅| 8.17✅| 8.18✅|8.25✅ | |
| [236 lowest-common-ancestor-of-a-binary-tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)| 🟡 中等 | 泛型递归、树的递归 | - |  8.17✅|8.19✅ | 8.20✅|8.27✅ | |
| [105 construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)| 🟡 中等 | 泛型递归、树的递归 | - |8.18✅ |8.19✅ |8.20✅ |8.27✅| |
| [264 ugly-number-ii](https://leetcode.com/problems/ugly-number-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [丑数](https://leetcode-cn.com/problems/ugly-number-ii/)| 🟡 中等 | 泛型递归、树的递归 | - | 8.18✅| 8.19✅|8.20✅ |8.27✅ | |
| [347 top-k-frequent-elements](https://leetcode.com/problems/top-k-frequent-elements/discuss/?currentPage=1&orderBy=most_votes&query=) | [前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)| 🟡 中等 | 泛型递归、树的递归 | - |8.18✅ | 8.19✅|8.20✅ |8.27✅| |

  
