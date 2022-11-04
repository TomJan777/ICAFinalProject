import { EncodeObject, GeneratedType } from "@cosmjs/proto-signing"
import {
    MsgCreateRoadOperator,
    MsgCreateRoadOperatorResponse,
    MsgCreateUserVault,
    MsgCreateUserVaultResponse,
    MsgUpdateRoadOperator,
    MsgUpdateRoadOperatorResponse,
    MsgUpdateUserVault,
    MsgUpdateUserVaultResponse,
    MsgDeleteRoadOperator,
    MsgDeleteRoadOperatorResponse,
    MsgDeleteUserVault,
    MsgDeleteUserVaultResponse,
} from "../generated/tollroad/tx"

export const typeUrlMsgCreateRoadOperator =
    "/b9lab.tollroad.tollroad.MsgCreateRoadOperator"
export const typeUrlMsgCreateRoadOperatorResponse =
    "/b9lab.tollroad.tollroad.MsgCreateRoadOperatorResponse"
export const typeUrlMsgUpdateRoadOperator =
    "/b9lab.tollroad.tollroad.MsgUpdateRoadOperator"
export const typeUrlMsgUpdateRoadOperatorResponse =
    "/b9lab.tollroad.tollroad.MsgUpdateRoadOperatorResponse"
export const typeUrlMsgDeleteRoadOperator =
    "/b9lab.tollroad.tollroad.MsgDeleteRoadOperator"
export const typeUrlMsgDeleteRoadOperatorResponse =
    "/b9lab.tollroad.tollroad.MsgDeleteRoadOperatorResponse"
export const typeUrlMsgCreateUserVault =
    "/b9lab.tollroad.tollroad.MsgCreateUserVault"
export const typeUrlMsgCreateUserVaultResponse =
    "/b9lab.tollroad.tollroad.MsgCreateUserVaultResponse"
export const typeUrlMsgUpdateUserVault =
    "/b9lab.tollroad.tollroad.MsgUpdateUserVault"
export const typeUrlMsgUpdateUserVaultResponse =
    "/b9lab.tollroad.tollroad.MsgUpdateUserVaultResponse"
export const typeUrlMsgDeleteUserVault =
    "/b9lab.tollroad.tollroad.MsgDeleteUserVault"
export const typeUrlMsgDeleteUserVaultResponse =
    "/b9lab.tollroad.tollroad.MsgDeleteUserVaultResponse"

export const tollroadTypes: ReadonlyArray<[string, GeneratedType]> = [
    [typeUrlMsgCreateRoadOperator, MsgCreateRoadOperator],
    [typeUrlMsgCreateRoadOperatorResponse, MsgCreateRoadOperatorResponse],
    // TODO other types
]

export interface MsgCreateRoadOperatorEncodeObject extends EncodeObject {
    readonly typeUrl: "/b9lab.tollroad.tollroad.MsgCreateRoadOperator"
    readonly value: Partial<MsgCreateRoadOperator>
}

export function isMsgCreateRoadOperatorEncodeObject(
    encodeObject: EncodeObject,
): encodeObject is MsgCreateRoadOperatorEncodeObject {
    return (
        (encodeObject as MsgCreateRoadOperatorEncodeObject).typeUrl ===
        typeUrlMsgCreateRoadOperator
    )
}

// TODO other types
