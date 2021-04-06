# go-code-problem-day23

給定一個 integer array，從中找出 shortest unsorted subarray，並 return 其長度。若我們將該 unsorted subarray 由小到大排序，則整個 array 會變成 sorted array。

若原 array 是 sorted array，則 shortest unsorted subarray 的長度為 0
Example 1:

Input: nums = [1, 6, 3, 8, 10, 9, 10, 20]
Output: 5
Explanation: Shortest subarray 為 [6, 3, 8, 10, 9]，將此 subarray 由小到大排序以後，原本的 array 就會是 sorted array
Example 2:

Input: nums = [-3, -1, 3, 4, 10]
Output: 0
Explanation: nums 本身是 sorted array，所以 shortest unsorted subarray 長度為 0

## 題目分析

給定一個 array A

要找出其中unsorted subArray S

讓 S 排序之後 就可以讓 A變成 sorted

首先是 所謂 unsorted 就是存在 integer i, j  i < j

s.t. A[i] > A[j]

所以初步想法就是 一直計算目前累計的MAX 以及 S 最左邊的index 

然後確認目前的索引 currentIdx 的值 是否有 大於 MAX

如果沒有 則代表目前的 S length為 原本的 currentIdx - S 最左邊的index + 1

## verify

```sh===
cd sol
go test -v
```