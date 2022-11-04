/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../tollroad/params";
import { SystemInfo } from "../tollroad/system_info";
import { RoadOperator } from "../tollroad/road_operator";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "b9lab.tollroad.tollroad";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetSystemInfoRequest {}

export interface QueryGetSystemInfoResponse {
  SystemInfo: SystemInfo | undefined;
}

export interface QueryGetRoadOperatorRequest {
  index: string;
}

export interface QueryGetRoadOperatorResponse {
  roadOperator: RoadOperator | undefined;
}

export interface QueryAllRoadOperatorRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllRoadOperatorResponse {
  roadOperator: RoadOperator[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetSystemInfoRequest: object = {};

export const QueryGetSystemInfoRequest = {
  encode(
    _: QueryGetSystemInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },

  toJSON(_: QueryGetSystemInfoRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetSystemInfoRequest>
  ): QueryGetSystemInfoRequest {
    const message = {
      ...baseQueryGetSystemInfoRequest,
    } as QueryGetSystemInfoRequest;
    return message;
  },
};

const baseQueryGetSystemInfoResponse: object = {};

export const QueryGetSystemInfoResponse = {
  encode(
    message: QueryGetSystemInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.SystemInfo !== undefined) {
      SystemInfo.encode(message.SystemInfo, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetSystemInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.SystemInfo = SystemInfo.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromJSON(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetSystemInfoResponse): unknown {
    const obj: any = {};
    message.SystemInfo !== undefined &&
      (obj.SystemInfo = message.SystemInfo
        ? SystemInfo.toJSON(message.SystemInfo)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetSystemInfoResponse>
  ): QueryGetSystemInfoResponse {
    const message = {
      ...baseQueryGetSystemInfoResponse,
    } as QueryGetSystemInfoResponse;
    if (object.SystemInfo !== undefined && object.SystemInfo !== null) {
      message.SystemInfo = SystemInfo.fromPartial(object.SystemInfo);
    } else {
      message.SystemInfo = undefined;
    }
    return message;
  },
};

const baseQueryGetRoadOperatorRequest: object = { index: "" };

export const QueryGetRoadOperatorRequest = {
  encode(
    message: QueryGetRoadOperatorRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRoadOperatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRoadOperatorRequest,
    } as QueryGetRoadOperatorRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRoadOperatorRequest {
    const message = {
      ...baseQueryGetRoadOperatorRequest,
    } as QueryGetRoadOperatorRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetRoadOperatorRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRoadOperatorRequest>
  ): QueryGetRoadOperatorRequest {
    const message = {
      ...baseQueryGetRoadOperatorRequest,
    } as QueryGetRoadOperatorRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetRoadOperatorResponse: object = {};

export const QueryGetRoadOperatorResponse = {
  encode(
    message: QueryGetRoadOperatorResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.roadOperator !== undefined) {
      RoadOperator.encode(
        message.roadOperator,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetRoadOperatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetRoadOperatorResponse,
    } as QueryGetRoadOperatorResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roadOperator = RoadOperator.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetRoadOperatorResponse {
    const message = {
      ...baseQueryGetRoadOperatorResponse,
    } as QueryGetRoadOperatorResponse;
    if (object.roadOperator !== undefined && object.roadOperator !== null) {
      message.roadOperator = RoadOperator.fromJSON(object.roadOperator);
    } else {
      message.roadOperator = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetRoadOperatorResponse): unknown {
    const obj: any = {};
    message.roadOperator !== undefined &&
      (obj.roadOperator = message.roadOperator
        ? RoadOperator.toJSON(message.roadOperator)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetRoadOperatorResponse>
  ): QueryGetRoadOperatorResponse {
    const message = {
      ...baseQueryGetRoadOperatorResponse,
    } as QueryGetRoadOperatorResponse;
    if (object.roadOperator !== undefined && object.roadOperator !== null) {
      message.roadOperator = RoadOperator.fromPartial(object.roadOperator);
    } else {
      message.roadOperator = undefined;
    }
    return message;
  },
};

const baseQueryAllRoadOperatorRequest: object = {};

export const QueryAllRoadOperatorRequest = {
  encode(
    message: QueryAllRoadOperatorRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRoadOperatorRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRoadOperatorRequest,
    } as QueryAllRoadOperatorRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRoadOperatorRequest {
    const message = {
      ...baseQueryAllRoadOperatorRequest,
    } as QueryAllRoadOperatorRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRoadOperatorRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRoadOperatorRequest>
  ): QueryAllRoadOperatorRequest {
    const message = {
      ...baseQueryAllRoadOperatorRequest,
    } as QueryAllRoadOperatorRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllRoadOperatorResponse: object = {};

export const QueryAllRoadOperatorResponse = {
  encode(
    message: QueryAllRoadOperatorResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.roadOperator) {
      RoadOperator.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllRoadOperatorResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllRoadOperatorResponse,
    } as QueryAllRoadOperatorResponse;
    message.roadOperator = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.roadOperator.push(
            RoadOperator.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllRoadOperatorResponse {
    const message = {
      ...baseQueryAllRoadOperatorResponse,
    } as QueryAllRoadOperatorResponse;
    message.roadOperator = [];
    if (object.roadOperator !== undefined && object.roadOperator !== null) {
      for (const e of object.roadOperator) {
        message.roadOperator.push(RoadOperator.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllRoadOperatorResponse): unknown {
    const obj: any = {};
    if (message.roadOperator) {
      obj.roadOperator = message.roadOperator.map((e) =>
        e ? RoadOperator.toJSON(e) : undefined
      );
    } else {
      obj.roadOperator = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllRoadOperatorResponse>
  ): QueryAllRoadOperatorResponse {
    const message = {
      ...baseQueryAllRoadOperatorResponse,
    } as QueryAllRoadOperatorResponse;
    message.roadOperator = [];
    if (object.roadOperator !== undefined && object.roadOperator !== null) {
      for (const e of object.roadOperator) {
        message.roadOperator.push(RoadOperator.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a SystemInfo by index. */
  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse>;
  /** Queries a RoadOperator by index. */
  RoadOperator(
    request: QueryGetRoadOperatorRequest
  ): Promise<QueryGetRoadOperatorResponse>;
  /** Queries a list of RoadOperator items. */
  RoadOperatorAll(
    request: QueryAllRoadOperatorRequest
  ): Promise<QueryAllRoadOperatorResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.tollroad.tollroad.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  SystemInfo(
    request: QueryGetSystemInfoRequest
  ): Promise<QueryGetSystemInfoResponse> {
    const data = QueryGetSystemInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.tollroad.tollroad.Query",
      "SystemInfo",
      data
    );
    return promise.then((data) =>
      QueryGetSystemInfoResponse.decode(new Reader(data))
    );
  }

  RoadOperator(
    request: QueryGetRoadOperatorRequest
  ): Promise<QueryGetRoadOperatorResponse> {
    const data = QueryGetRoadOperatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.tollroad.tollroad.Query",
      "RoadOperator",
      data
    );
    return promise.then((data) =>
      QueryGetRoadOperatorResponse.decode(new Reader(data))
    );
  }

  RoadOperatorAll(
    request: QueryAllRoadOperatorRequest
  ): Promise<QueryAllRoadOperatorResponse> {
    const data = QueryAllRoadOperatorRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.tollroad.tollroad.Query",
      "RoadOperatorAll",
      data
    );
    return promise.then((data) =>
      QueryAllRoadOperatorResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
