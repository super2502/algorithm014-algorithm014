学习笔记

# 第一周

## 记录

- 283 move-zeroes: [双指针+swap令人恼火]
  - 双指针的边界实在是太难理解了。 为什么i和j都从0开始刚好能cover住所有的场景，这是怎么想出来的。
  - 我从i=0,j=1开始考虑，就需要处理OO,OX,XO,XX这四种情况，代码非常复杂
  - i=0,j=0然后还能swap(i,j),从常理很难去这么思考，这算是技巧还是什么？
  - 刻意使用swap在不熟练的时候非常难以理解
  - 比较朴素的遍历思想更清晰，实际上用于记录非零的指针就是第二个指针


## 五毒神掌 [偷笑]

### Array
| 题号 | 名称 | 难度 | 分类 | 第一掌 | 第二掌 | 第三掌 | 第四掌 | 第?掌 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| [283 move-zeroes](https://leetcode.com/problems/move-zeroes/discuss/?currentPage=1&orderBy=most_votes&query=) | [移动零](https://leetcode-cn.com/problems/move-zeroes/)| 🟢 简单 | 数组 | 08.10✅  | 08.10✅  | 08.11✅  | 08.18 | - |
| [70 climbing-stairs](https://leetcode.com/problems/climbing-stairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)| 🟢 简单 | 数组 | 08.10✅  | 08.10✅  | 08.12 | 08.19 | - |
| [15 3sum](https://leetcode.com/problems/3sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [三数之和](https://leetcode-cn.com/problems/3sum/)| 🟡 中等 | 数组 | 08.10✅  | 08.12 | 08.13 | 08.20 | - |
| [11 container-with-most-water](https://leetcode.com/problems/container-with-most-water/discuss/?currentPage=1&orderBy=most_votes&query=) | [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)| 🟡 中等 | 数组 | 08.10✅  | 08.10✅  | 08.12 | 08.19 | - |


### Linked list
| 题号 | 名称 | 难度 | 分类 | 第一掌 | 第二掌 | 第三掌 | 第四掌 | 第?掌 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| [146 lru-cache](https://leetcode.com/problems/lru-cache/discuss/?currentPage=1&orderBy=most_votes&query=) | [LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/)| 🟡 中等 | 链表 | 08.10✅ | 08.10✅ | 08.12 | 08.19 | - |
| [206 reverse-linked-list](https://leetcode.com/problems/reverse-linked-list/discuss/?currentPage=1&orderBy=most_votes&query=) | [反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)| 🟢 简单 | 链表 | 08.10✅  | 08.12 | 08.13 | 08.20 | - |
| [24 swap-nodes-in-pairs](https://leetcode.com/problems/swap-nodes-in-pairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)| 🟡 中等 | 链表 | 08.10✅ | 08.12 | 08.13 | 08.20 | - |
| [141 linked-list-cycle](https://leetcode.com/problems/linked-list-cycle/discuss/?currentPage=1&orderBy=most_votes&query=) | [环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)| 🟢 简单 | 链表 |08.10✅ | 08.12 | 08.13 | 08.20 | - |
| [142 linked-list-cycle-ii](https://leetcode.com/problems/linked-list-cycle-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)| 🟡 中等 | 链表 | 08.10✅  | 08.12 | 08.13 | 08.20 | - |
| [25 reverse-nodes-in-k-group](https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/?currentPage=1&orderBy=most_votes&query=) | [K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)| 🔴️ 困难 | 链表 | 08.13 | 08.13 | 08.14 | 08.21 | - |

### Stack & Queue
| 题号 | 名称 | 难度 | 分类 | 第一掌 | 第二掌 | 第三掌 | 第四掌 | 第?掌 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| [20 valid-parentheses](https://leetcode.com/problems/valid-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)| 🟢 简单 | 栈、队列 | 08.10✅ | 08.12✅  | 08.13 | 08.20 | - |
| [155 min-stack](https://leetcode.com/problems/min-stack/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小栈](https://leetcode-cn.com/problems/min-stack/)| 🟢 简单 | 栈、队列 | 08.10✅  | 08.12✅  | 08.13 | 08.20 | - |
| [84 largest-rectangle-in-histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/discuss/?currentPage=1&orderBy=most_votes&query=) | [柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)| 🔴️ 困难 | 栈、队列 | 08.13 | 08.13 | 08.14 | 08.21 | - |
| [239 sliding-window-maximum](https://leetcode.com/problems/sliding-window-maximum/discuss/?currentPage=1&orderBy=most_votes&query=) | [滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)| 🔴️ 困难 | 栈、队列 | 08.10✅  | 08.12 | 08.13 | 08.20 | - |

### Homework
| 题号 | 名称 | 难度 | 分类 | 第一掌 | 第二掌 | 第三掌 | 第四掌 | 第?掌 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| [26 remove-duplicates-from-sorted-array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)| 🟢 简单 | 数组、链表、跳表 | 08.10✅ | 08.10 | 08.11 | 08.18 | - |
| [189 rotate-array](https://leetcode.com/problems/rotate-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [旋转数组](https://leetcode-cn.com/problems/rotate-array/)| 🟢 简单 | 数组、链表、跳表 | 08.10 | 08.10 | 08.11 | 08.18 | - |
| [21 merge-two-sorted-lists](https://leetcode.com/problems/merge-two-sorted-lists/discuss/?currentPage=1&orderBy=most_votes&query=) | [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)| 🟢 简单 | 数组、链表、跳表 | 08.10 | 08.10 | 08.11 | 08.18 | - |
| [88 merge-sorted-array](https://leetcode.com/problems/merge-sorted-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)| 🟢 简单 | 数组、链表、跳表 | 08.10 | 08.10 | 08.11 | 08.18 | - |
| [1 two-sum](https://leetcode.com/problems/two-sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [两数之和](https://leetcode-cn.com/problems/two-sum/)| 🟢 简单 | 数组、链表、跳表 | 08.10 | 08.10 | 08.11 | 08.18 | - |
| [283 move-zeroes](https://leetcode.com/problems/move-zeroes/discuss/?currentPage=1&orderBy=most_votes&query=) | [移动零](https://leetcode-cn.com/problems/move-zeroes/)| 🟢 简单 | 数组、链表、跳表 | 08.10 | 08.10 | 08.11 | 08.18 | - |
| [66 plus-one](https://leetcode.com/problems/plus-one/discuss/?currentPage=1&orderBy=most_votes&query=) | [加一](https://leetcode-cn.com/problems/plus-one/)| 🟢 简单 | 数组、链表、跳表 | 08.10✅ | 08.10 | 08.11 | 08.18 | - |
| [641 design-circular-deque](https://leetcode.com/problems/design-circular-deque/discuss/?currentPage=1&orderBy=most_votes&query=) | [设计循环双端队列](https://leetcode-cn.com/problems/design-circular-deque/)| 🟡 中等 | 栈、队列 | 08.10 | 08.10 | 08.11 | 08.18 | - |
| [42 trapping-rain-water](https://leetcode.com/problems/trapping-rain-water/discuss/?currentPage=1&orderBy=most_votes&query=) | [接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)| 🔴️ 困难 | 栈、队列 | 08.10 | 08.10 | 08.11 | 08.18 | - |



## 讲义练习题

### Array
1. https://leetcode-cn.com/problems/container-with-most-water/ 
1. https://leetcode-cn.com/problems/move-zeroes/
1. https://leetcode-cn.com/problems/climbing-stairs/ 
1. https://leetcode-cn.com/problems/3sum/ (高频老题)

### Linked List
1. https://leetcode-cn.com/problems/reverse-linked-list/
1. https://leetcode-cn.com/problems/swap-nodes-in-pairs
1. https://leetcode-cn.com/problems/linked-list-cycle
1. https://leetcode-cn.com/problems/linked-list-cycle-ii
1. https://leetcode-cn.com/problems/reverse-nodes-in-k-group/

### Homework
1. https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/ 
1. https://leetcode-cn.com/problems/rotate-array/
1. https://leetcode-cn.com/problems/merge-two-sorted-lists/
1. https://leetcode-cn.com/problems/merge-sorted-array/
1. https://leetcode-cn.com/problems/two-sum/
1. https://leetcode-cn.com/problems/move-zeroes/ 
1. https://leetcode-cn.com/problems/plus-one/

### stack and queue
1. https://leetcode-cn.com/problems/valid-parentheses/ - 最近相关性 —> 栈!
1. https://leetcode-cn.com/problems/min-stack/
1. https://leetcode-cn.com/problems/largest-rectangle-in-histogram
1. https://leetcode-cn.com/problems/sliding-window-maximum

### Homework
1. https://leetcode.com/problems/design-circular-deque 
1. https://leetcode.com/problems/trapping-rain-water/


