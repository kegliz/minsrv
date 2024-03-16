FROM golang:1.21 as builder

LABEL maintainer="Zoltan Kegli <kegli.zoltan@gmail.com>"

WORKDIR /app

# Enable Go's DNS resolver to read from /etc/hosts
RUN echo "hosts: files dns" > /etc/nsswitch.conf.min

# Create a minimal passwd so we can run as non-root in the container
RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc/passwd.min

# Fetch latest CA certificates
RUN apt-get update && \
  apt-get install -y ca-certificates


#COPY go.mod go.sum ./
#RUN go mod download

COPY . .

RUN CGO_ENABLED=0 \
  go build -ldflags '-s -w' -tags 'osusergo netgo' -o minsrv

FROM scratch AS final

# Copy over the binary artifact
COPY --from=builder /app/minsrv /

# Copy configuration from builder
COPY --from=builder /etc/nsswitch.conf.min /etc/nsswitch.conf
COPY --from=builder /etc/passwd.min /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --chown=nobody --from=builder /tmp /tmp

USER nobody

ENTRYPOINT ["/minsrv"]
