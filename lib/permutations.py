def perm(rest: list, p: list = []) -> tuple:
    """再起による順列ジェネレータ

    長さ10の順列の生成に1秒ほどかかるが、長さ9以下では十分高速
    当然itertools.permutationsの方が高速

    Args:
        rest (list): 順列の元となるリスト
        p (list): 生成中の順列を格納

    Yields:
        tuple: 順列

    Example:
        perm([0,1,2]) → 012 021 102 120 201 210
    """

    if len(rest) == 0:
        yield tuple(p)
        return

    for i in range(len(rest)):
        p.append(rest[i])
        rest.pop(i)
        yield from perm(rest, p)
        rest.insert(i, p.pop())
