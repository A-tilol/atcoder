class Item:
    def __init__(self, _id, value):
        self.id = _id
        self.v = value


class PriorityQueue:
    def __init__(self, Descending=True):
        self.id_to_ind = {}
        self.q = []
        self.desc = Descending
        self.op = max if self.desc else min

    def __len__(self):
        return len(self.q)

    def push(self, item: Item):
        self.update(item)

    def update(self, item: Item):
        """Update the existing item. If Id is not found, push item."""
        ind = None
        if item.id in self.id_to_ind:
            ind = self.id_to_ind[item.id]
            self.q[ind].v = self.op(self.q[ind].v, item.v)
        else:
            self.q.append(item)
            ind = len(self.q) - 1
            self.id_to_ind[item.id] = ind

        while ind != 0:
            p_ind = (ind - 1) // 2
            # !!! Sort Conditions !!!
            if (self.desc and self.q[p_ind].v >= self.q[ind].v) or (
                not self.desc and self.q[p_ind].v <= self.q[ind].v
            ):
                break

            self.q[p_ind], self.q[ind] = self.q[ind], self.q[p_ind]
            self.id_to_ind[self.q[ind].id] = ind
            self.id_to_ind[self.q[p_ind].id] = p_ind

            ind = p_ind

    def pop(self) -> Item:
        if len(self.q) == 0:
            return None

        item = self.q[0]
        del self.id_to_ind[item.id]

        self.q[0] = self.q[len(self.q) - 1]
        self.id_to_ind[self.q[0].id] = 0
        self.q.pop()

        if len(self.q) == 0:
            self.id_to_ind = {}
            return item

        ind = 0
        while ind * 2 + 1 < len(self.q):
            c_ind = 2 * ind + 1
            c_ind2 = 2 * ind + 2
            if c_ind2 < len(self.q):
                # !!! Sort Conditions !!!
                if self.desc:
                    _, c_ind = max((self.q[c_ind].v, c_ind), (self.q[c_ind2].v, c_ind2))
                else:
                    _, c_ind = min((self.q[c_ind].v, c_ind), (self.q[c_ind2].v, c_ind2))

            # !!! Sort Conditions !!!
            if (self.desc and self.q[c_ind].v <= self.q[ind].v) or (
                not self.desc and self.q[c_ind].v >= self.q[ind].v
            ):
                break

            self.q[c_ind], self.q[ind] = self.q[ind], self.q[c_ind]
            self.id_to_ind[self.q[ind].id] = ind
            self.id_to_ind[self.q[c_ind].id] = c_ind

            ind = c_ind

        return item

    def remove(self, _id):
        pass

    def contains(self, _id):
        return _id in self.id_to_ind

    def get(self, _id) -> Item:
        if _id not in self.id_to_ind:
            return None
        ind = self.id_to_ind[_id]
        return self.q[ind]

    def top(self) -> Item:
        return self.q[0]
