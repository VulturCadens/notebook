import { v4 as uuidv4 } from "uuid"

export function UniversallyUniqueIdentifier() {
    const uuid = uuidv4()
    console.log(uuid)
}