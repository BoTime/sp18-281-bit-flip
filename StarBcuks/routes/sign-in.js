var express = require('express');
var router = express.Router();

//Controller to render sign-in page
router.post('/',
	function(req, res)
	{	
		console.log("SignIn User data:", req);
		res.redirect('/index.html');
	}
);

module.exports = router;
