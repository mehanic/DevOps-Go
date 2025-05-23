package example

import (
	"github.com/pkg/errors"
	"github.com/quasilyte/go-ruleguard/dsl"
)

// $ ruleguard -rules gorules/rules.go example.go

func foo() error {
	if err := bar(); err != nil {
		return err
	}
	return nil
}

func bar() error {
	return errors.Wrap(errors.New("bad request"), "do request")
}

// wrapWithPkgErrors - правило, которое предлагает врапить ошибки через github.com/pkg/errors.
func wrapWithPkgErrors(m dsl.Matcher) {
	m.Import("github.com/pkg/errors")

	m.Match(`if err := $x; err != nil { return err }`).
		Report(`err is not wrapped`).
		Suggest(`if err := $x; err != nil { return errors.Wrap($x, "FIXME: wrap the error") }`)

	// Прочие варианты.
	// ...
}
