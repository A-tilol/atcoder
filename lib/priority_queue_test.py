import unittest

from priority_queue import Item, PriorityQueue


class TestPriorityQueue(unittest.TestCase):
    def test_push_and_top(self):
        pq = PriorityQueue()
        item1 = Item(1, 10)
        item2 = Item(2, 5)
        item3 = Item(3, 8)
        item4 = Item(4, 15)

        pq.push(item1)
        self.assertEqual(pq.top(), item1)

        pq.push(item2)
        self.assertEqual(pq.top(), item1)

        pq.push(item3)
        self.assertEqual(pq.top(), item1)

        pq.push(item4)
        self.assertEqual(pq.top(), item4)

    def test_push_and_top_acs(self):
        pq = PriorityQueue(Descending=False)
        item1 = Item(1, 5)
        item2 = Item(2, 10)
        item3 = Item(3, 15)
        item4 = Item(4, 1)

        pq.push(item1)
        self.assertEqual(pq.top(), item1)

        pq.push(item2)
        self.assertEqual(pq.top(), item1)

        pq.push(item3)
        self.assertEqual(pq.top(), item1)

        pq.push(item4)
        self.assertEqual(pq.top(), item4)

    def test_pop(self):
        pq = PriorityQueue()
        item1 = Item(1, 1)
        item2 = Item(2, 3)
        item3 = Item(3, 8)
        item4 = Item(4, 4)
        item5 = Item(5, 0)
        item6 = Item(6, 6)
        item7 = Item(7, 5)
        item8 = Item(8, 9)
        item9 = Item(9, 2)
        item10 = Item(10, 7)

        pq.push(item1)
        pq.push(item2)
        pq.push(item3)
        pq.push(item4)
        pq.push(item5)
        pq.push(item6)
        pq.push(item7)
        pq.push(item8)
        pq.push(item9)
        pq.push(item10)

        self.assertEqual(pq.pop(), item8)
        self.assertEqual(pq.pop(), item3)
        self.assertEqual(pq.pop(), item10)
        self.assertEqual(pq.pop(), item6)
        self.assertEqual(pq.pop(), item7)
        self.assertEqual(pq.pop(), item4)
        self.assertEqual(pq.pop(), item2)
        self.assertEqual(pq.pop(), item9)
        self.assertEqual(pq.pop(), item1)
        self.assertEqual(pq.pop(), item5)

    def test_pop_asc(self):
        pq = PriorityQueue(Descending=False)
        item1 = Item(1, 1)
        item2 = Item(2, 3)
        item3 = Item(3, 8)
        item4 = Item(4, 4)
        item5 = Item(5, 0)
        item6 = Item(6, 6)
        item7 = Item(7, 5)
        item8 = Item(8, 9)
        item9 = Item(9, 2)
        item10 = Item(10, 7)

        pq.push(item1)
        pq.push(item2)
        pq.push(item3)
        pq.push(item4)
        pq.push(item5)
        pq.push(item6)
        pq.push(item7)
        pq.push(item8)
        pq.push(item9)
        pq.push(item10)

        self.assertEqual(pq.pop(), item5)
        self.assertEqual(pq.pop(), item1)
        self.assertEqual(pq.pop(), item9)
        self.assertEqual(pq.pop(), item2)
        self.assertEqual(pq.pop(), item4)
        self.assertEqual(pq.pop(), item7)
        self.assertEqual(pq.pop(), item6)
        self.assertEqual(pq.pop(), item10)
        self.assertEqual(pq.pop(), item3)
        self.assertEqual(pq.pop(), item8)

    def test_remove(self):
        pass

    def test_update(self):
        pq = PriorityQueue()
        item1 = Item(1, 10)
        item2 = Item(2, 5)

        pq.push(item1)
        pq.push(item2)
        self.assertEqual(pq.top(), item1)

        item2_updated = Item(2, 12)
        pq.update(item2_updated)
        self.assertEqual(pq.top(), item2_updated)

    def test_update_asc(self):
        pq = PriorityQueue(Descending=False)
        item1 = Item(1, 5)
        item2 = Item(2, 10)

        pq.push(item1)
        pq.push(item2)
        self.assertEqual(pq.top(), item1)

        item2_updated = Item(2, 4)
        pq.update(item2_updated)
        self.assertEqual(pq.top(), item2_updated)


if __name__ == "__main__":
    unittest.main()
