const merge = require("webpack-merge");
const common = require("./webpack.common.js");
const Dotenv = require("dotenv-webpack");

module.exports = merge(common, {
    mode: "development",
    devtool: "inline-source-map",
    devServer: {
        port: 3000,
        contentBase: "./dist",
        historyApiFallback: true
    },
    plugins: [...common.plugins, new Dotenv({
        path: "./.dev.env"
    })]
});
