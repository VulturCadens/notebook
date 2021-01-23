import { Example } from "./interfaces"

async function http<T>(path: string): Promise<T> {
    const response = await fetch(
        path, {
        method: "post",
        credentials: "omit"
    })

    if (!response.ok) {
        throw new Error(response.statusText)
    }

    try {
        const responseObject: T = await response.json()
        return responseObject
    } catch (error) {
        throw new Error(error.message)
    }
}

async function setup(): Promise<any> {
    let res: Example

    try {
        res = await http<Example>("/json/example.json")
    } catch (error) {
        console.error(error.message)
        return
    }

    if ("name" in res && "value" in res) {
        const textElement: HTMLDivElement = document.createElement("div")
        const textContent: Text = document.createTextNode(`Name: ${res.name} Value: ${res.value}`)

        textElement.appendChild(textContent)

        document.body.appendChild(textElement)
    }
}

window.onload = setup