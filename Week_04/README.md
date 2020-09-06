学习笔记

# 总结 

- 套路一： DFS/BFS
  - DFS：深度优先搜索， 递归或迭代， 用栈。
  - DFS迭代实际上是树的先序遍历
    - 在做树的染色法迭代遍历时，曾经发现过，只有先序遍历不需要把当前节点压栈两次，其简化版就和现在的dfs迭代模板一样
  - BFS：广度优先搜索， 迭代， 用队列
    - BFS使用队列时一次性pop出同层所有节点，处理完再一起丢进队列，从代码上来看要容易理解，并且更容易处理visited问题
    - 还可以使用成员变量level来保存层数而不用把level塞到队列的每个node里
    - 使用  if deque.Len()> 0 { if deque.Len() > 0 { allResult = pop all } ; for allResults {process; deque.Pushback(allChindrens)}; level++ }  这种模板
    - 这种方式因为用了额外的空间来清空队列，空间使用会多一点，但不超过O(n), 因为队列本身存储也是O(n)的，所以空间复杂度没有什么变化
  - 关于visited：
    - DFS的visited还没怎么接触，根据以前的做DAG的经验看，当一个节点递归回来之后要从visited里去除掉，这样其父节点的其他路径才有可能再次访问到它
    - BFS的visited的使用需要根据具体场景。
      - 剪兄弟节点：同样是单词接龙，如果仅找最小接龙次数，那么在同层的BFS时直接对兄弟节点进行visited剪枝。
      - 只剪父+节点：但如果要穷举所有的最小次数的路径，就要在每次处理完一整层的时候再一次性灌注visited，因为同层节点即使用过，但父路径是不一样的。
    
- 套路二： 二分查找，（最后发现这是一道证明题，证明模板的普适性，然后背就好了。。）
  - 最重要的，二分查找的边界问题
    - 只要记住一个模板，不要其他的
    - left <= right;  left = mid +1; right = mid - 1
    - 然后证明这个模板百搭。
  - 二分搜索魔鬼细节1，分到最后的边界问题
    - 最后的最后，只有一个数，肯定不是结果了，想清楚左右指针的走向
    - 倒数第二后，挨着两个数，mid=left，想清楚左右指针的走向，这个阶段问题最多，很多case也是直接就写俩奇葩数直接测这儿
  - 二分搜索魔鬼细节2，判断完mid == targe之后， 二择的对象是什么
    - 常规搜索，二择的对象还是 mid与target
    - 找有重复数组的最左目标，二择对象也是 mid与target
    - 带旋转的搜索，二择的对象是 left与mid （这个可以理解，但是需要背诵，因为非常绕）
    - 带旋转的最小值，二择的对象是？
  - 二分搜索魔鬼细节3，判断完mid == targe之后， return还是干啥
    - 搜索有重复的数组时，俩人等了不要return，继续往左搜
  - 二分变种1，单调数组，找头一个比目标小的鬼, 具体见表1
    - 魔鬼般的细节2：二择的目标还是mid与target的大小比较
    - 魔鬼般的细节1：查到就剩俩数挨着时，要查找的数如果是其中一个，总能找到，如果也不是，分三种情况
    - 不管什么情况， 拿right就对了
  - 二分变种2，单调数组但是有重复的，找到最左边的目标， 见表2
    - 这是一个奇葩，分析的方式不太一样
    - 魔鬼般的细节1：就剩俩数挨着时，是不是符合模板
  - 二分变种3，搜索单调数组旋转过的
    - 魔鬼般的细节2：判断完mid==target之后，二择的目标是  num[l]与num[mid]的大小，因此其中有个等号也是魔鬼
    - 魔鬼般的细节1：查到就剩俩数挨着时，一定要保证下面三点,具体见表3
      - 1 左小于右时，如果targe比左大，会被扔到右边
      - 2 左大于右时，如果targe比左小，会被扔到右边
    - 额外的魔鬼：因为targe要与左右界比较，所以要在左右界比较时加等号，为啥mid不加，因为mid专门判断了等不等。
  - 二分变种4，单调数组找最小值，待证明
  - 其他没见过的二分，慢慢来
    
