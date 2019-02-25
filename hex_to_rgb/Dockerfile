ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder


RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN mkdir -p /go/src/github.com/marcpar/HextoRGB

ADD . /go/src/github.com/marcpar/HextoRGB

RUN go get gopkg.in/go-playground/colors.v1

WORKDIR hex_to_rgb

ENV GO111MODULE=on


RUN go mod download
RUN go builder


# Final stage: the running container.
FROM scratch AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the compiled executable from the second stage.
COPY --from=builder /app /app

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

# Perform any further action as an unprivileged user.
USER nobody:nobody

ENTRYPOINT  hex_to_rgb