#Provided example

#Add
curl -X POST localhost:6969/add -H "Content-Type: application/json" -d '{"key": "A", "value": "a"}'
curl -X POST localhost:6969/add -H "Content-Type: application/json" -d '{"key": "B", "value": "b"}'
curl -X POST localhost:6969/add -H "Content-Type: application/json" -d '{"key": "D", "value": "d"}'
curl -X POST localhost:6969/add -H "Content-Type: application/json" -d '{"key": "E", "value": "e"}'
curl -X POST localhost:6969/add -H "Content-Type: application/json" -d '{"key": "C", "value": "c"}'
  
#Get all
curl localhost:6969/

#Remove
curl -X POST localhost:6969/remove -H "Content-Type: application/json" -d '{"key": "D"}'

#Get key and index
curl localhost:6969/key/B
curl localhost:6969/index/2

#Get all
curl localhost:6969/