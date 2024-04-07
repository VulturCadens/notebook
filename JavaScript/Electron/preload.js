//
// From Electron 20 onwards, preload scripts are sandboxed by default
// and no longer have access to a full Node.js environment.
// Practically, this means that you have a polyfilled require function
// that only has access to a limited set of APIs.
//
// Electron modules     -> Renderer process modules.
// Node.js modules      -> [ events, timers, url ]
// Polyfilled globals   -> [ buffer, process, clearImmediate, setImmediate ]
//
// https://www.electronjs.org/docs/latest/tutorial/sandbox
//

const { contextBridge, ipcRenderer } = require("electron")

contextBridge.exposeInMainWorld("versions", {
	node: () => process.versions.node,
	chrome: () => process.versions.chrome,
	electron: () => process.versions.electron
})

contextBridge.exposeInMainWorld("subprocess", {
	randomStr: () => ipcRenderer.invoke("randomHex")
})
