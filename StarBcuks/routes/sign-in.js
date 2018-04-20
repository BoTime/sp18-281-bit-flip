var express = require('express');
var router = express.Router();
var proxy = require('express-http-proxy');

const KONG_API_GATEWAY_URL = process.env.KONG_URL;


router.post('/', proxy(KONG_API_GATEWAY_URL,{
		proxyReqPathResolver: function(req) {
			return require('url').parse(req.url).path + 'login';
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

// Return sign in page
router.get('/', (req, res) => {
	res.render('signin');
});

module.exports = router;
