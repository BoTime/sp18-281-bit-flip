var express = require('express');
var router = express.Router();
var ejs = require('ejs');
var proxy = require('express-http-proxy');

// Return order page
router.post('/', proxy("http://localhost:3000",{
		proxyReqPathResolver: function(req) {
			console.log("ORDER POST");
			console.log(req.body);	
			return require('url').parse(req.url).path + 'order';
		},
		userResDecorator: function(proxyRes, proxyResData, userReq, userRes) {
		    data = JSON.parse(proxyResData.toString('utf8'));
		   	console.log('status code====', proxyRes.statusCode);
			if (proxyRes.statusCode === 200) {
				// Order updated sucessfully
				userRes.statusCode = 200;
				userRes.setHeader('Location', '/order');

			} else if (proxyRes.statusCode === 401 || proxyRes.statusCode === 400) {
				// Login failed, redirect to signin page
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/signin');
			}
		    return JSON.stringify(data);
	  	}
	})
);
/*router.post('/', (req, res) => {
	console.log("ORDER POST");
	console.log(req.body);
	
	res.render('order');
});*/
// Return order page
router.get('/', (req, res) => {
	res.render('order');
});

module.exports = router;
 