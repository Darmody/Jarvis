var mount = require('koa-mount');
var tyrion = require('tyrion');

/* istanbul ignore if  */
if(process.env.NODE_ENV !== 'test'){
  require('./lib/logger.js');
}

(function(){
  var app = require('./');
  app.use(mount('/tyrion', tyrion));

  module.exports = app.listen(3000, '0.0.0.0');
})();
