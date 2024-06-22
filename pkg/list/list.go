package list
import (
    "fmt"
    "GoMod/pkg/linked_list"
)

type List struct {
    listItems *linked_list.LinkedList
    name string
}

var lists []*List 

func Sayhello() {
    fmt.Println("Hello world!");
}

func CreateList(listName string) {
    if len(listName) > 0 {
        list := &linked_list.LinkedList{}  
        lists = append(lists, &List{listItems: list, name: listName})
    }
}

func PrintLists() {
    fmt.Println("<------Lists----->")
    for i, list := range lists {
        fmt.Println(i+1, ":", list.name)
    }
    fmt.Println("------------------")
}

func GetList(listNum int64)(*List) {
    if listNum > 0 {
        return lists[listNum-1]
    }
    return nil
}
