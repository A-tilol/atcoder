class SegTree:
    """1-indexed Segment Tree"""

    def __init__(self, n: int, init_v, operator):
        """セグメントツリーを初期化

        Args:
            n (int): 対象の個数
            init_v (int, float): 初期値
            operator (function): 2項演算子
                max: max
                min: min
                sum: lambda a, b: a + b
        """
        self.n = n
        self.init_v = init_v
        self.op = operator
        self.a = [init_v] * (2 * 2 ** (n - 1).bit_length())
        self.l_n = len(self.a) // 2  # 葉の数

    def update(self, i: int, v: int):
        """一点更新

        Args:
            i (int): 更新する対象のインデックス
            v (int): 更新する値
            vis (bool, optional): Trueのときツリーを可視化する. Defaults to False.
        """

        i += self.l_n - 1
        self.a[i] = v

        while i > 1:
            i //= 2
            self.a[i] = self.op(self.a[2 * i], self.a[2 * i + 1])

    def query(self, L: int, r: int):
        """閉区間 [L, r] に対するクエリ結果を返す"""

        ret = self.init_v
        L += self.l_n - 1
        r += self.l_n - 1

        while L <= r:
            if L & 1:
                ret = self.op(ret, self.a[L])
                L += 1
            if not r & 1:
                ret = self.op(ret, self.a[r])
                r -= 1
            L //= 2
            r //= 2

        return ret

    def vis(self):
        print()
        s, L = 1, 1
        while s + L <= len(self.a):
            print(self.a[s : s + L])
            s += L
            L *= 2
        print()


class LazySegTree:
    """1-indexed Lazy Segment Tree"""

    def __init__(self, n: int, init_v, operator):
        """セグメントツリーを初期化

        Args:
            n (int): 対象の個数
            init_v: 初期値
            operator: 構築する木に対応する2項演算子
                最大値: max
                最小値: min
                合計値: lambda a, b: a + b
        """
        self.n = n
        self.init_v = init_v
        self.op = operator
        self.a = [init_v] * (2 * 2 ** (n - 1).bit_length())
        self.l_n = len(self.a) // 2

    def set_leaves(self, arr: list):
        """葉の値を設定"""

        for i in range(len(arr)):
            self.a[self.l_n + i] = arr[i]

    def get_leaves(self, L: int, r: int) -> list:
        """閉区間 [L, r] の葉の値を取得"""
        return self.a[self.l_n : self.l_n + self.n]

    def set_leaf_v(self, i: int, v):
        """指定した葉の値を更新"""

        i += self.l_n - 1
        self.a[i] = v

    def get_leaf_v(self, i: int):
        """上位ノードを考慮した葉の値を取得"""

        ret = 0
        i += self.l_n - 1
        while i > 0:
            ret += self.a[i]
            i //= 2
        return ret

    def get_leaf_v_as_is(self, i: int):
        """現時点で葉に設定されている値を取得"""

        return self.a[self.l_n - 1 + i]

    def update_range(self, L: int, r: int, v):
        """閉区間 [L, r] に対して値を更新する。該当ノードのみ更新される"""

        L += self.l_n - 1
        r += self.l_n - 1
        while L <= r:
            if L & 1:
                self.a[L] = self.op(self.a[L], v)
                L += 1
            if not r & 1:
                self.a[r] = self.op(self.a[r], v)
                r -= 1
            L //= 2
            r //= 2

    def update_whole_leaves(self):
        """内部ノードに蓄積された更新をすべての葉に反映し、内部ノードを初期化"""

        for i in range(1, self.n + 1):
            self.a[self.l_n + i - 1] = self.get_leaf_v(i)

        self.a[: self.l_n] = [self.init_v] * self.l_n

    def vis(self):
        print()
        s, L = 1, 1
        while s + L <= len(self.a):
            print(self.a[s : s + L])
            s += L
            L *= 2
        print()
