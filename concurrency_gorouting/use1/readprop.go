package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
)

type UpperString string

func (u *UpperString) UnmarshalProp(b []byte) error {
	*u = UpperString(strings.ToUpper(string(b)))
	return nil
}

func main() {
	r := strings.NewReader(
		"\n# comment, ignore\nkey1: 10.5\nkey2: some string" +
			"\nkey3: 42\nkey4: false\nspecial: another string\n")
	var v struct {
		Key1 float32
		Key2 string
		Key3 uint64
		Key4 bool
		Key5 UpperString `prop:"special"`
		key6 int
	}
	if err := NewDecoder(r).Decode(&v); err != nil {
		log.Fatal(r)
	}
	log.Printf("%+v", v)
}

func Unmarshal(data []byte, v interface{}) error {
	return NewDecoder(bytes.NewReader(data)).Decode(v)
}

type Decoder struct {
	scanner *bufio.Scanner
}

type decodeError struct {
	line int
	err  error
}

var errNoSep = errors.New("no separator")

func (d decodeError) Error() string {
	return fmt.Sprintf("line %d: %s", d.line, d.err)
}

var cache = make(map[reflect.Type]map[string]int)

func findIndex(t reflect.Type, k string) (int, bool) {
	if v, ok := cache[t]; ok {
		n, ok := v[k]
		return n, ok
	}
	m := make(map[string]int)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if s := f.Name[:1]; strings.ToLower(s) == s {
			continue
		}
		name := strings.ToLower(f.Name)
		if tag := f.Tag.Get("prop"); tag != "" {
			name = tag
		}
		m[name] = i
	}
	cache[t] = m
	return findIndex(t, k)
}

func (d *Decoder) Decode(v interface{}) error {
	val := reflect.ValueOf(v)
	t := val.Type()
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("%v not a struct pointer", t)
	}
	val = val.Elem()
	t = t.Elem()
	line := 0
	for d.scanner.Scan() {
		line++
		b := d.scanner.Bytes()
		if len(b) == 0 || b[0] == '#' {
			continue
		}
		parts := bytes.SplitN(b, []byte{':'}, 2)
		if len(parts) != 2 {
			return decodeError{line: line, err: errNoSep}
		}
		index, ok := findIndex(t, string(parts[0]))
		if !ok {
			continue
		}
		value := bytes.TrimSpace(parts[1])
		if err := d.decodeValue(val.Field(index), value); err != nil {
			return decodeError{line: line, err: err}
		}
	}
	return d.scanner.Err()
}

func (d *Decoder) decodeValue(v reflect.Value, value []byte) error {
	if v, ok := v.Addr().Interface().(Unmarshaller); ok {
		return v.UnmarshalProp(value)
	}
	switch valStr := string(value); v.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(valStr, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(valStr, 10, 64)
		if err != nil {
			return err
		}
		v.SetUint(i)
	case reflect.Float32, reflect.Float64:
		i, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return err
		}
		v.SetFloat(i)
	case reflect.String:
		v.SetString(valStr)
	case reflect.Bool:
		switch value := valStr; value {
		case "true":
			v.SetBool(true)
		case "false":
			v.SetBool(false)
		default:
			return fmt.Errorf("invalid bool: %s", value)
		}
	default:
		return fmt.Errorf("invalid type: %s", v.Type())
	}
	return nil
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{scanner: bufio.NewScanner(r)}
}

type Unmarshaller interface {
	UnmarshalProp([]byte) error
}



//2025/01/28 21:05:27 {Key1:10.5 Key2:some string Key3:42 Key4:false Key5:ANOTHER STRING key6:0}