|表1|情况|分析|模板: l <= r; l = mid +1; r = mid -1| |
|---|---|---|---|---|
|1|l<target<r|倒数第二次查找， 一定会查left， 查不着 left进一和right重合，查完right以后，会令right-1|right会指向比要查的数小的位置|
|2|l<r<target|倒数第二次查找， 一定会查left， 查不着 left进一和right重合，查完right以后，会令left+1， right不变|right会指向比要查的数小的位置|
|3|target<l<r|倒数第二次查找， 一定会查left， 查不着 right-1和left重合， 查完right以后，继续令right-1 |right会指向比要查的数小的位置|
   
|表2|情况|特殊的分析|模板：l <= r; l = mid +1; r = mid | 魔鬼细节：因为没有target，所以只要比较mid和r， 因为要取左中，所以r左移不能超过mid，这里没办法只能改模板了|
|---|---|---|---|---|
|1|形如4567801 左<中>右| 在右边 |||
|1.1|形如321 左>中>右|不存在|||
|1.2|形如21 左=中>右|在右边|||
|1.3|形如1中=右|已无所谓，只决定l r最后在哪儿|||
|2|形如56701234 左>中<右| 在左边，包含中|||
|3|形如12 左=中<右| 在左边|||
|4|形如012345 左<中<右| 在左边|||
|5|没了|从上边看出不管左和中什么关系，只要中比右小，最小值就在左(包含中)，或者说只要中比右大，就在右边，还不包含中|因此丢到左边时，得带着中丢，所以r=mid, 而丢到右边时不带中，l=mid+1|
|6|返回谁 形如12|mid=left 中<右，走左，r左移至mid与l重合，之后最后一次循环l走掉了，r留在最小值上||不管啥情况，最后r指向最小值|
|7|返回谁 形如21|mid=left 中不小于右，走右，l右移与r重合，之后最后一次循环l走掉了，r留在最小值上||不管啥情况，最后r指向最小值|

|表3|情况|分析|模板：l <= r; l = mid +1; r = mid -1| 细节处理： l <= mid { l <= target<mid? to left: to right} 或者 l > mid {mid < target <= r? to right:to left} |
|---|---|---|---|---|
|1|target>l<r target<l<r| 只剩两个元素永远符合l<=mid, 不可能满足target在两者之间，一定会丢到右边|总之不管怎么样target都会被丢到右边去比一次，就比全了|   |
|2|l<r<target l<r\>target| 跟上边是一样的|这个细节处理 l<=mid加上l<=target<mid配套的太牛逼，直接把所有情况的target全部一脚踢到右边，这里有个问题，如果是找谁比target小就得瞎|这个魔鬼操作只能硬背了|
|3|其他情况已被标准情况cover|||


- 模板

|模板|笔记|理解|历史|
|---|---|---|---|
|DFS递归|https://shimo.im/docs/9CYPpdcPGwXT93QV/|实际上回溯是带数据处理的DFS，在递归前后调用数据的处理和revert就成回溯了，另外回溯更侧重考虑剪枝|week04|
|DFS迭代|https://shimo.im/docs/JqXDvhW9jt6Y9hQV/ |这就是树的先序遍历模板 |week04|
|BFS|https://shimo.im/docs/VkVGpccqqxwvqtgR/|用队列|week04|
|二分查找|https://shimo.im/docs/vkPwvRcktgHdWWdW/|细节需要研究，都是边界问题|week04|
|回溯|https://shimo.im/docs/3c3VYtCW9kyCcGYc|模板很好用|week03|
|单调栈单调队列|https://shimo.im/docs/hyWwqQ39xVcwk3xG/| 理解的还凑合，记得不牢|week02|

