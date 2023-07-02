# A Little Bit About TypeScript

## Types

---

### Primitives

* number
* string
* boolean
* any
* void
* undefined
* null
* unknown
* never
* bigint
* symbol

```typescript
const s: string = "Cat"

type UID = number   // A unique identifier.
const catID: UID = 42
```

### Build-in Objects

* Date
* Error
* Array
* Map
* Set
* Regexp
* Promise

```typescript
async function setup(): Promise<any> { ... }

const regex: RegExp = /^[ABC]/
```

### Union

```typescript
let u: number | string
```

### Literal Types

```typescript
type ColorDepth = 8 | 16 | 32

type Button = "UP" | "DOWN" | "HOLD"

let c: ColorDepth = 16
```

### Array

Array types can be written in one of two ways: __T[]__ or __Array\<T\>__.

```typescript
let a1: number[] = [1, 2]
let a2: Array<string> = ["A", "B"]
```

### Tuple

Tuple is an array with a fixed number of elements whose types are known.

```typescript
type Data = [string, boolean]

let x: Data = ["A", true]
```

### Enum

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

### Type

```typescript
type User = {
    readonly    ID:       number    // Readonly property.
                name:     string
                password: string
                room?:    number    // Optional property.
                login:    boolean
}

let u: User = {
    ID: 123, name: "John", password: "Smith", room: 42, login: false,
}
```

### Function

Function type: **(str: string) => number**

---

## Optional Chaining

The question mark dot (?.) syntax is called optional chaining in TypeScript.

```typescript
type Cat = {
  name: string;
  behaviour?: {
    nice: boolean;
  };
};

const c: Cat = { name: "Mirre", behaviour: { nice: true }};

console.log(c.name)
console.log(c.behaviour?.nice)

// You can also use optional chaining when attempting to call a method
// which may not exist (only calling a function if it exists).

const getFunc = (): Function | undefined => {
    const n: number = Math.floor(Math.random() * 2); // 0 or 1

    if (n == 0 ) {
        return () => { console.log("OK")}
    }

    return undefined;
}

const func = getFunc()

func?.()
```

## Declare

Sometimes you have to tell the TypeScript compiler that a variable or an object exists, even the compiler knows nothing about it. The declare keyword tells the compiler that an object exists.

```typescript
// The script coming from other domain.
const externalResource: any = {
    func: (n: number): string => {
        return "cat"
    }
}

// Another domain.
declare const externalResource : { func: (n: number) => string }

const str: string = externalResource.func(42)
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