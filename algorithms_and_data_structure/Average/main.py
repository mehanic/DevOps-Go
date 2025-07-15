def avg(*num: float) -> float:
    print(num)
    print(type(num))
    total = sum(num)
    return total / len(num)

def main():
    av = avg(12, 23, 45, 45)
    print(av)

if __name__ == "__main__":
    main()


def avg_list(numbers: list[float]) -> float:
    print(numbers)
    print(type(numbers))
    total = sum(numbers)
    return total / len(numbers)

avg_list([12, 23, 45, 45])


import statistics

def avg_stat(*num: float) -> float:
    print(num)
    print(type(num))
    return statistics.mean(num)


