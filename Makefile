sensor = System Temp

ipmitool_exporter:
	go build -ldflags="-X 'main.iPMI_TEMP_SENSOR=$(sensor)'" -o ipmitool_exporter main.go
