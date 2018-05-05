/**
 * To verify whether browser will automatically attach all cookies to its requests
 * @param { object } req: htt request object
 * @return {} [description]
 */
function read (req, cookieName) {
    
    return req.cookies[cookieName];
}

/**
 * Write cookies to browser
 * @param  {[type]} req            [description]
 * @param  {[type]} res            [description]
 * @param  {[type]} cookieName     [description]
 * @param  {[type]} cookieValue    [description]
 * @param  {[type]} durationInMils [description]
 * @return {[type]}                [description]
 */
function write (req, res, cookieName, cookieValue, durationInMils) {
	res.cookie(cookieName, cookieValue, {expires: new Date(Date.now() + durationInMils)});
}


const Utils = {
    write: write,
    read: read
};

module.exports = Utils;
