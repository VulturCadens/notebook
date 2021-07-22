import * as PIXI from "pixi.js"

var application, sprite

function init() {
    application = new PIXI.Application({
        width: 800,
        height: 600,
        antialias: true,
        backgroundAlpha: 0.2,
        resolution: 1,
    })

    const bg = new PIXI.Sprite(PIXI.Texture.EMPTY)

    bg.width = application.screen.width
    bg.height = application.screen.height
    bg.interactive = true

    bg.on("pointerdown", downEvent)

    application.stage.addChild(bg)

    document.body.appendChild(application.view)

    application.loader
        .add("SHIP", "/static/image/ship.png")
        .load(setup)
}

function downEvent(event) {
    const x = event.data.global.x
    const y = event.data.global.y

    console.log("Click " + x + ":" + y)
}

function setup(_loader, resources) {
    sprite = new PIXI.Sprite(resources.SHIP.texture)

    sprite.anchor.set(0.5)

    sprite.x = 100
    sprite.y = 100

    sprite.velocity = {}
    sprite.velocity.x = 0.1
    sprite.velocity.y = 0.05

    application.stage.addChild(sprite)

    const text = new PIXI.Text("Text Object", new PIXI.TextStyle({
        dropShadow: true,
        dropShadowAlpha: 0.4,
        dropShadowBlur: 3,
        fill: "#7ebdcd",
        fontFamily: "Times New Roman",
        fontSize: 40,
    	fontVariant: "small-caps",
    	fontWeight: "bold",
    }))

    text.x = 200
    text.y = 200

    application.stage.addChild(text)

    application.ticker.add(loop)
}

function loop(delta) {
    sprite.rotation += (0.01 * delta)

    sprite.x += (sprite.velocity.x * delta)
    sprite.y += (sprite.velocity.y * delta)
}

window.onload = init
