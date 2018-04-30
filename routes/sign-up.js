const express = require('express');
const router = express.Router();
const proxy = require('express-http-proxy');
const RequestModifier = require('../utils/RequestModifier');
const JwtToken = require('../utils/JwtToken');

const KONG_API_GATEWAY_URL = process.env.KONG_URL;


router.post('/', RequestModifier, JwtToken.attachTokenToHeader, proxy(KONG_API_GATEWAY_URL,{
		proxyReqPathResolver: function(req) {
			let newUrl = '';
			if (KONG_API_GATEWAY_URL.indexOf('localhost') !== -1) {
				// request through local server
				newUrl = require('url').parse(req.url).path + 'signup';
			} else {
				// request through Kong API Gateway
				newUrl = require('url').parse(req.url).path + 'users/v1/signup';
			}
			console.log('auth header from sign up====',  req.headers['Authorization']);
			return newUrl;
		},
		userResDecorator: function(proxyRes, proxyResData, userReq, userRes) {
			data = JSON.parse(proxyResData.toString('utf8'));
			console.log('status code====', proxyRes.statusCode);
			if (proxyRes.statusCode === 200) {
				// Login success, redirect to home page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/home');

			} else if (proxyRes.statusCode === 401 || proxyRes.statusCode === 400) {
				// Login failed, redirect to signin page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/signin');
			}
			return JSON.stringify(data);
		}
	})
);



module.exports = router;
