#!/usr/bin/python

import unittest
from day3 import manhattan_distance, first_largest_number


class TestManhattanDistance(unittest.TestCase):

    def test_manhattan_distance_1(self):
        self.assertEqual(0, manhattan_distance(1))

    def test_manhattan_distance_12(self):
        self.assertEqual(3, manhattan_distance(12))

    def test_manhattan_distance_23(self):
        self.assertEqual(2, manhattan_distance(23))

    def test_manhattan_distance_1024(self):
        self.assertEqual(31, manhattan_distance(1024))

    # low numbers and square numbers (two tests?)
    def test_manhattan_distance_edge(self):
        self.assertEqual(0, manhattan_distance(1))
        self.assertEqual(1, manhattan_distance(2))
        self.assertEqual(2, manhattan_distance(3))
        self.assertEqual(1, manhattan_distance(4))
        self.assertEqual(2, manhattan_distance(9))
        self.assertEqual(3, manhattan_distance(16))


class TestFirstLargestNumber(unittest.TestCase):

    @unittest.skip("not implemented yet")
    def test_first_largest_number_1(self):
        self.assertEqual(2, first_largest_number(1))

    @unittest.skip("not implemented yet")
    def test_first_largest_number_2(self):
        self.assertEqual(4, first_largest_number(2))

    @unittest.skip("not implemented yet")
    def test_first_largest_number_3(self):
        self.assertEqual(4, first_largest_number(3))

    @unittest.skip("not implemented yet")
    def test_first_largest_number_4(self):
        self.assertEqual(5, first_largest_number(4))

    @unittest.skip("not implemented yet")
    def test_first_largest_number_large(self):
        self.assertEqual(122, first_largest_number(100))
        self.assertEqual(304, first_largest_number(300))
        self.assertEqual(747, first_largest_number(500))
