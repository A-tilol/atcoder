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

N, M = read_ints()
a = read_ints()
b = read_ints()

ad = [set() for _ in range(N)]
for i in range(M):
    ad[a[i] - 1].add(b[i] - 1)
    ad[b[i] - 1].add(a[i] - 1)


def dfs(i, t):
    log.debug((i, t, types))
    if types[i] != 0 and types[i] != t:
        return False
    if types[i] != 0:
        return True
    types[i] = t

    for ni in ad[i]:
        if not dfs(ni, t * -1):
            return False

    return True


types = [0] * N
for i in range(N):
    if types[i] != 0:
        continue

    log.debug(i)
    if not dfs(i, 1):
        print("No")
        exit()
print("Yes")
