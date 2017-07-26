package methods

import (
	"math"
	"github.com/ashigirl96/moretypes"
	"fmt"
	"time"
)

type Vertex moretypes.Vertex
type MyFloat64 float64

type Abser interface {
	Abs() float64
}
type I interface {
	M()
}

type T struct {
	S string
}


type Person struct {
	Name string
	Age int
}

type IPAddr [4]byte

type MyError struct {
	When time.Time
	What string
}

type ErrNegativeSqrt float64


func (f MyFloat64) M() {
	fmt.Println(float64(f))
}


func (t *T) M() {
	fmt.Print(t.S)
}


func (v Vertex) Abs2() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}


func (v *Vertex) Scale(f float64) {
	// bang method
	v.X *= f
	v.Y *= f
}



func ScaleFunc(v *Vertex, f float64) {
	v.X *= f
	v.Y *= f
}


func (f MyFloat64) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}


func Describe(i interface{}) {
	fmt.Printf("(%v, %T)", i, i)
}


func TypeSwitchesDo(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v * 2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}


func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (ipaddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, \n\t%s",
		e.When, e.What)
}


func Run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}


func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}


const delta = 1e-10

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // When cast x into ErrNegativeSqrt,
									// it will return String fmt.Sprintf..
									// but, it's also child of Error()
	}
	z := x
	for {
		n := z - (z*z-x)/(2*z)
		if math.Abs(n-z) < delta {
			break
		}
		z = n
	}
	return z, nil
}
