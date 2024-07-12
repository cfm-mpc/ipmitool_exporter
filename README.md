# Description

Prometheus exporter to collect the inlet temperature of the server using the `ipmitool` command.

Information about the prometheus Python client can be found at:
http://prometheus.github.io/client_python/

# Requirements

- `python 3.6` with:
    - prometheus_client

- `ipmitool`

# Instructions

1. Clone the repository

```bash
git clone git@github.com:cfm-mpc/ipmitool_exporter.git
```

2. Copy `exporter.py` into the server you want to monitor

3. Run the exporter

```bash
[root@dave ipmitool_exporter] python3 exporter.py --port 5000
```

Optional arguments:
- `--port`: The exporter's server port (default 8000)
- `--interval`: The polling interval for the metrics (default 5s)

4. View the metrics

```bash
[root@dave ipmitool_exporter] curl -L http://localhost:5000

# HELP python_gc_objects_collected_total Objects collected during gc
# TYPE python_gc_objects_collected_total counter
python_gc_objects_collected_total{generation="0"} 194.0
python_gc_objects_collected_total{generation="1"} 0.0
python_gc_objects_collected_total{generation="2"} 0.0
# HELP python_gc_objects_uncollectable_total Uncollectable objects found during GC
# TYPE python_gc_objects_uncollectable_total counter
python_gc_objects_uncollectable_total{generation="0"} 0.0
python_gc_objects_uncollectable_total{generation="1"} 0.0
python_gc_objects_uncollectable_total{generation="2"} 0.0
# HELP python_gc_collections_total Number of times this generation was collected
# TYPE python_gc_collections_total counter
python_gc_collections_total{generation="0"} 48.0
python_gc_collections_total{generation="1"} 4.0
python_gc_collections_total{generation="2"} 0.0
# HELP python_info Python platform information
# TYPE python_info gauge
python_info{implementation="CPython",major="3",minor="6",patchlevel="8",version="3.6.8"} 1.0
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 4.71646208e+08
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 2.2646784e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.72076645414e+09
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.24000000000000002
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 6.0
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1024.0
# HELP my_inprogress_request Inlet Temperature
# TYPE my_inprogress_request gauge
my_inprogress_request 33.0
```