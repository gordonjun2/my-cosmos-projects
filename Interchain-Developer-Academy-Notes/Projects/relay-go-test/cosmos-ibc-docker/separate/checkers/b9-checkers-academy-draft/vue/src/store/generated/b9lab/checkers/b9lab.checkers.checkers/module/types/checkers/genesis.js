/* eslint-disable */
import { Leaderboard } from "../checkers/leaderboard";
import { PlayerInfo } from "../checkers/player_info";
import { StoredGame } from "../checkers/stored_game";
import { NextGame } from "../checkers/next_game";
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "b9lab.checkers.checkers";
const baseGenesisState = { portId: "" };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.leaderboard !== undefined) {
            Leaderboard.encode(message.leaderboard, writer.uint32(34).fork()).ldelim();
        }
        for (const v of message.playerInfoList) {
            PlayerInfo.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.storedGameList) {
            StoredGame.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.nextGame !== undefined) {
            NextGame.encode(message.nextGame, writer.uint32(10).fork()).ldelim();
        }
        if (message.portId !== "") {
            writer.uint32(42).string(message.portId);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.playerInfoList = [];
        message.storedGameList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 4:
                    message.leaderboard = Leaderboard.decode(reader, reader.uint32());
                    break;
                case 3:
                    message.playerInfoList.push(PlayerInfo.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.storedGameList.push(StoredGame.decode(reader, reader.uint32()));
                    break;
                case 1:
                    message.nextGame = NextGame.decode(reader, reader.uint32());
                    break;
                case 5:
                    message.portId = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.playerInfoList = [];
        message.storedGameList = [];
        if (object.leaderboard !== undefined && object.leaderboard !== null) {
            message.leaderboard = Leaderboard.fromJSON(object.leaderboard);
        }
        else {
            message.leaderboard = undefined;
        }
        if (object.playerInfoList !== undefined && object.playerInfoList !== null) {
            for (const e of object.playerInfoList) {
                message.playerInfoList.push(PlayerInfo.fromJSON(e));
            }
        }
        if (object.storedGameList !== undefined && object.storedGameList !== null) {
            for (const e of object.storedGameList) {
                message.storedGameList.push(StoredGame.fromJSON(e));
            }
        }
        if (object.nextGame !== undefined && object.nextGame !== null) {
            message.nextGame = NextGame.fromJSON(object.nextGame);
        }
        else {
            message.nextGame = undefined;
        }
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = String(object.portId);
        }
        else {
            message.portId = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.leaderboard !== undefined &&
            (obj.leaderboard = message.leaderboard
                ? Leaderboard.toJSON(message.leaderboard)
                : undefined);
        if (message.playerInfoList) {
            obj.playerInfoList = message.playerInfoList.map((e) => e ? PlayerInfo.toJSON(e) : undefined);
        }
        else {
            obj.playerInfoList = [];
        }
        if (message.storedGameList) {
            obj.storedGameList = message.storedGameList.map((e) => e ? StoredGame.toJSON(e) : undefined);
        }
        else {
            obj.storedGameList = [];
        }
        message.nextGame !== undefined &&
            (obj.nextGame = message.nextGame
                ? NextGame.toJSON(message.nextGame)
                : undefined);
        message.portId !== undefined && (obj.portId = message.portId);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.playerInfoList = [];
        message.storedGameList = [];
        if (object.leaderboard !== undefined && object.leaderboard !== null) {
            message.leaderboard = Leaderboard.fromPartial(object.leaderboard);
        }
        else {
            message.leaderboard = undefined;
        }
        if (object.playerInfoList !== undefined && object.playerInfoList !== null) {
            for (const e of object.playerInfoList) {
                message.playerInfoList.push(PlayerInfo.fromPartial(e));
            }
        }
        if (object.storedGameList !== undefined && object.storedGameList !== null) {
            for (const e of object.storedGameList) {
                message.storedGameList.push(StoredGame.fromPartial(e));
            }
        }
        if (object.nextGame !== undefined && object.nextGame !== null) {
            message.nextGame = NextGame.fromPartial(object.nextGame);
        }
        else {
            message.nextGame = undefined;
        }
        if (object.portId !== undefined && object.portId !== null) {
            message.portId = object.portId;
        }
        else {
            message.portId = "";
        }
        return message;
    },
};
