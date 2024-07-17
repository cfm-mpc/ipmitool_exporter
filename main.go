/* Prometheus exporter to collect the inlet temperature of the host
using the ipmitool command */

package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func fetch() string {
	
	/* Fetch the inlet temperature running shell commands */

	var IPMI_TEMP_SENSOR string = "System Temp"

	sensor := exec.Command("/usr/bin/ipmitool", "sdr", "get", IPMI_TEMP_SENSOR)
	output,_ := sensor.Output() 

	re,_ := regexp.Compile("Sensor Reading.*") //Regex to match
	inlet_temp:=strings.Fields(re.FindString( string(output) ))[3]

	return inlet_temp

}

func main() {

	fmt.Println( fetch() )

}