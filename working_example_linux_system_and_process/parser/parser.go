package otus_lesson2

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(hash string) (string, error) {
	hash = strings.TrimSpace(hash)

	if hash == "" {
		return "", nil
	}

	hashRunes := []rune(hash)
	var resultBuilder strings.Builder
	var err error
	ctx := parseContext{}

	for i := range hashRunes {
		singleRune := hashRunes[i]

		if unicode.IsDigit(singleRune)  {
			if !ctx.isRuneSet() {
				if ctx.escape {
					ctx.setRune(singleRune)
					ctx.escape = false
				} else {
					return "", fmt.Errorf("incorrect rune at index %d: expected letter or escape, actual: %d", i, singleRune)
				}
			} else {
				err = ctx.addRetryDigit(singleRune)

				if err != nil {
					return "", err
				}
			}
		} else if singleRune == '\\' {
			if ctx.isRuneSet() {
				err = appendToResult(&ctx, &resultBuilder)

				if err != nil {
					return "", err
				}

				ctx.escape = true
			} else {
				if ctx.escape {
					ctx.escape = false
					ctx.setRune(singleRune)
				} else {
					ctx.escape = true
				}
			}
		} else {
			if ctx.isRuneSet() {
				err = appendToResult(&ctx, &resultBuilder)

				if err != nil {
					return "", err
				}

				ctx.setRune(singleRune)
			} else {
				ctx.setRune(singleRune)
			}
		}
	}

	if ctx.isRuneSet() {
		err = appendToResult(&ctx, &resultBuilder)

		if err != nil {
			return "", err
		}
	}

	return resultBuilder.String(), nil
}

func appendToResult(context *parseContext, builder *strings.Builder) error {
	retryCount, err := context.retryCount()

	if err != nil {
		return err
	}

	for i := 0; i < retryCount; i++ {
		_, err = builder.WriteRune(context.r)

		if err != nil {
			return err
		}
	}

	context.reset()
	return nil
}

type parseContext struct {
	escape bool
	runeSet bool
	r rune
	retryBuilder strings.Builder
}

func (ctx *parseContext) reset() {
	ctx.escape = false
	ctx.runeSet = false
	ctx.r = 0
	ctx.retryBuilder.Reset()
}

func (ctx *parseContext) setRune(r rune) {
	ctx.r = r
	ctx.runeSet = true
}

func (ctx *parseContext) isRuneSet() bool {
	return ctx.runeSet
}

func (ctx *parseContext) addRetryDigit(digit rune) error {
	_, err := ctx.retryBuilder.WriteRune(digit)
	return err
}

func (ctx *parseContext) retryCount() (int, error) {
	if ctx.retryBuilder.Len() == 0 {
		return 1, nil
	} else {
		return strconv.Atoi(ctx.retryBuilder.String())
	}
}