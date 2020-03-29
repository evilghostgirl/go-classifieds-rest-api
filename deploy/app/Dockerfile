FROM golang
WORKDIR /go/src/classifieds-rest-api
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix .
EXPOSE 3000
ENTRYPOINT ./classifieds-rest-api