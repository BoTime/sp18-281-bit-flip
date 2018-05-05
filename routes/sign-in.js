const express = require('express');
const router = express.Router();
const proxy = require('express-http-proxy');
const RequestModifier = require('../utils/RequestModifier');
const LocalStorageUtils = require('../utils/LocalStorageUtils');
const CookieUtils = require('../utils/CookieUtils');

const KONG_API_GATEWAY_URL = process.env.KONG_URL;

router.post('/', RequestModifier, proxy(KONG_API_GATEWAY_URL,{
		proxyErrorHandler: function(err, res, next) {
			// switch (err.code) {
		    // 	case 'ECONNRESET':    { return res.status(405).send('504 became 405'); }
		    // 	case 'ECONNREFUSED':  { return res.status(200).send('gotcher back'); }
		    // 	default:              { next(err); }
		    // }
		    console.log('proxy error code ', err)
			next()
		},
		proxyReqPathResolver: function(req) {
			// Modify urls before redirecting requests
			console.log('post signin req body ====', req.body);
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
			// console.log('&&&&&& 1', proxyRes);
			// console.log('&&&&&& 2', JSON.parse(proxyResData));
			// console.log('&&&&&& 3', userRes);
			// redirect user based on status code
			if (proxyRes.statusCode === 200) {
				// Login success, redirect to home page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/index');

				let jwtTokenString = userRes.getHeaders()['authorization'];
				let name = JSON.parse(proxyResData)['name'];
				let durationInMils = 1000 * 600;

				CookieUtils.write(userReq, userRes, 'jwtToken', jwtTokenString, durationInMils);
				CookieUtils.write(userReq, userRes, 'name', name, durationInMils);

			} else if (proxyRes.statusCode >= 400 && proxyRes.statusCode < 500) {
				// Login failed, redirect to signin page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/signin');

			} else {
				// Redirect to "oops" page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/oops');
			}
		    return proxyResData;
	  	}
	})
);

// Return sign in page
router.get('/', (req, res) => {
	res.render('signin');
});

module.exports = router;
