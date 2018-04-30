var express = require('express');
var router = express.Router();
var html_file_name ='./public/index.html';
var path = require("path");
const LocalStorage = require('node-localstorage').LocalStorage;
const localStorage = new LocalStorage('./token');


//Controller to render application home page
router.get('/', (req, res) => {
    res.render('index');
});

router.get('/index', (req, res) => {
    // let name = localStorage.getItem('name');
    // if (name !== undefined) name = name.toUpperCase();
    // res.render('index', { name: name });
    res.render('index');
});

router.get('/menu', (req, res) => {
    res.render('menu');
});

router.get('/contact', (req, res) => {
    res.render('contact');
});
module.exports = router;
