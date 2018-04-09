var express = require('express');
var router = express.Router();
var html_file_name ='./public/index.html';
var path = require("path");
//Controller to render application home page
router.get('/' ,
   function(req, res)
   {
       res.sendFile(path.join(__dirname +'/../public/index.html'));
   }
);
module.exports = router;
