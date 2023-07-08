# A Little Bit About TypeScript

## Types

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

### Utility Types

* Awaited\<Type>
* Partial\<Type>
* Required\<Type>
* Readonly\<Type>
* Record<Keys, Type>
* Pick<Type, Keys>
* Omit<Type, Keys>
* Exclude<UnionType, ExcludedMembers>
* Extract<Type, Union>
* Et Cetera...

[TypeScript: Documentation - Utility Types.](https://www.typescriptlang.org/docs/handbook/utility-types.html)

TypeScript Record is Map-like object. [Objects vs. Maps](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Map#objects_vs._maps)

```typescript
type UID = number;

type User = {
    id: UID;
    name: string;
}

// Record

const usersRecord: Record<string, User> = {
    "43-AA-34": { id: 112233, name: "Jack" },
    "42-BB-24": { id: 445566, name: "John" },
};

// Map

const usersMap = new Map<string, User>([
    ["43-AA-34", { id: 112233, name: "Jack" }],
    ["42-BB-24", { id: 112233, name: "John" }],
]);

// Another example.

const animal: Record<string, (n: number) => string> = {
    "cat": (n: number): string => {
       return "Meow ".repeat(n);
    },

    "dog": (n: number): string => {
        return "Wuff ".repeat(n);
    }
}

console.log(animal.cat(2));
console.log(animal.dog(5));
```

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
    public abstract say(): string;

    protected sound: string = "";
    private static _count: number = 0;

    constructor() {
        Animal._count += 1;
    }

    get count(): number {
        return Animal._count;
    }
}

class Cat extends Animal {
    private _colour: string = "";

    constructor(s: string, c: string) {
        super()
        this.sound = s + " "
        this._colour = c
    }

    public say(): string {
        return (this.sound).repeat(this.count).trimEnd()
    }

    get colour(): string {
        return this._colour
    }
}

let firstCat: Cat = new Cat("Meow", "orange")

console.log(`We have ${firstCat.count} cat. First cat sounds like ${firstCat.say()}.`)

let secondCat: Cat = new Cat("Zzz", "black")
let anotherCat: Cat = new Cat("Purrr", "white")

console.log(`We have ${firstCat.count} cats. Second cat sounds like ${secondCat.say()}.`)
console.log(`Colours are ${firstCat.colour}, ${secondCat.colour}, and ${anotherCat.colour}`)
```