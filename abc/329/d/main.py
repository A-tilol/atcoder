def read_int():
    return int(input())


def read_ints():
    params = input().split(" ")
    return [int(p) for p in params]


class Item:
    def __init__(self, _id, value):
        self.id = _id
        self.v = value


class PriorityQueue:
    def __init__(self):
        self.id_to_ind = {}
        self.q = []

    def push(self, item: Item):
        self.update(item)

    def update(self, item: Item):
        ind = None
        if item.id in self.id_to_ind:
            ind = self.id_to_ind[item.id]
            self.q[ind] = item
        else:
            self.q.append(item)
            ind = len(self.q) - 1
            self.id_to_ind[item.id] = ind

        while ind != 0:
            p_ind = (ind - 1) // 2
            # !!! Sort Conditions !!!
            if self.q[p_ind].v > self.q[ind].v or (
                self.q[p_ind].v == self.q[ind].v and self.q[p_ind].id < self.q[ind].id
            ):
                break

            self.q[p_ind], self.q[ind] = self.q[ind], self.q[p_ind]
            self.id_to_ind[self.q[ind].id] = ind
            self.id_to_ind[self.q[p_ind].id] = p_ind

            ind = p_ind

    # def pop(self):
    #     item = self.q[0]
    #     del self.id_to_ind[item.id]

    #     self.q[0] = self.q[len(self.q) - 1]
    #     self.id_to_ind[self.q[0].id] = 0
    #     self.q.pop()

    #     ind = 0
    #     while ind >= len(self.q):
    #         c_ind1 = 2 * ind + 1
    #         c_ind2 = 2 * ind + 2
    #         c_ind = None
    #         # !!! Sort Conditions !!!
    #         if self.q[c_ind1].v > self.q[c_ind2]:
    #             c_ind = c_ind1
    #         else:
    #             c_ind = c_ind2

    #         # !!! Sort Conditions !!!
    #         if self.q[c_ind].v <= self.q[ind].v:
    #             break

    #         self.q[c_ind], self.q[ind] = self.q[ind], self.q[c_ind]
    #         self.id_to_ind[self.q[ind].id] = ind
    #         self.id_to_ind[self.q[c_ind].id] = c_ind

    #         ind = c_ind

    #     self.q = self.q[1:]
    #     return item

    def remove(self, _id):
        pass

    def get(self, _id) -> Item:
        if _id not in self.id_to_ind:
            return None
        ind = self.id_to_ind[_id]
        return self.q[ind]

    def get_top(self) -> Item:
        return self.q[0]


# -------------------------------

N, M = read_ints()
a = read_ints()


pq = PriorityQueue()
for i in range(M):
    item = pq.get(a[i])
    if item is None:
        pq.update(Item(a[i], 1))
    else:
        pq.update(Item(a[i], item.v + 1))

    print(pq.get_top().id)

# for i, item in enumerate(pq.q):
#     print(i, item.id, item.v)

""" test cases
3 7
1 2 2 3 1 3 3

1
1
2
2
1
1
3

100 5
100 90 80 70 60

100
90
80
70
60

9 8
8 8 2 2 8 8 2 2

8
8
8
2
8
8
8
2
"""
