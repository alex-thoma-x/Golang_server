# Golang_server
Simple Server using golang and reading yaml file in Golang

This program converts a simple Nested yaml file into Go Struct and
stores thr values in BoltDB **https://github.com/boltdb/bolt**

This code uses viper package to unmarshall yaml file into GO Struct
Viper **https://github.com/spf13/viper**

#before running the program 
1.open the folder in Vs code
2.run :go mod init <folder name> in Terminal
3.And now install the packages from Git (ie; viper and Boltdb)
4.run :go get github.com/spf13/viper and go get github.com/boltdb/bolt
5.Now run the command in requirements_&_env to run the program (given file names are same)


emp_detail.html is the Html Template for displayin the output
