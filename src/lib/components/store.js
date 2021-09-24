import { writable } from 'svelte/store';

const store = writable({
    circlesize: 50,
    otherdata: "gugus"
});

export default store;