package generator

import (
  "log"
)

type Cell struct {
  X int
  Y int
}

type stack struct {
 cells []Cell
 size int
}

func (this *stack) Push(v Cell) {
  this.cells = append(this.cells, v)
  this.size++
}

func (this *stack) Pop() Cell {
  // get top element
  top := this.cells[len(this.cells) - 1]

  // slice array
  this.cells = this.cells[:len(this.cells) - 1]
  this.size--

  // return element at the top
  return top
}

func (this *stack) Len() int {
  return this.size
}

func (this *stack) Peek() Cell {
  return this.cells[len(this.cells) - 1]
}

func (this *stack) Print() {
  log.Println(this.cells[len(this.cells) - this.size:])
}

func newStack(size int) *stack {
    cells := make([]Cell, size)
    return &stack{ cells: cells, size: 0 }  // enforce the default value here
}
