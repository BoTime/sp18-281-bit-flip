var express = require('express');
var router = express.Router();

// Return order page
router.get('/', (req, res) => {
	res.render('created');
});

module.exports = router;
 