package persons

import "fmt"

type User struct {
	Name        string
	Username    string
	Password    string
	Age         int
	YearOfBirth int
}

func (u User) ShowName() {
	fmt.Println(u.Name)
}

func (u User) ShowUsernameAndPassword() {
	fmt.Println(u.Username, u.Password)
}

func (u *User) SetYearOfBirth() {
	u.YearOfBirth = 2021 - u.Age
}

func (u *User) SetName(name string) {
	u.Name = name
}


func main() {
	u1 := User{
		Name:     "max",
		Username: "george",
		Password: "123212312",
		Age:      23,
	}
	u1.ShowName()
	u1.ShowUsernameAndPassword()
	fmt.Println(u1.YearOfBirth)
	u1.SetYearOfBirth()
	fmt.Println(u1.YearOfBirth)

	u2 := User{
		Name:     "nick",
		Username: "second",
		Password: "123212312",
		Age:      23,
	}
	u2.ShowName()
	u2.SetName("linda")
	u2.ShowName()

	
}
