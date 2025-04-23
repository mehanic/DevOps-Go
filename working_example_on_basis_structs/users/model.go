package users

import "fmt"

type AddressModel struct {
	District   string
	Street     string
	HomeNumber string
}

func (a *AddressModel) getInfo() string {
	message := fmt.Sprintf("District:%s,Street:%s,HomeNumber:%s", a.District, a.Street, a.HomeNumber)
	return message
}

type User struct {
	Id       string
	Username string
	Password string
	Marks    []int
	Address  AddressModel
}

func (u *User) ShowInfo() {
	fmt.Println(u.Id, u.Username, u.Password, u.getAverageMark(), u.Address.getInfo())
}

func (u *User) getAverageMark() int {
	sumi := 0
	n := len(u.Marks)
	for _, v := range u.Marks {
		sumi += v
	}
	return sumi / n
}


func main() {
	user1 := User{
		Id:       "1",
		Username: "kirito",
		Password: "12123123",
		Marks:    []int{1, 2, 3, 4},
		Address: AddressModel{
			District:   "13",
			Street:     "asdsadas",
			HomeNumber: "34",
		},
	}
	marks2 := []int{3, 4, 5, 3}
	user2 := User{
		Id:       "2",
		Username: "asdsada",
		Password: "13123123",
		Marks:    marks2,
		Address: AddressModel{
			District:   "1",
			Street:     "33333",
			HomeNumber: "123",
		},
	}
	usersData := [] User{user1, user2}
	for _, v := range usersData {
		v.ShowInfo()
	}
}
