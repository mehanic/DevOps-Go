def binary_search_recursive(array, target):
    def helper(left, right):
        if left > right:
            return -1
        mid = (left + right) // 2
        if array[mid] == target:
            return mid
        elif array[mid] > target:
            return helper(left, mid - 1)
        else:
            return helper(mid + 1, right)
    return helper(0, len(array) - 1)

