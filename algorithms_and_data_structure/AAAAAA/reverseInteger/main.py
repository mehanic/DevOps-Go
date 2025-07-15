# class Solution:
#     def reverse(self, x: int) -> int:
#         INT_MIN, INT_MAX = -(2**31), 2**31 - 1
#         sign = -1 if x < 0 else 1
#         x_abs = abs(x)
#         reversed_num = 0
#         while x_abs != 0:
#             digit = x_abs % 10
#             x_abs //= 10
#             if reversed_num > INT_MAX // 10 or (
#                 (reversed_num == INT_MAX // 10 and digit > INT_MAX % 10)
#             ):
#                 return 0

#             reversed_num = reversed_num * 10 + digit

#         return sign * reversed_num
    

def reverse(x):
    INT_MIN, INT_MAX = -(2**31), 2**31 - 1
    sign = -1 if x < 0 else 1
    x_abs = abs(x)
    reversed_num = 0

    while x_abs != 0:
        digit = x_abs % 10
        x_abs //= 10
        if reversed_num > INT_MAX // 10 or (
            reversed_num == INT_MAX // 10 and digit > INT_MAX % 10
        ):
            return 0
        reversed_num = reversed_num * 10 + digit
    return sign * reversed_num


print(reverse(123))  # Output: 321
print(reverse(-123))  # Output: -321
print(reverse(120))  # Output: 21
print(reverse(0))  # Output: 0