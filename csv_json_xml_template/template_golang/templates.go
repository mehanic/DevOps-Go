package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)


func main() {
	if err := RunTemplate(); err != nil {
		panic(err)
	}
}


const sampleTemplate = `
    This template demonstrates printing a {{ .Variable | printf "%#v" }}.
    {{if .Condition}}
    If condition is set, we'll print this
    {{else}}
    Otherwise, we'll print this instead
    {{end}}
    Next we'll iterate over an array of strings:
    {{range $index, $item := .Items}}
        {{$index}}: {{$item}}
    {{end}}
    We can also easily import other functions like strings.Split
    then immediately used the array created as a result:
    {{ range $index, $item := split .Words ","}}
        {{$index}}: {{$item}}
    {{end}}
    Blocks are a way to embed templates into one another
    {{ block "block_example" .}}
        No Block defined!
    {{end}}

    {{/*
        This is a way
        to insert a multi-line comment
    */}}
`

const secondTemplate = `
    {{ define "block_example" }}
        {{.OtherVariable}}
    {{end}}
`

// RunTemplate initializes a template and demonstrates a variety
// of template helper functions
func RunTemplate() error {
	data := struct {
		Condition     bool
		Variable      string
		Items         []string
		Words         string
		OtherVariable string
	}{
		Condition:     true,
		Variable:      "variable",
		Items:         []string{"item1", "item2", "item3"},
		Words:         "another_item1,another_item2,another_item3",
		OtherVariable: "I'm defined in a second template!",
	}

	funcmap := template.FuncMap{
		"split": strings.Split,
	}

	// these can also be chained
	t := template.New("example")
	t = t.Funcs(funcmap)

	// We could use Must instead to panic on error
	// template.Must(t.Parse(sampleTemplate))
	t, err := t.Parse(sampleTemplate)
	if err != nil {
		return err
	}

	// to demonstrate blocks we'll create another template
	// by cloning the first template, then parsing a second
	t2, err := t.Clone()
	if err != nil {
		return err
	}

	t2, err = t2.Parse(secondTemplate)
	if err != nil {
		return err
	}

	// write the template to stdout and populate it
	// with data
	err = t2.Execute(os.Stdout, &data)
	if err != nil {
		return err
	}

	return nil
}

//-----

// CreateTemplate will creat a template file that contains data
func CreateTemplate(path string, data string) error {
	return ioutil.WriteFile(path, []byte(data), os.FileMode(0755))
}

// InitTemplates sets up templates from a directory
func InitTemplates() error {
	tempdir, err := ioutil.TempDir("", "temp")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempdir)

	err = CreateTemplate(filepath.Join(tempdir, "t1.tmpl"), `
        Template 1! {{ .Var1 }}
        {{ block "template2" .}} {{end}}
        {{ block "template3" .}} {{end}}
    `)
	if err != nil {
		return err
	}

	err = CreateTemplate(filepath.Join(tempdir, "t2.tmpl"), `
        {{ define "template2"}}Template 2! {{ .Var2 }}{{end}}
    `)
	if err != nil {
		return err
	}

	err = CreateTemplate(filepath.Join(tempdir, "t3.tmpl"), `
        {{ define "template3"}}Template 3! {{ .Var3 }}{{end}}
    `)
	if err != nil {
		return err
	}

	pattern := filepath.Join(tempdir, "*.tmpl")

	// Parse glob will combine all the files that match glob
	// and combine them into a single template
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}

	// Execute can also work with a map instead
	// of a struct
	tmpl.Execute(os.Stdout, map[string]string{
		"Var1": "Var1!!",
		"Var2": "Var2!!",
		"Var3": "Var3!!",
	})

	return nil
}

//----

// HTMLDifferences highlights some of the differences
// between html/template and text/template
func HTMLDifferences() error {
	t := template.New("html")
	t, err := t.Parse("<h1>Hello! {{.Name}}</h1>\n")
	if err != nil {
		return err
	}

	// html/template auto-escapes unsafe operations like javascript injection
	// this is contextually aware and will behave differently
	// depending on where a variable is rendered
	err = t.Execute(os.Stdout, map[string]string{"Name": "<script>alert('Can you see me?')</script>"})
	if err != nil {
		return err
	}

	// you can also manually call the escapers
	fmt.Println(template.JSEscaper(`example <example@example.com>`))
	fmt.Println(template.HTMLEscaper(`example <example@example.com>`))
	fmt.Println(template.URLQueryEscaper(`example <example@example.com>`))

	return nil
}




// This template demonstrates printing a "variable".
    
//     If condition is set, we'll print this
    
//     Next we'll iterate over an array of strings:
    
//         0: item1
    
//         1: item2
    
//         2: item3
    
//     We can also easily import other functions like strings.Split
//     then immediately used the array created as a result:
    
//         0: another_item1
    
//         1: another_item2
    
//         2: another_item3
    
//     Blocks are a way to embed templates into one another
    
//         I'm defined in a second template!
    

