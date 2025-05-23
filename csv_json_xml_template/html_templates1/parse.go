package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseGlob("../../csv_json_xml/html_templates1/*.html")
	if err != nil {
		log.Fatal("Error:", err)
	}
	data := map[string]string{
		"name":       "Jin Kazama",
		"style":      "Karate",
		"appearance": "Tekken 3",
	}
	if err := tpl.Execute(os.Stdout, data); err != nil {
		log.Fatal("Error:", err)
	}
}

// <html>
//     <body>
//         <h1>Jin Kazama</h1>
//         <ul>
//             <li>First appearance: Tekken 3</li>
//             <li>Style: Karate</li>
//         </ul>
//     </body>
// </html>⏎ 