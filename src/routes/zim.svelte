<script>
	import Test from '$lib/components/Test/index.svelte';
	import store from '$lib/components/store.js';
</script>

<h1>Welcome to SvelteKit</h1>

<Test />

<h3>Start date: {$store.startdate}</h3>

<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
