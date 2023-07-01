# A Little Bit About TypeScript

## Basic Types

* number
* string
* boolean
* any
* void
* undefined
* null
* unknown
* never

## Union

```typescript
let u: number | string
```

## Literal Types

```typescript
type ColorDepth = 8 | 16 | 32

type Button = "UP" | "DOWN" | "HOLD"

let c: ColorDepth = 16
```

## Array

Array types can be written in one of two ways: __T[]__ or __Array\<T\>__.

```typescript
let a1: number[] = [1, 2]
let a2: Array<string> = ["A", "B"]
```

Tuple is an array with a fixed number of elements whose types are known.

```typescript
type Data = [string, boolean]

let x: Data = ["A", true]
```

## Enum

```typescript
enum City {
  Berlin,
  Paris,
  London,
}

let c: City = City.Berlin

enum Color {
  Red   = "rgb(255,0,0)",
  Green = "rgb(0,255,0)",
  Blue  = "rgb(0,0,255)",
}

const element: HTMLDivElement = document.createElement("div")

element.style.backgroundColor = Color.Green;
```

## Declare

Sometimes you have to tell the TypeScript compiler that a variable or an object exists, even the compiler knows nothing about it. The declare keyword tells the compiler that an object exists.

```typescript
declare externalResource : { func: (n : number) => string }

const value : number = externalResource.func(42)
```

## Interface

Optional property: __?__

Readonly property: __readonly__

```typescript
interface User {
  readonly ID:       number
           name:     string
           password: string
           room?:    number
           login:    boolean
}

let u: User = {
    ID: 123, name: "John", password: "Smith", room: 42, login: false,
}
```
## Functions

```typescript
function add(x: number, y: number = 2): number {
  return x + y
}

function add(x: number, y?: number): number | undefined {
  if (y == undefined) {
    return undefined
  }

  return x + y
}
```

## Type assertions

Type assertions have two forms.

```typescript
let str: any = "foobar"

let len: number = (str as string).length
let len: number = (<string>str).length
```

## Class

```typescript
abstract class Animal {
    protected abstract sound: string
    protected static count: number = 0

    constructor() {
        Animal.count += 1
    }

    get count(): number {
        return Animal.count
    }
}

class Cat extends Animal {
    public sound: string

    constructor(s: string) {
        super()
        this.sound = s
    }

    public say(): string {
        return (this.sound).repeat(this.count)
    }
}

let _1 : Cat = new Cat(" Meow!")

console.log(`We have ${_1.count} cat(s):${_1.say()}\n`)

let _2 : Cat = new Cat(" Zzz")
let _3 : Cat = new Cat(" Purrr...")

console.log(`We have ${_1.count} cat(s):${_2.say()}\n`)
```