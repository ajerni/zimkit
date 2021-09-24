# https://api.particle.io/v1/devices/' + myDevice + '/' + variable + '?access_token=' + myApiKey

import requests.get
# ploads = {'things':2,'total':25}
r = requests.get('https://api.particle.io/v1/devices/e00fce6839e00714b083442d/temperature?access_token=906d5e4a9041e4c0773cad80ccf23490fe83e76c')
print(r.text)
print(r.url)