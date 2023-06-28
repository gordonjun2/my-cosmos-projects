package types

import (
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type winningPlayerParsed struct {
	PlayerAddress string
	WonCount      uint64
	DateAdded     time.Time
}

func (winningPlayer *WinningPlayer) GetDateAddedAsTime() (DateAdded time.Time, err error) {
	DateAdded, errDateAdded := time.Parse(DateAddedLayout, winningPlayer.DateAdded)
	return DateAdded, sdkerrors.Wrapf(errDateAdded, ErrInvalidDateAdded.Error(), winningPlayer.DateAdded)
}

func GetDateAdded(ctx sdk.Context) time.Time {
	return ctx.BlockTime()
}

func FormatDateAdded(DateAdded time.Time) string {
	return DateAdded.UTC().Format(DateAddedLayout)
}

func (winningPlayer *WinningPlayer) parse() (parsed *winningPlayerParsed, err error) {
	DateAdded, err := winningPlayer.GetDateAddedAsTime()
	if err != nil {
		return nil, err
	}
	return &winningPlayerParsed{
		PlayerAddress: winningPlayer.PlayerAddress,
		WonCount:      winningPlayer.WonCount,
		DateAdded:     DateAdded,
	}, nil
}

func (parsed *winningPlayerParsed) stringify() (stringified *WinningPlayer) {
	return &WinningPlayer{
		PlayerAddress: parsed.PlayerAddress,
		WonCount:      parsed.WonCount,
		DateAdded:     FormatDateAdded(parsed.DateAdded),
	}
}

func (leaderboard *Leaderboard) parseWinners() (winners []*winningPlayerParsed, err error) {
	winners = make([]*winningPlayerParsed, len(leaderboard.Winners))
	var parsed *winningPlayerParsed
	for index, winningPlayer := range leaderboard.Winners {
		parsed, err = winningPlayer.parse()
		if err != nil {
			return nil, err
		}
		winners[index] = parsed
	}
	return winners, nil
}

func stringifyWinners(winners []*winningPlayerParsed) (stringified []*WinningPlayer) {
	stringified = make([]*WinningPlayer, len(winners))
	for index, winner := range winners {
		stringified[index] = winner.stringify()
	}
	return stringified
}

// The goal is to sort with the highest score at index 0 and then descending lower scores.
// Plus, for equal scores, to have new ones closer to index 0 than older ones.
func sortWinners(winners []*winningPlayerParsed) {
	sort.SliceStable(winners[:], func(i, j int) bool {
		if winners[i].WonCount > winners[j].WonCount {
			return true
		}
		if winners[i].WonCount < winners[j].WonCount {
			return false
		}
		return winners[i].DateAdded.After(winners[j].DateAdded)
	})
}

func AddParsedCandidatesAndSort(parsedWinners []*winningPlayerParsed, candidates []*winningPlayerParsed) (updated []*winningPlayerParsed) {
	
	found:= false
	// if it is end of a game instead of a genesis import
	if (len(candidates) == 1) {
		for i := range parsedWinners {
	    if parsedWinners[i].PlayerAddress == candidates[0].PlayerAddress {
	        // found the player, update the fields
	        parsedWinners[i].WonCount = candidates[0].WonCount
	        parsedWinners[i].DateAdded = candidates[0].DateAdded
	        found= true
	        break
	    }
		}
	}

	if(found) {
		updated = parsedWinners
	} else {
		updated = append(parsedWinners, candidates...)
	}
	
	sortWinners(updated)
	if LeaderboardWinnerLength < len(updated) {
		updated = updated[:LeaderboardWinnerLength]
	}
	return updated
}

func (leaderboard *Leaderboard) AddCandidatesAndSortAtNow(now time.Time, playerInfos []*PlayerInfo) (err error) {
	parsedWinners, err := leaderboard.parseWinners()
	if err != nil {
		return err
	}
	parsedPlayers := make([]*winningPlayerParsed, len(playerInfos))
	for index, playerInfo := range playerInfos {
		parsedPlayers[index] = &winningPlayerParsed{
			PlayerAddress: playerInfo.Index,
			WonCount:      playerInfo.WonCount,
			DateAdded:     now,
		}
	}
	parsedWinners = AddParsedCandidatesAndSort(parsedWinners, parsedPlayers)
	leaderboard.Winners = stringifyWinners(parsedWinners)
	return nil
}

func (leaderboard *Leaderboard) AddCandidatesAndSort(ctx sdk.Context, playerInfos []*PlayerInfo) (err error) {
	return leaderboard.AddCandidatesAndSortAtNow(GetDateAdded(ctx), playerInfos)
}

func (leaderboard *Leaderboard) AddCandidateAndSort(ctx sdk.Context, playerInfo PlayerInfo) (err error) {
	return leaderboard.AddCandidatesAndSort(ctx, []*PlayerInfo{&playerInfo})
}
