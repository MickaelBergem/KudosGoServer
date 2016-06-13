# Kudos server

A test Kudos Server, able to track the number of Kudos given with a given
button.

    go get github.com/MickaelBergem/KudosPlease
    ./KudosPlease

The server is now running on port 8090.

## Usage

Once the server is running, you can send HTTP request:

    # To retrieve the kudo value of an existing button
    curl http://localhost:8090/new-awesome-blogpost
    # Increase the kudo value
    curl -X POST http://localhost:8090/new-awesome-blogpost
    # Create a new button
    curl -X PUT -d 'URL=http://blog.securem.eu/new-awesome-blogpost/' http://localhost:8090/new-awesome-blogpost

## Docker

A lightweigh Docker container is available on DockerHub:

    docker pull suixo/kudosplease
    docker run --rm -it -v `pwd`/kudos_count.sqlite3:/kudos_count.sqlite3 --name kudosplease -p 80:8090 suixo/kudosplease

You can build the container image yourself with the provided Dockerfile:

    # First, statically compile the server
    go build --ldflags '-extldflags "-static"' -o main .
    # Then, copy it inside the image
    docker build -t kudosplease .

The resulting container image will weigh only ~12MB.

## Testing

To test the code and generate the coverage:

    go test -coverprofile cover.out
    go tool cover -html=cover.out -o cover.html
