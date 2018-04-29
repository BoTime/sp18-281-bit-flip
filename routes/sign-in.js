const express = require('express');
const router = express.Router();
const proxy = require('express-http-proxy');
const RequestModifier = require('../utils/RequestModifier');

const KONG_API_GATEWAY_URL = process.env.KONG_URL;

const testMiddleware = function(req, res, next) {
	next();
}

router.post('/', RequestModifier, proxy(KONG_API_GATEWAY_URL,{
		// proxyReqOptDecorator: function(proxyReqOpts, srcReq) {
		//     // you can update headers
		//     proxyReqOpts.headers['Content-Type'] = 'application/json';
		//     return proxyReqOpts;
		// },
		proxyReqPathResolver: function(req) {
			let newUrl = '';
			console.log('request body====', req.body);
			if (KONG_API_GATEWAY_URL.indexOf('localhost') !== -1) {
				// request through local server
				newUrl = require('url').parse(req.url).path + 'signin';
			} else {
				// request through Kong API Gateway
				newUrl = require('url').parse(req.url).path + 'users/v1/signin';
			}
			console.log('=====', newUrl)
			return newUrl;
		},
		userResDecorator: function(proxyRes, proxyResData, userReq, userRes) {
		    data = JSON.parse(proxyResData.toString('utf8'));
		   	console.log('status code====', proxyRes.statusCode);

			if (proxyRes.statusCode === 200) {
				// Login success, redirect to home page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/home');

			} else if (proxyRes.statusCode === 401 || proxyRes.statusCode === 400 || proxyRes.statusCode === 510 || proxyRes.statusCode === 404) {
				// Login failed, redirect to signin page
				console.log('$$$$$$$$');
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/signin');

			} else {
				console.log('********');
			}
		    return JSON.stringify(data);
	  	}
	})
);

// Return sign in page
router.get('/', (req, res) => {
	res.render('signin');
});

module.exports = router;
