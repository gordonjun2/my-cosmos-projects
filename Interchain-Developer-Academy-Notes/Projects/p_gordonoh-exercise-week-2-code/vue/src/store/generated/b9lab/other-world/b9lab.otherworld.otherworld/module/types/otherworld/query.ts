/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../otherworld/params";
import { WorldParams } from "../otherworld/world_params";

export const protobufPackage = "b9lab.otherworld.otherworld";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetWorldParamsRequest {}

export interface QueryGetWorldParamsResponse {
  WorldParams: WorldParams | undefined;
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

const baseQueryGetWorldParamsRequest: object = {};

export const QueryGetWorldParamsRequest = {
  encode(
    _: QueryGetWorldParamsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetWorldParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetWorldParamsRequest,
    } as QueryGetWorldParamsRequest;
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

  fromJSON(_: any): QueryGetWorldParamsRequest {
    const message = {
      ...baseQueryGetWorldParamsRequest,
    } as QueryGetWorldParamsRequest;
    return message;
  },

  toJSON(_: QueryGetWorldParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<QueryGetWorldParamsRequest>
  ): QueryGetWorldParamsRequest {
    const message = {
      ...baseQueryGetWorldParamsRequest,
    } as QueryGetWorldParamsRequest;
    return message;
  },
};

const baseQueryGetWorldParamsResponse: object = {};

export const QueryGetWorldParamsResponse = {
  encode(
    message: QueryGetWorldParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.WorldParams !== undefined) {
      WorldParams.encode(
        message.WorldParams,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetWorldParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetWorldParamsResponse,
    } as QueryGetWorldParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.WorldParams = WorldParams.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetWorldParamsResponse {
    const message = {
      ...baseQueryGetWorldParamsResponse,
    } as QueryGetWorldParamsResponse;
    if (object.WorldParams !== undefined && object.WorldParams !== null) {
      message.WorldParams = WorldParams.fromJSON(object.WorldParams);
    } else {
      message.WorldParams = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetWorldParamsResponse): unknown {
    const obj: any = {};
    message.WorldParams !== undefined &&
      (obj.WorldParams = message.WorldParams
        ? WorldParams.toJSON(message.WorldParams)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetWorldParamsResponse>
  ): QueryGetWorldParamsResponse {
    const message = {
      ...baseQueryGetWorldParamsResponse,
    } as QueryGetWorldParamsResponse;
    if (object.WorldParams !== undefined && object.WorldParams !== null) {
      message.WorldParams = WorldParams.fromPartial(object.WorldParams);
    } else {
      message.WorldParams = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a WorldParams by index. */
  WorldParams(
    request: QueryGetWorldParamsRequest
  ): Promise<QueryGetWorldParamsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.otherworld.otherworld.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  WorldParams(
    request: QueryGetWorldParamsRequest
  ): Promise<QueryGetWorldParamsResponse> {
    const data = QueryGetWorldParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "b9lab.otherworld.otherworld.Query",
      "WorldParams",
      data
    );
    return promise.then((data) =>
      QueryGetWorldParamsResponse.decode(new Reader(data))
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
