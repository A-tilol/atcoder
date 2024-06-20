def modpow(a: int, n: int, mod: int) -> int:
    """二分累乗法"""

    res = 1
    while n > 0:
        if n & 1:
            res = res * a % mod
        a = a * a % mod
        n >>= 1
    return res


def modinv(a: int, mod: int) -> int:
    """逆元の計算

    オイラーの小定理より
    法modにおけるaの逆元a^-1 = a^(mod-2) % mod
    """

    return modpow(a, mod - 2, mod)


def factorial(n: int, mod: int) -> int:
    ret = 1
    for i in range(2, n + 1):
        ret = ret * i % mod
    return ret
