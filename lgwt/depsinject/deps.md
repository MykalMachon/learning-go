# Dependency Injection

This seems like a complicated topic, but the site does a pretty good job of going over what it is?

To explain it in my own words, it's basically making an underlying requirement of the function more generic so you can *inject* that underlying piece *the dependency* into the function making it more generic. 

## A simple example

This is a very basic and *strict* Greet function.

```go
func Greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```

It directly uses `fmt.Printf()` which prints to STDOUT. This is hard to test, and it might be easier to test if we could have it write *somewhere else* 

If you look at Printf, it uses `FPrintF()` under the hood. This function expects a io.Writer, a string, and the additional args for formatting.

What is io.Writer? 

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

So why don't we also expect an io.Writer in our function? this would make it so we can pass in anything that implements the Writer interface for printing. This would make it a lot easier to test too! 

Here's our app code:

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
```

And here's a test that works the same way.

```go
package main 

import (
    "testing", 
    "bytes"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}
    Greet(&buffer, "Chris")

    got := buffer.String()
    want := "Hello, Chris"

    if got != want {
        t.ErrorF("got %q, wanted %q", got, want)
    }
}
```

## A more practical example

For an example of this in the internet lands, see `./http/http.go`

# Wrapping up

In general, code ends up being hard to test because it wrote data to somewhere we couldn't control or couldn't test (think database connection handlers, http libs, etc)

injecting dependencies will fix that problem and a few more:

- Test our code: If you can't test a function easily, it's usually because of dependencies hard-wired into a function or global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency (via an interface) which you can then mock out with something you can control in your tests.

- Separate our concerns: decoupling where the data goes from how to generate it. If you ever feel like a method/function has too many responsibilities (generating data and writing to a db? handling HTTP requests and doing domain level logic?) DI is probably going to be the tool you need.

- Allow our code to be re-used in different contexts: The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.

## How does mocking get involved?

Mocking is basically just faking an interace and using DI to control what we're testing. Will cover later but the biggest concern with mocking is that by mocking you're actually removing relevancy of tests as you're over-mocking the functionality.

## Learn the go standard lib

By having some familiarity with the `io.Writer` interface we are able to use `bytes.Buffer` in our test as our Writer and then we can use other Writers from the standard library to use our function in a command line app or in web server.

Cool! The more familiar you are with the standard library the more you'll see these general purpose interfaces which you can then re-use in your own code to make your software reusable in a number of contexts.

# Further reading

- [LGWT Dependency Injection](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection)
- [The Go Programming Language](https://www.amazon.co.uk/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440)