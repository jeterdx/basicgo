go build httpBasic.go

./httpBasic &

#open http://localhost:8000/list
#open http://localhost:8000/price?item=socks

#sleep 2
#Create new item
open "http://localhost:8000/create?item=shirt&price=30"
open "http://localhost:8000/price?item=shirt"

#sleep 2
#Update the price
#open http://localhost:8000/update?item=socks&price=100
#open http://localhost:8000/price?item=socks

#sleep 2
#Delete the item
#open http://localhost:8000/delete?item=socks
#open http://localhost:8000/list

#ps ax | grep http
#kill -9 PID