const path = require("path")

module.exports = {
  entry: "./src/main.js",

  output: {
    filename: "main.js",
    path: path.resolve(__dirname, "www/javascript")
  },

  devServer: {
    contentBase: path.join(__dirname, "www"), // The root directory.
    publicPath: "/javascript/",               // The base path where dev-server store the bundle in memory.
    port: 8000
  }
}