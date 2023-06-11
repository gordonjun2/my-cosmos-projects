/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "b9lab.otherworld.otherworld";

export interface WorldParams {
  name: string;
  gravity: number;
  landSupply: number;
}

const baseWorldParams: object = { name: "", gravity: 0, landSupply: 0 };

export const WorldParams = {
  encode(message: WorldParams, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.gravity !== 0) {
      writer.uint32(16).uint64(message.gravity);
    }
    if (message.landSupply !== 0) {
      writer.uint32(24).uint64(message.landSupply);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): WorldParams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseWorldParams } as WorldParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.gravity = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.landSupply = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WorldParams {
    const message = { ...baseWorldParams } as WorldParams;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.gravity !== undefined && object.gravity !== null) {
      message.gravity = Number(object.gravity);
    } else {
      message.gravity = 0;
    }
    if (object.landSupply !== undefined && object.landSupply !== null) {
      message.landSupply = Number(object.landSupply);
    } else {
      message.landSupply = 0;
    }
    return message;
  },

  toJSON(message: WorldParams): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.gravity !== undefined && (obj.gravity = message.gravity);
    message.landSupply !== undefined && (obj.landSupply = message.landSupply);
    return obj;
  },

  fromPartial(object: DeepPartial<WorldParams>): WorldParams {
    const message = { ...baseWorldParams } as WorldParams;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.gravity !== undefined && object.gravity !== null) {
      message.gravity = object.gravity;
    } else {
      message.gravity = 0;
    }
    if (object.landSupply !== undefined && object.landSupply !== null) {
      message.landSupply = object.landSupply;
    } else {
      message.landSupply = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
