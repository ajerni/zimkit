<!-- src/routes/__layout.svelte wird für alle pages angewandt -->
<!-- https://rodneylab.com/getting-started-with-sveltekit/ -->

<script
	src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
	integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
	crossorigin="anonymous">
</script>

<title>Andi's Toolkit</title>

    <div class="bg-dark text-secondary px-4 py-5 text-center">
      <div class="py-5">
        <h1 class="display-5 fw-bold text-white">Andi's Web-Kit</h1>
        <div class="col-lg-6 mx-auto">
          <p class="fs-5 mb-4">
            Svelte Frontend - Go serverless functions backend
          </p>
		  <p class="fs-5 mb-4">
            Reference. Template. Toolkit. Playground.
          </p>
          <div class="d-grid gap-2 d-sm-flex justify-content-sm-center">


            <a href="/"><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold">Home</button></a>
			<a href="/zim"><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold">ZIM</button></a>
            <a href="/error"><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold">Error</button></a>
			<a href="/temperatur"><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold">Sensor</button></a>
			<a href="/form"><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold">Form</button></a>
			<a href="/heroku"><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold">Database</button></a>
			<a href="/ajax"><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold">AJAX</button></a>
	
			
			<a href="https://github.com/ajerni/zimkit/"
              ><button type="button" class="btn btn-outline-light btn-lg px-4">
                Code
              </button></a
            >
          </div>
        </div>
      </div>
    </div>

<main>
	<slot />
</main>

<footer class="d-grid gap-2 d-sm-flex justify-content-sm-center">
	
	<p>A project by <a href="https://www.andierni.ch/" target="_blank" rel="nofollow noopener noreferrer">ERNILABS</a>.</p>

</footer>

<svelte:head>
	<link
		href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css"
		rel="stylesheet"
		integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU"
		crossorigin="anonymous"
	/>
</svelte:head>
