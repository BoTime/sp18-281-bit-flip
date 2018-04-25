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
			console.log(proxyResData.toString('utf8'));
		    data = JSON.parse(proxyResData.toString('utf8'));
		   	console.log('status code====', proxyRes.statusCode);
			if (proxyRes.statusCode === 200) {
				console.log("200");
				// Order updated sucessfully
				userRes.statusCode = 200;
				userRes.redirect('order');	

			} else if (proxyRes.statusCode === 401 || proxyRes.statusCode === 400) {
				// Order placing failed, redirect to signin page
				console.log("400");
				userRes.statusCode = 302;
				userRes.setHeader('Location', '/signin');
			}else {
				// Order placing failed, redirect to oops page
				console.log("500");
				userRes.statusCode = 500;
				userRes.setHeader('Location', '/oops');
			}	
		    return userRes;
	  	}
	})
);

// Return order page
router.get('/', (req, res) => {
	res.render('order');
});

module.exports = router;
 