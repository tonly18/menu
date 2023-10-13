package menu_test

import (
	"fmt"
	"github.com/tonly18/menu"
	"sort"
	"testing"
)

var menus []*menu.Menu = []*menu.Menu{
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
	{Id: 13, Pid: 11, Name: "小天鹅", Select: 1},
}

func TestMenu(t *testing.T) {
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].Id < menus[j].Id
	})

	menuPid := 0
	treeMenu := menu.NewMenu(menus, "|--")
	mList := treeMenu.GenMenu(menuPid)

	mtree := make([]string, 0, 50)
	treeMenu.ShowTreeMenu(menuPid, mList, treeMenu.Separator, &mtree)
	for _, v := range mtree {
		fmt.Println(v)
	}
}
