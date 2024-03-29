import { str } from "./module.js"

const information = document.getElementById("info")
information.innerText = `${str} Chrome (v${versions.chrome()}), Node.js (v${versions.node()}), and Electron (v${versions.electron()})`

const element = document.getElementById("randomBytes")
const randomStr = await subprocess.randomStr()
element.innerHTML = `Random bytes: <b>${randomStr}</b>`