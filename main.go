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

type tempCollector struct {

	tempMetric *prometheus.Desc

}

func (collector *tempCollector) Describe(ch chan<- *prometheus.Desc) {

	ch <- collector.tempMetric

}

func (collector *tempCollector) Collect(ch chan <- prometheus.Metric) {

	/* Collect the metric */
	metric := fetch()
	
	//Write latest value for each metric in the prometheus metric channel.
	metric_latest := prometheus.MustNewConstMetric(collector.tempMetric, prometheus.GaugeValue, metric)
	ch <- metric_latest
}

func newTempCollector() *tempCollector{

	return &tempCollector{
		tempMetric: prometheus.NewDesc("my_inprogress_request", "Inlet Temperature", nil, nil),
	}

}

func main() {

	// Create and register the metrics
	inlet_temperature := newTempCollector()
	prometheus.MustRegister(inlet_temperature)

	// Expose the metrics
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8000", nil))

}