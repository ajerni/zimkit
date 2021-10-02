<script>
	import store from '$lib/components/store.js';
	import Zim from '$lib/components/Zim/index.svelte';
	import callVercel from '$lib/components/utils/fetch.js';

	let wert = $store.otherdata;
	let sizeA = $store.circlesize;
	let resp;

	$: {
		console.log($store.otherdata); //just used to trigger this "function" whenever $store.otherdata changes
		resp = $store.otherdata;
	}

	let zimcomponent;

	function makeZim(e) {
		store.update((data) => {
			return {
				circlesize: sizeA,
				otherdata: wert
			};
		});

		//direkter Funktionsaufruf möglich dank binding mit :this
		zimcomponent.startZim();
	}
</script>

<h1>ZIM als Svelte Komponente:</h1>

<h2>Size: {sizeA}</h2>
<h2>Wert: {wert}</h2>

<label for="size">Size:</label>
<input type="number" bind:value={sizeA} label="size" />

<label for="input">Wert:</label>
<input id="wert" type="text" bind:value={wert} label="input" />

{#if resp != 'gugus'}
	<h3>{resp}</h3>
	<!-- <a href="https://zimkit.vercel.app/api/redirect"><p>Also see Photon message changed on IOT Dashboard</p></a> -->
	<!-- Umweg über serverless go function rediredt zu Demo Zwecken (Link geht auf iot.andierni.ch) -->
	<p on:click={() => callVercel('https://zimkit.vercel.app/api/redirect')}>
		Also see Photon message changed on IOT Dashboard>
	</p>
{/if}

<button on:click={makeZim}>Make ZIM</button>

<!-- erlaubt direkten Funktionsaufruf (anstelle von on:event) -->
<Zim bind:this={zimcomponent} />
