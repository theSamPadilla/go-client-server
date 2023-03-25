go run client.go add -key A -value a
go run client.go add -key B -value b
go run client.go add -key D -value d
go run client.go add -key E -value e
go run client.go add -key C -value c
go run client.go remove -key D
go run client.go add -key X -value x
go run client.go add -key X -value y
go run client.go remove -key B
go run client.go add -key Z -value z
go run client.go add -key F -value f
go run client.go getall

#Expected A:a, E:e, C:c, X:y, Z:z, F:f