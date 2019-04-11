# This Dockerfile builds everything at once.
# It can be used for an automated build on the Docker Hub.

# Build the backend
FROM golang:latest AS backend-build-env
RUN go get -u github.com/golang/dep/cmd/dep
RUN mkdir -p /go/src/github.com/zottelchin/Votierungstracker
WORKDIR /go/src/github.com/zottelchin/Votierungstracker
COPY *.go /go/src/github.com/zottelchin/Votierungstracker/
COPY Gopkg.* /go/src/github.com/zottelchin/Votierungstracker/
RUN dep ensure
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static" -s' -installsuffix cgo -o Votierungstracker -v .

# Build the frontend
FROM node AS frontend-build-env
RUN mkdir -p /frontend
WORKDIR /frontend
COPY frontend/ /frontend/
RUN npm run build

# Put everything together
FROM scratch

COPY --from=backend-build-env /go/src/github.com/zottelchin/Votierungstracker/Votierungstracker /
COPY --from=frontend-build-env /frontend/dist/ /frontend/dist/

ENV GIN_MODE=release
WORKDIR /
VOLUME /db
EXPOSE 8900

ENTRYPOINT [ "/Votierungstracker" ]
