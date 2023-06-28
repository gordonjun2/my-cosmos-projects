/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'cosmonaut.leaderboard.leaderboard';
const baseLeaderboardPacketData = {};
export const LeaderboardPacketData = {
    encode(message, writer = Writer.create()) {
        if (message.noData !== undefined) {
            NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
        }
        if (message.ibcTopRankPacket !== undefined) {
            IbcTopRankPacketData.encode(message.ibcTopRankPacket, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseLeaderboardPacketData };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.noData = NoData.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.ibcTopRankPacket = IbcTopRankPacketData.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseLeaderboardPacketData };
        if (object.noData !== undefined && object.noData !== null) {
            message.noData = NoData.fromJSON(object.noData);
        }
        else {
            message.noData = undefined;
        }
        if (object.ibcTopRankPacket !== undefined && object.ibcTopRankPacket !== null) {
            message.ibcTopRankPacket = IbcTopRankPacketData.fromJSON(object.ibcTopRankPacket);
        }
        else {
            message.ibcTopRankPacket = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.noData !== undefined && (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
        message.ibcTopRankPacket !== undefined &&
            (obj.ibcTopRankPacket = message.ibcTopRankPacket ? IbcTopRankPacketData.toJSON(message.ibcTopRankPacket) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseLeaderboardPacketData };
        if (object.noData !== undefined && object.noData !== null) {
            message.noData = NoData.fromPartial(object.noData);
        }
        else {
            message.noData = undefined;
        }
        if (object.ibcTopRankPacket !== undefined && object.ibcTopRankPacket !== null) {
            message.ibcTopRankPacket = IbcTopRankPacketData.fromPartial(object.ibcTopRankPacket);
        }
        else {
            message.ibcTopRankPacket = undefined;
        }
        return message;
    }
};
const baseNoData = {};
export const NoData = {
    encode(_, writer = Writer.create()) {
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNoData };
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
    fromJSON(_) {
        const message = { ...baseNoData };
        return message;
    },
    toJSON(_) {
        const obj = {};
        return obj;
    },
    fromPartial(_) {
        const message = { ...baseNoData };
        return message;
    }
};
const baseIbcTopRankPacketData = { playerId: '', score: 0, dateAdded: '', gameId: '' };
export const IbcTopRankPacketData = {
    encode(message, writer = Writer.create()) {
        if (message.playerId !== '') {
            writer.uint32(10).string(message.playerId);
        }
        if (message.score !== 0) {
            writer.uint32(16).uint64(message.score);
        }
        if (message.dateAdded !== '') {
            writer.uint32(26).string(message.dateAdded);
        }
        if (message.gameId !== '') {
            writer.uint32(34).string(message.gameId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseIbcTopRankPacketData };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.playerId = reader.string();
                    break;
                case 2:
                    message.score = longToNumber(reader.uint64());
                    break;
                case 3:
                    message.dateAdded = reader.string();
                    break;
                case 4:
                    message.gameId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseIbcTopRankPacketData };
        if (object.playerId !== undefined && object.playerId !== null) {
            message.playerId = String(object.playerId);
        }
        else {
            message.playerId = '';
        }
        if (object.score !== undefined && object.score !== null) {
            message.score = Number(object.score);
        }
        else {
            message.score = 0;
        }
        if (object.dateAdded !== undefined && object.dateAdded !== null) {
            message.dateAdded = String(object.dateAdded);
        }
        else {
            message.dateAdded = '';
        }
        if (object.gameId !== undefined && object.gameId !== null) {
            message.gameId = String(object.gameId);
        }
        else {
            message.gameId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.playerId !== undefined && (obj.playerId = message.playerId);
        message.score !== undefined && (obj.score = message.score);
        message.dateAdded !== undefined && (obj.dateAdded = message.dateAdded);
        message.gameId !== undefined && (obj.gameId = message.gameId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseIbcTopRankPacketData };
        if (object.playerId !== undefined && object.playerId !== null) {
            message.playerId = object.playerId;
        }
        else {
            message.playerId = '';
        }
        if (object.score !== undefined && object.score !== null) {
            message.score = object.score;
        }
        else {
            message.score = 0;
        }
        if (object.dateAdded !== undefined && object.dateAdded !== null) {
            message.dateAdded = object.dateAdded;
        }
        else {
            message.dateAdded = '';
        }
        if (object.gameId !== undefined && object.gameId !== null) {
            message.gameId = object.gameId;
        }
        else {
            message.gameId = '';
        }
        return message;
    }
};
const baseIbcTopRankPacketAck = { playerId: '' };
export const IbcTopRankPacketAck = {
    encode(message, writer = Writer.create()) {
        if (message.playerId !== '') {
            writer.uint32(10).string(message.playerId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseIbcTopRankPacketAck };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.playerId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseIbcTopRankPacketAck };
        if (object.playerId !== undefined && object.playerId !== null) {
            message.playerId = String(object.playerId);
        }
        else {
            message.playerId = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.playerId !== undefined && (obj.playerId = message.playerId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseIbcTopRankPacketAck };
        if (object.playerId !== undefined && object.playerId !== null) {
            message.playerId = object.playerId;
        }
        else {
            message.playerId = '';
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
