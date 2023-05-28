const envCf = require('dotenv');

envCf.config();

let env = {
	development: {
		VERSION: '1.0.0(beta)',
		BASE_URL:'http://ln.ichigozdata.win:8888',
		// BASE_URL:'http://localhost:8888',
	},
	production: {
		VERSION: '1.0.1',
		BASE_URL:'http://ln.ichigozdata.win:8888',
	},
}

module.exports = env[process.env.API_ENV || process.env.NODE_ENV] || {};
