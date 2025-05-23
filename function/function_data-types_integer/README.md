This Go code contains a series of functions that perform basic arithmetic operations and number manipulations on `int64` values. I'll explain each function and its purpose step by step:

### Integer1
```go
func Integer1(l int64) (n int64) {
    n = l / 100
    return
}
```
- **Purpose**: Divides the input `l` by 100 and returns the result. Essentially, it returns the number of hundreds in the given number.

### Integer2
```go
func Integer2(kg int64) (tn int64) {
    tn = kg / 1000
    return
}
```
- **Purpose**: Converts kilograms to tons by dividing by 1000.

### Integer3
```go
func Integer3(bait int64) (kbait int64) {
    kbait = bait / 1024
    return
}
```
- **Purpose**: Converts bytes to kilobytes by dividing by 1024.

### Integer4
```go
func Integer4(a, b int64) (z int64) {
    z = a / b
    return
}
```
- **Purpose**: Divides `a` by `b` and returns the result.

### Integer5
```go
func Integer5(a, b int64) (z int64) {
    z = a % b
    return
}
```
- **Purpose**: Returns the remainder when `a` is divided by `b`.

### Integer6
```go
func Integer6(a int64) (d, e int64) {
    d = a / 10
    e = a % 10
    return
}
```
- **Purpose**: Returns the quotient and the remainder when `a` is divided by 10. Essentially, it separates the last digit (`e`) and the rest of the number (`d`).

### Integer7
```go
func Integer7(a int64) (s, p int64) {
    d := a / 10
    e := a % 10
    s = d + e
    p = d * e
    return
}
```
- **Purpose**: Returns the sum (`s`) and the product (`p`) of the quotient (`d`) and remainder (`e`) when dividing `a` by 10.

### Integer8
```go
func Integer8(a int64) (a1 int64) {
    d := a / 10
    e := a % 10
    a1 = e*10 + d
    return
}
```
- **Purpose**: Swaps the last digit of `a` and the rest of the number. It essentially reverses the last two digits.

### Integer9
```go
func Integer9(a int64) (s int64) {
    s = a / 100
    return
}
```
- **Purpose**: Divides `a` by 100 and returns the result, essentially removing the last two digits of `a`.

### Integer10
```go
func Integer10(a int64) (e, d int64) {
    e = a % 10
    d = (a / 10) % 10
    return
}
```
- **Purpose**: Returns the last two digits of `a` as `e` (ones digit) and `d` (tens digit).

### Integer11
```go
func Integer11(a int64) (s, p int64) {
    e := a % 10
    d := (a / 10) % 10
    c := a / 100
    s = e + d + c
    p = e * d * c
    return
}
```
- **Purpose**: Calculates the sum (`s`) and product (`p`) of the last three digits of `a`.

### Integer12
```go
func Integer12(a int64) (a1 int64) {
    e := a % 10
    d := (a / 10) % 10
    c := a / 100
    a1 = e*100 + d*10 + c
    return
}
```
- **Purpose**: Reverses the first three digits of `a`. It takes the last digit (`e`), second last (`d`), and the rest (`c`), and rearranges them to form a new number.

### Integer13
```go
func Integer13(a int64) (a1 int64) {
    e := a % 10
    d := (a / 10) % 10
    c := a / 100
    a1 = d*100 + e*10 + c
    return
}
```
- **Purpose**: Rearranges the first three digits of `a` in a specific order, swapping the tens and ones places.

### Integer14
```go
func Integer14(a int64) (a1 int64) {
    e := a % 10
    d := (a / 10) % 10
    c := a / 100
    a1 = e*100 + c*10 + d
    return
}
```
- **Purpose**: Rearranges the first three digits of `a` in a specific order, swapping the ones, hundreds, and tens places.

### Integer15
```go
func Integer15(a int64) (a1 int64) {
    e := a % 10
    d := (a / 10) % 10
    c := a / 100
    a1 = d*100 + c*10 + e
    return
}
```
- **Purpose**: Rearranges the first three digits of `a` in a specific order, swapping the tens, hundreds, and ones places.

### Integer16
```go
func Integer16(a int64) (a1 int64) {
    e := a % 10
    d := (a / 10) % 10
    c := a / 100
    a1 = c*100 + e*10 + d
    return
}
```
- **Purpose**: Rearranges the first three digits of `a` in a specific order, swapping the hundreds, ones, and tens places.

### Integer17
```go
func Integer17(a int64) (c int64) {
    c = a / 100 % 10
    return
}
```
- **Purpose**: Extracts the hundreds digit of `a`.

### Integer18
```go
func Integer18(a int64) (ts int64) {
    ts = a / 1000 % 10
    return
}
```
- **Purpose**: Extracts the thousands digit of `a`.

### Integer19
```go
func Integer19(s int64) (m int64) {
    m = s / 60
    return
}
```
- **Purpose**: Converts seconds (`s`) into minutes (`m`).

### Integer20
```go
func Integer20(s int64) (h int64) {
    h = s / 3600
    return
}
```
- **Purpose**: Converts seconds (`s`) into hours (`h`).

### Integer21
```go
func Integer21(s int64) (s1 int64) {
    s1 = s % 60
    return
}
```
- **Purpose**: Returns the remaining seconds (`s1`) after converting `s` to minutes.

### Integer22
```go
func Integer22(s int64) (s1 int64) {
    s1 = s % 3600
    return
}
```
- **Purpose**: Returns the remaining seconds (`s1`) after converting `s` to hours.

### Integer23
```go
func Integer23(s int64) (m int64) {
    m = s / 60 % 60
    return
}
```
- **Purpose**: Converts seconds (`s`) to minutes, but only returns the minutes part (not the total minutes).

### Integer24
```go
func Integer24(d int64) (vn int64) {
    vn = d % 7
    return
}
```
- **Purpose**: Finds the remainder when `d` is divided by 7 (essentially finding the day of the week, assuming `d` is the day number).

### Integer25
```go
func Integer25(d int64) (vn int64) {
    vn = d%7 + 4
    return
}
```
- **Purpose**: Modifies the result of `d % 7` by adding 4.

### Integer26
```go
func Integer26(d int64) (vn int64) {
    vn = d%7 + 1
    return
}
```
- **Purpose**: Modifies the result of `d % 7` by adding 1.

### Integer27
```go
func Integer27(d int64) (vn int64) {
    vn = d%7 + 5
    return
}
```
- **Purpose**: Modifies the result of `d % 7` by adding 5.

### Integer28
```go
func Integer28(d, vn int64) (nn int64) {
    nn = ((d + vn - 2) % 7) + 1
    return
}
```
- **Purpose**: A more complex operation based on the sum of `d` and `vn` and the modulo operation.

### Integer29
```go
func Integer29(a, b, c int64) (n, so int64) {
    sp := a * b
    sd := c * c
    n = sp / sd
    so = sp % sd
    return
}
```
- **Purpose**: Performs a calculation based on the multiplication and division of `a`, `b`, and `c` and returns two values: the quotient (`n`) and the remainder (`so`).

### Integer30
```go
func Integer30(y int64) (c int64) {
    c = ((y - 1) / 100) + 1
    return
}
```
- **Purpose**: Performs a century calculation for a given year `y`, returning the century number.

---

### Conclusion:
Each of these functions operates on `int64` numbers and performs a variety of mathematical operations, such as division, modulo, and digit manipulation, to return specific results based on the input values. Many of these functions deal with extracting specific digits or performing conversions, like converting between units of time (seconds to minutes, etc.) or units of size (bytes to kilobytes).