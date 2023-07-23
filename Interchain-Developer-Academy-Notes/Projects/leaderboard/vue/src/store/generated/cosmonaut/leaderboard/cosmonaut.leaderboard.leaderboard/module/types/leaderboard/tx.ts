/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "cosmonaut.leaderboard.leaderboard";

export interface MsgSendIbcTopRank {
  creator: string;
  port: string;
  channelID: string;
  timeoutTimestamp: number;
  playerId: string;
  rank: number;
  score: number;
}

export interface MsgSendIbcTopRankResponse {}

const baseMsgSendIbcTopRank: object = {
  creator: "",
  port: "",
  channelID: "",
  timeoutTimestamp: 0,
  playerId: "",
  rank: 0,
  score: 0,
};

export const MsgSendIbcTopRank = {
  encode(message: MsgSendIbcTopRank, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.port !== "") {
      writer.uint32(18).string(message.port);
    }
    if (message.channelID !== "") {
      writer.uint32(26).string(message.channelID);
    }
    if (message.timeoutTimestamp !== 0) {
      writer.uint32(32).uint64(message.timeoutTimestamp);
    }
    if (message.playerId !== "") {
      writer.uint32(42).string(message.playerId);
    }
    if (message.rank !== 0) {
      writer.uint32(48).uint64(message.rank);
    }
    if (message.score !== 0) {
      writer.uint32(56).uint64(message.score);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSendIbcTopRank {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSendIbcTopRank } as MsgSendIbcTopRank;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.port = reader.string();
          break;
        case 3:
          message.channelID = reader.string();
          break;
        case 4:
          message.timeoutTimestamp = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.playerId = reader.string();
          break;
        case 6:
          message.rank = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.score = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendIbcTopRank {
    const message = { ...baseMsgSendIbcTopRank } as MsgSendIbcTopRank;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = String(object.port);
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = String(object.channelID);
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = Number(object.timeoutTimestamp);
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.playerId !== undefined && object.playerId !== null) {
      message.playerId = String(object.playerId);
    } else {
      message.playerId = "";
    }
    if (object.rank !== undefined && object.rank !== null) {
      message.rank = Number(object.rank);
    } else {
      message.rank = 0;
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = Number(object.score);
    } else {
      message.score = 0;
    }
    return message;
  },

  toJSON(message: MsgSendIbcTopRank): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.port !== undefined && (obj.port = message.port);
    message.channelID !== undefined && (obj.channelID = message.channelID);
    message.timeoutTimestamp !== undefined &&
      (obj.timeoutTimestamp = message.timeoutTimestamp);
    message.playerId !== undefined && (obj.playerId = message.playerId);
    message.rank !== undefined && (obj.rank = message.rank);
    message.score !== undefined && (obj.score = message.score);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSendIbcTopRank>): MsgSendIbcTopRank {
    const message = { ...baseMsgSendIbcTopRank } as MsgSendIbcTopRank;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.port !== undefined && object.port !== null) {
      message.port = object.port;
    } else {
      message.port = "";
    }
    if (object.channelID !== undefined && object.channelID !== null) {
      message.channelID = object.channelID;
    } else {
      message.channelID = "";
    }
    if (
      object.timeoutTimestamp !== undefined &&
      object.timeoutTimestamp !== null
    ) {
      message.timeoutTimestamp = object.timeoutTimestamp;
    } else {
      message.timeoutTimestamp = 0;
    }
    if (object.playerId !== undefined && object.playerId !== null) {
      message.playerId = object.playerId;
    } else {
      message.playerId = "";
    }
    if (object.rank !== undefined && object.rank !== null) {
      message.rank = object.rank;
    } else {
      message.rank = 0;
    }
    if (object.score !== undefined && object.score !== null) {
      message.score = object.score;
    } else {
      message.score = 0;
    }
    return message;
  },
};

const baseMsgSendIbcTopRankResponse: object = {};

export const MsgSendIbcTopRankResponse = {
  encode(
    _: MsgSendIbcTopRankResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgSendIbcTopRankResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgSendIbcTopRankResponse,
    } as MsgSendIbcTopRankResponse;
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

  fromJSON(_: any): MsgSendIbcTopRankResponse {
    const message = {
      ...baseMsgSendIbcTopRankResponse,
    } as MsgSendIbcTopRankResponse;
    return message;
  },

  toJSON(_: MsgSendIbcTopRankResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgSendIbcTopRankResponse>
  ): MsgSendIbcTopRankResponse {
    const message = {
      ...baseMsgSendIbcTopRankResponse,
    } as MsgSendIbcTopRankResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendIbcTopRank(
    request: MsgSendIbcTopRank
  ): Promise<MsgSendIbcTopRankResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  SendIbcTopRank(
    request: MsgSendIbcTopRank
  ): Promise<MsgSendIbcTopRankResponse> {
    const data = MsgSendIbcTopRank.encode(request).finish();
    const promise = this.rpc.request(
      "cosmonaut.leaderboard.leaderboard.Msg",
      "SendIbcTopRank",
      data
    );
    return promise.then((data) =>
      MsgSendIbcTopRankResponse.decode(new Reader(data))
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
