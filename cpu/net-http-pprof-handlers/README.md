# go-performance-examples

How to run this example:
1. Run `go get` if needed
2. Run `go run .` and access `localhost:8080/debug/pprof/profile?seconds=30` to start profiling
3. Access `localhost:8080/api/fib` a couple of times to run the code on this route to increase CPU usage.
4. After 30 seconds, this will download a `profile` file to your local computer
5. Use the `go tool pprof` command to read the profile - examples of the `top10` command can be found [here](Profiling Go Programs)
- Example: Run `go tool pprof main.go profile` - where `profile` is the name of the profile downloaded to your computer, relative to where this command is being ran.
