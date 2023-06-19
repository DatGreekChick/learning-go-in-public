# Day 3: Complete [Tour of Go]

**Date**: 19 June 2023

## What I did

### Wrapped up [Tour of Go]

Again, all Go snippets are lifted from this tutorial sometimes with minor
modifications and/or commentary from me.

Today I wrapped up this tour and was happy to see more exercises towards the
end. My commentary is in real-time, so I occasionally make assumptions or
comparisons that prove to be incorrect as I dive deeper. It's part of my
process!

#### Maps

Enter maps: something that appears, once again, like a dict or object-literal
in Python and JavaScript, respectively. The definition the tour is giving of a
map is as such:

> A map maps keys to values.

Maps can be `nil` if there are no keys within it. Further, a `nil` map doesn't
allow the addition of keys! This is interesting. When I started working in
Python back in 2018, I _loved_ the fact that you didn't have to check `len()`
in order to see if there were any values in a list or dict. In practice, the
times I used such checking...

```python
if not some_dict:
  # do something
```

...I always had the option to add key-value pairs to the dict. In Go, at least
from this small introduction, that option appears to be nonexistent. 😱

```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

func main() {
  // here's make again!
  var m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])
}
```

We can use `make` for maps as well as for arrays in Go. I love a good
overloaded function (one of my favorite C++ features!).

Maps also can be map literals, but their keys are required. Remember when we
looked at struct literals? We didn't have to instantiate a struct literal with
named keys if we didn't want or need to; the struct literal would just set the
default values for the keys in the struct literal. This is not the case with
map literals. We need to assign keys. But what if I assign an empty string as
the key? Does Go even allow us to access empty strings like that?

The answer? Yes!

```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "Bell Labs": Vertex{
    40.68433, -74.39967,
  },
  "Google": Vertex{
    37.42202, -122.08408,
  },
  "": Vertex{ // this is totally valid!
    1.0, -1.0,
  },
}

func main() {
  fmt.Println(m) // map[:{1 -1} Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
  fmt.Println(m[""]) // {1 -1}
}
```

It shouldn't be all that surprising... bad data exists everywhere. 😎

Onwards to some helper functionality with maps:

```go
package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"])

  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"])

  // you can delete keys by passing the map and the key
  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"])

  // you can pass in `ok` as a second return value when accessing a key
  // this will be `false` if the value is not there, and `true` if it is
  // v does hold a value, however: the default for that type (in this case 0)
  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok)
}
```

I've worked with many programming languages in my time as a software engineer.
I can't recall a time when attempting to access a key that was not in the
object didn't throw an error. In Go, you just get the default! That's wild! You
can get similar behavior in Python, but you have to explicitly set the default
value if the key doesn't exist. Go just wants to keep going and gives you some
of your time back!

#### [Map Exercise]

New package alert: `"strings"`! As suggested by the tour, we can use
`strings.Fields` to give us an array of all the words in a string. Like
`string.split` in Python or `String.prototype.split` in JavaScript.

#### Function values

Functions are values and can be passed around as return values and function
parameters. Go functions may also be closures.

> A closure is a function value that references variables from outside its
> body. The function may access and assign to the referenced variables; in this
> sense the function is "bound" to the variables.

```go
package main

import "fmt"

func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {
  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
    )
  }
}
```

To wrap up functions, I worked on a classic [fibonacci closure exercise].

#### Methods

> Go does not have classes. However, you can define methods on types.

WHAT?! That is bananas!!!! Everything is a type?! So, how are we able to call
methods on types? Turns out, we use something called a `receiver` argument.

In the below example, we see that there's a function called `Abs` that takes in
a receiver argument of `v Vertex`. In `main()` we create an instance of
`Vertex` (`v`) and then can call `Abs` using dot notation: `v.Abs()`. You can
define methods on any type so long as the type is defined in the same package
as the method (including any built-in types such as `int`). I agree with this
last bit's strictness, otherwise it would be mayhem:

![mayhem]

IYKYK. Now if only HBO would wrap up the series with one last test... 😭

```go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
}
```

This is all well and good, but that likely means any method on a type needs to
live in its own function closure and not as part of some structure like:

```go
// pseudocode example
type Vertex struct {
  X, Y float64

  func Abs(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
}
```

I prefer class hierarchies like this pseudocode example for simplicity and
readability. When I first started learning JavaScript, I was using ES4 syntax.
That syntax required you to create a class using object syntax and then adding
methods to that object using `.prototype`. It was so cumbersome to me, so when
I began using ES2015+ syntax, I was thrilled with the syntactic sugar of
`class`. This made things so simple and emulated other languages in which I was
interested. But even with that kind of syntax, a program or module can get out
of hand, so the problems are the same: reducing churn, decreasing complexity,
writing manageable and maintainable code that can be well tested, etc.

