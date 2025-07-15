package main

func NewGender(value string) Gender {
    return Gender{strings.TrimSpace(value)}
}

func NewGender(value string) *Gender {
    return &Gender{strings.TrimSpace(value)}
}


func (g *Gender) HandlePost(w http.ResponseWriter, req *http.Request) {
  fmt.Fprint(w, g.value)
}

func NewGender(value string) MyGender {
  gender := Gender{strings.TrimSpace(value)}
  http.Handle("/", http.HandlerFunc(gender.HandlePost))
  return gender
}

gen := NewGender("male")
gen.value = "female"
// http POST will still return "male"