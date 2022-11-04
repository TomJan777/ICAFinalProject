import Long from "long"
import { QueryClient } from "@cosmjs/stargate"
import { PageResponse } from "../../types/generated/cosmos/base/query/v1beta1/pagination"
import { RoadOperator } from "../../types/generated/tollroad/road_operator"
import { SystemInfo } from "../../types/generated/tollroad/system_info"
import { UserVault } from "../../types/generated/tollroad/user_vault"

export interface AllRoadOperatorResponse {
    roadOperators: RoadOperator[]
    pagination?: PageResponse
}

export interface AllUserVaultResponse {
    userVaults: UserVault[]
    pagination?: PageResponse
}

export interface TollroadExtension {
    readonly tollroad: {
        readonly getSystemInfo: () => Promise<SystemInfo>

        readonly getRoadOperator: (
            index: string,
        ) => Promise<RoadOperator | undefined>
        readonly getAllRoadOperators: (
            key: Uint8Array,
            offset: Long,
            limit: Long,
            countTotal: boolean,
        ) => Promise<AllRoadOperatorResponse>

        readonly getUserVault: (
            owner: string,
            roadOperatorIndex: string,
            token: string,
        ) => Promise<UserVault | undefined>
        readonly getAllUserVaults: (
            key: Uint8Array,
            offset: Long,
            limit: Long,
            countTotal: boolean,
        ) => Promise<AllUserVaultResponse>
    }
}

export function setupTollroadExtension(base: QueryClient): TollroadExtension {
    throw "Not implemented"
}
