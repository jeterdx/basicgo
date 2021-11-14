go build dirFindLinks/findlinks.go
go build dirFetch/fetch.go

./fetch https://xkcd.com/ | ./findlinks

rm ./findlinks
rm ./fetch