package list

import "strings"

func (l List) String() string{
    var b strings.Builder
    b.WriteString("[")
    for p := l.first; p != nil; p = p.next {
        if p != l.first {
                b.WriteString(", ")
        }
        b.WriteString(p.payload.String())
    }
    b.WriteString("]")
    return b.String()
}

func (l List) Length() int{
    c := 0
    for p := l.first; p != nil; p = p.next {
        c++
    }
    return c
}

func (l List) Avg() float64 {
    c:=0
    sum:=0
    for p:=l.first; p!=nil; p=p.next {
        c++
        sum+=p.payload.Length()
    }
    return float64(sum)/float64(c)
}

