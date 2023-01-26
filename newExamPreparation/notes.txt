	// for node != nil {
	// 	if a == 5 {
	// 		break
	// 	}
	// 	fmt.Println("nodo numero: ", node.Val, "nodo prossimo: ", node.Next.Val, "nodo precedente: ", node.Previous.Val)
	// 	node = node.Previous
	// 	a++
	// }

	// func f(p *Node, k int) int {
	// 	a := 1
	// 	start := p
	// 	end := p.Next
	// 	fmt.Println(start, end)
	// 	for start.Val != end.Val {
	// 		a++
	// 		if a == k {
	// 			fmt.Println(p.Val)
	// 		}
	// 	}
	// 	return a
	// }

	// sta roba sulla lista circolare va in overflow perché p non é mai == nil

	// func addNewNode(p *linkedList, k int) {
	// 	node := newNode(k)
	// 	if p.Head == nil {
	// 		p.Head = node
	// 		p.Tail = node
	// 	} else {
	// 		node.Previous = p.Tail
	// 		p.Tail.Next = node
	// 		p.Tail = node
	// 		p.Tail.Next = p.Head
	// 	}
	// }