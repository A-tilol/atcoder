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

b = read_int()
for i in range(1, 18):
    if i**i == b:
        print(i)
        exit()
print(-1)
