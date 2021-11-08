go build dirGetAllContents/getAllContents.go
go build dirFetch/fetch.go

./fetch https://xkcd.com/ | ./getAllContents

rm ./getAllContents
rm ./fetch