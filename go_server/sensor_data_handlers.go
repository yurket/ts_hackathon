package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type Sensor struct {
	Time     string `json:"time"`
	Value string `json:"value"`
}

var sensor_data  = []Sensor{Sensor{Time : "time", Value : "bestval"}}


func getSensorDataHandler(w http.ResponseWriter, r *http.Request) {
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

func putSensorDataHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of Bird
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
	sensor.Time = r.Form.Get("time")
	sensor.Value = r.Form.Get("value")

	// Append our existing list of birds with a new entry
	sensor_data = append(sensor_data, sensor)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}