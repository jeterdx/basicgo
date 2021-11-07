go build dirCountNodeType/countNodeType.go
go build dirFetch/fetch.go

./fetch https://xkcd.com/ | ./countNodeType

rm ./countNodeType
rm ./fetch