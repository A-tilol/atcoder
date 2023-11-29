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

N, X = read_ints()
s = read_ints()

ans = 0
for i in range(N):
    if s[i] <= X:
        ans += s[i]
print(ans)
