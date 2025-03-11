# **Working with Arrays in Go**

## **Algorithm Explanation**

1. **Array Initialization**
   ```go
   var o [5]int = [5]int{1, 2, 3, 4, 5}
   ```
   - An array `o` of type `int` with 5 elements is created and initialized with the values `{1, 2, 3, 4, 5}`.

2. **Printing the Array**
   ```go
   fmt.Println("o=", o)
   ```
   - The array `o` is printed to the console, showing its contents.

## **Example Output**
```
Array:
o= [1 2 3 4 5]
```

---

## **Key Points**
- Arrays in Go are fixed in size and hold elements of the same type.
- Initialization is done at the point of declaration.