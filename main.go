package main

import (
    "fmt"
    "GoMod/pkg/list"
    "strconv"
)

func main() {
    var userInput string

    for userInput != "q" {
        fmt.Println()
        fmt.Println("Welcome to GOMOD")
        fmt.Println("These are your created lists")
        list.PrintList()
        fmt.Println("Enter the list number to view the list")
        fmt.Println("To create a new list enter 'a'")
        fmt.Println("To quit enter 'q'") 
        fmt.Scanln(&userInput)

        fmt.Println("==================")

        switch {
        case userInput == "q":
            fmt.Println("quit")
        case userInput == "a":
            var newListName string
            fmt.Println("Enter the list name")
            fmt.Scanln(&newListName)
            list.CreateList(newListName)
            list.PrintLists()
        default:
            if listNumber, err := strconv.ParseFloat(userInput, 64); err == nil {
                var selectedList *list.List = list.GetList(int64(listNumber))
                for _, item := range selectedList.listItems {
                    fmt.Println(item)
                }
            }
        }
    }
}
