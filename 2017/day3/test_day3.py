#!/usr/bin/python

import unittest
from day3 import manhattan_distance


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