def binary_search(array, target):
    return helper(array, target, 0, len(array) - 1)

def helper(array, target, left, right):
    if left > right:
        return -1

    middle = (left + right) // 2
    potential_match = array[middle]

    if target == potential_match:
        return middle
    elif target < potential_match:
        return helper(array, target, left, middle - 1)
    else:
        return helper(array, target, middle + 1, right)

def main():
    array = [1, 3, 5, 7, 9, 11, 13, 15, 17]
    targets = [7, 15, 2, 11]
    for target in targets:
        result = binary_search(array, target)
        if result != -1:
            print(f"Target {target} found at index {result}")
        else:
            print(f"Target {target} not found in the array")

# Optional simple variant
def main_simple():
    array = [1, 3, 5, 7, 9, 11, 13, 15, 17]
    targets = [7, 15, 2, 11]
    for target in targets:
        print(binary_search(array, target))

if __name__ == "__main__":
    main()
    # main_simple()  # Uncomment to run the simpler output version
