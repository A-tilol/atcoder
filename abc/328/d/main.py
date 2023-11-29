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

A, B, C = "A", "B", "C"
s = list(input())


def left_index(i):
    if i < 0 or s[i] != 0:
        return i
    li[i] = left_index(li[i])
    return li[i]


def right_index(i):
    if i >= len(s) or s[i] != 0:
        return i
    ri[i] = right_index(ri[i])
    return ri[i]


def replace(i):
    if i < 0 or i >= len(s) or s[i] != A:
        return

    ai = i
    bi = right_index(ai + 1)
    ci = right_index(bi + 1)
    # print(ai, bi, ci)
    if ci >= len(s):
        return

    if s[ai] != A or s[bi] != B or s[ci] != C:
        return

    s[ai], s[bi], s[ci] = 0, 0, 0
    # print(ai, bi, ci, s)

    l1 = left_index(ai - 1)
    r1 = right_index(ci + 1)
    li[ai], li[bi], li[ci] = l1, l1, l1
    ri[ai], ri[bi], ri[ci] = r1, r1, r1

    l2 = left_index(l1 - 1)
    replace(l2)
    replace(l1)


li = [-1] * len(s)
ri = [-1] * len(s)
for i in range(len(s)):
    replace(i)

ans = ""
for i in range(len(s)):
    if s[i] != 0:
        ans += s[i]
# print(li)
# print(ri)
# print(s)
print(ans)
