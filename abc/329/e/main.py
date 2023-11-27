import sys


def read_int():
    return int(input())


def read_ints():
    params = input().split(" ")
    return [int(p) for p in params]


# -------------------------------
sys.setrecursionlimit(10**7)

N, M = read_ints()
s = list(input())
t = list(input())

# visit = [False] * N
# q = list(range(N))
# while len(q) > 0:
#     i = q[0]
#     q = q[1:]
#     # print(len(q))

#     if visit[i]:
#         continue

#     ok = True
#     for j in range(M):
#         if i + j >= N or s[i + j] != t[j] and s[i + j] != "#":
#             ok = False
#             break
#     if not ok:
#         continue

#     for j in range(M):
#         s[i + j] = "#"
#     visit[i] = True

#     for j in range(M - 1):
#         ni = i - 1 - j
#         if ni >= 0 and not visit[ni]:
#             q.append(ni)
#         ni = i + 1 + j
#         if ni < N and not visit[ni]:
#             q.append(ni)


def sharpify(i):
    if visit[i]:
        return

    for j in range(M):
        if i + j >= N or s[i + j] != t[j] and s[i + j] != "#":
            return

    for j in range(M):
        s[i + j] = "#"
    visit[i] = True

    for j in range(M - 1):
        if i - 1 - j >= 0:
            sharpify(i - 1 - j)
        if i + 1 + j < N:
            sharpify(i + 1 + j)


visit = [False] * N
for i in range(N):
    sharpify(i)

for i, ss in enumerate(s):
    if ss != "#":
        print("No")
        exit()
print("Yes")

""" test cases
7 3
ABCBABC
ABC

Yes

7 3
ABBCABC
ABC

No

12 2
XYXXYXXYYYXY
XY

Yes

"""
