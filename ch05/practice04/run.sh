go build dirGetAllLinks/getAllLinks.go
go build dirFetch/fetch.go

./fetch https://xkcd.com/ | ./getAllLinks

rm ./getAllLinks
rm ./fetch