const path = require("path")

module.exports = {
  entry: "./src/main.js",

  output: {
    filename: "javascript/main.js",
    path: path.resolve(__dirname, "www")
  },

  devServer: {
    contentBase: path.join(__dirname, "www"),
    publicPath: "/javascript/",
    port: 8000
  }
}
