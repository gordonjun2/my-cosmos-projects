// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgSendScore } from "./types/checkers/tx";
import { MsgCreateGame } from "./types/checkers/tx";
import { MsgRejectGame } from "./types/checkers/tx";
import { MsgPlayMove } from "./types/checkers/tx";
const types = [
    ["/b9lab.checkers.checkers.MsgSendScore", MsgSendScore],
    ["/b9lab.checkers.checkers.MsgCreateGame", MsgCreateGame],
    ["/b9lab.checkers.checkers.MsgRejectGame", MsgRejectGame],
    ["/b9lab.checkers.checkers.MsgPlayMove", MsgPlayMove],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgSendScore: (data) => ({ typeUrl: "/b9lab.checkers.checkers.MsgSendScore", value: data }),
        msgCreateGame: (data) => ({ typeUrl: "/b9lab.checkers.checkers.MsgCreateGame", value: data }),
        msgRejectGame: (data) => ({ typeUrl: "/b9lab.checkers.checkers.MsgRejectGame", value: data }),
        msgPlayMove: (data) => ({ typeUrl: "/b9lab.checkers.checkers.MsgPlayMove", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
