# Sorting a Slice of Structs in Go

This Go program sorts a slice of `person` structs by **name** and **age** in both **ascending** and **descending** order using the `sort` package.

## **Algorithm Explanation**

### **Sorting by Name (`SortByName`)**
1. Implements `sort.Interface` with:
   - `Len()`: Returns the length of the slice.
   - `Less(i, j)`: Compares `name` fields lexicographically.
   - `Swap(i, j)`: Swaps elements at positions `i` and `j`.
2. Uses `sort.Sort(SortByName(xs))` for ascending order.
3. Uses `sort.Sort(sort.Reverse(SortByName(xs)))` for descending order.

### **Sorting by Age (`SortByAge`)**
1. Implements `sort.Interface` similarly but compares by `age`.
2. Uses `sort.Sort(SortByAge(xs))` for ascending order.
3. Uses `sort.Sort(sort.Reverse(SortByAge(xs)))` for descending order.

## **Final Output**
```
sort by name
[{AAB 29} {ABC 28} {AAC 27}]
reverse
[{AAC 27} {ABC 28} {AAB 29}]
sort by age
[{AAC 27} {ABC 28} {AAB 29}]
reverse
[{AAB 29} {ABC 28} {AAC 27}]
```