- 想不出来的问题整理，只能背了的整理
  - 二叉树公共祖先节点
  - 二叉树展开成右单链表
  - 丑数
  
- 第三周题目第四掌总结
  - 隔周做了第三周的题目，发现好多细节记不清了，比如子集问题回溯法的遍历目标是什么，全排列ii那个visit-1是怎么回事，组合ii里的i>start 都不记得了，调试了才想起来
  - 细思极恐
  - 易错的细节也要整理一下

- 贪心算法
  - 这周没怎么敢动贪心算法
  - 感觉贪心算法才是最难的，根本没有套路，得先知道怎么贪，还得能证明贪的对
  
# 思考

- DFS和回溯
  - 回溯也是dfs，这两个玩意模板还不一样，具体区别在哪？还是没有区别
    - 已总结
- BFS和回溯
  - 广度优先不需要回溯，一个节点的哥哥节点被探索过，之后，visited里永远都有了，后面的弟弟节点也不需要再走哥哥节点，因此跟回溯不同，不会撤销visited
  - 广度优先遍历使用队列和visited时，可以一次pop出所有节点，处理完之后全塞进队尾，之后再统一加visited
  - 上面两种情况已总结，是不同的使用场景
  

# 随笔

- 第四周的问题明显的需要背诵的更多了，根本想不出来
- 二分查找，x平方根， 需要取整，这里已经懵了， mid小的时候有可能确是解，所以是不是要cache一下这个mid
    - 已总结，就是套模板，在最后两个数的时候仔细辨识一下
- 还完全不能理解二分的各种边界问题
  - 旋转排序数组的最小值， 边界边界边界搞不清楚
  - 总有 left <= right,  right = mid - 1, left = mid + 1的标准解法
  - 怎么统一这个模板？ 才能一定用到上面的边界？
  - 持续总结中

