const envCf = require('dotenv');

envCf.config();

let env = {
	development: {
		VERSION: '1.0.0(beta)',
		BASE_URL:'https://ln-api.ichigozdata.win',
	},
	production: {
		VERSION: '1.0.1',
		BASE_URL:'https://ln-api.ichigozdata.win',
	},
}

module.exports = env[process.env.API_ENV || process.env.NODE_ENV] || {};
