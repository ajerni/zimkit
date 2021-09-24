//import adapter from '@sveltejs/adapter-static';
		// adapter: adapter({
		//	pages: 'build',
		//	assets: 'build',
		//	fallback: null
	
import adapter from '@sveltejs/adapter-netlify';
 
const config = {
  kit: {
    adapter: adapter(),
    // hydrate the <div id="svelte"> element in src/app.html
    target: '#svelte'
	},
	vite: {
		define: {
			'process.env': process.env
		}
	},
};
 
export default config
