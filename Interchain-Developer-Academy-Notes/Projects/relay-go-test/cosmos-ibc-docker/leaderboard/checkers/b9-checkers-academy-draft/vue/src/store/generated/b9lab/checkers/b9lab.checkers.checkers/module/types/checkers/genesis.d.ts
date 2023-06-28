import { Leaderboard } from '../checkers/leaderboard';
import { PlayerInfo } from '../checkers/player_info';
import { StoredGame } from '../checkers/stored_game';
import { NextGame } from '../checkers/next_game';
import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "b9lab.checkers.checkers";
/** GenesisState defines the checkers module's genesis state. */
export interface GenesisState {
    /** this line is used by starport scaffolding # genesis/proto/state */
    leaderboard: Leaderboard | undefined;
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    playerInfoList: PlayerInfo[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    storedGameList: StoredGame[];
    /** this line is used by starport scaffolding # genesis/proto/stateField */
    nextGame: NextGame | undefined;
    /** this line is used by starport scaffolding # ibc/genesis/proto */
    portId: string;
}
export declare const GenesisState: {
    encode(message: GenesisState, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): GenesisState;
    fromJSON(object: any): GenesisState;
    toJSON(message: GenesisState): unknown;
    fromPartial(object: DeepPartial<GenesisState>): GenesisState;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
