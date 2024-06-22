package linked_list
import "fmt"

type Node struct {
  Data  interface{} 
  Next  *Node      
}

type LinkedList struct {
    Head *Node
}

func (list *LinkedList) Insert(data interface{}) {
    newNode := &Node{Data: data, Next: list.Head}
    list.Head = newNode
}

func (list *LinkedList) Pop()(interface {}, error) {
    if (list.Head == nil) {
        return nil, fmt.Errorf("Linked list is empty")
    }
    poppedNode := list.Head
    list.Head = list.Head.Next
    return poppedNode.Data, nil
}
