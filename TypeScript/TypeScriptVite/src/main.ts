import { Example, isExample } from "./interfaces"

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
        return <T>await response.json()
    } catch (error: any) {
        throw new Error(error.message)
    }
}

async function setup(): Promise<any> {
    let res: unknown

    try {
        res = await http<Example>("/json/example.json")

        if (!isExample(res)) {
            throw new Error("Response is not <Example>")
        }
    } catch (error: any) {
        console.error(error.message)
        return
    }

    const textElement: HTMLDivElement = document.createElement("div")

    textElement.innerHTML = `{<br>&nbsp;&nbsp;&nbsp;&nbsp;Name: ${res.name}<br>&nbsp;&nbsp;&nbsp;&nbsp;Value: ${res.value}<br>}`

    document.body.appendChild(textElement)
}

window.onload = setup