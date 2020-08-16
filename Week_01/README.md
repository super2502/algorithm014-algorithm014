学习笔记

# 第一周

## 总结

### 学习内容

- 最基本的数据结构就是数组/链表
  - 其他的高级数据结构的底层数据结构都是这两个东西（？），因此牢记数组和链表的基本操作就是整个数据结构中最基础的步骤
  - 数组和链表在使用上各有优劣，能在恰当的场景使用恰当的数据结构是需要修炼的，因此需要运用以下原则，并加以训练
    - 复杂度原则， 数组在随机查找上优于链表，而链表在删改时优于数组
  - 数组和链表从习题整理上看都有以下相同的处理行为，即算法（题目做的少，还不能列出所有常见行为）
    - 遍历： 数组按照索引+1进行遍历，链表按照node.Next进行遍历,找到1个满足条件的值
    - 遍历进阶： 在同一个数据结构中找到满足条件的 2/3/4/...N组合
    - 交换： 数组可以直接对换任何两个索引的值，而链表一般对换相邻两个node
    - 交换进阶，反转： 这个练习最多的是链表，因为链表只能两两反转，对于k反转需要遍历k； 而数组反转只要直接双指针对调就可以了，没有太多需要练习的地方
    - 两操作，merge：最常见的两有序数组/链表合并，
    - 两操作，对应位置操作：对位相加
    - 多操作： 一般就是对两操作的归并运算
  - 当一个数据结构具有一定的规律时，可以产生更多的算法
    - 有序： 非常强的条件，以至于会有专门的算法来进行排序， 有序的数据结构可以应对很多问题(还需要总结)
    - 双指针： 这东西对有序数组非常好用
  
- 个人理解
  - 递归不是一个算法，而是一种思想或范式，不是特指用在某个具体数据结构上的
    - 递归本身并不会解决问题，解决问题的实际上是在处理递归结果的那段代码里（这很重要）
    - 容易想出的递归： 从头开始的，去掉第一个节点之后一模一样的
    - 难一点的： 从尾开始的，前面都做完再处理最后一个
    - 还不太能理解的： 两边都有的，这个估计要等到把树弄熟练了自然就理解了
    - 最难的：
  - 每当学习一个高级数据结构的时候
    - 它一定对应一套API
    - 想过它底层是什么基础数据结构吗？
    - 用了什么算法来提供的API，有多少是基础算法，有多少是特制的算法
    - 它的复杂度是怎样的
    - 它最常用的场景是什么，或者为什么发明这么个结构出来
  - 掌握数组和链表的基本操作究竟能做什么
    - 看起来这些基本操作，就是未来要快速应用到高级数据结构中的前置能力
    - 后面的高级数据结构和算法中要随时能发现用到了那些基本操作
  
  - 一个初步理解的问题
    - 当目标是一个基础数据结构时，如数组和链表，其每个索引或node对应的值就是要处理的内容，直接对着它执行算法就好了
    - 当目标是一个高级数据结构，在使用其底层的基础数据结构时，要注意基础数据结构里存储的数据是索引，要用算法拿出索引进一步运算
    - 举个例子，例如heap，底层如果用数组存储的话，这个数组的索引只是用于指示其值在heap中的位置，而heap作为一个数据结构来讲，存储的内容是其他数据结构要使用的索引，这个东西相对于底层数组来说实际是底层数组的值，这个要多做一些才能适应，否则会绕进去很容易迷惑
    
- 何谓基础
  - 模板方式学习
    - 模板应该是最重要的部分，这个就是需要背诵的地方
    - 我整理了哪些模板，我整理模板了吗，我在用什么方式整理模板，在用什么方式记住模板？
    - 我掌握基础数据结构的所有操作模板了吗？

  - 典型问题的 （典型/最优）数据结构和算法
    - 这也是个能背诵多少就背诵多少的东西
    - 我整理了哪些典型问题？
    - 我怎么整理的典型问题，怎么记住它的对应最优解？
  
  - 能不能把数组和链表的基本操作整的像99表一样张嘴就来？
    - 儿子为什么能记住161不是一个质数？ 还不是每天都问他一遍

