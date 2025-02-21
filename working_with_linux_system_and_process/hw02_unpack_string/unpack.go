package hw02unpackstring

import (
    "errors"
    "strconv"
    "strings"
    "unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
    if input == "" {
        return "", nil
    }

    var result strings.Builder
    var prev rune
    escaped := false

    for _, r := range input {
        switch {
        case unicode.IsDigit(r):
            if prev == 0 || unicode.IsDigit(prev) && !escaped {
                return "", ErrInvalidString
            }
            if !escaped {
                count, _ := strconv.Atoi(string(r))
                result.WriteString(strings.Repeat(string(prev), count-1))
            } else {
                result.WriteRune(r)
                escaped = false
            }
        case r == '\\':
            if escaped {
                result.WriteRune(r)
                escaped = false
            } else {
                escaped = true
            }
        default:
            if prev != 0 && !unicode.IsDigit(prev) {
                result.WriteRune(prev)
            }
            if escaped && unicode.IsDigit(r) {
                return "", ErrInvalidString
            }
            prev = r
            escaped = false
        }
    }

    if prev != 0 {
        result.WriteRune(prev)
    }

    return result.String(), nil
}