- 非常惊诧的发现了这样一个模板，和标准二分模板一模一样，而且能获取比要搜索的值小的最大的那个。。。, 完全搞不懂为什么
  - ok 今天稍微懂了点， return right也行 @2020-08-31
  - 已入总结 @2020-09-02
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
| [102 binary-tree-level-order-traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/discuss/?currentPage=1&orderBy=most_votes&query=) | [二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)| 🟡 中等 | 深度优先、广度优先 | - | 8.30✅|9.1✅|9.4✅|9.11||
| [433 minimum-genetic-mutation](https://leetcode.com/problems/minimum-genetic-mutation/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小基因变化](https://leetcode-cn.com/problems/minimum-genetic-mutation/)| 🟡 中等 | 深度优先、广度优先 | - | 8.30✅|9.1✅|9.4✅|9.11||
| [22 generate-parentheses](https://leetcode.com/problems/generate-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [括号生成](https://leetcode-cn.com/problems/generate-parentheses/)| 🟡 中等 | 深度优先、广度优先 | - | 8.30✅|9.1✅|9.4✅|9.11||
| [515 find-largest-value-in-each-tree-row](https://leetcode.com/problems/find-largest-value-in-each-tree-row/discuss/?currentPage=1&orderBy=most_votes&query=) | [在每个树行中找最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)| 🟡 中等 | 深度优先、广度优先 | - | 8.30✅|9.1✅|9.4✅|9.11||
| [322 coin-change](https://leetcode.com/problems/coin-change/discuss/?currentPage=1&orderBy=most_votes&query=) | [零钱兑换](https://leetcode-cn.com/problems/coin-change/)| 🟡 中等 | 动态规划 | - |8.30✅|9.1✅|9.4✅|9.11||
| [69 sqrtx](https://leetcode.com/problems/sqrtx/discuss/?currentPage=1&orderBy=most_votes&query=) | [x 的平方根](https://leetcode-cn.com/problems/sqrtx/)| 🟢 简单 | 二分查找 | - | 8.30✅|9.1✅|9.4✅|9.11||
| [367 valid-perfect-square](https://leetcode.com/problems/valid-perfect-square/discuss/?currentPage=1&orderBy=most_votes&query=) | [有效的完全平方数](https://leetcode-cn.com/problems/valid-perfect-square/)| 🟢 简单 | 二分查找 | - |8.30✅|9.1✅|9.4✅|9.11||

### 课后作业

| 题号 | 名称 | 难度 | 分类 | 备注 |#1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- |--- | --- | --- | --- | --- |
| [860 lemonade-change](https://leetcode.com/problems/lemonade-change/discuss/?currentPage=1&orderBy=most_votes&query=) | [柠檬水找零](https://leetcode-cn.com/problems/lemonade-change/)| 🟢 简单 | 贪心算法 | - |8.30✅|9.2✅|9.7|9.18||
| [122 best-time-to-buy-and-sell-stock-ii](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)| 🟢 简单 | 贪心算法 | - |8.30✅|9.2✅|9.7|9.18||
| [455 assign-cookies](https://leetcode.com/problems/assign-cookies/discuss/?currentPage=1&orderBy=most_votes&query=) | [分发饼干](https://leetcode-cn.com/problems/assign-cookies/)| 🟢 简单 | 贪心算法 | - | 8.30✅|9.2✅|9.7|9.18||
| [874 walking-robot-simulation](https://leetcode.com/problems/walking-robot-simulation/discuss/?currentPage=1&orderBy=most_votes&query=) | [模拟行走机器人](https://leetcode-cn.com/problems/walking-robot-simulation/)| 🟢 简单 | 贪心算法 | - |8.30✅|9.2✅|9.7|9.18||
| [127 word-ladder](https://leetcode.com/problems/word-ladder/discuss/?currentPage=1&orderBy=most_votes&query=) | [单词接龙](https://leetcode-cn.com/problems/word-ladder/)| 🟡 中等 | 深度优先、广度优先 | - |8.30✅|9.2✅|9.7|9.18||
| [200 number-of-islands](https://leetcode.com/problems/number-of-islands/discuss/?currentPage=1&orderBy=most_votes&query=) | [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)| 🟡 中等 | 深度优先、广度优先 | - |8.30✅|9.2✅|9.7|9.18||
| [529 minesweeper](https://leetcode.com/problems/minesweeper/discuss/?currentPage=1&orderBy=most_votes&query=) | [扫雷游戏](https://leetcode-cn.com/problems/minesweeper/)| 🟡 中等 | 深度优先、广度优先 | - |8.30✅|9.2✅|9.7|9.18||
| [55 jump-game](https://leetcode.com/problems/jump-game/discuss/?currentPage=1&orderBy=most_votes&query=) | [跳跃游戏](https://leetcode-cn.com/problems/jump-game/)| 🟡 中等 | 贪心算法 | - |8.30✅|9.3✅|9.5|9.11||
| [33 search-in-rotated-sorted-array](https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)| 🟡 中等 | 二分查找 | - |8.30✅|9.3✅|9.6✅|9.13||
| [74 search-a-2d-matrix](https://leetcode.com/problems/search-a-2d-matrix/discuss/?currentPage=1&orderBy=most_votes&query=) | [搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)| 🟡 中等 | 二分查找 | - |8.30✅|9.3✅|9.7✅|9.18||
| [153 find-minimum-in-rotated-sorted-array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/discuss/?currentPage=1&orderBy=most_votes&query=) | [寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)| 🟡 中等 | 二分查找 | - |8.30✅|9.3✅|9.6✅|9.13||
| [126 word-ladder-ii](https://leetcode.com/problems/word-ladder-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [单词接龙 II](https://leetcode-cn.com/problems/word-ladder-ii/)| 🔴 困难 | 深度优先、广度优先 | - | 9.3✅|||||
| [45 jump-game-ii](https://leetcode.com/problems/jump-game-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/)| 🔴 困难 | 贪心算法 | - |9.3✅ |||||
