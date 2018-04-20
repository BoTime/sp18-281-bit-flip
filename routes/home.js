const express = require('express');
const router = express.Router();

//Controller to render application home page
router.get('/', (req, res) => {
  res.render('home', { name: 'Bob Marley' })
});

module.exports = router;
