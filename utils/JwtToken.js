const attachTokenToHeader = (req, res, next) => {
    let jwtToken = readTokenFromBrowser();
    req.headers['Authorization'] = jwtToken;
    next()
};

if (typeof localStorage === "undefined" || localStorage === null) {
  var LocalStorage = require('node-localstorage').LocalStorage;
  localStorage = new LocalStorage('./token');
}

const readTokenFromBrowser = () => {
    let jwtToken = localStorage.getItem('jwtToken');
    return jwtToken;
};

/**
 * format: "jwt xxxxxxx"
 * @param  { string } tokenString [description]
 * @return { void }             [description]
 */
const writeTokenToBrowser = (tokenString) => {
    localStorage.setItem('jwtToken', tokenString);
};

const JwtUtil = {
    readTokenFromBrowser: readTokenFromBrowser,
    attachTokenToHeader: attachTokenToHeader,
    writeTokenToBrowser: writeTokenToBrowser
};

module.exports = JwtUtil;
