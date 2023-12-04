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
a = sorted(a)

ans = 1
cnt = 1
si = 0
for i in range(1, N):
    if a[i] < a[si] + M:
        cnt += 1
        ans = max(ans, cnt)
    else:
        while a[si] + M <= a[i]:
            si += 1
            cnt -= 1
        cnt += 1
    log.debug((si, i, cnt))
print(ans)

"""
5 4
1 5 6 7 8

4
"""
