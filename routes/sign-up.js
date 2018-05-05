const express = require('express');
const router = express.Router();
const proxy = require('express-http-proxy');
const RequestModifier = require('../utils/RequestModifier');
// const JwtToken = require('../utils/JwtUtils');
const CookieUtils = require('../utils/CookieUtils');

const KONG_API_GATEWAY_URL = process.env.KONG_URL;


router.post('/', RequestModifier, proxy(KONG_API_GATEWAY_URL,{
		proxyReqPathResolver: function(req) {
			let newUrl = '';
			if (KONG_API_GATEWAY_URL.indexOf('localhost') !== -1) {
				// request through local server
				newUrl = require('url').parse(req.url).path + 'signup';
			} else {
				// request through Kong API Gateway
				newUrl = require('url').parse(req.url).path + 'users/v1/signup';
			}

			return newUrl;
		},
		userResDecorator: function(proxyRes, proxyResData, userReq, userRes) {
			// redirect user based on status code
			if (proxyRes.statusCode === 200) {
				// Login success, redirect to home page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/index');

				let jwtTokenString = userRes.getHeaders()['authorization'];
				let name = JSON.parse(proxyResData)['name'];
				let durationInMils = 1000 * 60;

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

module.exports = router;