#### Pointer receivers

Methods can be declared with pointer receivers.

> Since methods often need to modify their receiver, pointer receivers are more
> common than value receivers.

Makes sense.

```go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  v.Scale(10)
  fmt.Println(v.Abs()) // 50
}
```

Now if you change the `Scale` method's receiver to not contain a pointer like
so:

```go
func (v Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}
```

The output of `fmt.Println(v.Abs())` is no longer 50. Instead, it's 5. That's
because the method takes a copy of what you pass in instead of mutating the
object. This is the same behavior for any other function argument.

#### Methods and pointer receivers

Say you have a method that takes in a pointer receiver. If you pass in a value
instead of a pointer, Go will conveniently call the pointer version itself.

```go
func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

var v Vertex
v.Scale(5) // OK
p := &v
p.Scale(10) // OK
```

Under the hood, it's doing this: `(&v).Scale(5)` so it compiles both
invocations of the `Scale` method. This happens in the reverse (with value
receivers) as well!

```go
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

Under the hood with value receivers that are passed a pointer, Go interprets
the call as this: `(*p).Abs()`.

#### Choosing a value or pointer receiver

So when should we use value receivers over pointer receivers and vice versa?
Apparently, there are two reasons you'd want to use a pointer receiver:

1. So the method can modify the value to which its receiver points
2. To avoid copying the value on each method call

This second reason helps keep things more efficient for large receivers (a
large struct, for example). No matter the choice, we need to keep things
consistent and not mix value and pointer receivers in the same method.

#### Interfaces

> An interface type is defined as a set of method signatures. A value of
> interface type can hold any value that implements those methods.

This is as close to a class we're going to get with Go.

```go
package main

import (
  "fmt"
  "math"
)

type Abser interface {
  Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main() {
  var a Abser
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3, 4}

  a = f  // a MyFloat implements Abser
  a = &v // a *Vertex implements Abser

  // In the following line, v is a Vertex (not *Vertex)
  // and does NOT implement Abser.
  // a = v

  fmt.Println(a.Abs()) // 5
}
```

Interfaces are implemented implicitly. What does that mean? Let's use C++ as an
example. If you want to derive one class from another in C++, you need to
specify the base class from which the derived class is being derived. In other
words, you need to state the inheritance explicitly.

```cpp
struct Base
{
    int a, b, c;
};

// every object of type Derived includes Base as a subobject
struct Derived : Base
{
    int b;
};
```

You don't need to do this in Go.

```go
package main

import "fmt"

type I interface {
  M()
}

type T struct {
  S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
  fmt.Println(t.S)
}

func main() {
  // here `i` is of  type `I` (an interface)
  // we assign it to `T` a struct that has a key `S` of type string
  var i I = T{"hello"}
  i.M() // hello
}
```

When we assign `i` to `T{"hello"}`, we're gaining access to the keys in `T`
while also maintaining any methods on `T`. `T` is kind of like a derived class,
and `I` is sort of like a base class (or the opposite). Basically, Go gives you
access to any methods and keys, etc. for both the type and the interface and
doesn't quite care about the organization behind it. Why?

> Implicit interfaces decouple the definition of an interface from its
> implementation, which could then appear in any package without
> prearrangement.

This gives us some wiggle room in that strict requirement of only being able to
create methods in the same package a type is defined.

You can also have an empty interface: `interface{}`. Empty interfaces can hold
values of any type, and every type implements at least zero methods. Why might
this be useful? If the type you pass in is unknown. `fmt.Print` is a great
example in that it can take any number of arguments of type `interface{}`.

```go
package main

import "fmt"

func describe(i interface{}) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  var i interface{}
  describe(i) // (<nil>, <nil>)

  i = 42
  describe(i) // (42, int)

  i = "hello"
  describe(i) // (hello, string)
}
```

#### Interface values

An interface value is similar to a tuple of a value and a concrete type:
`(value, type)`. An interface value holds this value of a particular type. So,
when you call a method on an interface value, Go will execute the method of the
same name on its underlying type.

```go
package main

import (
  "fmt"
  "math"
)

type I interface {
  M()
}

type T struct {
  S string
}

func (t *T) M() {
  fmt.Println(t.S)
}

type F float64

