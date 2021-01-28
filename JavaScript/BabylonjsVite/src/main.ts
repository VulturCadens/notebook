import { Engine } from "@babylonjs/core/Engines/engine"
import { Scene } from "@babylonjs/core/scene"

import { createScene } from "./scene"

function createEngine(canvas: HTMLCanvasElement): Engine {
    return new Engine(
        canvas,
        true,
        {
            preserveDrawingBuffer: true,
            stencil: true,
            disableWebGL2Support: false
        })
}

function start(): void {
    const canvas = document.getElementById("renderCanvas") as HTMLCanvasElement

    const engine: Engine = createEngine(canvas)

    if (!engine) {
        console.error("Engine is null!")
        return
    }

    const scene: Scene = createScene(engine)

    engine.runRenderLoop(function () {
        scene.render()
    })

    /*
     *  window.addEventListener("resize", function () {
     *     engine.resize()
     *  })
     */
}

window.onload = start