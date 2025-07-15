

import bisect

def binary_search_bisect(array, target):
    index = bisect.bisect_left(array, target)
    if index != len(array) and array[index] == target:
        return index
    return -1


def binary_search_tail_recursive(array, target, left=0, right=None):
    if right is None:
        right = len(array) - 1
    if left > right:
        return -1
    mid = (left + right) // 2
    if array[mid] == target:
        return mid
    if array[mid] > target:
        return binary_search_tail_recursive(array, target, left, mid - 1)
    return binary_search_tail_recursive(array, target, mid + 1, right)
