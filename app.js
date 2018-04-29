'use strict';
// Load environment variables
require('dotenv').config()

console.log('[*] KONG URL: ' + process.env.KONG_URL);

const bodyParser = require('body-parser');
const express = require('express');
var session = require('express-session');
var cookieSession = require('cookie-session');
var cookieParser = require('cookie-parser');
const path = require('path');
const ejs = require('ejs');
const lineReader = require('line-reader');
const querystring = require('querystring');
const monk = require('monk');
const randomstring = require("randomstring");
const morgan = require('morgan')

const signinRouter = require('./routes/sign-in');
const signupRouter = require('./routes/sign-up');
const homeRouter = require('./routes/home');
const indexRouter = require('./routes/index');
const orderRouter = require('./routes/order');
const ordersRouter = require('./routes/orders');
const ordersHistory = require('./routes/history');
const oops = require('./routes/oops');
const created = require('./routes/created');
// Create the app.
var app = express();
// app.use(morgan('tiny'));
/*
app.use(cookieSession({
    secret: 'post-it',
    name: 'session',
    keys: [randomstring.generate()],
    // Cookie Options
    maxAge: 24 * 60 * 60 * 1000 // 24 hours
}));
*/

//To store valid user credentials
var valid_password="xxxx";
var valid_user="xxxx";


// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

// Use the bodyParser() middleware for all routes.
// parse json
app.use(express.static(path.join(__dirname, 'public')));
app.use(bodyParser.urlencoded({
    extended: true
}));
// parse form
const upload = require('multer')();
app.use(cookieParser());
app.use(upload.array());

app.use('/', indexRouter);
app.use('/signin', signinRouter);
app.use('/signup', signupRouter);
app.use('/logout', (req, res) => res.redirect('/signin'));
app.use('/home', homeRouter);
app.use('/order', orderRouter);
app.use('/orders', ordersRouter);
app.use('/history', ordersHistory);
app.use('/oops', oops);
app.use('/created', created);

var port = process.env.PORT || 8000;
app.listen(port);
console.log("Listening on port 8000");
