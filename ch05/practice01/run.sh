go build dirfindlinks/findlinks.go
go build dirfetch/fetch.go

./fetch https://xkcd.com/ | ./findlinks

rm ./findlinks
rm ./fetch