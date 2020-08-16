学习笔记


## 其他

### 还不能理解

- 为什么用前序遍历？

- 为什么用中序遍历？

- 为什么用后序遍历？
  - [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)  为什么？？？

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
  
## 练习记录

### 实战
| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #5 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [226 invert-binary-tree](https://leetcode.com/problems/invert-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [翻转二叉树](https://leetcode-cn.com/problems/invert-binary-tree/)| 🟢 简单 | 泛型递归、树的递归 | - |8.16✅  | 8.17| 8.18| 8.25| |
| [98 validate-binary-search-tree](https://leetcode.com/problems/validate-binary-search-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)| 🟡 中等 | 泛型递归、树的递归 | - |8.16✅  | 8.17|8.18 |8.25 | |
| [104 maximum-depth-of-binary-tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)| 🟢 简单 | 泛型递归、树的递归 | - |8.17✅  | 8.18 |8.19 | 8.26||
| [111 minimum-depth-of-binary-tree](https://leetcode.com/problems/minimum-depth-of-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)| 🟢 简单 | 泛型递归、树的递归 | - | 8.17✅ |8.18 |8.19 |8.26 | |
| [297 serialize-and-deserialize-binary-tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)| 🔴️ 困难 | 泛型递归、树的递归 | - | | | | | |

### 课后作业
| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #5 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [242 valid-anagram](https://leetcode.com/problems/valid-anagram/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的字母异位词](https://leetcode-cn.com/problems/valid-anagram/)| 🟢 简单 | 哈希表、映射、集合 | ⭐️ 也是实战题目 | 8.16✅  |8.17 |8.18 | 8.25| |
| [49 group-anagrams](https://leetcode.com/problems/group-anagrams/discuss/?currentPage=1&orderBy=most_votes&query=) | [字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/)| 🟡 中等 | 哈希表、映射、集合 | ⭐️ 也是实战题目 | 8.16✅ |8.17 | 8.18|8.25 | |
| [1 two-sum](https://leetcode.com/problems/two-sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [两数之和](https://leetcode-cn.com/problems/two-sum/)| 🟢 简单 | 哈希表、映射、集合 | ⭐️ 也是实战题目 |8.16✅  | 8.17| 8.18|8.25 | |
| [94 binary-tree-inorder-traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/)| 🟡 中等 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 | 8.16✅|8.17 | 8.18| 8.25| |
| [144 binary-tree-preorder-traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的前序遍历](https://leetcode-cn.com/problems/binary-tree-preorder-traversal/)| 🟡 中等 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 | 8.16✅|8.17 |8.18 |8.25 | |
| [590 n-ary-tree-postorder-traversal](https://leetcode.com/problems/n-ary-tree-postorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [N叉树的后序遍历](https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/)| 🟢 简单 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 |8.16✅|8.16✅ |8.17|8.24| |
| [589 n-ary-tree-preorder-traversal](https://leetcode.com/problems/n-ary-tree-preorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [N叉树的前序遍历](https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/)| 🟢 简单 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 | 8.16✅|8.16✅| 8.17|8.24| |
| [429 n-ary-tree-level-order-traversal](https://leetcode.com/problems/n-ary-tree-level-order-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [N叉树的层序遍历](https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/)| 🟡 中等 | 树、二叉树、二叉搜索树 | ⭐️ 也是实战题目 | 8.16✅| 8.17| 8.18|8.25 | |
| [236 lowest-common-ancestor-of-a-binary-tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)| 🟡 中等 | 泛型递归、树的递归 | - |  8.17✅| | | | |
| [105 construct-binary-tree-from-preorder-and-inorder-traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)| 🟡 中等 | 泛型递归、树的递归 | - | | | | | |
| [264 ugly-number-ii](https://leetcode.com/problems/ugly-number-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [丑数](https://leetcode-cn.com/problems/ugly-number-ii/)| 🟡 中等 | 泛型递归、树的递归 | - | | | | | |
| [347 top-k-frequent-elements](https://leetcode.com/problems/top-k-frequent-elements/discuss/?currentPage=1&orderBy=most_votes&query=) | [前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)| 🟡 中等 | 泛型递归、树的递归 | - | | | | | |

  
