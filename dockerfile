#build Backend 
FROM golang:latest AS backend-build-env
RUN go get -u github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/github.com/zottelchin/Votierungstracker
WORKDIR /go/src/github.com/zottelchin/Votierungstracker
COPY main.go /go/src/github.com/zottelchin/Votierungstracker
COPY Gopkg.toml /go/src/github.com/zottelchin/Votierungstracker
COPY Gopkg.lock /go/src/github.com/zottelchin/Votierungstracker
RUN dep ensure
RUN go build -o server

#build frontend 
FROM node AS frontend-build-env
RUN mkdir -p /frontend
WORKDIR /frontend
COPY frontend/ /frontend/
RUN npm build

#final build 
FROM ubuntu:18.04
RUN mkdir -p /app/db
WORKDIR /app/
COPY --from=backend-build-env /go/src/github.com/zottelchin/Votierungstracker/server /app/
RUN mkdir -p frontend/dist
COPY --from=frontend-build-env /frontend/dist/ frontend/dist/
EXPOSE 8900
ENTRYPOINT ./server
