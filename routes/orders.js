var express = require('express');
var router = express.Router();
var ejs = require('ejs');
var proxy = require('express-http-proxy');
const JwtUtils = require('../utils/JwtToken');
const goAPI =  process.env.KONG_URL;
/**
 * Fxn that returns a JSON stringified version of an object.
 * This fxn uses a custom replacer function to handle circular references
 * see http://stackoverflow.com/a/11616993/3043369
 * param {object} object - object to stringify
 * returns {string} JSON stringified version of object
 */
function JSONStringifyTweeked(object) {
    var cache = [];
    var str = JSON.stringify(object,
        // custom replacer fxn - gets around "TypeError: Converting circular structure to JSON"
        function(key, value) {
            if (typeof value === 'object' && value !== null) {
                if (cache.indexOf(value) !== -1) {
                    // Circular reference found, discard key
                    return;
                }
                // Store value in our collection
                cache.push(value);
            }
            return value;
        }, 4);
    cache = null; // enable garbage collection
    return str;
};

// Delete order pid
router.delete('/', JwtUtils.attachTokenToHeader, proxy(goAPI,{
		proxyReqPathResolver: function(req) {
			console.log("ORDERS DELETE");
			console.log(req.body);
			return require('url').parse(req.url).path + 'orders/v1/order';
		},
		userResDecorator: function(proxyRes, proxyResData, userReq, userRes) {
		    data = JSON.parse(proxyResData);
		   	console.log('status code====', proxyRes.statusCode);
			if (proxyRes.statusCode === 200 || proxyRes.statusCode === 202) {
				// Order updated sucessfully
				console.log('Delete sucess in route');
				userRes.statusCode = 200;
				return;

			} else if (proxyRes.statusCode === 401 || proxyRes.statusCode === 400 || proxyRes.statusCode === 404) {
				// Order placing failed, redirect to signin page
				userRes.statusCode = 401;
				return;
			}
		    return data;
	  	}
	})
);
//Get orders of a user
router.get('/', JwtUtils.attachTokenToHeader, proxy(goAPI,{
		proxyReqPathResolver: function(req) {
			console.log("ORDERS GET");
			console.log(req.body);
			return require('url').parse(req.url).path + 'orders/v1/order';
		},
		userResDecorator: function(proxyRes, proxyResData, userReq, userRes) {
			console.log("BEFORE");
		    data = JSON.parse(proxyResData);
			//data = JSONStringifyTweeked(proxyResData);
			console.log('status code====', proxyRes.statusCode);

			if (proxyRes.statusCode === 200) {
				// Order updated sucessfully
				userRes.statusCode = 200;
				console.log("Sucess");
				userRes.setHeader('Location', '/history');
			} else if (proxyRes.statusCode === 401 || proxyRes.statusCode === 400 || proxyRes.statusCode === 404) {
				// Order placing failed, redirect to signin page
				userRes.statusCode = 401;
				userRes.setHeader('Location', '/signin');
				return;
			}else{
				// Order placing failed, redirect to oops page
				console.log("500");
				userRes.statusCode = 500;
				userRes.setHeader('Location', '/oops');
				return;
			}
		    return  data;
	  	}
	})
);


module.exports = router;
