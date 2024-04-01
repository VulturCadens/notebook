const { app, BrowserWindow, ipcMain } = require("electron")

const path = require("node:path")
const nodejsCrypto = require("node:crypto")

const createWindow = () => {
	const win = new BrowserWindow({
		width: 960,
		height: 540,
		x: 100,
		y: 100,
		icon: "./assets/favicon-32x32.png",
		webPreferences: {
			preload: path.join(__dirname, "preload.js")
		}
	})

	win.setMenuBarVisibility(false)

	win.loadFile("index.html")
}

const randomHex = () => {
    return nodejsCrypto.randomBytes(10).toString("hex")
}

app.whenReady().then(() => {
    ipcMain.handle("randomHex", () => randomHex())

	createWindow()

	app.on("activate", () => {
		if (BrowserWindow.getAllWindows().length === 0) {
			createWindow()
		}
	})
})

app.on("window-all-closed", () => {
	if (process.platform !== "darwin") app.quit()
})
