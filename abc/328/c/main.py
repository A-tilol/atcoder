import sys

sys.setrecursionlimit(10**7)


def read_int():
    return int(input())


def read_ints():
    return list(map(int, input().split(" ")))


def read_float():
    return float(input())


def read_floats():
    return list(map(float, input().split(" ")))


# -------------------------------

n, q = read_ints()
s = input()

cum = [0] * n
for i in range(1, n):
    cum[i] = cum[i - 1]
    if s[i - 1] == s[i]:
        cum[i] += 1

for i in range(q):
    l, r = read_ints()
    l, r = l - 1, r - 1
    print(cum[r] - cum[l])
