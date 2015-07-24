FROM ribeirofelix/golang-godep


# Setting up working directory
WORKDIR $GOPATH/src/github.com/vtex/recipes/
COPY . $GOPATH/src/github.com/vtex/recipes/

# Install
RUN godep go build -o /go/bin/recipes


# Exposing port
EXPOSE 8000

ENTRYPOINT ["/go/bin/recipes"]
