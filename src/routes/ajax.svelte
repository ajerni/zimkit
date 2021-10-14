<script>
	var show = false;
	var user = '';

	var xhr; //Ajax Objekt

	$: {
		console.log(user);
		xhr = new XMLHttpRequest();
		xhr.open('POST', 'https://zimkit.vercel.app/api/ajax', true);
		xhr.addEventListener('readystatechange', function () {
			//callback function für Ajax
			if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
				var resp = xhr.responseText; //wird true oder false retournieren --> siehe api/ajax.go
                console.log(resp);
				if (resp === 'true') {
					show = true;
				} else {
					show = false;
				}
			}
		});
		xhr.send(user);
	}

</script>

<p>Checkt mit jeder Eingabe über AJAX request, ob username schon vergeben ist.</p>
<p>"Andi" + "Max" sind die einzigen vergebenen Namen in diesem Beispiel.</p>

<form method="post" id="form-create-user">
	<input type="text" id="username" name="username" placeholder="username" bind:value={user} />
</form>
<br />
<p>Username: {user}</p>
<br />
{#if show}
	<p class="warning">Username {user} ist bereits vergeben</p>
{/if}

<style>
	.warning {
		color: red;
	}
</style>
