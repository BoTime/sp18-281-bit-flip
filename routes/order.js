var express = require('express');
var router = express.Router();
var ejs = require('ejs');
var proxy = require('express-http-proxy');
const JwtUtils = require('../utils/JwtToken');

const goAPI =  process.env.KONG_URL;
// Return order page
router.post('/', JwtUtils.attachTokenToHeader, proxy(goAPI,{
		proxyReqPathResolver: function(req) {
			console.log("ORDER POST");
			console.log(req.body);
			return require('url').parse(req.url).path + 'orders/v1/order';
		},
		userResDecorator: function(proxyRes, proxyResData, userReq, userRes) {
			console.log("Back",proxyRes);
		    data = JSON.parse(proxyResData.toString('utf8'));
		   	console.log('status code====', proxyRes.statusCode);
			if (proxyRes.statusCode === 200 || proxyRes.statusCode === 201) {
				console.log("Sucess");
				// Order updated sucessfully
				//userRes.statusCode = 201;
				userRes.redirect('created');

			} else if (proxyRes.statusCode === 401 || proxyRes.statusCode === 400 || proxyRes.statusCode === 404) {
				// Order placing failed, redirect to signin page
				console.log("400");
				userRes.statusCode = 401;
				userRes.redirect('signin');
			}else {
				// Order placing failed, redirect to oops page
				console.log("500");
				userRes.statusCode = 500;
				userRes.redirect('oops');
			}
		    return userRes;
	  	}
	})
);

// Return order page
router.get('/', JwtUtils.attachTokenToHeader, (req, res) => {
	console.log("order get",req);
	res.render('order');
});

module.exports = router;
