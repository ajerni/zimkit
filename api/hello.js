module.exports = (req, res) => {
	const { name = 'World' } = req.query;
	res.status(200).send(`Hello ${name}!`);
};

// https://zimkit.vercel.app/api/hello?name=andi