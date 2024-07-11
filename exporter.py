"""
Prometheus exporter to collect the inlet temperature of the host
using the ipmitool command 
"""

import time
from subprocess import Popen, PIPE

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

    print(output)

def run_metrics():
    """
    Run the metrics in a loop
    """

    POLLING_INTERVAL=5

    while True:
        fetch()
        time.sleep(POLLING_INTERVAL)

run_metrics()