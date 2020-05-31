// This Servertemplate uses Webpack joins js files to single file named bundled.js .
// npm run webpack
// to make bundled.js file.

const path = require('path');
var glob = require('glob');

module.exports = {
    mode: "production",
    entry: glob.sync('./dist/**/**.js'),
    output: {
        path: path.join(__dirname, 'resources/static/js'),
        filename: 'bundled.js'
    },
    module:{
        rules: [
            {
                test: /\.js?$/,
                exclude: /node_modules/,
            },
        ],
    }
}

