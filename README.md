# go-performance-examples

This repository contains some examples of how to profiling and take heap dumps with Go:

- `cpu`
    - `net-http-pprof`: This uses the `net/http/pprof` package with the default `ServeMux` to automatically registry `pprof` handlers
    - `net-http-pprof-hanlders`: This uses the `net/http/pprof` package with a custom `ServeMux` to register `pprof` handlers
    - `pkg-profile`: This uses the `profile` package to profile CPU through code
- `memory`