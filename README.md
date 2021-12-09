# kube-tmux

## Installation

```
go get -u github.com/marcsauter/kube-tmux
```

## Configuration

Add `kube-tmux` to your status bar in `.tmux.conf`:
```
set -g status-right "#[fg=black] #(kube-tmux) #H %H:%M %d.%m.%Y"
```
or with a custom template:
```
set -g status-right "#[fg=black] #(kube-tmux '<{{.Context}}::{{.Namespace}}>') #H %H:%M %d.%m.%Y"
```

Consider reducing the status-interval for more current information:
```
set -g status-interval 5
```

## Benchmark

```plaintext
$ go test -v -bench . -benchtime=10s -test.benchmem
goos: linux
goarch: amd64
pkg: github.com/marcsauter/kube-tmux
cpu: Intel(R) Core(TM) i7-8650U CPU @ 1.90GHz
BenchmarkTemplate
BenchmarkTemplate-8   	  327462	     33208 ns/op	    4607 B/op	      47 allocs/op
BenchmarkPrinting
BenchmarkPrinting-8   	45178626	       241.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/marcsauter/kube-tmux	22.445s
```
