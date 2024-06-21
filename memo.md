# Memo

## General

- 3つの要素に対する操作を考える際には、真ん中の要素に注目すると良い場合がある

## Python Tips

### library

- permutation -> itertools
- queueから先頭要素を取り出す際にはcollections.dequeを使う. a = q[0] q = q[1:]は遅い
  - 単にpop()はlistの方が速そう

### speed

- 文字列結合よりint配列のappendの方が10倍高速
- pypyでは再起の速度、メモリ効率が悪い。深さ10^5程度でTLE/MLEする。
  - JITコンパイル時に再起部分のコード展開を行っているためっぽい
  - `import pypyjit; pypyjit.set_param("max_unroll_recursion=-1")` で展開の深さを制限することで対応可能
  - Cpythonで提出した方が速いこともある
  - 深さ10^6以上はスタックを使用したDFSしないとTLEする
- 文字列はイミュータブルなので結合の繰り返しは避ける

## Check List

- 最小ケース、境界値
- 時間、空間計算量
- 入力（整数/非整数/正負）
- 特殊ケースがないか

## Algorithm

### ニム和 (Nim Sum)

山くずしのような２人交互型ゲームの勝敗判定法。ニム和を計算し0になった方の負けが確定する

**ニム和の計算方法**: すべての山に積まれているものの個数について排他的論理和を計算する

```python
a = [10, 5, 8]  # それぞれの山に積まれているものの個数
nim = 0
for ai in a:
    nim ^= ai
print(nim)  # 0なので現局面で手番になっている方の負け
```

例題

- [A33 - Game 2](https://atcoder.jp/contests/tessoku-book/tasks/tessoku_book_ag)