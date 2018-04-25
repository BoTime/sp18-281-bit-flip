var express = require('express');
var router = express.Router();

// Return order page
router.get('/', (req, res) => {
	console.log("HSITORY CALLED");
	res.render('history');
});

module.exports = router;
 