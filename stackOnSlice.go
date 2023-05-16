package main

import "fmt"

//Создание своего типа - "стэк"
type stack struct {
	data []int
}

//Создание методов для работы с типом "стэк"
//Если вызвать "data" напрямую, то можно использовать методы для работы с этим типом. В нашем слачае со слайсом.

//Добавление в конец стэка элемента.
func (stack *stack) push(number int) {
	stack.data = append(stack.data, number)
}

//Получение из конца элемента с его удалением из стэка.
func (stack *stack) pop() int {
	n := stack.data[(len(stack.data) - 1)]
	stack.data = stack.data[:(len(stack.data) - 1)]
	return n
}
func main() {
	var myStack stack
	fmt.Println(`Стэк в начале: `, myStack)

	myStack.push(1)
	fmt.Println(`Стэк после добавления одного элемента`, myStack)

	myStack.push(2)
	myStack.push(3)
	fmt.Println(`Стэк после добавления двух элементов`, myStack)

	myStack.pop()
	fmt.Println(`Стэк после удаления одного элемента`, myStack)

}
