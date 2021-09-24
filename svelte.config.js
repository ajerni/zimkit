//import adapter from '@sveltejs/adapter-static';
import adapter from '@sveltejs/adapter-netlify';

export default {
	kit: {
		target: '#svelte',
		vite: {
			define: {
				'process.env': process.env
			}
		},
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			fallback: null
		})
	}
};
