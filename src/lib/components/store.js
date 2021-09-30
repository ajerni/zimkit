import { writable } from 'svelte/store';

const store = writable({
	circlesize: 50,
	otherdata: 'gugus',
	startdate: 'not set yet'
});

export default store;
