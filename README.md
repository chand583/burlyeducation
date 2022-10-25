#   Setup Guide

## Check go version
* $ go version   //go version go1.17

## Set GOPATH and GOROOT (You can either go to your project directory and EXPORT it or edit below file to add.)
* sudo nano ~/.bashrc and add below line in the end of file
* export GOROOT=/usr/local/go   OR <your/gopath/dir> 
* export GOPATH=$HOME/go OR <your/gopath/dir>
* export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

## Go to project dir

```cd burlyeducation```

## Download dependencies

```go mod tidy```

## Create app.conf
* Create a file name "app.conf" inside directory "conf" and do config/db related changes here [ create and copy it from file "app-local.conf"]
* For Developer only [Optional and only if first approach is not followed] - Create a symlink for e.g ln -s Your-project-path/conf/app-local.conf Your-project-path/conf/app.conf

## DB changes
* Config Changes - Go to config/app.conf and change the DB and Cache related parameters. 
* Create Table in Postgres - burlyeducation schema file "conf/burlyeducation.sql"

## DB migration
* bee migrate -driver=<DRIVERNAME> -conn="<CONNECTION-STRING>"
* example - bee migrate -driver=postgres -conn="postgres://postgres:postgres@localhost:5454/burlyeducation?sslmode=disable"

## Start the server 

```bee run```

## Restart the server (beego creates routes using notation so it needs a restart for changes to take effect )

```CTRL + C```
```bee run```

## To run the server with Swagger documentation [Recommended for DEV environment only]
* bee run -downdoc=true -gendoc=true
* Access swagger documentation from - http://localhost:8080/swagger/

## Test

```http://localhost:8080/burlyed/v1/question```

#SET server time zone in UTC format



# How to perform Unit testing

 First Create file with _test.go for example article_test.go after that goto tests folder in CLI 

 ```
  cd tests 
 
 ```

then that run command 

``` 
go test  -v 
OR  
if you want to run the unit test across your project then run below command from your root directory
go test -v ./...

``` 

if you are getting conf/app.conf not found error copy conf/app.conf folder in tests folder then run that command 

if you want to run single test function run this command 

```
go test -v -run=TestQuestionGetAll #replace TestQuestionGetAll with our function name

```

Migrate 
bee generate migration user_table
