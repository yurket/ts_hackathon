#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import requests
import random
import time
import argparse

def randint():
    return int(random.random() * 10)

def send_data():
    url = 'http://localhost:8080/'
    myobj = f'{{"Time":"{time.time()}","Temperature":"{randint()}","Pressure":"{randint()}","Humidity":"{randint()}"}}'
    headers = {'Content-Type': 'application/json'}
    x = requests.post(url, data=myobj, headers=headers)
    print(x.status_code)

def get_data():
    url = 'http://localhost:8080/'
    headers = {'Content-Type': 'application/json'}
    x = requests.get(url, headers=headers)
    print(x.status_code, x.text)

parser = argparse.ArgumentParser()
parser.add_argument('-c', choices=['send', 'get'])
args = parser.parse_args()

if args.c == "get":
    get_data()
elif args.c == 'send':
    send_data()