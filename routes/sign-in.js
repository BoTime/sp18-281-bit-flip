const express = require('express');
const router = express.Router();
const proxy = require('express-http-proxy');
const RequestModifier = require('../utils/RequestModifier');
const JwtUtils = require('../utils/JwtToken');
const LocalStorageUtils = require('../utils/LocalStorageUtils');

const KONG_API_GATEWAY_URL = process.env.KONG_URL;

router.post('/', RequestModifier, proxy(KONG_API_GATEWAY_URL,{
		proxyReqPathResolver: function(req) {
			// Modify urls before redirecting requests
			console.log('req body ====', req.body);
			let newUrl = '';
			if (KONG_API_GATEWAY_URL.indexOf('localhost') !== -1) {
				// request through local server
				console.log('redirect to local host ====');
				newUrl = require('url').parse(req.url).path + 'signin';

			} else {
				// request through Kong API Gateway
				newUrl = require('url').parse(req.url).path + 'users/v1/signin';
			}
			return newUrl;
		},
		userResDecorator: function(proxyRes, proxyResData, userReq, userRes) {
		    data = JSON.parse(proxyResData.toString('utf8'));
		   	console.log('status code====', proxyRes.statusCode);

			if (proxyRes.statusCode === 200) {
				// Login success, redirect to home page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/index');

				// Read jwt token from response header
				let jwtToken = userRes.getHeaders()['authorization'];
				console.log("jwt from response header=====", jwtToken);

				// Write jwt token to local storage
				JwtUtils.writeTokenToBrowser(jwtToken);
				console.log('jwt from local storage====', JwtUtils.readTokenFromBrowser());

				LocalStorageUtils.write('name', userRes);
				console.log('read from local storage====', LocalStorageUtils.read('name'));

			} else if (proxyRes.statusCode === 401 || proxyRes.statusCode === 400 || proxyRes.statusCode === 510 || proxyRes.statusCode === 404) {
				// Login failed, redirect to signin page
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
