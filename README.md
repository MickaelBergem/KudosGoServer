# Kudos server

A test Kudos Server, able to track count of the number of Kudos given with a
given button.

    go get github.com/MickaelBergem/KudosPlease
    ./KudosPlease

## Testing

To test the code and generate the coverage:

    go test -coverprofile cover.out
    go tool cover -html=cover.out -o cover.html
