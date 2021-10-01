<script>

import store from '$lib/components/store.js'

let val = "";
let resp = "";

//updates the store whenever val changes (einfach weil val rechts gebraucht wird - $: dynamschee Werte oder wie hier ganze Funktionen!)
$: store.update((data) => {
					return {
                        circlesize: data.circlesize,
                        otherdata: data.otherdata,
                        startdate: val
					};
				});

function handleSubmit(){

    let form = document.getElementById("myForm");
    let formData = new FormData(form);

    fetch('https://zimkit.vercel.app/api/form-json', {
        method: "post",
        body: formData,
    })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            console.log(data.message);
            resp = data.sonstiges;
        })
}

function handle2(){
	let form = document.getElementById("myForm");
    let formData = new FormData(form);

	fetch('https://zimkit.vercel.app/api/form-redirect')
			.then((res) => {
				if (!res.ok) {
					throw new Error('Failed!');
				}
				return res.json();
			})
			.then((data) => {
				console.log(data);
			})
			.catch((err) => {
				console.log(err);
			});
}

</script>

<div class="container">
    <h1>Form</h1>
    <form id="myForm" on:submit|preventDefault={handleSubmit}>
	<!-- <form action="https://zimkit.vercel.app/api/form" method="post"> -->
		<div class="mb-3">
			<label for="date-from" class="form-label">From:</label>
			<input
				type="date"
				class="form-control"
				id="date-from"
                name="start"
				aria-describedby="dateFromHelp"
                bind:value="{val}"
			/>
		</div>
		<div class="mb-3">
			<label for="date-to" class="form-label">To:</label>
			<input type="date" class="form-control" id="date-to" name="end"/>
		</div>
		<button type="submit" class="mb-4 btn btn-primary">Submit</button>
		<button class="mb-4 btn btn-primary" on:click={handle2}>Submit with redirect</button>
	</form>
	
</div>

<h3>See the start date again: {val}</h3>
<h3>Response aus JSON: {resp}</h3>


