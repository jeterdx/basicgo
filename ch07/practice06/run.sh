go build tempflag.go

./tempflag
./tempflag -temp -18C
./tempflag -temp 212°F
./tempflag -temp 300K
./tempflag -temp 300°K
./tempflag -temp 100

rm ./tempflag