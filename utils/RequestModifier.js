const RequestModifier = (req, res, next) => {
    req.headers['Content-Type'] = 'application/json';
    next()
}

module.exports = RequestModifier;
