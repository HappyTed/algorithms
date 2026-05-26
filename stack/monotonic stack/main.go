package main

import "fmt"

type Stack[T any] struct {
	value []T
}

func (s *Stack[T]) Push(value T) {
	s.value = append(s.value, value)
}

func (s *Stack[T]) Pop() T {
	pop := s.value[len(s.value)-1]
	s.value = s.value[:len(s.value)-1]
	return pop
}

func (s *Stack[T]) Peek() T {
	return s.value[len(s.value)-1]
}
func (s *Stack[T]) Len() int {
	return len(s.value)
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		value: make([]T, 0),
	}
}

/**
* Находит следующий больший элемент для каждого элемента массива.
* Алгоритм:
*  Используем стек для хранения индексов элементов, для которых ещё не найден следующий больший.
*  Стек поддерживается в монотонно убывающем порядке значений (чем глубже в стеке — тем больше значение).
*  При встрече элемента, большего, чем элемент на вершине стека, мы "закрываем" этот элемент, его следующий больший — текущий элемент.
*
* @param nums входной массив целых чисел
* @return массив, где result[i] — следующий больший элемент для nums[i], или -1, если его нет
 */
func nextGreaterElement(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	for i := range result {
		result[i] = -1
	}

	// Стек хранит индексы элементов (не сами значения!), чтобы знать, куда записывать результат
	var stack = NewStack[int]()

	fmt.Println("Len:", stack.Len())

	for i := 0; i < n; i++ {
		fmt.Println("i:", i)
		fmt.Println("stack:", stack)
		fmt.Println("result:", result, "\n")

		// while (!stack.isEmpty() && nums[i] > nums[stack.peek()])
		for stack.Len() > 0 && nums[i] > nums[stack.Peek()] {
			// Извлекаем индекс из стека — для этого элемента мы нашли следующий больший
			index := stack.Pop() // stack pop
			// Записываем текущий элемент как следующий больший для nums[index]
			result[index] = nums[i]
		}
		// Текущий индекс всегда помещается в стек — возможно, для него найдётся больший элемент позже
		stack.Push(i) // stack push
	}

	return result
}

func main() {
	var input = []int{2, 1, 2, 4, 3}
	var result = nextGreaterElement(input)

	fmt.Println("Вход:  ", input)
	fmt.Println("Выход: ", result) // Ожидаем: [4, 2, 4, -1, -1]
}
