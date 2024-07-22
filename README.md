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

- Build the exporter.

By default, the ipmisensor is "System Temp". If you need to change it, override the `sensor` variable in the Makefile with the option `-e`. E.g: 

```bash
[root@dave ipmitool_exporter] make -e sensor="Inlet Temp" 
```

- Run the exporter:

```bash
./ipmitool_exporter -address=:5000 -path=/metrics
```

- View the metrics:

```bash
[root@dave ~] curl -L localhost:5000/metrics
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 7
# HELP my_inprogress_request Inlet Temperature
# TYPE my_inprogress_request gauge
my_inprogress_request 33
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 262144
```
