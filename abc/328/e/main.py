import logging
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


def get_logger(debug=True) -> logging.Logger:
    FORMAT = "[%(funcName)s:%(lineno)s] %(message)s"
    log_level = logging.CRITICAL
    if debug:
        log_level = logging.DEBUG

    logger = logging.getLogger("")
    logger.setLevel(log_level)
    logger.propagate = False

    handler = logging.StreamHandler(sys.stdout)
    handler.setLevel(logging.DEBUG)
    handler.setFormatter(logging.Formatter(FORMAT))
    logger.addHandler(handler)

    return logger


# -------------------------------
log = get_logger(False)

N, M, K = read_ints()
u, v, w = [0] * M, [0] * M, [0] * M
for i in range(M):
    u[i], v[i], w[i] = read_ints()
    u[i] -= 1
    v[i] -= 1

ans = K + 1
visit = None
ad = None


def traverse(i, pre):
    if visit[i]:
        return False
    visit[i] = True

    for ni in ad[i]:
        if ni == pre:
            continue
        if not traverse(ni, i):
            return False

    return True


def ok(inds):
    global visit, ad
    visit = [False] * N
    ad = [[] for _ in range(N)]
    for i in inds:
        ad[u[i]].append(v[i])
        ad[v[i]].append(u[i])

    return traverse(0, -1) and all(visit)


def comb(i, edges: list, cost):
    global ans
    if len(edges) == N - 1:
        if cost < ans and ok(edges):
            ans = min(ans, cost)
        return

    if i >= M or len(edges) + M - i < N - 1:
        return

    comb(i + 1, edges.copy(), cost)
    edges.append(i)
    comb(i + 1, edges.copy(), (cost + w[i]) % K)


comb(0, [], 0)

print(ans)
