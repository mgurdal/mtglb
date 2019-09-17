// Queue tipi için aşağıdaki kullanımı sağlayan metodları yazın.

package main

type Queue struct {
	arr []int
}

// Append adds the given element at the end of the queue
func Append(q *Queue) {
	Queue = append(Queue, e)
}

// AppendLeft adds the given element at the head of the queue
func AppendLeft(q *Queue) {
	Queue = append([]int{"i"},Queue...)
}

// Pop removes the latest element in the queue
func Pop(q *Queue) {
	e, Queue := Queue[len(Queue)-1], Queue[:len(Queue)-1]
}

// Pop removes the first element in the queue
func PopLeft(qu *Queue) {
	e, Queue := Queue[0], Queue[:len(Queue)-1]
}

func main() {
	q := &Queue{}
	q.Append(4)     // q: [4]
	q.AppendLeft(1) // q: [1, 4]
	e := q.Pop()    // q: [1], e: 4
	e = q.PopLeft() // q: [0], e: 1
}
