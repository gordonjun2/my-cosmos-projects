import { txClient, queryClient, MissingWalletError } from './module';
// @ts-ignore
import { SpVuexError } from '@starport/vuex';
import { Leaderboard } from "./module/types/checkers/leaderboard";
import { NextGame } from "./module/types/checkers/next_game";
import { CheckersPacketData } from "./module/types/checkers/packet";
import { NoData } from "./module/types/checkers/packet";
import { ScorePacketData } from "./module/types/checkers/packet";
import { ScorePacketAck } from "./module/types/checkers/packet";
import { PlayerInfo } from "./module/types/checkers/player_info";
import { StoredGame } from "./module/types/checkers/stored_game";
import { WinningPlayer } from "./module/types/checkers/winning_player";
export { Leaderboard, NextGame, CheckersPacketData, NoData, ScorePacketData, ScorePacketAck, PlayerInfo, StoredGame, WinningPlayer };
async function initTxClient(vuexGetters) {
    return await txClient(vuexGetters['common/wallet/signer'], {
        addr: vuexGetters['common/env/apiTendermint']
    });
}
async function initQueryClient(vuexGetters) {
    return await queryClient({
        addr: vuexGetters['common/env/apiCosmos']
    });
}
function mergeResults(value, next_values) {
    for (let prop of Object.keys(next_values)) {
        if (Array.isArray(next_values[prop])) {
            value[prop] = [...value[prop], ...next_values[prop]];
        }
        else {
            value[prop] = next_values[prop];
        }
    }
    return value;
}
function getStructure(template) {
    let structure = { fields: [] };
    for (const [key, value] of Object.entries(template)) {
        let field = {};
        field.name = key;
        field.type = typeof value;
        structure.fields.push(field);
    }
    return structure;
}
const getDefaultState = () => {
    return {
        Leaderboard: {},
        PlayerInfo: {},
        PlayerInfoAll: {},
        CanPlayMove: {},
        StoredGame: {},
        StoredGameAll: {},
        NextGame: {},
        _Structure: {
            Leaderboard: getStructure(Leaderboard.fromPartial({})),
            NextGame: getStructure(NextGame.fromPartial({})),
            CheckersPacketData: getStructure(CheckersPacketData.fromPartial({})),
            NoData: getStructure(NoData.fromPartial({})),
            ScorePacketData: getStructure(ScorePacketData.fromPartial({})),
            ScorePacketAck: getStructure(ScorePacketAck.fromPartial({})),
            PlayerInfo: getStructure(PlayerInfo.fromPartial({})),
            StoredGame: getStructure(StoredGame.fromPartial({})),
            WinningPlayer: getStructure(WinningPlayer.fromPartial({})),
        },
        _Subscriptions: new Set(),
    };
};
// initial state
const state = getDefaultState();
export default {
    namespaced: true,
    state,
    mutations: {
        RESET_STATE(state) {
            Object.assign(state, getDefaultState());
        },
        QUERY(state, { query, key, value }) {
            state[query][JSON.stringify(key)] = value;
        },
        SUBSCRIBE(state, subscription) {
            state._Subscriptions.add(subscription);
        },
        UNSUBSCRIBE(state, subscription) {
            state._Subscriptions.delete(subscription);
        }
    },
    getters: {
        getLeaderboard: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.Leaderboard[JSON.stringify(params)] ?? {};
        },
        getPlayerInfo: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.PlayerInfo[JSON.stringify(params)] ?? {};
        },
        getPlayerInfoAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.PlayerInfoAll[JSON.stringify(params)] ?? {};
        },
        getCanPlayMove: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.CanPlayMove[JSON.stringify(params)] ?? {};
        },
        getStoredGame: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.StoredGame[JSON.stringify(params)] ?? {};
        },
        getStoredGameAll: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.StoredGameAll[JSON.stringify(params)] ?? {};
        },
        getNextGame: (state) => (params = { params: {} }) => {
            if (!params.query) {
                params.query = null;
            }
            return state.NextGame[JSON.stringify(params)] ?? {};
        },
        getTypeStructure: (state) => (type) => {
            return state._Structure[type].fields;
        }
    },
    actions: {
        init({ dispatch, rootGetters }) {
            console.log('Vuex module: b9lab.checkers.checkers initialized!');
            if (rootGetters['common/env/client']) {
                rootGetters['common/env/client'].on('newblock', () => {
                    dispatch('StoreUpdate');
                });
            }
        },
        resetState({ commit }) {
            commit('RESET_STATE');
        },
        unsubscribe({ commit }, subscription) {
            commit('UNSUBSCRIBE', subscription);
        },
        async StoreUpdate({ state, dispatch }) {
            state._Subscriptions.forEach(async (subscription) => {
                try {
                    await dispatch(subscription.action, subscription.payload);
                }
                catch (e) {
                    throw new SpVuexError('Subscriptions: ' + e.message);
                }
            });
        },
        async QueryLeaderboard({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryLeaderboard()).data;
                commit('QUERY', { query: 'Leaderboard', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryLeaderboard', payload: { options: { all }, params: { ...key }, query } });
                return getters['getLeaderboard']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryLeaderboard', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryPlayerInfo({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryPlayerInfo(key.index)).data;
                commit('QUERY', { query: 'PlayerInfo', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryPlayerInfo', payload: { options: { all }, params: { ...key }, query } });
                return getters['getPlayerInfo']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryPlayerInfo', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryPlayerInfoAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryPlayerInfoAll(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.queryPlayerInfoAll({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'PlayerInfoAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryPlayerInfoAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getPlayerInfoAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryPlayerInfoAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryCanPlayMove({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryCanPlayMove(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.queryCanPlayMove({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'CanPlayMove', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryCanPlayMove', payload: { options: { all }, params: { ...key }, query } });
                return getters['getCanPlayMove']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryCanPlayMove', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryStoredGame({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryStoredGame(key.index)).data;
                commit('QUERY', { query: 'StoredGame', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryStoredGame', payload: { options: { all }, params: { ...key }, query } });
                return getters['getStoredGame']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryStoredGame', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryStoredGameAll({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryStoredGameAll(query)).data;
                while (all && value.pagination && value.pagination.nextKey != null) {
                    let next_values = (await queryClient.queryStoredGameAll({ ...query, 'pagination.key': value.pagination.nextKey })).data;
                    value = mergeResults(value, next_values);
                }
                commit('QUERY', { query: 'StoredGameAll', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryStoredGameAll', payload: { options: { all }, params: { ...key }, query } });
                return getters['getStoredGameAll']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryStoredGameAll', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async QueryNextGame({ commit, rootGetters, getters }, { options: { subscribe, all } = { subscribe: false, all: false }, params: { ...key }, query = null }) {
            try {
                const queryClient = await initQueryClient(rootGetters);
                let value = (await queryClient.queryNextGame()).data;
                commit('QUERY', { query: 'NextGame', key: { params: { ...key }, query }, value });
                if (subscribe)
                    commit('SUBSCRIBE', { action: 'QueryNextGame', payload: { options: { all }, params: { ...key }, query } });
                return getters['getNextGame']({ params: { ...key }, query }) ?? {};
            }
            catch (e) {
                throw new SpVuexError('QueryClient:QueryNextGame', 'API Node Unavailable. Could not perform query: ' + e.message);
            }
        },
        async sendMsgRejectGame({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRejectGame(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRejectGame:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRejectGame:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgSendScore({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSendScore(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSendScore:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSendScore:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgPlayMove({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgPlayMove(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgPlayMove:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgPlayMove:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async sendMsgCreateGame({ rootGetters }, { value, fee = [], memo = '' }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateGame(value);
                const result = await txClient.signAndBroadcast([msg], { fee: { amount: fee,
                        gas: "200000" }, memo });
                return result;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateGame:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateGame:Send', 'Could not broadcast Tx: ' + e.message);
                }
            }
        },
        async MsgRejectGame({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgRejectGame(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgRejectGame:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgRejectGame:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgSendScore({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgSendScore(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgSendScore:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgSendScore:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgPlayMove({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgPlayMove(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgPlayMove:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgPlayMove:Create', 'Could not create message: ' + e.message);
                }
            }
        },
        async MsgCreateGame({ rootGetters }, { value }) {
            try {
                const txClient = await initTxClient(rootGetters);
                const msg = await txClient.msgCreateGame(value);
                return msg;
            }
            catch (e) {
                if (e == MissingWalletError) {
                    throw new SpVuexError('TxClient:MsgCreateGame:Init', 'Could not initialize signing client. Wallet is required.');
                }
                else {
                    throw new SpVuexError('TxClient:MsgCreateGame:Create', 'Could not create message: ' + e.message);
                }
            }
        },
    }
};
