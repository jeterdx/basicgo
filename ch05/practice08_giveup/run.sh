go build dirFindFirstElement/findFirstElement.go
go build dirFetch/fetch.go

./fetch https://xkcd.com/ | ./findFirstElement

rm ./findFirstElement
rm ./fetch