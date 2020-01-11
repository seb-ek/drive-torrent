FROM golang:1.10.4
LABEL maintainer="Prasad Kavinda <pp.kavinda@gmail.com>"

WORKDIR $GOPATH/src/github.com/ppkavinda/drive-torrent
COPY . .


RUN chmod +x ./main
EXPOSE 3000
CMD ./main


