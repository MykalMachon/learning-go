# Learn Go with Tests

[See the site for Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

## Reasoning 

I chose to start here as it was recommended to me by a friend and because I saw it in a youtube video that suggested getting started with go. 

## Fundamentals 

### Installing Go

[See how to do the install here](https://go.dev/doc/install). The linux guide was pretty easy which is nice. just had to unzip the zip and add some stuff to my bashrc

### Running code 
`go run hello.go`

### Running tests
`go test`

### Running benchmarks 

to run benchmakrs you need to create a benchmark and run them from within the correct folder with `go test -bench="."`

example would be: 

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}
```

### Documentation 

SEe more on `godoc` here [on the site](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#go-doc)

- if you run `godoc` in the same folder as your `go.mod` file it will load your package in as part of the standard lib section.
- comments above each function show up in the documentation as well.
- you can setup examples (see `integers/adder_test.go:ExampleAdd` for an example). These show up in docs. 

## Stopping points: 
- [Stopping point 2024-01-16](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#keep-going-more-requirements)
- [Stopping point 2024-01-17](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)