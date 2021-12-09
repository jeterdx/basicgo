go build httpHTML.go

./httpHTML &

open "http://localhost:8000/list"

#ps ax | grep http
#kill -9 PID