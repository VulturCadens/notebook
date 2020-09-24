const mat4 = glMatrix.mat4

const global = {
    rotation: 0,
    vertexArrayObject: null
}

const loadShader = (webgl, shaderType, source) => {
    const shader = webgl.createShader(shaderType)

    webgl.shaderSource(shader, source)

    webgl.compileShader(shader)

    if (!webgl.getShaderParameter(shader, webgl.COMPILE_STATUS)) {
        const info = webgl.getShaderInfoLog(shader)
        alert("Cannot compile the shader: " + info)
        webgl.deleteShader(shader)
        return null
    }

    return shader
}

const initShaderProgram = (webgl, vertexShaderSource, fragmentShaderSource) => {
    const vertexShader = loadShader(webgl, webgl.VERTEX_SHADER, vertexShaderSource)
    const fragmentShader = loadShader(webgl, webgl.FRAGMENT_SHADER, fragmentShaderSource)

    const shaderProgram = webgl.createProgram()

    webgl.attachShader(shaderProgram, vertexShader)
    webgl.attachShader(shaderProgram, fragmentShader)
    webgl.linkProgram(shaderProgram)

    if (!webgl.getProgramParameter(shaderProgram, webgl.LINK_STATUS)) {
        const err = webgl.getProgramInfoLog(shaderProgram)
        alert("Cannot initialize the shader program: " + err)
        return null
    }

    return shaderProgram
}

const initBuffer = (webgl) => {
    const positions = new Float32Array([
        0, 0, 0,
        0, 0.5, 0,
        0.8, 0, 0,

        0, 0, 0,
        0, 0, 1,
        0.8, 0, 0,

        0, 0, 1,
        0.8, 0, 0,
        0, 0.5, 0,

        0, 0, 0,
        0, 0, 1,
        0, 0.5, 0
    ])

    const positionBuffer = webgl.createBuffer()
    webgl.bindBuffer(webgl.ARRAY_BUFFER, positionBuffer)
    webgl.bufferData(
        webgl.ARRAY_BUFFER,
        positions,
        webgl.STATIC_DRAW
    )

    const index = new Uint16Array([
        0, 1, 2,
        3, 4, 5,
        6, 7, 8,
        9, 10, 11
    ]) // Uint16 = webGL.UNSIGNED_SHORT

    const indexBuffer = webgl.createBuffer()
    webgl.bindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)
    webgl.bufferData(
        webgl.ELEMENT_ARRAY_BUFFER,
        index,
        webgl.STATIC_DRAW
    )

    const faceColors = [
        [1.0, 0.0, 0.0, 1.0],  // RED
        [0.0, 1.0, 0.0, 1.0],  // GREEN
        [0.0, 0.0, 1.0, 1.0],  // BLUE
        [0.5, 0.5, 0.5, 1.0]   // GREY
    ]

    let colors = []

    for (let i = 0; i < faceColors.length; i = i + 1) {
        const color = faceColors[i]
        colors = colors.concat(color, color, color)
    }

    const colorBuffer = webgl.createBuffer()
    webgl.bindBuffer(webgl.ARRAY_BUFFER, colorBuffer)
    webgl.bufferData(
        webgl.ARRAY_BUFFER,
        new Float32Array(colors),
        webgl.STATIC_DRAW
    )

    return {
        position: positionBuffer,
        color: colorBuffer,
        index: indexBuffer
    }
}

