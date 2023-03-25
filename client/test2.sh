#Add
go run client.go add -key A -value a
go run client.go add -key B -value b
go run client.go remove -key A
go run client.go getall
go run client.go add -key D -value d
go run client.go getall
go run client.go remove -key B
go run client.go getall
go run client.go add -key E -value e
go run client.go add -key C -value c

#Get all
go run client.go getall

#Remove
go run client.go remove -key D
go run client.go getall
go run client.go remove -key B
go run client.go getall
go run client.go add -key Sam -value 'is awesome' 
go run client.go getall
go run client.go remove -key C
go run client.go add -key bloXroute -value 'is fire'
go run client.go get -key Sam
go run client.go add -key 'bloXroute & Sam' -value 'Great Match'

#Get all
go run client.go getall
