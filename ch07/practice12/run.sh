go build httpHTML.go

./httpHTML &

open "http://localhost:8000/list"

rm ./httpHTML

#ps ax | grep http
#kill -9 PID