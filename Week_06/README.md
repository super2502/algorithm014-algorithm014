学习笔记

# 总结

- 为什么难
  - 主要是没有标准套路，dp的状态和转移方程只是框架，但是状态定义本身和状态转移并没有标准可循
  - 需要按照类型总结，如果熟练了是不是才有可能在题目变种时联想到相关的类型？
    - 如何总结每个类型问题的特点？ 有了特点是不是就可以套用已知dp问题解法？
    - 总结了几个系列？
  - 问题种类发散，很难系统的整理成一个完整的体系
  - dp问题本身从理解到解题本身都比简单问题时间要长很多，时间不够的话过遍数也成了问题
    - 平均一小时过2道题，有时一个半小时也搞不出来一个，包括第二遍做也是
- 背包套题
  - 为何没有背包套题的训练？  
- 还有哪些典型的dp系列？
- 系列1，各种数组系列
  - 严格来说这不是一个系列，类型也不一样，但是特点就是只有一个数组，希望时间复杂度尽量低的找到什么目标
  - 1.1 最大子序和，最大子序乘积： 要用O(n)复杂度求解，关键点是max比较的对象是什么，是当前索引值和其之前的max比
  - 1.2 有跳跃的数组，一般数组的某个位置的值会决定其他位置的一些结果
    - 零钱兑换，先记套路吧，有可能本质还是青蛙过河，但要求解的目标不同
    - 青蛙过河，注意这个问题通过预测下一跳的所有可能落点再求石头是否存在来大量剪枝
    - 跳房子系列，这类虽然贪心法解，但是其实也是个青蛙过河的问题，石头都挨着而已，只是贪心的效率最高
    - 所以跳房子1就是青蛙过河，跳房子2就是零钱兑换？？？
  - 1.3 打家劫舍，股票问题，涉及到其有代表性，单独总结
  - 1.4 子数组问题，因为特点和解法都不同，单独总结各种类型系列  
- 系列2，杨辉三角系列
  - 相对最好理解的dp
  - 特点：dp数组的形态和三角的形态是一样的，对应点保存最优值即可
  - 像二叉树，但是在下一层相邻子节点又重合在一起，一般dp\[i]\[j]和dp\[i-1]\[j]与dp\[i]\[j-1]有关
  - 也可能是dp\[i-1]\[j-1],dp\[i-1]\[j]有关，和三角怎么放置有关系
  - 最小路径和，不同路径I,II，三角倒下变成矩形
  - 三角形最小路径和，标准的三角
- 系列3，能变成杨辉三角形态的，并且和dp\[i-1]\[j-1]也有关系列
  - 特点：dp一般也是个二维数组，目标是个现成的二维数组，或者是两个一维数组
  - 这个就只能记题了
  - 两个一维数组： 两个字符串系列
    - 看到这种先搞一个m*n dp数组，再仔细研究某个ij点和左，上，左上三个点的关系
    - 有时边界条件可以用m+1*n+1引入空字串简化代码写法，类似哨兵用法
    - 1143 longest-common-subsequence 最长公共子序列
    - 72 edit-distance 编辑距离
    - 记忆中好像还有两字符串问题，但想不出来了，也许就这几个
  - 二维数组目标
    - 221 maximal-square 最大正方形
- 系列4，某个状态n和前(或自身状态)有限状态数(一般是n, n-1)有关的
  - 严格说，什么类型的都得和之前的有关
  - 这里说的关系性更强，一般前一个状态取值就能直接决定当前状态是否的
  - 这种就在是的时候考虑前一个状态，否的时候跳过前一个状态
  - 打家劫舍系列
    - 当前状态取值就能决定当前是否考虑
  - 解码方法
    - 当前状态和前一个状态同时决定是否考虑
  - 强行总结的特点，描述的不太好，也还没有完全理解
