const path = require("path")

module.exports = {
    entry: "./src/main.js",

    output: {
        path: path.resolve(__dirname, "./dist"),
    },

    resolve: {
        alias: {
            "vue$": "vue/dist/vue.esm-bundler.js"
        }
    },

    devServer: {
        inline: true,
        hot: true,
        contentBase: path.join(__dirname, "src/public"),
        stats: "minimal",
        overlay: true
    }
}