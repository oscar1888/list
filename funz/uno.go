package list

import "fmt"

type Lengthable interface {
    fmt.Stringer
    Length() int
}

type Node struct {
    next *Node
    payload Lengthable
}

type List struct {
    first *Node
}

