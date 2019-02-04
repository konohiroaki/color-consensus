const path = require("path");
const CleanWebpackPlugin = require("clean-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
    entry: {
        app: "./src/index.js"
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                exclude: /node_modules/,
                use: ["babel-loader"]
            }, {
                test: /\.css$/,
                use: [
                    "style-loader",
                    "css-loader"
                ],
            }
        ]
    },
    resolve: {
        extensions: ["*", ".js", ".jsx"]
    },
    plugins: [
        new CleanWebpackPlugin(["dist"]),
        new HtmlWebpackPlugin({template: "src/template/template.html"})
    ],
    output: {
        path: path.resolve(__dirname, "dist"),
        filename: "[name].bundle.js"
    }
};