- 本周的高级数据结构
  - 栈
    - 栈主要的操作（API）： push/pop/top/isEmpty，对应的就是数组的最后一个元素以及数组长度的操作， 栈是不是大多数情况下不考虑full？
    - 一般用一个数组就能实现一个栈
    - 复杂度是什么： 
    - 什么场景用：
  - 队列
    - 队列的主要操作（API）: push/pop/top/isEmpty/isFull，是不是大多数情况下不会search一个队列？
    - 一般也用一个数组实现队列，最好用循环数组节省空间；链表的话要用一个双向链表才能方便push，还要记录size
    - 复杂度是什么：
    - 什么场景用：
  
  - 单调栈
    - 模板 https://shimo.im/docs/hyWwqQ39xVcwk3xG/
    - 还处于难以理解状态
    - 经常用于一次遍历一个未排序的数据结构，需求本身却有对数据大小的需求，遍历一次同时得到结果
  - 优先级队列
    - 还没研究
  - linkedHashMap
    - LRU写过几遍了，这东西还有哪些用处？
  - 堆
    - 看了golang的堆实现
    - API： pop，出来的就是极值， push，丢一个数据进去，而且丢完立马被排序好
    - 高级API： remove，指定一个索引删除数据，删完之后排好序，这个得知道数据的index；fix，在一个index为i的数据的value发生变化后修理它的位置；这俩可能都要配合一个hashMap才能搞定； 
    - 底层是个数组，因为是完全二叉树，因此按顺序排列所有节点即可
    - 时间复杂度：不管什么操作都是按照树的高度走一次 O(logn),非常稳定
    - 典型场景： 目前看  TopK问题，配合实现那些需要pop API的数据结构
    
  
## 其他
### 弱项

- 从同一侧出发的双指针，非常难搞清楚起点的边界条件
  - 从同一侧出发，可以先假想是两个数组，为了节省空间映射到了一个上面，先拿俩数组整，整好了之后在想办法落到一个数组里，看看对应的位置会不会覆盖到需要使用的数据
    - move-zero和merge-two-array，一个从前面，一个从后面，两个同侧双指针非常好的例子
- 单调栈，接触太少，不熟悉用法，为什么用以及到底是用递增还是递减的都还没搞清楚
  - 从左往右比， 如果新来的元素在大于栈顶元素时需要处理，这就是个递减栈，  模板就是 while !s.isEmpty() && s.peek() < nums[i] {s.pop(); 主要处理逻}; s.push(i)
  - 反之，如果新来的元素小于栈顶元素时需要处理，就是个递增栈 
  - 模板模板

### 随笔

- 什么问题用栈求解？ - 看起来像剥洋葱的问题。。。
  - 怎么把接雨水想成剥洋葱的？？

## 记录

- 283 move-zeroes: [双指针+swap令人恼火]
  - 双指针的边界实在是太难理解了。 为什么i和j都从0开始刚好能cover住所有的场景，这是怎么想出来的。
  - 我从i=0,j=1开始考虑，就需要处理OO,OX,XO,XX这四种情况，代码非常复杂
  - i=0,j=0然后还能swap(i,j),从常理很难去这么思考，这算是技巧还是什么？
  - 刻意使用swap在不熟练的时候非常难以理解
  - 比较朴素的遍历思想更清晰，实际上用于记录非零的指针就是第二个指针
- 1 two-sum
  - 注意没有排序，不能直接用双指针夹逼，一上来就错了，双指针一定是要处理排过序的 
  - 双指针求和才是需要排过序的
- 641 design-circular-deque
  - 不就是LRU的前置，使劲做LRU就好了
  - 理解的有问题，好像不是一个东西
  - 用数组实现了一版，关键只要记住 head和rear的模型
    - arr长度要定义为k+1
    - 计算IsFull时 (rear+1) % (k+1) == head
    - head和rear增加时 对(k+1)求模
    - head和rear减少时 用 (x+k) 对(k+1)求模 
    - 二刷发现了一个遗漏点， rear的上一个才是尾元素，一刷莫名其妙写对了根本就没记住
  - 加了一个点： rear取值的时候要找rear-1的位置的，rear本身是个空