func (f F) M() {
  fmt.Println(f)
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  var i I

  i = &T{"Hello"}
  describe(i) // (&{Hello}, *main.T)
  i.M() // Hello

  i = F(math.Pi)
  describe(i) // (3.141592653589793, main.F)
  i.M() // 3.141592653589793
}
```

Note that in `describe` we pass in `%v` as the value and `%T` as the type while
we only need to `Printf` on `i`. If you try to split this up like so,

```go
val, iType := i
fmt.Println(val, iType)
// ./prog.go:28:16: assignment mismatch: 2 variables but 1 value
```

it doesn't work. There's only one value for `var i I`.

#### Interface values with nil underlying values

The ease of Go and `nil` continues! If the concrete value inside the interface
is `nil`, any method will be called with a `nil receiver`. Some languages -- in
fact all the ones with which I've worked -- would throw their version of a null
pointer exception. In Go, it's a common occurrence to gracefully handle methods
that are called with nil receivers.

```go
package main

import "fmt"

type I interface {
  M()
}

type T struct {
  S string
}

func (t *T) M() {
  if t == nil {
    fmt.Println("<nil>")
    return
  }
  fmt.Println(t.S)
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  var i I
  var t *T

  i = t
  describe(i) // (<nil>, *main.T)
  i.M() // <nil>

  i = &T{"hello"}
  describe(i) // (&{hello}, *main.T)
  i.M() // hello
}
```

However, an interface value that holds a nil concrete value is itself non-nil,
while a nil interface value neither has value nor a concrete type.

```go
package main

import "fmt"

type I interface {
  M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  var i I
  describe(i) // (<nil>, <nil>)
  i.M() // run-time error!
}
```

`i.M()` is a run-time error. We called a method on a nil interface, and because
there's no type inside the interface tuple, Go can't reason with which concrete
method to invoke.

```bash
# so you get this panic error instead
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x482c41]

goroutine 1 [running]:
main.main()
	/tmp/sandbox1525303985/prog.go:12 +0x61
```

#### Type assertions

> A type assertion provides access to an interface value's underlying concrete
> value.

```go
t := i.(T)
```

You can use the `.(T)` syntax for type assertions. Another way you can write
this is how we read from a map:

```go
t, ok := i.(T)
```

The biggest difference between reading from a map and asserting a type, is that
if you don't supply `ok` and aren't recovering your program, you'll run into a
panic:

```bash
hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64

goroutine 1 [running]:
main.main()
	/tmp/sandbox424597700/prog.go:17 +0x14a
```

This is the result of the following snippet:

```go
package main

import "fmt"

func main() {
  var i interface{} = "hello"

  s := i.(string)
  fmt.Println(s) // hello

  s, ok := i.(string)
  fmt.Println(s, ok) // hello true

  f, ok := i.(float64)
  fmt.Println(f, ok) // 0 false

  f = i.(float64) // panic!
  fmt.Println(f)
}
```

When you supply `ok`, Go will again assign `f` in this example to the default
value of the type you're trying to assert. That's why `f` equates to 0.
Omitting the second value from a type assertion (`ok`) is what causes the
panic. How might the program avoid a panic? We can use a type switch for that.

```go
package main

import "fmt"

func do(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("I don't know about type %T!\n", v)
  }
}

func main() {
  do(21) // Twice 21 is 42
  do("hello") // "hello" is 5 bytes long
  do(true) // I don't know about type bool!
}
```

In a type switch statement, which is a construct that permits type assertions
in a series, you can use the `type` keyword instead of passing in a specific
type `T` like in a type assertion. The default case will take whatever type is
prevalent that might not be prevalent elsewhere in the switch.

#### [Stringer exercise]

A `Stringer` is "one of the most ubiquitous interfaces" and defined by the
`fmt` package. You can use a `Stringer` to print values in a particular way.

#### [Errors exercise]

Similar to `fmt.Stringer`, the `error` type is a built-in interface in Go.
Previously, I've assumed that Go could [overload] functions like C++. That was
an incorrect assumption that I discovered with this exercise (sadly).

#### [Readers exercise]

There is an `io` package which specifies the `io.Reader` interface. This
represents the read end of a stream of data. There are many implementations of
this interface in Go.

Notably, there's a `Read` method as part of the `io.Reader` interface.

```go
func (T) Read(b []byte) (n int, err error)
```

Its purpose is to populate the given byte slice with data and return the number
of bytes populated along with an error value. When the stream ends, it returns
an `io.EOF` error.

The exercise is to accomplish the following:

> Implement a Reader type that emits an infinite stream of the ASCII character
> 'A'.

### Remaining exercises in this section

- [rot13Reader]
- [images]

This is what I got for one iteration of the image exercise:

![bluescale]

#### Type parameters

```go
func Index[T comparable](s []T, x T) int
```

Few things to note here: `comparable` allows you to compare values (using `==`
and `!=` operators) between objects of the same type. As for a type parameter,
that just means that you have a Go function that is written to handle multiple
types. The type parameter appears in brackets before the function's arguments
like in the above example: `[T comparable]`.

```go
package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
  for i, v := range s {
    // v and x are type T, which has the comparable
    // constraint, so we can use == here.
    if v == x {
      return i
    }
  }
  return -1
}

