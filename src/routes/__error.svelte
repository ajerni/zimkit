<script context="module">
	export function load({ error, status }) {
		return {
			props: {
				title: `${status}: ${error.message}`
			}
		};
	}
</script>

<script>
	export let title;
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>
<h1>Ooopsy... ...this is my standard error page.</h1>
<h2>Error message: {title}</h2>
