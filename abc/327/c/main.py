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

a = [None] * 9
for i in range(9):
    a[i] = read_ints()

for i in range(9):
    if len(set(a[i])) != 9:
        print("No")
        exit()

for i in range(9):
    s = set()
    for j in range(9):
        s.add(a[j][i])
    if len(s) != 9:
        print("No")
        exit()

for i in range(0, 9, 3):
    for j in range(0, 9, 3):
        s = set()
        for k in range(3):
            s.update(a[i + k][j : j + 3])
        if len(s) != 9:
            print("No")
            exit()
print("Yes")
