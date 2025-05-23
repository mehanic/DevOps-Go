package main

import (
	"fmt"
	"github.com/kirigaikabuto/lesson20acl"
)

func main() {
	cfg := lesson20acl.PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "setdatauser",
		Password: "123456789",
		Database: "lesson20acl",
		Params:   "sslmode=disable",
	}
	store, err := lesson20acl.NewPostgresRoleStore(cfg)
	if err != nil {
		panic(err)
		return
	}
	service := lesson20acl.NewRoleService(store)
	cmdCreate := lesson20acl.CreateRoleCommand{Name: "1234"}
	response, err := cmdCreate.Exec(service)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(response)
	cmdList := lesson20acl.ListRoleCommand{}
	roles, err := cmdList.Exec(service)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(roles)

	//role := &lesson20acl.Role{
	//	Id:   "2",
	//	Name: "role2",
	//}
	//_, err = store.Create(role)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//roles, err := store.List()
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fmt.Println(roles)
	//_, err = store.Get("1")
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//err = store.Delete("1")
	//if err != nil {
	//	panic(err)
	//	return
	//}
}
