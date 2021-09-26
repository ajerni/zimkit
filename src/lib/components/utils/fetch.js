let url;

//GET Request (used to call serverless functions on own api on Vercel)
function myFetch(url){

    fetch(url)
	.then((res) => {
		if (!res.ok) {
			throw new Error('Failed!');
		}
		return res.json();
	})
	.then((data) => {
		console.log(data.result);
		return data.result;
	})
	.catch((err) => {
		console.log(err);
	})

} 

export default myFetch;
