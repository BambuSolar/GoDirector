# iron/go is the alpine image with only ca-certificates added
FROM golang

WORKDIR /go/src/github.com/BambuSolar/GoDirector/

# Now just add the binary
ADD . /go/src/github.com/BambuSolar/GoDirector/

CMD ./GoDirector download

ENTRYPOINT ./GoDirector