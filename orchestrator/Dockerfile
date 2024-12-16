FROM golang:1.22-bullseye AS builder
LABEL authors="jrourke"

WORKDIR /dist
COPY ./ ./

RUN apt install -y ca-certificates

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /orchestrator

# Install pigpio
RUN <<-EOF
    # If platform == "rpi" then
    if [ "$PLATFORM" == "rpi" ]; then
        wget https://github.com/joan2937/pigpio/archive/master.zip
        unzip master.zip
        cd pigpio-master
        make
        sudo make install
    fi
EOF

# Install and run click.c
RUN <<-EOF
    if [ "$PLATFORM" == "rpi" ]; then
        git clone https://github.com/dupontgu/retro-ipod-spotify-client.git
        cd retro-ipod-spotify-client/clickwheel/
        gcc -Wall -pthread -o click click.c -lpigpio -lrt
        gcc -o click click.c -lpigpio -lrt
        ./click &
    fi
EOF

FROM scratch

EXPOSE 9091

COPY --from=builder /orchestrator /orchestrator
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder ./piper-db ./

# Run
CMD ["/orchestrator"]