const { app, BrowserWindow } = require("electron")
const path = require("node:path")

const createWindow = () => {
	const win = new BrowserWindow({
		width: 800,
		height: 300,
		icon: "./assets/favicon-32x32.png",
		webPreferences: {
			preload: path.join(__dirname, "preload.js")
		}
	})

	win.setMenuBarVisibility(false)

	win.loadFile("index.html")
}

app.whenReady().then(() => {
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
