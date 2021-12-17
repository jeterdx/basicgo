go build httpBasic.go

./httpBasic &

open "http://localhost:8000/list"
open "http://localhost:8000/price?item=socks"

#Create new item
open "http://localhost:8000/create?item=shirt&price=30"
open "http://localhost:8000/price?item=shirt"

#Update the price
open "http://localhost:8000/update?item=socks&price=100"
open "http://localhost:8000/price?item=socks"

#Delete the item
open "http://localhost:8000/delete?item=socks"
open "http://localhost:8000/list"


rm ./httpBasic
#ps ax | grep http
#kill -9 PID