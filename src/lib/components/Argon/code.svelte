<script>
//temptinker code in build.particle.io
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
`
</script>

<h2>The code running on Argon (C++)</h2>
<h3>...show code here mit ibm library</h3>