# my-cosmos-projects
A repository to document my learning about Cosmos (aka Interchain).

## Learning Journey
- [Interchain Developer Academy](https://ida.interchain.io/)
    - The content of Interchain Developer Academy is not fully available to the public. However, some of the content that is available to the public can be found in [Cosmos Developer Portal Tutorial](https://tutorials.cosmos.network/).
    - This repository contains the raw course content from the website (converted from *.mhtml* to *.docx*), projects and quizzes.
    - As Week 1 quiz report is not available yet, it is not certain that all the answers in the quiz are correct. Thus, not academic advice (NAA) XD!

## General Installation

## Useful APIs / CLI References

## Useful Commands to Set Up Projects

## Useful Videos
- [Blockchain 101 - A Visual Demo](https://www.youtube.com/watch?v=_160oMzblY8)
- [Alternatives to Proof-of-Work with Tendermint's Jae Kwon | EDCON Toronto 2018](https://www.youtube.com/watch?v=T_FYBKJxlbk&t=155s)
- [What Is Cosmos SDK and How to Use It to Customize Your Chain w/ Marko Baricevic of Interchain.](https://www.youtube.com/watch?v=1_ottIKPfI4&t=2s)
- [Introduction to app.go, Julien Robert, Developer Relations Engineer for the Cosmos SDK.](https://www.youtube.com/watch?v=G6QUIUwYaSU)
- [Cosmos Academy - Running a node Screencast](https://www.youtube.com/watch?v=wNUjkp2PFQI)
- [Cosmos Academy - Your own chain, with Ignite](https://www.youtube.com/watch?v=z1HDh2KdiGI)

## Useful Information and Websites
- [Tendermint](https://tendermint.com/)
- [Light Clients in Tendermint Consensus](https://blog.cosmos.network/light-clients-in-tendermint-consensus-1237cfbda104)
- [Cosmos SDK Documentation](https://docs.cosmos.network/)
- [CometBFT Documentation](https://docs.cometbft.com/)
- [Cosmos Hub Documentaion](https://hub.cosmos.network/)
- [Inter-Blockchain Communication Protocol](https://ibcprotocol.dev/)
- [IBC Documentation](https://ibc.cosmos.network/)
- [Protocol Buffers Documentation](https://protobuf.dev/)
- [gRPC](https://grpc.io/)
- [gRPC-Gateway Documentation](https://grpc-ecosystem.github.io/grpc-gateway/)
- [Tendermint Core (BFT Consensus) in Go](https://github.com/tendermint/tendermint)
- [Tendermint Go Amino Documentation](https://github.com/tendermint/go-amino)
- [Docker Hub](https://hub.docker.com/)
- [The Ultimate Docker Cheat Sheet](https://dockerlabs.collabnix.com/docker/cheatsheet/)
- [Interchain Market Capitalization Overview](https://cosmos.network/ecosystem/tokens/)
- [Interchain Ecosystem Overview](https://cosmos.network/ecosystem/apps/?ref=cosmonautsworld)
- [Interchain Wallets](https://cosmos.network/ecosystem/wallets/)
- [Map of Zones - Cosmos Network Explorer](https://mapofzones.com/)
- [Gravity Bridge - Peg-Zone Implementation for Ethereum](https://github.com/cosmos/gravity-bridge)
- [Ignite CLI Documentation](https://docs.ignite.com/)
- [CosmWasm](https://cosmwasm.com/)
- [CosmWasm Documentation](https://docs.cosmwasm.com/docs/)
- [Ethermint](https://github.com/evmos/ethermint)
- [MintScan - Interchain Explorer](https://hub.mintscan.io/chains/overview)
- [MintScan's Validator List](https://www.mintscan.io/cosmos/validators)
- [The latest gossip on BFT consensus - Practical Byzantine Fault Tolerance](https://arxiv.org/abs/1807.04938)
- [Checkers Game in Go](https://github.com/batkinson/checkers-go)
- [Mnemonic-Code-Converter (BIP32/BIP39/BIP44/BIP49)](https://www.bip32.net/)
- [Chain Registry](https://github.com/cosmos/chain-registry)
- [Cosmos Hub Forum](https://forum.cosmos.network/)
- [Vue.js](https://vuejs.org/)

## Testnet Faucet

## DEX

## Whitepapers
- [Bitcoin: A Peer-to-Peer Electronic Cash System](https://bitcoin.org/bitcoin.pdf)
- [Tendermint: Consensus without Mining](https://tendermint.com/static/docs/tendermint.pdf)
- [A Network of Distributed Ledgers (Cosmos Whitepaper)](https://v1.cosmos.network/resources/whitepaper)

## Tips
- If project's binary cannot be called after using ```ignite chain serve``` in the ignite project folder, it may be because the *GOPATH* is not set. Ensure that *PATH* contains *GOPATH* by appending the following to your profile, bashrc, zshrc or equivalent:
    ```
    export PATH=$PATH:$(go env GOPATH)/bin
    ```
- To reset ```ignite chain serve```, use
    ```
    ignite chain serve --reset-once
    ```
- To installing Node.js dependencies and ```npm install``` hangs at ```sill: idealTree build```, try deleting *node_modules* and *package-lock.json* and then run
    ```
    npm config set registry http://registry.npmjs.org/
    npm install
    ```
