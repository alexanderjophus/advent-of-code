#!/usr/bin/python
from math import ceil, floor, sqrt


def manhattan_distance(number):
    i = 0
    while (i+1)**2 < number:
        i += 1

    v = (i**2)+ceil(i/2)
    h = ((i+1)**2)-floor(i/2)
    return (min(abs(number-v), abs(number-h)))+ceil(i/2)


if __name__ == '__main__':
    print("Result: ", manhattan_distance(347991))
