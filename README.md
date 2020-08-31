# challenge-go
# You must executed the following sentences
go get github.com/gin-gonic/gin
go get github.com/gin-gonic/contrib/sessions

#Server initialized at port :7007
#Example to use :
#first, you try to login
POST - > http://localhost:7007/user-api/login
#second, query for this endpoint
GET -> http://localhost:7007/user-api/pictures

#Postman collections (to be import)
Challenge-go.postman_collection.json

#To run and build
git clone https://github.com/jproza/challenge-go.git

cd /challenge-go
go build main.go
./main

#or only run
go run main.go
