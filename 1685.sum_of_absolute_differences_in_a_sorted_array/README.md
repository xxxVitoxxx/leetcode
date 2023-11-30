# Solution

## Approach 1: Prefix Sum

要計算 nums[i] 與其他數值的差的絕對值，可以將問題分成以下兩部分，只要將兩部分的結果加總就是答案。
1. 計算比 nums[i] 小的數字與 nums[i] 差的絕對值的和
2. 計算比 nums[i] 大的數字與 nums[i] 差的絕對值的和

答案會是 (nums[i]-nums[0])+(nums[i]-nums[1])+...+(nums[i]-nums[i-1])+...+(nums[i+1]-nums[i])+(nums[i+2]-nums[i])+...+(nums[n-1]-nums[i]) ，且因為 nums 是遞升排序，所以可以確定 nums[0] < nums[i] < nums[n-1] 所以只要拿後面的數字減前面的數字就不用特別計算絕對值。

簡化後會變成 `(nums[i]*left_count - left_sum) + (right_sum - nums[i]*right_count)`

## Approach 2

觀察數組 [1, 4, 6, 8, 10] ，會發現一定的規律。

```
nums 1 4 6 8 10
     ^ 3 5 7 9 = 24 -> diff sum
     3 ^ 2 4 6 = 15 -> diff sum
     5 2 ^ 2 4 = 13 -> diff sum
     7 4 2 ^ 2 = 15 -> diff sum
     9 6 4 2 ^ = 21 -> diff sum
```

* 第一個 diff sum 為各數值到 1 的差的絕對值，因此可以知道 answer[0] = 24
* 觀察第二個 diff sum 會發現兩個 diff sum 相差 9 再往下看第三個 diff sum 會發現有一定的規律

要計算第二個 diff sum 時可以透過 second_diff_sum = first_diff_sum + "^ 左邊的個數" * (nums[1]-nums[0]) - "^ 到最後一位的個數" * (nums[1]-nums[0])

簡化後，先計算出 index 0 的答案後，要求得後面 index i 的答案，其過程為  
`diffSum += i * (nums[i]-nums[i-1]) - (len(nums)-i) * (nums[i]-nums[i-1])`

與上述方法類似，但相對更精簡的方法。  

```
nums 1 4 6 8 10 -> sum = 29
answer[0] -> sum = 24
answer[1] -> sum = 15
answer[2] -> sum = 13
answer[3] -> sum = 15
answer[4] -> sum = 21
```

透過計算所有數值的總和再去做左右邊的相減  
`sum += i * (nums[i]-nums[i-1]) - (len(nums)-i) * (nums[i]-nums[i-1])`

再簡化一下會得到 `sum += (nums[i]-nums[i-1]) * (2*i - len(nums))`
