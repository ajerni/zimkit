require 'httparty'

response = HTTParty.get('https://api.particle.io/v1/devices/e00fce6839e00714b083442d/temperature?access_token=906d5e4a9041e4c0773cad80ccf23490fe83e76c')
puts response.body if response.code == 200
p response