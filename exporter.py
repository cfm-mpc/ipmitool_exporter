"""
Prometheus exporter to collect the inlet temperature of the host
using the ipmitool command 
"""

import time
from subprocess import Popen, PIPE
from prometheus_client import Gauge, start_http_server

def collect():
    """
    Populate the metric
    """

    inlet_temperature.set(fetch())

def fetch():
    """
    Fetch the inlet temp
    """

    IPMI_TEMP_SENSOR="System Temp"

    system_temp=Popen(["/usr/bin/ipmitool", "sdr", "get", IPMI_TEMP_SENSOR], stdout=PIPE)
    sensor_reading=Popen(["/usr/bin/grep", "Sensor Reading"], stdin=system_temp.stdout, stdout=PIPE, universal_newlines=True)
    inlet_temp=Popen(["/usr/bin/awk", "{print $4}"], stdin=sensor_reading.stdout, stdout=PIPE, universal_newlines=True)

    output,errors=inlet_temp.communicate()

    system_temp.kill()
    sensor_reading.kill()
    inlet_temp.kill()

    return(output)

"""
Start the server
"""

POLLING_INTERVAL=5

start_http_server(8000)

# Generate the requests
inlet_temperature=Gauge('my_inprogress_request', 'Inlet Temperature')

while True:
    collect()
    time.sleep(POLLING_INTERVAL)