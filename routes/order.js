var express = require('express');
var router = express.Router();
var ejs = require('ejs');

// Return order page
router.post('/', (req, res) => {
	console.log("ORDER POST");
	console.log(req.body);
	res.render('order');
});
// Return order page
router.get('/', (req, res) => {
	res.render('order');
});

module.exports = router;
 