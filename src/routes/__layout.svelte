<!-- src/routes/__layout.svelte wird für alle pages angewandt -->
<!-- https://rodneylab.com/getting-started-with-sveltekit/ -->

<script
	src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
	integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
	crossorigin="anonymous">
</script>

<header>
	<nav>
		<a href="/.">Home</a>
		<a href="/error">Error</a>
		<a href="/temperatur">Temperatur</a>
		<a href="/form">Form</a>
		<a href="/heroku">DB Heroku</a>
	</nav>
</header>
<main>
	<button><a href="https://zimkit.vercel.app/api/dbadd">Add temp</a></button>
	<slot />
</main>
<footer>
	<p>
		A project by
		<a
			aria-label="Open the Rodney Lab site"
			href="https://www.andierni.ch/"
			target="_blank"
			rel="nofollow noopener noreferrer">ERNILABS</a
		>.
	</p>
</footer>

<svelte:head>
	<link
		href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css"
		rel="stylesheet"
		integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU"
		crossorigin="anonymous"
	/>
</svelte:head>
