import { Engine } from "@babylonjs/core/Engines/engine"
import { Scene } from "@babylonjs/core/scene"

import { Vector3, Color3 } from "@babylonjs/core/Maths/math"

import { ArcRotateCamera } from "@babylonjs/core/Cameras/arcRotateCamera"

import { DirectionalLight } from "@babylonjs/core/Lights/directionalLight"
import { HemisphericLight } from "@babylonjs/core/Lights/hemisphericLight"

import { MeshBuilder } from "@babylonjs/core/Meshes/meshBuilder"

import { StandardMaterial } from "@babylonjs/core/Materials/standardMaterial"
import { Texture } from "@babylonjs/core/Materials/Textures/texture"

import { Animation } from "@babylonjs/core/Animations/animation"
import { EasingFunction, SineEase } from "@babylonjs/core/Animations/easing"

import { ShadowGenerator } from "@babylonjs/core/Lights/Shadows/shadowGenerator"

// For side effects only.
import "@babylonjs/core/Animations/animatable"
import "@babylonjs/core/Lights/Shadows/shadowGeneratorSceneComponent"

export function createScene(engine: Engine): Scene {
    const scene = new Scene(engine)

    /*
     *  CAMERA
     */
    const camera = new ArcRotateCamera(
        "camera",
        -Math.PI / 2,
        Math.PI / 2.2,
        4,
        new Vector3(0, 0.5, 0),
        scene
    )

    camera.fov = 0.7 // Default 0.8 radians.

    /*
     *  LIGHTS
     */
    const hemiLight = new HemisphericLight(
        "Hemi_Light",
        new Vector3(0, -3, 3),
        scene
    )
    hemiLight.intensity = 0.8

    const dirLight = new DirectionalLight(
        "Dir_Light",
        new Vector3(-5, -10, 5),
        scene
    )

    dirLight.diffuse = new Color3(1, 1, 0.5)
    dirLight.intensity = 0.8

    /*
     *  CUBE
     */
    const box = MeshBuilder.CreateBox("box", {})
    const boxMat = new StandardMaterial("Box_Material", scene)

    boxMat.diffuseColor = new Color3(0.2, 0.8, 0.2)

    box.scaling = new Vector3(0.5, 0.5, 0.5)
    box.position.y = 0.6
    box.material = boxMat

    /*
     *  CUBE ANIMATION
     */

    const boxAnimation = new Animation(
        "Box_Animation",
        "position.x",
        30,
        Animation.ANIMATIONTYPE_FLOAT,
        Animation.ANIMATIONLOOPMODE_CYCLE
    )

    const boxKeys = []

    boxKeys.push({
        frame: 0,
        value: -2
    })
    boxKeys.push({
        frame: 100,
        value: 2
    })
    boxKeys.push({
        frame: 200,
        value: -2
    })

    boxAnimation.setKeys(boxKeys)
    box.animations = []
    box.animations.push(boxAnimation)

    const easingFunc = new SineEase()
    easingFunc.setEasingMode(EasingFunction.EASINGMODE_EASEINOUT)
    boxAnimation.setEasingFunction(easingFunc)

    scene.beginAnimation(box, 0, 210, true)

    /*
     *  GROUND
     */

    const ground = MeshBuilder.CreateGround("ground", { width: 10, height: 10 })

    const groundTexture = new Texture("/image/texture-1024x1024.png", scene)

    groundTexture.onLoadObservable.add(function (texture) {
        console.log("Texture loaded: " + texture.url)
    })

    groundTexture.uScale = 5
    groundTexture.vScale = 5

    const groundMaterial = new StandardMaterial("Ground_Material", scene)
    groundMaterial.diffuseTexture = groundTexture

    ground.material = groundMaterial

    /*
     *  SHADOW
     */
    const shadowGenerator = new ShadowGenerator(1024, dirLight)
    shadowGenerator.addShadowCaster(box)
    shadowGenerator.useExponentialShadowMap = true

    ground.receiveShadows = true

    /*
     *  OBSERVABLE
     */
    let roty: number = 0

    const observ = scene.onAfterRenderObservable.add(function () {
        roty += 0.01
        box.rotation.y = roty
        box.rotation.x = roty / 2
    })

    window.setTimeout(function () {
        scene.onAfterRenderObservable.remove(observ)
    }.bind(observ), 5000)

    return scene
}