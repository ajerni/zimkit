<script
	src="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/js/bootstrap.bundle.min.js"
	integrity="sha384-/bQdsTh/da6pkI1MST/rWKFNjaCP5gBSY4sEBT38Q/9RBh9AH40zEOg7Hlq2THRZ"
	crossorigin="anonymous"></script>

<div class="container">
    <h1>Form</h1>
	<form action="https://zimkit.vercel.app/api/form" method="post">
		<div class="mb-3">
			<label for="date-from" class="form-label">From:</label>
			<input
				type="date"
				class="form-control"
				id="date-from"
                name="start"
				aria-describedby="dateFromHelp"
			/>
		</div>
		<div class="mb-3">
			<label for="date-to" class="form-label">To:</label>
			<input type="date" class="form-control" id="date-to" name="end"/>
		</div>
		<button type="submit" class="mb-4 btn btn-primary">Submit</button>
	</form>
</div>

<svelte:head>
	<link
		href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.1/dist/css/bootstrap.min.css"
		rel="stylesheet"
		integrity="sha384-F3w7mX95PdgyTmZZMECAngseQB83DfGTowi0iMjiWaeVhAn4FJkqJByhZMI3AhiU"
		crossorigin="anonymous"
	/>
</svelte:head>
