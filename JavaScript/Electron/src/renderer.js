import { str } from "./module.js"

const information = document.getElementById("info")

information.innerText = `${str} Chrome (v${versions.chrome()}), Node.js (v${versions.node()}), and Electron (v${versions.electron()})`
