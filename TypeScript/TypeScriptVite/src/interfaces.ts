export interface Example {
    name: string
    value: number
}

export function isExample(x: unknown): x is Example {
    if ("name" in <Example>x && "value" in <Example>x) {
        if (typeof (<Example>x).name === "string" && typeof (<Example>x).value === "number") {
            return true
        }
    }

    return false
}