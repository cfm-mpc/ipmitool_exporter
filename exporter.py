"""
Prometheus exporter to collect the inlet temperature of the host
using the ipmitool command 
"""

import argparse
import re
import time
from subprocess import Popen, PIPE
from prometheus_client import Gauge, start_http_server

def fetch():
    """
    Fetch the inlet temperature running shell commands
    """

    IPMI_TEMP_SENSOR="System Temp"

    sensor=Popen(["/usr/bin/ipmitool", "sdr", "get", IPMI_TEMP_SENSOR], stdout=PIPE, universal_newlines=True)
    output,errors=sensor.communicate() # store the output and close the pipe

    inlet_temp=re.findall("Sensor Reading.*", output)[0].split()[3]

    return(inlet_temp)
    
def main():
    """
    Command line arguments
    """

    parser=argparse.ArgumentParser()
    parser.add_argument("--port", type=int, default=8000)
    parser.add_argument("--interval", type=int, default=5, choices=[5,10,15,30,60])
    args=parser.parse_args()

    """
    Start the http server and expose the metrics
    """

    start_http_server(args.port)

    # Generate the requests
    inlet_temperature=Gauge('my_inprogress_request', 'Inlet Temperature')

    while True:
        inlet_temperature.set(fetch())
        time.sleep(args.interval)

if __name__ == "__main__":
    main()