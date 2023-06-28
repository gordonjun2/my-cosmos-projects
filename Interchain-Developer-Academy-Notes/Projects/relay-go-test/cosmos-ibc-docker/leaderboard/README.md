## Build Docker images

to build the **images for the checkers game with the leaderboard**, clone the repository and run:

```
$ cd cosmos-ibc-docker/leaderboard/checkers
$ ./build-images.sh
```

In addition, you will need to build the **relayer image**:

```
$ cd cosmos-ibc-docker/leaderboard/relayer
$ docker build -f Dockerfile . -t relayer --no-cache
```

## Start the network

You can use the provided compose file to spin up a network with two checkers blockchains(checkersa and checkersb) and a relayer:

```
$ cd cosmos-ibc-docker/leaderboard
$ docker-compose -f leaderboard.yml up

```

Observe the output of `docker-compose` until the chains are ready - it will take some time for the chains to be ready. 

## Start the relayer

If the chains are ready, you can start the relayer process:

```
$ docker exec relayer ./run-relayer.sh 
```

wait till the connection is etablished and a channel is created. 

## Create a game and play it till ende

The easiest way to do so is to start the test script, jump into container:

```
$ docker exec -it checkersb bash
```

and run:

```
$ ./createandplaygame.sh 
```

this will crate and play a game in the checkersb chain. 

You can check the leaderboard of checkersa and checkersb via(on your host):

```
$ docker exec checkersa checkersd q checkers show-leaderboard
```

and

```
$ docker exec checkersb checkersd q checkers show-leaderboard
```


After a game is over, you can send the score of the player to chaina. First jump into chainb container again:

```
$ docker exec -it checkersb bash
```

and run in it:

```
$ checkersd tx checkers send-score checkers channel-0 --from cosmos173czeq76k0lh0m6zcz72yu6zj8c6d0tf294w5k --chain-id checkersb
```

Check again the leaderboard of chaina(in another terminal on your host):

```
$ docker exec checkersa checkersd q checkers show-leaderboard
```

For another test on the same chain, you will need to call `./test.sh` instead of `./createandplaygame.sh` because the latter will use game index 1.

## Information about changes

There are several steps involved in enabling IBC and creating and sending a packet:

use ignite:
- ibc.go: ignite template
- module_ibc.go: ignite template
- events_ibc_go: ignite template
- expected_keeper_ibc.go: ignite template
- handler.go: ignite template, adjusted by ignite

mimic ignite:
- keys.go: add Version and PortID and PortKey
- genesis.go: bind port 
- app.go: add route via `ibcRouter.AddRoute(checkersmoduletypes.ModuleName, checkersModule)`
- keeper.go: add `channelKeeper`, `portKeeper` and `scopedKeeper`
- keeper_test.go: adjust for the keepers
- packet.proto: include `ScorePacketData`(bc of type WinningPlayer, not possible via ignite)
- tx.proto: ignite template, adjusted to fetch the player info
- messages_score.go: ignite template, adjusted to fetch the player info

implement sendScore:
- tx_score.go: adjust to fetch the player info instead of getting it via User
- packet_score.go: include logic to validate score data
- msg_server_score.go: include logic for sending the player score
- score.go: include logic for receiving the player score