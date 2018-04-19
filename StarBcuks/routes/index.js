var express = require('express');
var router = express.Router();
var html_file_name ='./public/index.html';
var path = require("path");

//Controller to render application home page
router.get('/', (req, res) => {
    res.render('index');
});

router.get('/index', (req, res) => {
    res.render('index');
});

router.get('/menu', (req, res) => {
    res.render('menu');
});

router.get('/contact', (req, res) => {
    res.render('contact');
});
module.exports = router;
