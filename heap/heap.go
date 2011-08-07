package main

import (
  // "fmt"
)

type PriorityQueue struct {
	q []int
	length int
}

func (queue PriorityQueue) parent(n int) int {
	if n == 0 {
		return -1
	}
	return (int)(n / 2)
}

func (queue PriorityQueue) youngChild(n int) int {
	return n * 2
}

func (queue PriorityQueue) insert(x int) {
	queue.length += 1
	queue.q[queue.length] = x

	queue.bubbleUp(queue.length)
}

func (queue PriorityQueue) bubbleUp(n int) {
	if queue.parent(n) == 1 { return }

	if queue.q[queue.parent(n)] > queue.q[n] {
		queue.swap(n, queue.parent(n))
		queue.bubbleUp(queue.parent(n))
	}
}

func (queue PriorityQueue) swap(x int, y int) {
	temp := queue.q[x]
	queue.q[x] = queue.q[y]
	queue.q[y] = temp
}

func (queue PriorityQueue) build(s int[], length int) {

}

func main() {

}
