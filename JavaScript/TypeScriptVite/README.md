# A Little Bit About TypeScript

## Basic Types

* number
* string
* boolean
* any
* void
* undefined
* null

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
let x: [string, boolean] = ["A", true]
```

## Enum

```typescript
enum City {
  Berlin,
  Paris,
  London,
}

let c: City = City.Berlin
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

