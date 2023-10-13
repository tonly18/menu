package menu

import (
	"fmt"
)

// 菜单
type Menu struct {
	Id     int     //当前id
	Pid    int     //父级id
	Name   string  //名字
	Select int8    //是否选中
	Child  []*Menu //子菜单
}

type treeMenu struct {
	Separator string
	menus     []*Menu
}

func NewMenu(m []*Menu, sep string) *treeMenu {
	return &treeMenu{
		Separator: sep,
		menus:     m,
	}
}

// GenMenu 递归生成树形菜单(数组)
//
// @params:
//
//	id int	父级菜单id
//
// @return:
//
//	[]*Menu
func (m *treeMenu) GenMenu(id int) []*Menu {
	mList := make([]*Menu, 0, 100)
	for _, v := range m.menus {
		if v.Pid == id {
			mList = append(mList, &Menu{
				Id:     v.Id,
				Pid:    v.Pid,
				Name:   v.Name,
				Select: v.Select,
				Child:  m.GenMenu(v.Id),
			})
		}
	}

	return mList
}

// ShowMenu 生成树形菜单
//
// @params:
//
//	id 	 int	 父级菜单id
//	menu []*Menu 处理过的menu列表
//	sep	string	 分隔符
//	mtree *[]string	树形menu
//
// @return:
func (m *treeMenu) ShowTreeMenu(id int, menu []*Menu, sep string, mtree *[]string) {
	for _, v := range menu {
		if v.Pid == id {
			//fmt.Println(fmt.Sprintf(`%v%v,%v%v(%v)`, sep, v.Pid, v.Id, v.Name, v.Select))
			//*mtree = append(*mtree, fmt.Sprintf(`%v%v,%v%v(%v)`, sep, v.Pid, v.Id, v.Name, v.Select))
			*mtree = append(*mtree, fmt.Sprintf(`%v%v`, sep, v.Name))
			m.ShowTreeMenu(v.Id, v.Child, fmt.Sprintf(`%v%v`, sep, m.Separator), mtree)
		}
	}
}
