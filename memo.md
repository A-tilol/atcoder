# Memo

## General

- 3つの要素に対する操作を考える際には、真ん中の要素に注目すると良い場合がある
- 再起を実装する際には深さに注意しDPへの置き換えを検討する。10^5程度の深さでもメモリ操作が多いとTLEになる

## Python Tips

### library

- permutation -> itertools
- queueから先頭要素を取り出す際にはcollections.dequeを使う. a = q[0] q = q[1:]は遅い

### speed

- 文字列結合よりint配列のappendの方が10倍高速
