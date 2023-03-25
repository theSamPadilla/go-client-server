go run client.go add -key A -value a
go run client.go getall
go run client.go add -key B -value b
go run client.go getall
go run client.go remove -key A
go run client.go getall
go run client.go add -key D -value d
go run client.go getall
go run client.go remove -key B
go run client.go getall
go run client.go add -key E -value e
go run client.go getall
go run client.go add -key C -value c
go run client.go getall
go run client.go remove -key D
go run client.go getall
go run client.go remove -key E
go run client.go getall
go run client.go remove -key C
go run client.go getall