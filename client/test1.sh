#Provided example

#Add
go run client.go add -key A -value a
go run client.go add -key B -value b
go run client.go add -key D -value d
go run client.go add -key E -value e
go run client.go add -key C -value c

#Get all
go run client.go getall

#Remove
go run client.go remove -key D

#Get all
go run client.go getall
