/* Prometheus exporter to collect the inlet temperature of the host
using the ipmitool command */

package main

import (
	"os/exec"
	"regexp"
	"strings"
	"strconv"
	"log"
	"net/http"	
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func fetch() float64 {
	
	/* Fetch the inlet temperature running shell commands */

	var IPMI_TEMP_SENSOR string = "System Temp"

	sensor := exec.Command("/usr/bin/ipmitool", "sdr", "get", IPMI_TEMP_SENSOR)
	output,_ := sensor.Output() 

	re,_ := regexp.Compile("Sensor Reading.*") // regex to match
	inlet := strings.Fields(re.FindString( string(output) ))[3]
	inlet_temp,_ := strconv.ParseFloat(inlet, 64) // convert into float
	
	return inlet_temp

}

func main() {

	// Create and register the metrics
	inlet_temperature := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_inprogress_request",
		Help: "Inlet Temperature",
	})
	prometheus.MustRegister(inlet_temperature)

	// Expose the metrics
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8000", nil))
	
	// Populate the metrics
	for true {
		inlet_temperature.Set(fetch())
	}

}