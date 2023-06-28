import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "cosmonaut.leaderboard.leaderboard";
export interface PlayerInfo {
    playerID: string;
    score: number;
    dateAdded: string;
    gameId: string;
}
export declare const PlayerInfo: {
    encode(message: PlayerInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): PlayerInfo;
    fromJSON(object: any): PlayerInfo;
    toJSON(message: PlayerInfo): unknown;
    fromPartial(object: DeepPartial<PlayerInfo>): PlayerInfo;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
