/* Prometheus exporter to collect the inlet temperature of the host
using the ipmitool command */

package main

import (
	"flag"
	"os/exec"
	"regexp"
	"strings"
	"strconv"
	"log"
	"net/http"	
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func fetch(sensor_id string) float64 {
	
	/* Fetch the inlet temperature running shell commands */

	sensor := exec.Command("/usr/bin/ipmitool", "sdr", "get", sensor_id)
	output,_ := sensor.Output() 

	// regex to match
	re,_ := regexp.Compile("Sensor Reading.*")
	inlet := strings.Fields(re.FindString( string(output) ))[3]
	// convert the match into a float
	inlet_temp,_ := strconv.ParseFloat(inlet, 64)
	
	return inlet_temp
}

type tempCollector struct {

	/* Define a structure for the collector */

	sensorID string
	tempMetric *prometheus.Desc
}

func (collector *tempCollector) Describe(ch chan<- *prometheus.Desc) {

	/* Implement the Describe function for the collector,
	which essentialy writes the descriptors to the desc channel */

	ch <- collector.tempMetric
}

func (collector *tempCollector) Collect(ch chan <- prometheus.Metric) {
	
	/* Implement the Collect function for the collector,
	which runs the logic to determine the value of the metric */

	// fetch the metric
	metric := fetch(collector.sensorID)
	
	// write the latest value for the metric in the metric channel
	metric_latest := prometheus.MustNewConstMetric(collector.tempMetric, prometheus.GaugeValue, metric)
	ch <- metric_latest
}

func newTempCollector(sensorID string) *tempCollector{

	/* Initialize the descriptor and return a pointer to the collector */

	return &tempCollector{
		sensorID: sensorID,
		tempMetric: prometheus.NewDesc("ipmitool_temp", "Inlet Temperature", nil, nil),
	}
}

func main() {

	/* Register the metric and start a httpd server to expose it */
	
	// command line arguments
	var (
		sensorID = flag.String("sensor", "System Temp",
		"Sensor ID to fetch the Inlet Temperature with ipmitool")
		listenAddress = flag.String("address", ":8000",
		"Address to listen on for this exporter")
		metricsPath = flag.String("path", "/metrics",
		"Path under which to expose metrics")
		)
	flag.Parse()

	// create and register the metric
	inlet_temperature := newTempCollector(*sensorID)
	prometheus.MustRegister(inlet_temperature)

	// expose the metric
	http.Handle(*metricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
