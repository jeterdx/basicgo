go build svg.go
./svg > ../templates/polygon.svg
rm ./svg

cd ..
go build server.go
./server

curl "http://localhost8989/"

rm ./server