- 系列5，股票系列套题
  - 严格说，这个系列也是系列4，只是更具代表性
  - 三维dp想明白之后，这个系列本身问题不大了，但是怎么推广到其他场景？？
    - 带交易次数k的就是在第三维上在售卖时+1
  - 唯一一个不好理解的是含冷冻期的问题，这个状态转换很难想清楚
    - 比如某天卖了，这天到底算持有？不持有？还是冷冻？
    - 无脑一点的记法就是冷冻的状态有+price的操作，持有的状态有-price的操作
  - 细节问题
    - 初始化
      - 三维dp不持有时所有次数交易全0
      - 持有时0次交易-price0, k次交易-INF
    - 注意dp的交易次数维度的长度是k+1
  - 压缩问题
    - 二维降一维能理解，k=2的三维降一维也能理解
    - k = k的三维降维想不清楚了，因为k这个维度本身又不能降，剩下的二维带n怎么降想不明白 
- 子数组系列1：
  - 647 palindromic-substrings，回文字串，马拉车算法一时半会搞不定
  - 还不知道类似的问题有什么
- 子数组系列2：
  - 621 task-scheduler 任务调度器
  - 这怎么看都是个优先级队列解决的问题
- 子数组系列3:
  - 32 longest-valid-parentheses 最长有效括号
  - 括号总离不开栈，好像也不是dp的菜
- 子数组系列4:
  - 从题解上看到两句话，现在还不能太理解，这也是一种特点吧。
    - 「将数组分割为 mm 段，求……」是动态规划题目常见的问法。
    - 「使……最大值尽可能小」是二分搜索题目常见的问法。
  - 410_split-array-largest-sum 分割数组的最大值
  - 二分用通用模版搞不定
  - 应该先学前置的二分，就是比xx小的最大值找法，只能left < right; mid <= target then right = mid 这个模版做
- 子数组系列5， ij卡子串
  - dp维度行表示遍历i，列表示遍历j
  - 最终结果是dp\[0]\[n-1]
  - 这一定是一个上三角矩阵，因为j一定在i后边
  - 只能背case
  - 上三角矩阵从对角线斜向上递推公式, 参考312_burst-balloons，戳气球
```
    // k为对角线方向 向右上角一层一层递推
    for k := 0; k < n; k++ {
        for i := 0; i < n - k ; i++ {
            j := k + i   // 注意这里 每层的j 就是i加上层数
            dp[i][j] = ...
            //for l:= i+1;l<j;l++{
            //    dp[i][j] = max(dp[i][j], dp[i][l] + dp[l][j] + nums[i]*nums[l]*nums[j])
            //}
        }
    }
```

# 随笔

## 做dp的题太慢了
  - 理解就很慢，看答案理解的也慢
  - 做一遍更慢
  - 还没搞清楚套路，只能一个类型题记一下
  - 两周的dp做的非常崩溃，大部分都想不出来，看了答案自己也写不出来
  
## 大概的一点感觉

- 子问题就是从0开始到n-1, n-2等问题的某种求和或求最值，这种dp状态数组和转移方程虽然也不好想，但最容易理解，dp初始化一般也都是一个二维数组（？）其dp\[i]\[0], dp\[0]\[j] 都是0，这些也可能是额外补出来辅助的，时间复杂度一般都是O(n),只要遍历一次就行了，dp一般也可以压缩成一个一维数组，运气好的还能压缩成几个变量
  - 这种问题首先是靠倒着(自顶向下)推算出来的，因为子问题很容易总结，然后再正着(自底向上)给出dp数组
  - 爬楼梯，斐波那契数列
  - 偷房子问题
  - 最大矩形

- 同上，自顶向下的思路一般在dp是一维数组时更容易想到，而dp是二维数组时， 自底向上是不是更容易想
  - 字符串匹配相关的， 都是先拿0个字符匹配，1个字符匹配，最后推到出n个字符
  - 但这样的思考逻辑是不是本质上也是因为曾经自顶向下考虑过才能想出来的？？ 所以有可能跟上面的是一回事
  
- 子问题是一个数组里i, j区间之内的最优解的某种求和方式，是从中间向两边扩散的形态，这种一般会以ij作为维度制作二维dp数组，而数组初始化是以i=j,i+1=j这种形态开始，一般是一个上三角矩阵，从对角线开始递推，这种问题是要穷举ij的，所以时间复杂度至少是O(n^2)，然后因为递推的过程是斜着并且反复遍历的，所以dp数组也基本上不能压缩
  - 戳气球


  

