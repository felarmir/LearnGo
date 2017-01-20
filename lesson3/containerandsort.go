package lesson3

import (
	"container/list"
	"fmt"
	"sort"
)

func LearnList() {
	var lst list.List
	lst.PushBack(1)
	lst.PushBack(2)
	lst.PushBack(4)
	lst.PushBack(3)
	lst.PushBack(7)
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println("V:", e.Value.(int))
	}
}

// ===============================================================
// sort list
type User struct {
	ID   int
	Name string
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Name < u[j].Name
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func LearnSort() {
	us := []User{
		{10, "Denis"},
		{11, "Jenifer"},
		{12, "Adam"},
	}
	sort.Sort(Users(us))
	fmt.Println(us)
}
