go build dirElemByTagName/elemByTagName.go
go build dirFetch/fetch.go

./fetch https://golang.org/| ./elemByTagName

rm ./elemByTagName
rm ./fetch