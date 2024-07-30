class UnionFind:
    def __init__(self, n):
        self.par = [i for i in range(n)]

    def root(self, v):
        if self.par[v] == v:
            return v
        self.par[v] = self.root(self.par[v])
        return self.par[v]

    def same(self, x, y):
        return self.root(x) == self.root(y)

    def unite(self, x, y):
        self.par[self.root(x)] = self.root(y)
