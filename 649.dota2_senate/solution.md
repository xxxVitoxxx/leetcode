# Solution

## Two Queue
建立兩個 queue 分別儲存兩個政黨的 index ，讓兩邊比較其第一位參議員的 index 大小，比較大的一方會被禁止投票權，勝利的一方會將其 index 加上原先 senate 長度並加到所屬政黨的 queue 中繼續參加下一輪的比較，只要兩邊政黨還有參議員就要一直互相比較，直到某一方政黨的參議員都被禁止投票。  

### 比較方式
index 較小的代表可以先行使權利讓 index 較大的無法投票，贏的一方，因為還擁有投票權，所以要將其 index 加上原先 senate 長度並加到所屬政黨的 queue 中。  

**當勝利的一方要異動 index 的原因**：假設 senate = RDD 則勝利的政黨應為 Dire 。  
因為一開始 senate[0] -> R 禁止 senate[1] -> D，再來 senate[2] -> D 因後面沒有參議員了，所以往前看有一個 Radiant 的參議員，所以他可以禁止 senate[0] -> R ，所以結果會是 Dire 勝利。  

這就是我們要異動 senate[0] -> R 的 index 原因，只要勝利的一方，我將其 index 加上原先 senate 長度並加到所屬的政黨 queue 中，就相當於我將 senate[0] -> R 排到 senate[2] -> D 的後面，這樣我就不需要往前判斷，只需要往後判斷只要 index 較小的一方就勝利，直到某一方政黨的參議員都被禁止。  

```
senate = RDD

radiant queue = [0]
dire queue    = [1, 2]

first round: because 0 < 1, add 0 + len(senate) to radiant queue.
radiant queue = [3]
dire queue    = [2]

second round: because 2 < 3, add 2 + len(senate) to dire queue.
radiant queue = []
dire queue    = [5]

because all senators of radiant is ban, so dire party is victory.
```

## One Queue

與 Two Queue 概念一樣，只是將兩個政黨的參議員放在同一個 queue 中，先取出 queue[0] 的參議員，在檢查 queue 中是否有其他參議員，如果有就禁止他（將他從 queue 中移除），直到其中一方的參議員都被禁止。  
