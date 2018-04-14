'use strict';
const bodyParser = require('body-parser');
const express = require('express');
var session = require('express-session');
var cookieSession = require('cookie-session');
var cookieParser = require('cookie-parser');
const path = require('path');
const ejs = require('ejs');
const fs = require('fs');
const lineReader = require('line-reader');
const querystring = require('querystring');
const monk = require('monk');
//var db = ;

var randomstring = require("randomstring");
var sign_in = require('./routes/sign-in');
var sign_up = require('./routes/sign-up');
var index = require('./routes/index');


// Create the app.
var app = express();
/*
app.use(cookieSession({
    secret: 'post-it',
    name: 'session',
    keys: [randomstring.generate()],
    // Cookie Options
    maxAge: 24 * 60 * 60 * 1000 // 24 hours
}));
*/

var port = 8000;
app.listen(port);
console.log("Listening on port 8000");

/*app.use(function(req, res, next)
    {
        req.db = db;
        next();
    }
);*/



var html_file_name ='./public/index.html';

//To store valid user credentials
var valid_password="xxxx";
var valid_user="xxxx";



// view engine setup
//app.set('views', path.join(__dirname, 'views'));
//app.set('view engine', 'jade');

// Use the bodyParser() middleware for all routes.
app.use(express.static("public"));
app.use(bodyParser());
app.use(cookieParser());

app.use('/', index);
app.use('/sign-in', sign_in);
app.use('/sign-up', sign_up);

