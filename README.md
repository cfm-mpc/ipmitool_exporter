# Description

Prometheus exporter to collect the inlet temperature of the server using the `ipmitool` command.

# Requirements

- A `Go` development environment to compile the program.
- `ipmitool` in the host we want to monitor. 

# Instructions

## Clone the repository

```bash
[root@localhost ~] git clone git@github.com:cfm-mpc/ipmitool_exporter.git
[root@localhost ~] cd ipmitool_exporter
```

## Build the exporter

Compile the exporter: 

```bash
[root@localhost ipmitool_exporter] make 
```

## Run the exporter

Run the binary in the host you want to monitor. You can cofigure a system service for this. 

```bash
[root@localhost ipmitool_exporter] ./ipmitool_exporter -sensor "Inlet Temp" -address=":5000" -path="/metrics"
```

## View the metrics

Check that the metrics are being exposed succesfully:

```bash
[root@localhost ipmitool_exporter] curl -L localhost:5000/metrics
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

# Integration with Grafana

## Requirements

- A `Prometheus server` (accessible through http by the `ipmitool_exporter`)
- A `Grafana server`(accessible through http by Prometheus)

## Instructions

- Configure Prometheus to scrape the metrics from the host we want to monitor:
```bash
[root@prometheus ~] sudo vim /etc/prometheus/prometheus.yml

(...)

   - job_name: "inlet-temperature"
     metrics_path: "/metrics"
     static_configs:
      - targets: ["localhost:5000"]
```

- Import the dashboard (`grafana/inlet-temperature.json`) into Grafana:
`Dashboards`->`New`->`Import`->`Upload JSON file`
