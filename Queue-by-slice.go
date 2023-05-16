package main

import "fmt"

//Создание своего типа - "очередь"
type queue struct {
	data []int
}

//Создание методов для работы с типом "очередь"
//Если вызвать "data" напрямую, то можно использовать методы для работы с этим типом. В нашем слачае со слайсом.

//Добавление в конец очереди элемента.
func (queue *queue) enter(number int) {
	queue.data = append(queue.data, number)
}

//Получение из начала элемента с его удалением из очереди.
func (queue *queue) leave() int {
	n := queue.data[0]
	queue.data = queue.data[1:]
	return n
}
func main() {
	var myStack queue
	fmt.Println(`Очередь в начале: `, myStack)

	myStack.enter(1)
	fmt.Println(`Очередь после добавления одного элемента`, myStack)

	myStack.enter(2)
	myStack.enter(3)
	fmt.Println(`Очередь после добавления двух элемента`, myStack)

	myStack.leave()
	fmt.Println(`Очередь после удаления одного элемента`, myStack)

}