const drawScene = (webgl, programInfo, delta) => {
    const projectionMatrix = mat4.create()
    const modelViewMatrix = mat4.create()

    webgl.clear(webgl.COLOR_BUFFER_BIT | webgl.DEPTH_BUFFER_BIT)

    mat4.perspective(
        projectionMatrix,
        45 * Math.PI / 180, // field of view
        webgl.canvas.clientWidth / webgl.canvas.clientHeight, // aspect
        0.1, // zNear,
        100.0 // zFar
    )

    mat4.translate(
        modelViewMatrix,   // destination
        modelViewMatrix,   // matrix to translate
        [-0.0, 0.0, -6.0]  // direction
    )

    mat4.rotate(
        modelViewMatrix,  // destination matrix
        modelViewMatrix,  // matrix to rotate
        global.rotation,  // radians
        [1, 0, 0]         // X axis
    )

    mat4.rotate(
        modelViewMatrix,      // destination matrix
        modelViewMatrix,      // matrix to rotate
        global.rotation / 2,  // radians
        [0, 1, 0]             // Y axis
    )

    global.rotation = global.rotation + delta

    webgl.useProgram(programInfo.program)

    /*
     * Void gl.uniformMatrix4fv(
     *   location WebGLUniformLocation,
     *   transpose GLboolean,
     *   value Float32Array)
     * 
     * Specify matrix value for uniform variable.
     */
    webgl.uniformMatrix4fv(
        programInfo.uniformLocations.projectionMatrix, // WebGLUniformLocation
        false,
        projectionMatrix
    )

    webgl.uniformMatrix4fv(
        programInfo.uniformLocations.modelViewMatrix, // WebGLUniformLocation
        false,
        modelViewMatrix
    )

    webgl.bindVertexArray(global.vertexArrayObject)

    webgl.drawElements(
        webgl.TRIANGLES,      // type primitive to render
        12,                   // vertex count
        webgl.UNSIGNED_SHORT, // type
        0                     // offset
    )
}

window.onload = async () => {
    const canvas = document.querySelector("#canvas")

    const webgl = canvas.getContext("webgl2")

    if (webgl === null) {
        alert("Cannot initialize WebGL.")
        return
    }

    webgl.clearColor(0.0, 0.0, 0.0, 1.0)
    webgl.clearDepth(1.0)
    webgl.enable(webgl.DEPTH_TEST)
    webgl.depthFunc(webgl.LEQUAL)

    const vsResponse = await fetch("shaders/vertex.glsl")
    const vsSource = await vsResponse.text()

    const fsResponse = await fetch("shaders/fragment.glsl")
    const fsSource = await fsResponse.text()

    const shaderProgram = initShaderProgram(webgl, vsSource, fsSource)

    const programInfo = {
        program: shaderProgram,
        attribLocations: {
            /*
             *  GLint gl.getAttribLocation(program WebGLProgram, name DOMString)
             */
            vertexPosition: webgl.getAttribLocation(shaderProgram, "aVertexPosition"),
            vertexColor: webgl.getAttribLocation(shaderProgram, "aVertexColor")
        },
        uniformLocations: {
            /*
             *  WebGLUniformLocation gl.getUniformLocation(program WebGLProgram, name DOMString)
             */
            projectionMatrix: webgl.getUniformLocation(shaderProgram, "uProjectionMatrix"),
            modelViewMatrix: webgl.getUniformLocation(shaderProgram, "uModelViewMatrix")
        }
    }

    {
        global.vertexArrayObject = webgl.createVertexArray()
        webgl.bindVertexArray(global.vertexArrayObject)

        const buffer = initBuffer(webgl)

        webgl.bindBuffer(webgl.ARRAY_BUFFER, buffer.position)
        webgl.vertexAttribPointer(
            programInfo.attribLocations.vertexPosition,  // attribute index
            3,           // number of components (X, Y, Z)
            webgl.FLOAT, // type
            false,       // normalize
            0,           // bytes to skip (stride)
            0            // offset
        )
        /*
         *  Void gl.enableVertexAttribArray(index GLint)
         */
        webgl.enableVertexAttribArray(programInfo.attribLocations.vertexPosition)

        webgl.bindBuffer(webgl.ARRAY_BUFFER, buffer.color);
        webgl.vertexAttribPointer(
            programInfo.attribLocations.vertexColor,  // attribute index
            4,           // number of components (R, G, B, A)
            webgl.FLOAT, // type
            false,       // normalize
            0,           // bytes to skip (stride)
            0            // offset
        )
        webgl.enableVertexAttribArray(programInfo.attribLocations.vertexColor)

        webgl.bindBuffer(webgl.ELEMENT_ARRAY_BUFFER, buffer.index)

        webgl.bindVertexArray(null)
    }

    let then = 0;

    const render = (now) => {
        now = now * 0.001
        const delta = now - then
        then = now

        drawScene(webgl, programInfo, delta)

        requestAnimationFrame(render)
    }

    requestAnimationFrame(render)
}