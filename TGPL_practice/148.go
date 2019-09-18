package main

import (
    "fmt"
)

/* Definition for singly-linked list.*/
type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    //merge sort bottom up
    length := 0
    for ptr := head; ptr != nil; ptr = ptr.Next {
        length++
    }
    if length < 2 {
        return head
    }
    
    
    var head1, end1, head2, end2, prevend *ListNode
    for iter :=1 ; iter < length; iter *= 2 {
        head1 = head
        firstIter := true
        fmt.Println("iter:%v\n+++++++++++++++++++++++++++++", iter)
        for i, ptr:=0, head; i < 6 ;  i++{
            if ptr == nil {
                fmt.Printf("nil\n")
                break;
            } else {
                fmt.Printf("%v-> ", ptr.Val)
                ptr = ptr.Next
            }
        }
        for head1 != nil {
            end1 = head1
            for i := 1; i < iter && end1.Next != nil; i, end1 =  i+1, end1.Next{
            }
            if end1.Next == nil {
                break;
            }
            head2, end2 = end1.Next, end1.Next
            for i := 1; i < iter && end2.Next != nil; i, end2 =  i+1, end2.Next{
            }
            fmt.Println("before sort, h1:%v e1:%v h2:%v e2:%v", head1.Val, end1.Val, head2.Val, end2.Val)
            for i, ptr:=0, head; i < 6 ;  i++{
                if ptr == nil {
                    fmt.Printf("nil\n")
                    break;
                } else {
                    fmt.Printf("%v-> ", ptr.Val)
                    ptr = ptr.Next
                }
            }
            
            //Inorder to sort the next chunk, we need to save a pointer points to the next node of end2
            nextHead := end2.Next
            head1, end1, head2, end2 = merge(head1, end1, head2, end2)
            fmt.Println("after sort, h1:%v e1:%v h2:%v e2:%v", head1.Val, end1.Val, head2.Val, end2.Val)
            for i, ptr:=0, head; i < 6 ;  i++{
                if ptr == nil {
                    fmt.Printf("nil\n")
                    break;
                } else {
                    fmt.Printf("%v-> ", ptr.Val)
                    ptr = ptr.Next
                }
            }
            
            //after merge
            if firstIter {
                head = head1
                firstIter = false
            } else {
                //connect previous sorted last node to this chunk's first node: head1
                prevend.Next = head1
            }
            prevend = end2
            head1 = nextHead

            fmt.Println()
            fmt.Println("afterwards: ")
            for i, ptr:=0, head; i < 6 ;  i++{
                if ptr == nil {
                    fmt.Printf("nil\n")
                    break;
                } else {
                    fmt.Printf("%v-> ", ptr.Val)
                    ptr = ptr.Next
                }
            }
            fmt.Println()
            fmt.Println("====")
        }
        //head2 == nil, no need for sorting
        //set prevend points to head1
        prevend.Next = head1
    }
    
    return head
}

func merge(head1 *ListNode, end1 *ListNode, head2 *ListNode, end2 *ListNode) ( *ListNode, *ListNode, *ListNode, *ListNode){

    //swap (head1, end1) and (head2, end2) if head1.Val > head2.Val
    if head1.Val > head2.Val {
        head1, end1, head2, end2 = head2, end2, head1, end1
    }
    ptr1, ptr2, end2Next := head1, head2, end2.Next
    var ptr1Next, ptr2Next *ListNode
    for ptr1 != nil && ptr1 != end1 && ptr1.Next != nil && ptr2 != nil && ptr2 != end2Next {
        if ptr1.Next.Val > ptr2.Val {
            //ptr1.Next.Val > ptr2.Val
            //insert ptr2 node after ptr1
            //set ptr2 = ptr1.Next
            ptr1Next, ptr2Next = ptr1.Next, ptr2.Next
            ptr1.Next = ptr2         
            ptr2.Next = ptr1Next
            ptr2 = ptr2Next
        }
        /*Two patterns:
            1. ptr1.Next.Val < ptr2.Val
                => move ptr1 to the next one only
            2. else
                => after the upper block, we insert ptr2 node behind ptr1 and move ptr2 forward
                   we need to check the value inequation between former ptr1.Next(the next node after new inserted ptr2 head node) and new ptr2
                   so we move ptr1 ONE STEP FORWARD, to where we inserted earler
            Therefore, no matter what we only need to move ptr1 one step forward.
        */
        ptr1 = ptr1.Next
    }
    if ptr1 == end1{
        //ptr1 reaches end, concat remaining chunck of ptr2 behind ptr1
        ptr1.Next = ptr2
    } else {
        //ptr2 reaches end, set end2 points to
        end2 = end1
    }
    return head1, end1, head2, end2
}

func main() {
    var head, prev *ListNode
    prev = nil
    list := []int{5,1,3,2,0}
    for i := len(list) -1; i>=0; i-- {
        fmt.Print(list[i])
        newNode := &ListNode{Val:list[i], Next: prev}
        prev = newNode
        if i == 0 {
            head = newNode
        }
    }
    for ptr:=head; ptr != nil; ptr = ptr.Next {
        fmt.Printf("%v-> ", ptr.Val)
    }
    fmt.Printf("nil\n After sort List\n")
    sortList(head)
    for ptr:=head; ptr != nil; ptr = ptr.Next {
        fmt.Printf("%v-> ", ptr.Val)
    }
    fmt.Printf("nil")
}
