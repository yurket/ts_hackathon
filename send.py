#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import requests
import random
import argparse

def randint():
    return int(random.random() * 10)

def send_data():
    url = 'http://localhost:8080/'
    myobj = f'{{"Time":"{randint()}","Temperature":"{randint()}","Pressure":"{randint()}","Humidity":"{randint()}"}}'
    headers = {'Content-Type': 'application/json'}

    x = requests.post(url, data=myobj, headers=headers)
    print(x.text)

send_data()
