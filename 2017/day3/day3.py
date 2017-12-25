#!/usr/bin/python
from math import ceil, floor


# Algorithm relies on square numbers. Designed for efficiency this algorithm avoids costly operations such as sqrt
# First, determine the floor of the root of the number-1. i.e. 12 -> 3, 21 -> 4, 1 -> 0
# Second, determine the numbers between i**2 and (i+1)**2 with the smallest manhattan distance,
# i.e. 1, 2, 11, 28 horizontally and 1, 4, 15, 34 vertically
# Third, find the minimum between the number and the v/h numbers and add that to the distance to the h/v number
def manhattan_distance(number):
    i = 0
    while (i+1)**2 < number:
        i += 1

    v = (i**2)+ceil(i/2)
    h = ((i+1)**2)-floor(i/2)
    return (min(abs(number-v), abs(number-h)))+ceil(i/2)


def first_largest_number(number):
    raise NotImplementedError


if __name__ == '__main__':
    user_input = 347991
    print("Manhattan Result: ", manhattan_distance(user_input))
    # print("First Largest Result: ", first_largest_number(user_input))
