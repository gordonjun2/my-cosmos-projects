// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgRejectGame } from "./types/checkers/tx";
import { MsgSendScore } from "./types/checkers/tx";
import { MsgPlayMove } from "./types/checkers/tx";
import { MsgCreateGame } from "./types/checkers/tx";


const types = [
  ["/b9lab.checkers.checkers.MsgRejectGame", MsgRejectGame],
  ["/b9lab.checkers.checkers.MsgSendScore", MsgSendScore],
  ["/b9lab.checkers.checkers.MsgPlayMove", MsgPlayMove],
  ["/b9lab.checkers.checkers.MsgCreateGame", MsgCreateGame],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgRejectGame: (data: MsgRejectGame): EncodeObject => ({ typeUrl: "/b9lab.checkers.checkers.MsgRejectGame", value: data }),
    msgSendScore: (data: MsgSendScore): EncodeObject => ({ typeUrl: "/b9lab.checkers.checkers.MsgSendScore", value: data }),
    msgPlayMove: (data: MsgPlayMove): EncodeObject => ({ typeUrl: "/b9lab.checkers.checkers.MsgPlayMove", value: data }),
    msgCreateGame: (data: MsgCreateGame): EncodeObject => ({ typeUrl: "/b9lab.checkers.checkers.MsgCreateGame", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
