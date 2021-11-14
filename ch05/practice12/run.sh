go build dirOutline/outline.go
go build dirFetch/fetch.go

./fetch https://xkcd.com/ | ./outline

rm ./outline
rm ./fetch