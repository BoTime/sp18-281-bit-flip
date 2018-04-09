var express = require('express');
var router = express.Router();

//Controller to render sign-out page
router.post('/',
	function(req, res)
	{	
		console.log("User data:", req.param);
		console.log("User data:", req.body.email, req.body.password1,req.body.name);
		res.redirect('/index.html');
	}
);

module.exports = router;
