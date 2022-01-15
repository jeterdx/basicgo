go build imageConverter.go
./imageConverter -format=gif < test.jpg > afterConvertion.gif
rm ./imageConverter
