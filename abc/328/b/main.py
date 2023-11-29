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

n = read_int()
d = read_ints()

ans = 0
for i in range(1, n + 1):
    for dd in range(1, d[i - 1] + 1):
        m = str(i)
        ok_m = True
        for k in range(1, len(m)):
            if m[k - 1] != m[k]:
                ok_m = False
                break

        ds = str(dd)
        ok_d = True
        for k in range(1, len(ds)):
            if ds[k - 1] != ds[k]:
                ok_d = False
                break

        if m[0] == ds[0] and ok_m and ok_d:
            ans += 1
print(ans)
