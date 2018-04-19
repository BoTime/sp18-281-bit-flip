var express = require('express');
var router = express.Router();

//Controller to render sign-in page
router.post('/',
	function(req, res)
	{
		console.log("SignIn User data:", req.body.email,req.body.password);
		res.redirect('/index.html');
	}
);

// Return sign in page
router.get('/', (req, res) => {
	res.render('signin');
});

module.exports = router;
