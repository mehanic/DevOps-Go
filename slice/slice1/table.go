package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	langs := [][]string{{"C#", "C", "Python"},
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "RUST", "Crystal", "OCAML"}}
	for _, v := range langs {
		for _, lang := range v {
			fmt.Print(lang, " ")
		}
		fmt.Println()
	}
	fmt.Println("---------------------------")
	//footer := []string{"ngôn ngữ", "ngữ ngôn", "test"}
	//header := []string{"test", "ngôn ngữ", "ngữ ngôn"}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	// Change table lines
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("╪")
	table.SetRowSeparator("-")

	table.SetAlignment(tablewriter.ALIGN_LEFT)

	/*if footer != nil {
		table.SetFooter(footer)
	}
	table.SetHeader(header)*/
	for _, v := range langs {
		table.Append(v)
	}
	table.Render()

}
