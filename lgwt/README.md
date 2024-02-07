# Learn Go with Tests

[See the site for Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

## Reasoning 

I chose to start here as it was recommended to me by a friend and because I saw it in a youtube video that suggested getting started with go. 

## Fundamentals 

### Installing Go

[See how to do the install here](https://go.dev/doc/install). The linux guide was pretty easy which is nice. just had to unzip the zip and add some stuff to my bashrc

### Creating modules 

To create a module you should create a new folder, cd into it, then run `go mod init <name-of-folder-or-module-or-repo-url>`. 
Each module can have multiple "packages" which are typically in folders that are *inside* the module. Each file has `package <package-name>` at the top. 

- module 
  - package1
    - package1.go
    - package1_test.go
  - package2
    - package2.go
    - package2_test.go
  - go.mod (keeps track of module info, deps, and other info)
  - go.sum (deps lock? seems like a `package.lock.json`?)

This could be wrong, but it's what I'm seeing so far while going through this course.

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

See more on `godoc` here [on the site](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#go-doc)

- if you run `godoc` in the same folder as your `go.mod` file it will load your package in as part of the standard lib section.
- comments above each function show up in the documentation as well.
- you can setup examples (see `integers/adder_test.go:ExampleAdd` for an example). These show up in docs. 

## Stopping points: 
- [Stopping point 2024-01-16](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#keep-going-more-requirements)
- [Stopping point 2024-01-17](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)
- [Stopping point on 2024-01-22](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#write-the-test-first-3)
- [Stopping point on 2024-01-23](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces)
- [Stopping point on 2024-02-06](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#refactor)
- [Stopping point on 2024-02-07](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps)