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

- 境界値(0, 上限, 下限)
- 最小・最大ケース
- 時間、空間計算量
- 入力（整数/非整数/正負）
- 特殊ケースの考慮
- 出力フォーマット（Yes or yes, No or -1）
