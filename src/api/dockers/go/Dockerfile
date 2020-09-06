# Use the golang alpine image as the base stage of a multi-stage routine
FROM golang:1.14-alpine as base

# Set the working directory to /api
WORKDIR /api

# Extend the base stage and create a new stage named dev
FROM base as dev

# Copy the go.mod and go.sum files to /api in the image's filesystem
COPY ../../go.* ./
COPY ../../.env ./

# Install go module dependencies in the image's filesystem
RUN go mod download

# ENV sets an environment variable
# Create GOPATH and PATH environment variables
ENV GOPATH /go
ENV PATH ${GOPATH}/bin/usr/local/go/bin:$PATH

# Print go environment for debugging purposes
RUN go env

# Install development dependencies to debug and live reload api
RUN go get ithub.com/go-delve/delve/cmd/dlv \
    && go get github.com/githubnemo/CompileDaemon

# Provide meta data about the ports the container must expose
# port 4000 -> api port
# port 2345 -> debugger port
EXPOSE 4000 2345

# Extend the dev stage and create a new stage named test
FROM dev as test

# Disable CGO and run unit tests
RUN export CGO_ENABLED=0 && \
    go test -v ./...

# Extend the test stage and create a new stage named build-stage
FROM test as build-stage

# Build Go executables for Multiple Platforms
# -s = disable symbol table
# -w = disable DWARF generation
RUN GOOS=linux go build -ldflags "-s -w" -o out ../../

# Extend the base stage and create a new stage named prod
FROM base as prod

# Copy only the files we want from a few stages into the prod stage
COPY --from=build-stage /api .

# Create a new group and user, recursively change directory ownership, then allow the binary to be executed
RUN addgroup gopher && adduser -D -G gopher gopher \
    && chown -R gopher:gopher /api/main && \
    chmod +x ./main

USER gopher
EXPOSE 4000

# Test the container to check that it is still working
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "wget", "-q", "0.0.0.0:4000" ]

# Provide the default command for the production container
CMD [ "./out" ]