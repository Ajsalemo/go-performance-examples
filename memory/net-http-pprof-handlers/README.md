# go-performance-examples

How to run this example:
1. Run `go get` if needed
2. Run `go run .` and access `localhost:8080/api/array` a couple of times to run the code on this route to increase memory usage.
3. Access `localhost:8080/debug/pprof/heap`, this will download a `heap` file to your local computer
5. Use the `go tool pprof` command to read the heap dump - examples of the `top10` command can be found [here](https://pkg.go.dev/net/http/pprof#section-documentation)
- Example: Run `go tool pprof heap` - where `heap` is the name of the profile downloaded to your computer, relative to where this command is being ran.
