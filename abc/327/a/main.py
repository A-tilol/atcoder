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
log = get_logger(True)

_ = input()
s = input()

for i in range(len(s) - 1):
    if (s[i] == "a" and s[i + 1] == "b") or (s[i] == "b" and s[i + 1] == "a"):
        print("Yes")
        exit()
print("No")
