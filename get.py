#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import requests

def get_data():
    url = 'http://localhost:8080/'
    headers = {'Content-Type': 'application/json'}

    x = requests.get(url, headers=headers)
    print(x.text)

get_data()