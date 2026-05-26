package main

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Len  int
}

/**
 * Разворачивает подсписок с позиции m по n (1-индексированные).
 *
 * @param head голова исходного списка
 * @param m    начало подсписка (включительно)
 * @param n    конец подсписка (включительно)
 * @return новая голова списка после разворота
 */
func reverseBetween[T any](head *Node, m, n int) *Node {
	if head == nil || m == n {
		return head
	}

	// Фиктивный узел для упрощения обработки случая, когда m = 1
	dummy := NewNode(0)
	dummy.Next = head

	return dummy
}

func main() {

}
