import { toAscii, toBech32 } from "@cosmjs/encoding"
import { sha256 } from "@cosmjs/crypto"

export function getModuleAddress(prefix: string, moduleName: string): string {
    const hash: Uint8Array = sha256(toAscii(moduleName))
    const truncated: Uint8Array = hash.slice(0, 20)
    return toBech32(prefix, truncated)
}
