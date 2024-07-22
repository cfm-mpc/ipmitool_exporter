# Description

Prometheus exporter to collect the inlet temperature of the server using the `ipmitool` command.

# Requirements

- A `Go` development environment

# Instructions

## Clone the repository

```bash
[root@dave ~] git clone git@github.com:cfm-mpc/ipmitool_exporter.git
[root@dave ~] cd ipmitool_exporter
```

## Build the exporter

By default, the ipmisensor is `System Temp`. If you need to change it, override the `sensor` variable in the Makefile with the option `-e`. E.g: 

```bash
[root@dave ipmitool_exporter] make -e sensor="Inlet Temp" 
```

## Run the exporter

```bash
[root@dave ipmitool_exporter] ./ipmitool_exporter -address=:5000 -path=/metrics
```

## View the metrics

```bash
[root@dave ipmitool_exporter] curl -L localhost:5000/metrics
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 7
# HELP ipmitool_temp Inlet Temperature
# TYPE ipmitool_temp gauge
ipmitool_temp 32
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 262144
```
