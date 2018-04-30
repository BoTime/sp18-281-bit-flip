if (typeof localStorage === "undefined" || localStorage === null) {
  var LocalStorage = require('node-localstorage').LocalStorage;
  localStorage = new LocalStorage('./token');
}

const read = (key) => {
    // 'jwtToken'
    let jwtToken = localStorage.getItem(key);
    return jwtToken;
};

/**
 * format: "jwt xxxxxxx"
 * @param  { string } key [description]
 * @return { string } value [description]
 */
const write = (key, value) => {
    localStorage.setItem(key, value);
};

const LocalStorageUtils = {
    read: read,
    write: write
};

module.exports = LocalStorageUtils;