# 提交作业
```
#学号:G20200343110147
#姓名:benben
#班级:14期1班3组
#语言:go
#作业&总结链接:https://github.com/super2502/algorithm014-algorithm014/tree/golang/Week_06

```

# 实战

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [62 unique-paths](https://leetcode.com/problems/unique-paths/discuss/?currentPage=1&orderBy=most_votes&query=) | [不同路径](https://leetcode-cn.com/problems/unique-paths/)| 🟡 中等 | 数组、动态规划 | - |9.8✅|9.13✅|9.14|9.21||
| [63 unique-paths-ii](https://leetcode.com/problems/unique-paths-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [不同路径II](https://leetcode-cn.com/problems/unique-paths-ii/)| 🟡 中等 | 数组、动态规划 | - |9.8✅|9.15✅|9.16|9.23||
| [1143 longest-common-subsequence](https://leetcode.com/problems/longest-common-subsequence/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长公共子序列](https://leetcode-cn.com/problems/longest-common-subsequence/)| 🟡 中等 | 动态规划 | - |9.8✅|9.15✅|9.16|9.23||
| [120 triangle](https://leetcode.com/problems/triangle/discuss/?currentPage=1&orderBy=most_votes&query=) | [三角形最小路径和](https://leetcode-cn.com/problems/triangle/)| 🟡 中等 | 数组、动态规划 | - |9.9✅|9.15✅|9.16|9.23||
| [53 maximum-subarray](https://leetcode.com/problems/maximum-subarray/discuss/?currentPage=1&orderBy=most_votes&query=) | [最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)| 简单 | 数组、动态规划 | - |9.9✅|9.16✅|9.17|9.22||
| [152 maximum-product-subarray](https://leetcode.com/problems/maximum-product-subarray/discuss/?currentPage=1&orderBy=most_votes&query=) | [乘积最大子数组](https://leetcode-cn.com/problems/maximum-product-subarray/)| 🟡 中等 | 数组、动态规划 | - |9.9✅|9.16✅|9.17|9.22||
| [322 coin-change](https://leetcode.com/problems/coin-change/discuss/?currentPage=1&orderBy=most_votes&query=) | [零钱兑换](https://leetcode-cn.com/problems/coin-change/)| 🟡 中等 | 动态规划 | - |9.10✅|9.16✅|9.17|9.22||
| [198 house-robber](https://leetcode.com/problems/house-robber/discuss/?currentPage=1&orderBy=most_votes&query=) | [打家劫舍](https://leetcode-cn.com/problems/house-robber/)| 简单 | 动态规划 | - |9.10✅|9.16✅|9.17|9.22||
| [213 house-robber-ii](https://leetcode.com/problems/house-robber-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [打家劫舍II](https://leetcode-cn.com/problems/house-robber-ii/)| 🟡 中等 | 动态规划 | - |9.10✅|9.16✅|9.17|9.22||
| [121 best-time-to-buy-and-sell-stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)| 简单 | 数组、动态规划 | - |9.10✅|9.16✅|9.17|9.22||
| [122 best-time-to-buy-and-sell-stock-ii](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/)| 简单 | 贪心算法、动态规划 | - |9.13✅|9.16✅|9.17|9.22||
| [123 best-time-to-buy-and-sell-stock-iii](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机III](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/)| 🔴 困难 | 数组、动态规划 | - |9.13✅|9.16✅|9.17|9.22||
| [309 best-time-to-buy-and-sell-stock-with-cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/discuss/?currentPage=1&orderBy=most_votes&query=) | [最佳买卖股票时机含冷冻期](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)| 🟡 中等 | 数组、动态规划 | - |9.13✅|9.18✅|9.19|9.26||
| [188 best-time-to-buy-and-sell-stock-iv](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机IV](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/)| 🔴 困难 | 数组、动态规划 | - |9.13✅|9.18✅|9.19|9.26||
| [714 best-time-to-buy-and-sell-stock-with-transaction-fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/?currentPage=1&orderBy=most_votes&query=) | [买卖股票的最佳时机含手续费](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)| 🟡 中等 | 数组、动态规划 | - |9.13✅|9.18✅|9.19|9.26||

# HomeWork

| 题号 | 名称 | 难度 | 分类 | 备注 | #1 | #2 | #3 | #4 | #？|
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| [64 minimum-path-sum](https://leetcode.com/problems/minimum-path-sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/)| 🟡 中等 | 数组、动态规划 | - |9.11✅|9.18✅|9.19|9.26||
| [91 decode-ways](https://leetcode.com/problems/decode-ways/discuss/?currentPage=1&orderBy=most_votes&query=) | [解码方法](https://leetcode-cn.com/problems/decode-ways/)| 🟡 中等 | 字符串、动态规划 | - |9.18✅|9.20|9.1|9.18||
| [221 maximal-square](https://leetcode.com/problems/maximal-square/discuss/?currentPage=1&orderBy=most_votes&query=) | [最大正方形](https://leetcode-cn.com/problems/maximal-square/)| 🟡 中等 | 动态规划 | - |9.19✅|9.20|9.22|9.29||
| [621 task-scheduler](https://leetcode.com/problems/task-scheduler/discuss/?currentPage=1&orderBy=most_votes&query=) | [任务调度器](https://leetcode-cn.com/problems/task-scheduler/)| 🟡 中等 | 贪心算法、队列、数组 | - |9.19✅|9.20|9.22|9.29||
| [647 palindromic-substrings](https://leetcode.com/problems/palindromic-substrings/discuss/?currentPage=1&orderBy=most_votes&query=) | [回文字串](https://leetcode-cn.com/problems/palindromic-substrings/)| 🟡 中等 | 字符串、动态规划 | - |9.19✅|9.20|9.22|9.29||
| [32 longest-valid-parentheses](https://leetcode.com/problems/longest-valid-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=) | [最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/)| 🔴 困难  | 字符串、动态规划 | - |9.19✅|9.20|9.22|9.29||
| [72 edit-distance](https://leetcode.com/problems/edit-distance/discuss/?currentPage=1&orderBy=most_votes&query=) | [编辑距离](https://leetcode-cn.com/problems/edit-distance/)| 🔴 困难  | 字符串、动态规划 | - |9.19✅|9.20|9.22|9.29||
| [363 max-sum-of-rectangle-no-larger-than-k](https://leetcode.com/problems/max-sum-of-rectangle-no-larger-than-k/discuss/?currentPage=1&orderBy=most_votes&query=) | [矩形区域不超过 K 的最大数值和](https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/)| 🔴 困难  | 队列、二分查找、动态规划 | - |9.7|9.1|9.4|9.11||
| [403 frog-jump](https://leetcode.com/problems/frog-jump/discuss/?currentPage=1&orderBy=most_votes&query=) | [青蛙过河](https://leetcode-cn.com/problems/frog-jump/)| 🔴 困难  | 动态规划 | - |9.19✅|9.20|9.22|9.29||
| [410 split-array-largest-sum](https://leetcode.com/problems/split-array-largest-sum/discuss/?currentPage=1&orderBy=most_votes&query=) | [分割数组的最大值](https://leetcode-cn.com/problems/split-array-largest-sum/)| 🔴 困难  | 二分查找、动态规划 | - |9.19✅|9.20|9.22|9.29||
| [552 student-attendance-record-ii](https://leetcode.com/problems/student-attendance-record-ii/discuss/?currentPage=1&orderBy=most_votes&query=) | [学生出勤记录II](https://leetcode-cn.com/problems/student-attendance-record-ii/)| 🔴 困难  | 动态规划 | - |9.7|9.1|9.4|9.11||
| [76 minimum-window-substring](https://leetcode.com/problems/minimum-window-substring/discuss/?currentPage=1&orderBy=most_votes&query=) | [最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/)| 🔴 困难  | 双指针、滑动窗口 | - |9.7|9.1|9.4|9.11||
| [312 burst-balloons](https://leetcode.com/problems/burst-balloons/discuss/?currentPage=1&orderBy=most_votes&query=) | [戳气球](https://leetcode-cn.com/problems/burst-balloons/)| 🔴 困难 | 分治算法、动态规划 | - |9.11✅|9.12|9.13|9.20||
