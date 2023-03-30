#Heap / memory
go tool pprof -png http://localhost:6060/debug/pprof/heap > pprof/heap.png

#CPU
go tool pprof -png http://localhost:6060/debug/pprof/profile > pprof/cpu.png

#Gorotuines
go tool pprof -png http://localhost:6060/debug/pprof/goroutine > pprof/goroutine.png