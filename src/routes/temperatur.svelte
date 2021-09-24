<script>
	import { MY_API_KEY, MY_DEVICE } from '$lib/Env';

	let myApiKey;
	let myDevice;

	let temp = 10;

	if (process.env.NODE_ENV === 'production') {
		// For production
		myApiKey = process.env.MY_API_KEY;
		myDevice = process.env.MY_DEVICE;
	} else {
		// For development
		myApiKey = MY_API_KEY;
		myDevice = MY_DEVICE;
	}

	//GET (um auf Particle.variable zuzugreifen)
	let variable = 'temperature';
	fetch(
		'https://api.particle.io/v1/devices/' + myDevice + '/' + variable + '?access_token=' + myApiKey
	)
		.then((res) => {
			if (!res.ok) {
				throw new Error('Failed!');
			}
			return res.json();
		})
		.then((data) => {
			console.log('Aktueller Wert: ' + data.result);
			temp = data.result;
		})
		.catch((err) => {
			console.log(err);
		});
</script>

<h1>Temperatur-Sensor an Particle Argon</h1>

<h2>Temperatur: {temp}</h2>