func main() {
  // Index works on a slice of ints
  intSlice := []int{10, 20, 15, -10}
  fmt.Println(Index(intSlice, 15)) // 2

  // Index also works on a slice of strings
  stringSlice := []string{"foo", "bar", "baz"}
  fmt.Println(Index(stringSlice, "hello")) // -1
}
```

#### Generic types

In Go you can have a generic function like the above example, and you can have
a generic type by using a type parameter once again. This is helpful to
implement any generic data structures in a program.

```go
// List represents a singly-linked list that holds values of any type
// where [T any] is the type parameter of the List type
type List[T any] struct {
	next *List[T]
	val  T
}
```

#### Concurrency

Note: I'll be skipping the exercises for these as I'll dive into this more as
my learning progresses.

We finally get to concurrency! This is what I've been hearing is so great about
Go. Let's start off with goroutines.

> A goroutine is a lightweight thread managed by the Go runtime.

```go
go f(x, y, z)
```

starts a new goroutine running

```go
f(x, y, z)
```

`f`, `x`, `y`, and `z` get evaluated in the current goroutine, while the
execution of `f` occurs in a new goroutine.

When you use a goroutine, you must synchronize access to shared memory because
goroutines run in the same address space.

Go eases the setup for you with something called channels.

> Channels are a typed conduit through which you can send and receive values
> with the channel operator, `<-`.

You have to define a channel before using it, similar to a map or a slice. You
can once again use `make` to do that and can define it using the `chan` type.

```go
// Note: the direction of the arrow is the direction of the data flow
ch := make(chan int)

// Send v to channel ch
ch <- v

// Receive from ch, and assign value to v
v := <-ch
```

I mentioned simplicity in setup before. That's because Go channels block sends
and receives until the other side is ready by default. This allows goroutines
to sync without explicit locks or conditions.

Channels can be buffered by passing in a second argument to `make` when
defining a channel:

```go
ch := make(chan int, 100)
```

This allows sends to be blocked if the buffer is full and receives to be
blocked when the buffer is empty. If you try to overfill a buffer (or the
buffer is empty), you'll get a fatal deadlock error:

```bash
# overfill (note chan send)
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox2178050203/prog.go:9 +0x5c

# empty (note chan receive)
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/tmp/sandbox2520458947/prog.go:7 +0x45
```

What if you want to close a channel before the buffer is overfilled? Yeah, you
can do that with our trusty friend `ok`: `v, ok := <-ch`. The value of `ok`
will be false if there are no more values for the channel to receive and the
channel is closed.

```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
  x, y := 0, 1
  for i := 0; i < n; i++ {
    c <- x
    x, y = y, x+y
  }
  close(c)
}

func main() {
  c := make(chan int, 10)
  go fibonacci(cap(c), c)
  for i := range c {
    fmt.Println(i)
  }
}
```

It's noted on the tour that only senders should close a channel, never
receivers. Sending on a closed channel causes a panic. But you don't normally
need to close a channel like you do a file. It's only necessary to do this when
the receiver needs to be told that no more values will be sent. One example of
that is to terminate a `range` loop.

You can add a `select` statement to a goroutine in order to wait on multiple
communications operations. This kind of statement will block until one of its
cases can run, then it will execute that case. If multiple cases are ready, the
goroutine will choose one at random to run. Like a `switch` statement, you can
have a `default` case in the `select`. The goroutine will run the default case
if no other case is ready in the statement.

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  tick := time.Tick(100 * time.Millisecond)
  boom := time.After(500 * time.Millisecond)
  for {
    select {
    case <-tick:
      fmt.Println("tick.")
    case <-boom:
      fmt.Println("BOOM!")
      return
    default:
      fmt.Println("    .")
      time.Sleep(50 * time.Millisecond)
    }
  }
}
```

The output of that is this:

```bash
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
tick.
    .
    .
BOOM!
```

💥

[tour of go]: https://go.dev/tour
[map exercise]: ../code/day03/exercise-maps.go
[fibonacci closure exercise]: ../code/day03/exercise-fibonacci-closure.go
[mayhem]:
  https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExaWw0bmg1MmswbmRqeXVscmxsOHJnZWNnb3I3aDNyeHNkODcyaWppOCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/l2JhDqhW1NGLtf6H6/giphy.gif
[stringer exercise]: ../code/day03/exercise-stringer.go
[errors exercise]: ../code/day03/exercise-errors.go
[overload]: https://go.dev/doc/faq#overloading
[readers exercise]: ../code/day03/exercise-readers.go
[rot13reader]: ../code/day03/exercise-rot-reader.go
[images]: ../code/day03/exercise-images.go
[bluescale]: ../img/day03/bluescale.png
