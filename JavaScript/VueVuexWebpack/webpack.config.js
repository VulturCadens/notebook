const path = require("path")

module.exports = {
  entry: "./src/main.js",

  output: {
    filename: "main.js",
    path: path.resolve(__dirname, "www/javascript")
  },

  devServer: {
    contentBase: path.join(__dirname, "www"),
    publicPath: "/javascript/",
    port: 8000
  }
}
