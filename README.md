Minimal server in go for experimenting with loadbalancers

# Build and run
```bash
docker image build -t minsrv .
docker run -p 3333:4112  minsrv a
docker run -p 3334:4112  minsrv b
curl localhost:3333
```
the output should be
```
Server Name: backend-a
Version: 0.0.1
Server IP: (Interface Name: eth0, IP Address: 172.17.0.2)

Date: 2024-03-16 09:47:47.999454
URI: /
```

# Environment variables
default SERVERPORT is 4112, but can be overridden with environment variable
```bash
docker run -e "SERVERPORT=3333" -p 3333:3333  minsrv
curl localhost:3333
```
