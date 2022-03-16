FROM golang:1.17 AS builder

ENV NOTEPATH="$GOPATH/src/github.com/VladMasarik/ef-notes/"
EXPOSE 8080

RUN mkdir -p $NOTEPATH
WORKDIR $NOTEPATH

COPY ./ ./

RUN go mod download
RUN go build -o bin/notes-app ./src

CMD [ "./bin/notes-app" ]