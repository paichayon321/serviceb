# serviceb


#To Build go on alpine
docker run -v ${pwd}:/go/src/example -it golang:alpine sh
---
cd src/example
go build goapp
--


docker build -t myserviceb:2 .

docker run --rm --name serviceb -p 32000:1323 -e HOST=http://httpbin.org -e URI=/status/503 myserviceb:2
