<!-- https://blog.hdks.org/Environment-Variables-in-SvelteKit-and-Vercel/  Achtung: bei import nicht Dateiendung .js anhängen!-->
<script>
	import { MY_API_KEY, MY_DEVICE } from '$lib/Env';
	import Code from '$lib/components/Argon/code.svelte';

	let myApiKey;
	let myDevice;

	let temp = 0;
	let hum = 0;

	if (process.env.NODE_ENV === 'production') {
		// For production
		myApiKey = process.env.MY_API_KEY;
		myDevice = process.env.MY_DEVICE;
	} else {
		// For development
		myApiKey = MY_API_KEY;
		myDevice = MY_DEVICE;
	}

	//GET temperature (um auf Particle.variable zuzugreifen)
	// let variable = 'temperature';
	// fetch(
	// 	'https://api.particle.io/v1/devices/' + myDevice + '/' + variable + '?access_token=' + myApiKey
	// )
	// 	.then((res) => {
	// 		if (!res.ok) {
	// 			throw new Error('Failed!');
	// 		}
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		console.log('Aktueller Wert: ' + data.result);
	// 		temp = data.result;
	// 	})
	// 	.catch((err) => {
	// 		console.log(err);
	// 	});
	// //GET humidity
	// let variable2 = 'humidity';
	// fetch(
	// 	'https://api.particle.io/v1/devices/' + myDevice + '/' + variable2 + '?access_token=' + myApiKey
	// )
	// 	.then((res) => {
	// 		if (!res.ok) {
	// 			throw new Error('Failed!');
	// 		}
	// 		return res.json();
	// 	})
	// 	.then((data) => {
	// 		console.log('Aktueller Wert: ' + data.result);
	// 		hum = data.result;
	// 	})
	// 	.catch((err) => {
	// 		console.log(err);
	// 	});

	function getTempHum() {
		//Aufruf (GET Request) der Server less function auf vercel, welche dann den GET Request macht zur Particle.variable
		fetch('https://zimkit.vercel.app/api/get?key=temperature')
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

		//GET Particle.variable humidity via serverless function get on Vercel:
		fetch('https://zimkit.vercel.app/api/get?key=humidity')
			.then((res) => {
				if (!res.ok) {
					throw new Error('Failed!');
				}
				return res.json();
			})
			.then((data) => {
				console.log('Aktueller Wert: ' + data.result);
				hum = data.result;
			})
			.catch((err) => {
				console.log(err);
			});
	}
</script>

<h1>Temperatur-Sensor an Particle Argon</h1>

<h2>Temperatur: {temp} °C</h2>
<h2>Humidity: {hum} %</h2>

<button on:click={getTempHum}>Get Data</button>

<br />
<br />

<Code />
