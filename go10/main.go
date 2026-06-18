package main

import (
	"fmt"
	"net"
	"sync"
)

// constant declarations
const SCALAR = 0.1 // bad
const scalar = 0.1 // good

const (
	name    = "App"
	version = "1.0"
)

// variables grouping
func Foo() int {
	// x := 100
	// y := 2
	// foo := "foo"

	var (
		x   = 100
		y   = 2
		foo = "foo"
	)

	fmt.Println(foo)

	return x + y
}

// functions that panic (name with "must")
func MustParseIntFromString(s string) int {
	// logic
	panic("oops")

	return 10
}

// struct initialization
type Vector struct {
	x int
	y int
}

// mutex grouping
type Server struct {
	listenAddr string
	isRunning  bool

	// always group mutext with protected value
	mu    sync.RWMutex
	peers map[string]net.Conn

	otherMu sync.RWMutex
	other   map[string]int
}

// interface declarations (-er, -or)

// type Storer interface {
// 	Get()
// 	Put()
// 	Delete()
// 	Patch()
// }

type Getter interface {
	Get()
}

type Putter interface {
	Put()
}

type Storer interface {
	Getter
	Putter
}

// funtion grouping

// func simpleUtil() {}
// func veryImportantFunc() {}
// func VeryImportantFuncExported() {}

func VeryImportantFuncExported() {} // 1st because important
func veryImportantFunc()         {} // 2nd important but private
func simpleUtil()                {} // 3rd util so less important

// http handler naming (start with "handle..."")
func handleGetUserById() {}
func handleResizeImage() {}

// enums (kinda!?)
type Suit byte

const (
	SuitHarts Suit = iota
	SuitClubs
	SuitDiamonds
	SuitSpades
)

// constructor (start with "New...")
type Order struct {
	Size float64
}

// inside the main package
func NewOrder(size float64) *Order {
	// logic here

	return &Order{
		Size: size,
	}
}

// inside separated package
func New(size float64) *Order {
	// logic here

	return &Order{
		Size: size,
	}
}

func main() {
	// constant declarations
	// variable grouping
	// functions that panic
	// struct initialization
	vector1 := Vector{10, 20}       // bad
	vector2 := Vector{x: 10, y: 20} //good

	fmt.Print(vector1)
	fmt.Print(vector2)

	// mutex grouping
	// interface declarations
	// function grouping
	// http handler naming
	// enums (kinda!?)
	// constructor
	order1 := NewOrder(3.14)
	fmt.Print(order1)

	// order2 := order.New(3.14)
	// fmt.Print(order2)
}
