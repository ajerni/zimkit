<script>
	import { CodeSnippet } from 'carbon-components-svelte/src';
	import argon from './argon.jpg';
	

	let code = `
#include <Grove_Temperature_And_Humidity_Sensor.h>

#define DHTPIN 	D2

DHT dht(DHTPIN);

float t;
char strT[16];

void setup()
{
	dht.begin();
	Spark.variable("temperature", strT);
}

void loop()
{
	delay(10000);
	t = dht.getTempCelcius();
    snprintf(strT, sizeof(strT), "%.2f", t);
}
`;
</script>

<img src={argon} alt="Argon"/>

<h3>The code running on Argon (C++)</h3>
<CodeSnippet type="multi" copy={() => {}} {code} hideCopyButton wrapText />

<style>
	img {
		width: 600px;
		height: 400px;
	}
</style>