- 15 3sum
  - 使用双指针处理twosum子问题时，两端去重不熟练，判定条件是 左指针只要next和当前相等，就往下++， 之后再++一次进入下一个循环，右方向同理
  - 细节要点： 排序和去重
- 142 linked-list-cycle-ii
  - 要使用快慢指针的方法，一定是要slow从head.Next fast从head.Next.Next出发，如果 slow从head，fast从head.Next出发，虽然会相遇，但之后再进行等速指针时就永远也遇不上了
  - 细节要点：如上
- 42 trapping-rain-water
  - 用自己的理解实现了一版，还不太能理解栈的解法
    - 双指针解法，先找到两侧最高，其外的不会存水
    - 从低的一侧向中间推进，遇到低的就加水，遇到高的，指针跳过去
    - 循环上一条直到两指针相遇

## 练习记录

### Array
| 题号 | 名称 | 难度 | 分类 | 第一掌 | 第二掌 | 第三掌 | 第四掌 | 第?掌 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| [283 move-zeroes](https://leetcode.com/problems/move-zeroes/discuss/?currentPage=1&orderBy=most_votes&query=) | [移动零](https://leetcode-cn.com/problems/move-zeroes/)| 🟢 简单 | 数组 | 08.10✅  | 08.10✅  | 08.11✅  | 08.18 | - |
| [70 climbing-stairs](https://leetcode.com/problems/climbing-stairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/)| 🟢 简单 | 数组 | 08.10✅  | 08.10✅  | 08.12✅  | 08.19 | - |
| [15 3sum](https://leetcode.com/problems/3sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [三数之和](https://leetcode-cn.com/problems/3sum/)| 🟡 中等 | 数组 | 08.10✅  | 08.12✅  | 08.13✅  | 08.20 | - |
| [11 container-with-most-water](https://leetcode.com/problems/container-with-most-water/discuss/?currentPage=1&orderBy=most_votes&query=) | [盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)| 🟡 中等 | 数组 | 08.10✅  | 08.10✅  | 08.12✅ | 08.19 | - |


### Linked list
| 题号 | 名称 | 难度 | 分类 | 第一掌 | 第二掌 | 第三掌 | 第四掌 | 第?掌 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| [146 lru-cache](https://leetcode.com/problems/lru-cache/discuss/?currentPage=1&orderBy=most_votes&query=) | [LRU缓存机制](https://leetcode-cn.com/problems/lru-cache/)| 🟡 中等 | 链表 | 08.10✅ | 08.10✅ | 08.12✅ | 08.19 | - |
| [206 reverse-linked-list](https://leetcode.com/problems/reverse-linked-list/discuss/?currentPage=1&orderBy=most_votes&query=) | [反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)| 🟢 简单 | 链表 | 08.10✅  | 08.12✅  | 08.13✅   | 08.20 | - |
| [24 swap-nodes-in-pairs](https://leetcode.com/problems/swap-nodes-in-pairs/discuss/?currentPage=1&orderBy=most_votes&query=) | [两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/)| 🟡 中等 | 链表 | 08.10✅ | 08.12✅  | 08.13✅   | 08.20 | - |
| [141 linked-list-cycle](https://leetcode.com/problems/linked-list-cycle/discuss/?currentPage=1&orderBy=most_votes&query=) | [环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)| 🟢 简单 | 链表 |08.10✅ | 08.12✅  | 08.13✅  | 08.20 | - |
| [142 linked-list-cycle-ii](https://leetcode.com/problems/linked-list-cycle-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)| 🟡 中等 | 链表 | 08.10✅  | 08.12✅  | 08.13✅  | 08.20 | - |
| [25 reverse-nodes-in-k-group](https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/?currentPage=1&orderBy=most_votes&query=) | [K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)| 🔴️ 困难 | 链表 | 08.13✅  | 08.14✅ | 08.15 | 08.21 | - |

### Stack & Queue
| 题号 | 名称 | 难度 | 分类 | 第一掌 | 第二掌 | 第三掌 | 第四掌 | 第?掌 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| [20 valid-parentheses](https://leetcode.com/problems/valid-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)| 🟢 简单 | 栈、队列 | 08.10✅ | 08.12✅  | 08.13 | 08.20 | - |
| [155 min-stack](https://leetcode.com/problems/min-stack/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小栈](https://leetcode-cn.com/problems/min-stack/)| 🟢 简单 | 栈、队列 | 08.10✅  | 08.12✅  | 08.13 | 08.20 | - |
| [84 largest-rectangle-in-histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/discuss/?currentPage=1&orderBy=most_votes&query=) | [柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)| 🔴️ 困难 | 栈、队列 | 08.13✅ | 08.13 | 08.14 | 08.21 | - |
| [239 sliding-window-maximum](https://leetcode.com/problems/sliding-window-maximum/discuss/?currentPage=1&orderBy=most_votes&query=) | [滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)| 🔴️ 困难 | 栈、队列 | 08.10✅  | 08.12✅  | 08.13 | 08.20 | - |

### Homework
| 题号 | 名称 | 难度 | 分类 | 第一掌 | 第二掌 | 第三掌 | 第四掌 | 第?掌 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | 
| [26 remove-duplicates-from-sorted-array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)| 🟢 简单 | 数组、链表、跳表 | 08.10✅ | 08.12✅  | 08.14✅  | 08.21 | - |
| [189 rotate-array](https://leetcode.com/problems/rotate-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [旋转数组](https://leetcode-cn.com/problems/rotate-array/)| 🟢 简单 | 数组、链表、跳表 | 08.11✅ | 08.12✅ | 08.14✅  | 08.21 | - |
| [21 merge-two-sorted-lists](https://leetcode.com/problems/merge-two-sorted-lists/discuss/?currentPage=1&orderBy=most_votes&query=) | [合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)| 🟢 简单 | 数组、链表、跳表 | 08.11✅  | 08.13✅  | 08.15✅  | 08.22 | - |
| [88 merge-sorted-array](https://leetcode.com/problems/merge-sorted-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)| 🟢 简单 | 数组、链表、跳表 | 08.11✅  | 08.13✅  | 08.15✅  | 08.22 | - |
| [1 two-sum](https://leetcode.com/problems/two-sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [两数之和](https://leetcode-cn.com/problems/two-sum/)| 🟢 简单 | 数组、链表、跳表 | 08.11✅ | 08.12✅ | 08.15✅  | 08.22 | - |
| [283 move-zeroes](https://leetcode.com/problems/move-zeroes/discuss/?currentPage=1&orderBy=most_votes&query=) | [移动零](https://leetcode-cn.com/problems/move-zeroes/)| 🟢 简单 | 数组、链表、跳表 | 08.10 | 08.10 | 08.11✅  | 08.18 | - |
| [66 plus-one](https://leetcode.com/problems/plus-one/discuss/?currentPage=1&orderBy=most_votes&query=) | [加一](https://leetcode-cn.com/problems/plus-one/)| 🟢 简单 | 数组、链表、跳表 | 08.11✅ | 08.13✅  | 08.14✅  | 08.21 | - |
| [641 design-circular-deque](https://leetcode.com/problems/design-circular-deque/discuss/?currentPage=1&orderBy=most_votes&query=) | [设计循环双端队列](https://leetcode-cn.com/problems/design-circular-deque/)| 🟡 中等 | 栈、队列 | 08.11✅  | 08.13✅  | 08.14✅  | 08.21 | - |
| [42 trapping-rain-water](https://leetcode.com/problems/trapping-rain-water/discuss/?currentPage=1&orderBy=most_votes&query=) | [接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)| 🔴️ 困难 | 栈、队列 | 08.12✅ | 08.13✅  | 08.14✅  | 08.21 | - |


