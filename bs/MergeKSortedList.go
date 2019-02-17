package bs

import "sort"

type ListNode struct {
	Val int
	Next *ListNode
}


type lists []*ListNode

func (lists lists) Len() int{
	return len(lists)
}

func (lists lists) Less(i,j int) bool{
	return lists[i].Val < lists[j].Val
}

func (lists lists) Swap(i,j int){
	tmp := lists[i]
	lists[i] = lists[j]
	lists[j] = tmp
}


func MergeKLists(lists []*ListNode) *ListNode {
	var root, cur *ListNode
	if len(lists)==0{
		return root
	}
	arr := initList(lists)
	sort.Sort(arr)
	///

	for len(arr)>0{
		nextNode := &ListNode{Val:arr[0].Val}
		if root==nil{
			root = nextNode
			cur = root
		}else{
			cur.Next = nextNode
			cur = cur.Next
		}
		arr = updateArr(arr)
	}
	return root
}

//这里才是难点
func updateArr(arr []*ListNode) lists{
	newNode := arr[0].Next
	arr = arr[1:]
	if newNode==nil{
		return arr
	}

	insertPos := bstNode(arr,0,len(arr)-1,newNode.Val)

	res := make([]*ListNode,0)
	res = append(res,arr[0:insertPos]...)
	res = append(res,newNode)
	if insertPos<len(arr){
		res = append(res,arr[insertPos:]...)
	}
	return res
}

func bstNode(arr []*ListNode,left,right,target int)int{
	if left>right{
		return left
	}
	mid := (left + right) / 2
	if arr[mid].Val<target{
		return bstNode(arr,mid+1,right,target)
	}else{
		return bstNode(arr,left,mid-1,target)
	}
}



func initList(lists []*ListNode) lists{
	res := make([]*ListNode,0)
	for _,list := range lists{
		if list!=nil{
			res = append(res,list)
		}
	}
	return res
}