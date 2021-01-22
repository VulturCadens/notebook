import { example } from "./iface"

window.onload = (): void => {
  const ex: example = {
    name: "John Smith",
    value: 42
  }

  const textElement: HTMLDivElement = document.createElement("div")
  const textContent: Text = document.createTextNode(`Name: ${ex.name} Value: ${ex.value}`)

  textElement.appendChild(textContent)
  
  document.body.appendChild(textElement)
}