import json
import smbus2
import bme280
import time
import os
import requests

API_ENDPOINT = "http://10.0.4.138:8080/sensor_post"

port = 1
address = 0x76
bus = smbus2.SMBus(port)

calibration_params = bme280.load_calibration_params(bus, address)

# the sample method will take a single reading and return a
# compensated_reading object


while True:
    data = bme280.sample(bus, address, calibration_params)
    

# the compensated_reading class has the following attributes
    #print(data.id)
    ticks = time.time()
    print('Time: ',ticks)
    print('Temperature: ', round(data.temperature, 1))
    print('Pressure: ', round(data.pressure, 1))
    print('Humidity: ', round(data.humidity, 1))


    data_frame = {'Time': ticks,
                  'Temperature': round(data.temperature, 2),
                  'Pressure': round(data.pressure, 1),
                  'Humidity': round(data.humidity, 1)
                  }
    
    

    data_json = json.dumps(data_frame)
    r = requests.post(url = API_ENDPOINT, data = data_frame)

# there is a handy string representation too
    #print(data)

    time.sleep(1)
    os.system('clear')