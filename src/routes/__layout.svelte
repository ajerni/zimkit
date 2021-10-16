<!-- Bootstrap import -->
<script>
	import { Accordion, AccordionItem, Collapse, Button, Card } from 'sveltestrap';
	let isOpen = false;
	('src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"');
</script>

<!-- src/routes/__layout.svelte wird für alle pages angewandt -->

<title>Andi's Toolkit</title>

<div class="bg-dark text-secondary px-4 py-5 text-center">
	<div class="py-5">
		<h1 class="display-5 fw-bold text-white">Andi's Web-Kit</h1>
		<div class="col-lg-6 mx-auto">
			<p class="fs-5 mb-4 text-white">Svelte Frontend - Go backend</p>
			<p class="fs-5 mb-4">Reference. Template. Toolkit. Playground.</p>
			<div class="custom-class-big">
				<div class="d-grid gap-2 d-sm-flex justify-content-sm-center">
					<a href="/"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Home</button
						></a
					>
					<a href="/zim"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>ZIM</button
						></a
					>
					<a href="/error"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Error</button
						></a
					>
					<a href="/temperatur"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Sensor</button
						></a
					>
					<a href="/form"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Form</button
						></a
					>
					<a href="/heroku"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Database</button
						></a
					>
					<a href="/ajax"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>AJAX</button
						></a
					>

					<a href="https://github.com/ajerni/zimkit/"
						><button type="button" class="btn btn-outline-light btn-lg px-4 custom">
							Code
						</button></a
					>
				</div>
			</div>

			<div class="custom-class">
				<Button class="mb-3" color="primary" on:click={() => (isOpen = !isOpen)}>Menu</Button>
				<Collapse {isOpen}>
					<a href="/"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Home</button
						></a
					>
					<a href="/zim"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>ZIM</button
						></a
					>
					<a href="/error"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Error</button
						></a
					>
					<a href="/temperatur"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Sensor</button
						></a
					>
					<a href="/form"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Form</button
						></a
					>
					<a href="/heroku"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>Database</button
						></a
					>
					<a href="/ajax"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>AJAX</button
						></a
					>

					<a href="https://github.com/ajerni/zimkit/"
						><button type="button" class="btn btn-outline-light btn-lg px-4 custom">
							Code
						</button></a
					>
					<a class="noshow" href="/ajax"
						><button type="button" class="btn btn-outline-info btn-lg px-4 me-sm-3 fw-bold custom"
							>empty</button
						></a
					>
				</Collapse>
			</div>
		</div>
	</div>
</div>

<main>
	<div class="min-vh-100 w-100 p-4">
		<slot />
	</div>
</main>

<footer>
	<div class="bg-dark text-secondary py-4 text-center">
		A project by <a
			href="https://www.andierni.ch/"
			target="_blank"
			rel="nofollow noopener noreferrer">ERNILABS</a
		>.
	</div>
</footer>

<svelte:head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta http-equiv="ScreenOrientation" content="autoRotate:disabled">

	<link
		rel="stylesheet"
		href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css"
	/>
	<link
		rel="stylesheet"
		href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.5.0/font/bootstrap-icons.css"
	/>
</svelte:head>

<style>
	.custom {
		width: 8rem !important;
	}

	.noshow {
		visibility: hidden;
		/* display: none; nimmte Element aus dem Fluss / visibility: hidden lässt es dort stehen, aber zeigt es einfach nicht an. Gut für Plazierung.*/
	}

	.custom-class {
		display: none;
	}

	@media only screen and (max-width: 576px) {
		.custom-class {
			display: block;
		}
	}

	.custom-class-big {
		display: none;
	}

	@media only screen and (min-width: 576px) {
		.custom-class-big {
			display: block;
		}
	}
</style>
