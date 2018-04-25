var express = require('express');
var router = express.Router();
var ejs = require('ejs');

// Return order page
router.get('/', (req, res) => {
	res.render('oops');
});

module.exports = router;
 