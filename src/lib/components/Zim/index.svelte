<script>
	import { onMount } from 'svelte';
	import store from '$lib/components/store.js';
	import callVercel from '$lib/components/utils/fetch.js';

	// let nodeRef;
	// nodeRef.parentNode.removeChild(nodeRef);

	let mounted = false;

	//wird dank export direkt zugänglich gemacht in Parent Komponente
	export function startZim() {
		if (mounted) {
			loadZim();
		}
	}

	onMount(() => {
		mounted = true;
		//loadZim(); ZIM könnte hier direkt geladen werden (mounted braucht es hier nur für den späteren Start über den button)
	});

	//Beispiel: sizeChanged wird immer aufgerufen, soblad $store.circlesize ändert
	$: {
		sizeChanged($store.circlesize);
	}

	function sizeChanged(val) {
		console.log(`size has changed to ${$store.circlesize}`);
		console.log('size has changed to (von argument): ' + val);
		//...do whatever you want here...
	}

	//ZIM START
	function loadZim() {
		var scaling = 'holder'; // use the ID of a tag to place the canvas in the tag
		var frame = new Frame({ scaling, allowDefault: true });
		frame.on('ready', function () {
			zog('ready from ZIM Frame');

			var stage = frame.stage;
			var stageW = frame.width;
			var stageH = frame.height;
			frame.color = pink;

			// put your ZIM code here (you have access to everything within this script tag i.e. the store object)

			//$store macht automatisch subscribe() und unsubscribe() - siehe Svelte store writable
			var circle = new Circle($store.circlesize, red)
				.addTo()
				.pos(stageW / 2 - 40, 70)
				.drag();

			var button = new Button(140, 60, $store.otherdata, frame.grey, frame.dark);
			// will position and scale the code in the resize function at bottom
			button.on('click', function () {
				store.update((data) => {
					return {
						circlesize: data.circlesize,
						otherdata: 'RESET FROM ZIM BUTTON'
					};
				});
				console.log($store.otherdata);
				//some Particle fun with go serverless functions on vercel:
				callVercel('https://zimkit.vercel.app/api/postevent?key=alarmall');
				callVercel('https://zimkit.vercel.app/api/post?f=setmessage&a=' + $store.otherdata);
			});

			// put resizing code in here
			frame.on('resize', resize);
			function resize(e) {
				stageW = frame.width;
				stageH = frame.height;
				button.scaleTo(stage, 30); // also see zim.Layout() for more complex scaling
				button.center(); // centers on stage - addTo() and others as well!
				stage.update();
			}
		}); // end of ready
	} //END ZIM
</script>

<svelte:head>
	<script src="https://zimjs.org/cdn/nft/00/zim.js"></script>
</svelte:head>

<!-- <div id="holder" bind:this={nodeRef} /> -->
<div id="holder" />

<style>
</style>
