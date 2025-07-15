def bubble_sort(input_list):
    is_sorted = False
    counter = 0
    while not is_sorted:
        is_sorted = True
        for i in range(len(input_list) - 1 - counter):
            if input_list[i] > input_list[i + 1]:
                input_list[i], input_list[i + 1] = input_list[i + 1], input_list[i]
                is_sorted = False
        counter += 1
    return input_list


def main():
    # Example 1: Unsorted array
    arr1 = [5, 3, 8, 4, 2]
    print("Sorted Array 1:", bubble_sort(arr1))

    # Example 2: Already sorted array
    arr2 = [1, 2, 3, 4, 5]
    print("Sorted Array 2:", bubble_sort(arr2))

    # Example 3: Array in reverse order
    arr3 = [9, 7, 5, 3, 1]
    print("Sorted Array 3:", bubble_sort(arr3))

    # Example 4: Array with repeated elements
    arr4 = [4, 4, 2, 1, 3, 2, 5, 1]
    print("Sorted Array 4:", bubble_sort(arr4))


if __name__ == "__main__":
    main()
