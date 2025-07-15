func isValid(s string) bool {
	for strings.Contains(s, "()") || strings.Contains(s, "{}") || strings.Contains(s, "[]") {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "[]", "")
	}
	return s == ""
}

fmt.Println(isValid("()[]{}"))      // true  — все пары корректны
fmt.Println(isValid("([{}])"))      // true  — вложенные скобки
fmt.Println(isValid("(]"))          // false — неправильная пара
fmt.Println(isValid("([)]"))        // false — скобки перекрещены
fmt.Println(isValid("{[]}"))        // true  — валидно
fmt.Println(isValid("(()"))         // false — не закрыта одна скобка
