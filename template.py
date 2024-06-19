import sys
from pprint import pprint

sys.setrecursionlimit(10**7)
read_int = lambda: int(sys.stdin.readline())
read_ints = lambda: list(map(int, sys.stdin.readline().split()))
read_float = lambda: float(sys.stdin.readline())
read_floats = lambda: list(map(float, sys.stdin.readline().split()))


def get_logger(debug=True):
    if not debug:
        return type("Dummy", (object,), {"debug": lambda self, a: None})()
    import logging

    logger = logging.getLogger("")
    logger.setLevel(logging.DEBUG)
    handler = logging.StreamHandler(sys.stdout)
    handler.setFormatter(logging.Formatter("[%(funcName)s:%(lineno)s] %(message)s"))
    logger.addHandler(handler)
    return logger


# -------------------------------
log = get_logger(False)

N = read_int()
a = read_ints()

s = sorted(set(a))
d = {}
for i, e in enumerate(s):
    d[e] = i + 1

for i in range(N):
    a[i] = d[a[i]]
print(*a)

"""test cases
"""
