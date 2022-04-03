package linktable

import (
    "reflect"
    "testing"
)

func TestCreate(t *testing.T) {
    emptyLinkTable := &LinkTable{
        head: nil,
        tail: nil,
        len:  0,
    }
    t.Run("empty linked list", func(t *testing.T) {
        if got := Create(); !reflect.DeepEqual(got, emptyLinkTable) {
            t.Errorf("Create() = %v, want %v", got, emptyLinkTable)
        }
    })
}

func TestPushFront(t *testing.T) {
    linkTable := Create()
    n1 := linkTable.PushFront(3)
    n2 := linkTable.PushFront(2)
    n3 := linkTable.PushFront(1)

    // should not be successfully added to front, because n4's val is not type of int
    n4 := linkTable.PushFront(true)
    t.Log(n4 == nil)

    t.Log(n1.Next() == nil, n1.Prev() == n2, n1 == linkTable.Back())
    t.Log(n2.Prev() == n3, n2.Next() == n1)
    t.Log(n3.Prev() == nil, n3.Next() == n2, n3 == linkTable.Front())
    t.Log(linkTable.len == 3)
}

func TestPushBack(t *testing.T) {
    linkTable := Create()
    n1 := linkTable.PushBack(3)
    n2 := linkTable.PushBack(2)
    n3 := linkTable.PushBack(1)
    t.Log(n1.Next() == n2, n1.Prev() == nil, n1 == linkTable.Front())
    t.Log(n2.Prev() == n1, n2.Next() == n3)
    t.Log(n3.Prev() == n2, n3.Next() == nil, n3 == linkTable.Back())
    t.Log(linkTable.Len() == 3)
}

func TestRemove(t *testing.T) {
    linkTable := Create()
    n1 := linkTable.PushBack(3)
    n2 := linkTable.PushBack(2)
    n3 := linkTable.PushBack(1)
    // should remove failed, because this node is not belong to linkTable
    linkTable.Remove(&Node{
        prev:      n3.prev,
        next:      n3.next,
        linkTable: nil,
        Val:       n3.Val,
    })
    t.Log(linkTable.Len() == 3)
    nodes := []*Node{n2, n1, n3}
    for _, n := range nodes {
        linkTable.Remove(n)
        t.Log(linkTable.Len(), linkTable.Front(), linkTable.Back())
    }
}
