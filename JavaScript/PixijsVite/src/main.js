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

    document.body.appendChild(application.view)

    application.loader
        .add("SHIP", "/static/image/ship.png")
        .load(setup)
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

    application.ticker.add(loop)
}

function loop(delta) {
    sprite.rotation += (0.01 * delta)

    sprite.x += (sprite.velocity.x * delta)
    sprite.y += (sprite.velocity.y * delta)
}

window.onload = init
