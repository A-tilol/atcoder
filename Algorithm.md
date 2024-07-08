# Algorithm

## グラフ

### 木の直径

木に存在する2つノード間の最大距離を木の直径という

#### アルゴリズム

> 1. 任意の頂点 s を選ぶ
> 2. s から DFS/BFS によって、最も遠くにある頂点 u を探索する
> 3. u から DFS/BFS によって、最も遠くにある頂点 v を探索する
> 4. u と v を結ぶパスが直径となる  
> [https://algo-logic.info/tree-diameter/](https://algo-logic.info/tree-diameter/)

#### 計算量: O(N)

#### 例題

- [E - Tree and Hamilton Path 2](https://atcoder.jp/contests/abc361/tasks/abc361_e)

## ニム和 (Nim Sum)

山くずしのような２人交互型ゲームの勝敗判定法。ニム和を計算し0になった方の負けが確定する

**ニム和の計算方法**: すべての山に積まれているものの個数について排他的論理和を計算する

```python
a = [10, 5, 8]  # それぞれの山に積まれているものの個数
nim = 0
for ai in a:
    nim ^= ai
print(nim)  # 0なので現局面で手番になっている方の負け
```

### 例題

- [A33 - Game 2](https://atcoder.jp/contests/tessoku-book/tasks/tessoku_book_ag)
