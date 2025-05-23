package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
)


func main() {
	if err := MarshalAll(); err != nil {
		panic(err)
	}

	if err := UnmarshalAll(); err != nil {
		panic(err)
	}

	if err := OtherJSONExamples(); err != nil {
		panic(err)
	}
}


func MarshalAll() error {
	t := TOMLData{
		Name: "Name1",
		Age:  20,
	}

	j := JSONData{
		Name: "Name2",
		Age:  30,
	}

	y := YAMLData{
		Name: "Name3",
		Age:  40,
	}

	tomlRes, err := t.ToTOML()
	if err != nil {
		return err
	}

	fmt.Println("TOML Marshal =", tomlRes.String())

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}

	fmt.Println("JSON Marshal=", jsonRes.String())

	yamlRes, err := y.ToYAML()
	if err != nil {
		return err
	}

	fmt.Println("YAML Marshal =", yamlRes.String())
	return nil
}

//---

const (
	exampleTOML = `name="Example1"
age=99
    `

	exampleJSON = `{"name":"Example2","age":98}`

	exampleYAML = `name: Example3
age: 97    
    `
)

// UnmarshalAll takes data in various formats
// and converts them into structs
func UnmarshalAll() error {
	t := TOMLData{}
	j := JSONData{}
	y := YAMLData{}

	if _, err := t.Decode([]byte(exampleTOML)); err != nil {
		return err
	}
	fmt.Println("TOML Unmarshal =", t)

	if err := j.Decode([]byte(exampleJSON)); err != nil {
		return err
	}
	fmt.Println("JSON Unmarshal =", j)

	if err := y.Decode([]byte(exampleYAML)); err != nil {
		return err
	}
	fmt.Println("Yaml Unmarshal =", y)
	return nil
}

//--

// TOMLData is our common data struct
// with TOML struct tags
type TOMLData struct {
	Name string `toml:"name"`
	Age  int    `toml:"age"`
}

// ToTOML dumps the TOMLData struct to
// a TOML format bytes.Buffer
func (t *TOMLData) ToTOML() (*bytes.Buffer, error) {
	b := &bytes.Buffer{}
	encoder := toml.NewEncoder(b)

	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return b, nil
}

// Decode will decode into TOMLData
func (t *TOMLData) Decode(data []byte) (toml.MetaData, error) {
	return toml.Decode(string(data), t)
}

//----

// JSONData is our common data struct
// with JSON struct tags
type JSONData struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ToJSON dumps the JSONData struct to
// a JSON format bytes.Buffer
func (t *JSONData) ToJSON() (*bytes.Buffer, error) {
	d, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode will decode into JSONData
func (t *JSONData) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}

// OtherJSONExamples shows ways to use types
// beyond structs and other useful functions
func OtherJSONExamples() error {
	res := make(map[string]string)
	err := json.Unmarshal([]byte(`{"key": "value"}`), &res)
	if err != nil {
		return err
	}

	fmt.Println("We can unmarshal into a map instead of a struct:", res)

	b := bytes.NewReader([]byte(`{"key2": "value2"}`))
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(&res); err != nil {
		return err
	}

	fmt.Println("we can also use decoders/encoders to work with streams:", res)

	return nil
}

// YAMLData is our common data struct
// with YAML struct tags
type YAMLData struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

// ToYAML dumps the YAMLData struct to
// a YAML format bytes.Buffer
func (t *YAMLData) ToYAML() (*bytes.Buffer, error) {
	d, err := yaml.Marshal(t)
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(d)

	return b, nil
}

// Decode will decode into TOMLData
func (t *YAMLData) Decode(data []byte) error {
	return yaml.Unmarshal(data, t)
}

// TOML Marshal = name = "Name1"
// age = 20

// JSON Marshal= {"name":"Name2","age":30}
// YAML Marshal = name: Name3
// age: 40

// TOML Unmarshal = {Example1 99}
// JSON Unmarshal = {Example2 98}
// Yaml Unmarshal = {Example3 97}
// We can unmarshal into a map instead of a struct: map[key:value]
// we can also use decoders/encoders to work with streams: map[key:value key2:value2]
