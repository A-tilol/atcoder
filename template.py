import sys

sys.setrecursionlimit(10**7)


def read_int():
    return int(sys.stdin.readline())


def read_ints():
    return list(map(int, sys.stdin.readline().split(" ")))


def read_float():
    return float(sys.stdin.readline())


def read_floats():
    return list(map(float, sys.stdin.readline().split(" ")))


def get_logger(debug=True):
    if not debug:

        class c:
            def debug(self, a):
                pass

        return c()

    import logging

    FORMAT = "[%(funcName)s:%(lineno)s] %(message)s"

    logger = logging.getLogger("")
    logger.setLevel(logging.DEBUG)
    logger.propagate = False

    handler = logging.StreamHandler(sys.stdout)
    handler.setLevel(logging.DEBUG)
    handler.setFormatter(logging.Formatter(FORMAT))
    logger.addHandler(handler)

    return logger


# -------------------------------
log = get_logger(True)


"""test cases
"""
