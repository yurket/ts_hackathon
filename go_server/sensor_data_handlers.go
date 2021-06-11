package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Sensor struct {
	Time        string `json:"Time"`
	Temperature string `json:"Temperature"`
	Pressure    string `json:"Pressure"`
	Humidity    string `json:"Humidity"`
}

var sensor_data = []Sensor{}

func getSensorDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//Convert the "birds" variable to json
	sensorListBytes, err := json.Marshal(sensor_data)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(sensorListBytes)
}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

func putSensorDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(formatRequest(r))

	sensor := Sensor{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the bird from the form info
	sensor.Time = r.Form.Get("Time")
	sensor.Temperature = r.Form.Get("Temperature")
	sensor.Pressure = r.Form.Get("Pressure")
	sensor.Humidity = r.Form.Get("Humidity")

	// Append our existing list of birds with a new entry
	sensor_data = append(sensor_data, sensor)
	fmt.Println(sensor_data)
}
