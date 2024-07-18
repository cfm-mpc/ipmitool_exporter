# Description

Prometheus exporter to collect the inlet temperature of the server using the `ipmitool` command.

# Requirements

- Go
- Prometheus client

# Instructions

- Create module file to track dependencies for this project:
```bash
[root@dave ipmitool_exporter] go mod init github.com/cfm-mpc/ipmitool_exporter
```

- Download the prometheus client:
```bash
[root@dave ipmitool_exporter] go get github.com/prometheus/client_golang/prometheus
[root@dave ipmitool_exporter] go get github.com/prometheus/client_golang/prometheus/promhttp
```

- Test:
```bash
[root@dave ipmitool_exporter] go run main.go
33
```
