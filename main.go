package main

import (
	"fmt"
)

type menu struct {
	Id     int
	Pid    int
	Name   string
	Select int8
}

var menus []menu = []menu{
	{Id: 1, Pid: 0, Name: "电子", Select: 0},
	{Id: 2, Pid: 0, Name: "家电", Select: 0},
	{Id: 3, Pid: 0, Name: "水果", Select: 0},
	{Id: 4, Pid: 0, Name: "建材", Select: 0},
	{Id: 5, Pid: 1, Name: "手机", Select: 0},
	{Id: 6, Pid: 1, Name: "电脑", Select: 0},
	{Id: 7, Pid: 5, Name: "Apple", Select: 0},
	{Id: 8, Pid: 5, Name: "HUAWEI", Select: 0},
	{Id: 9, Pid: 8, Name: "HUAWEI P50", Select: 0},
	{Id: 10, Pid: 2, Name: "电视", Select: 1},
	{Id: 11, Pid: 2, Name: "洗衣机", Select: 1},
	{Id: 12, Pid: 10, Name: "海尔", Select: 1},
	{Id: 13, Pid: 10, Name: "小天鹅", Select: 1},
}

type menuTree struct {
	Id     int
	Pid    int
	Name   string
	Select int8
	Child  []*menuTree
}

func GenMenu(id int, menus []menu) []*menuTree {
	mList := make([]*menuTree, 0, 100)
	for _, m := range menus {
		if m.Pid == id {
			mList = append(mList, &menuTree{
				Id:     m.Id,
				Pid:    m.Pid,
				Name:   m.Name,
				Select: m.Select,
				Child:  GenMenu(m.Id, menus),
			})
		}
	}

	return mList
}

//菜单分割符
var separator string = "|--"

//菜单显示
func showMenu(id int, menu []*menuTree, sep string) {
	for _, v := range menu {
		if v.Pid == id {
			fmt.Println(fmt.Sprintf(`%v%v,%v%v(%v)`, sep, v.Pid, v.Id, v.Name, v.Select))
			showMenu(v.Id, v.Child, fmt.Sprintf(`%v%v`, sep, separator))
		}
	}
}

func main() {
	//菜单
	menuId := 0
	mlist := GenMenu(menuId, menus)
	showMenu(menuId, mlist, separator)
	
	//for _, v := range mlist {
	//	fmt.Println("|--", v.Id, v.Pid, v.Name)
	//	for _, val := range v.Child {
	//		fmt.Println("|--|--", val.Id, val.Pid, val.Name)
	//		for _, value := range val.Child {
	//			fmt.Println("|--|--|--", value.Id, value.Pid, value.Name)
	//			for _, m := range value.Child {
	//				fmt.Println("|--|--|--|--", m.Id, m.Pid, m.Name)
	//			}
	//		}
	//	}
	//}

}
