package dota2

import (
	"fmt"
	"log"
	"math"
	"testing"
	"time"
)

func TestGetMatchHistory(t *testing.T) {
	dapi := NewApi(nil)
	dapi.SetApiKey("E09635A9F555CE8F0B0CCEECE8E40434")
	mh, err := dapi.GetMatchHistory("131900000d")
	if err != nil {
		t.Errorf("GetMatchHistory failed,%v\n", err)
	}
	log.Println(mh.Status)
	if mh.Status != 1 {
		t.Errorf("When invalid characters are in dota2 id , MatchHistory.Status should be 1, Got:%d\n", mh.Status)
	}

	mhabnor, err := dapi.GetMatchHistory("1531490000111111110")
	if err != nil {
		t.Errorf("GetMatchHistory failed,%v\n", err)
	}
	if mhabnor.Status != 15 {
		log.Printf("%+v\n", mhabnor)
		t.Errorf("Can't get matchhistory, with right dota2 id.\n")

	}

	mhnor, err := dapi.GetMatchHistory("131900000")
	if err != nil {
		t.Errorf("GetMatchHistory failed,%v\n", err)
	}
	if mhnor.Status != 1 {
		t.Errorf("When got the right results, MatchHistory.Status should be 1, Got:%d\n", mhnor.Status)
	}
}

func TestGetMatchDetails(t *testing.T) {
	dapi := NewApi(nil)
	dapi.SetApiKey("E09635A9F555CE8F0B0CCEECE8E40434")
	mtd, err := dapi.GetMatchDetails("4080856812") // Match:4080856812 -> the 5th match of TI8 Grand Final(BO5)
	if err != nil {
		t.Errorf("GetMatchDetails request failed.")
	}

	// frequently used statistics, eg: matchid, team name...
	TI8GrandFinal5th := &MatchDetail{
		MatchID:     4080856812,
		DireName:    "OG",
		RadiantName: "PSG.LGD",
	}

	if mtd.MatchID != TI8GrandFinal5th.MatchID {
		t.Errorf("MatchId not correct, Got:%d, Expected:%d.\n", mtd.MatchID, TI8GrandFinal5th.MatchID)
	}

	if mtd.DireName != TI8GrandFinal5th.DireName {
		t.Errorf("DireName not correct, Got:%s, Expected:%s.\n", mtd.DireName, TI8GrandFinal5th.DireName)
	}

	if mtd.RadiantName != TI8GrandFinal5th.RadiantName {
		t.Errorf("RadiantName not correct, Got:%s, Expected:%s.\n", mtd.RadiantName, TI8GrandFinal5th.RadiantName)
	}
}

func TestGetPlayerSummaries(t *testing.T) {
	dapi := NewApi(nil)
	dapi.SetApiKey("E09635A9F555CE8F0B0CCEECE8E40434")
	psummarylist, err := dapi.GetPlayerSummaries("76561198092165728,76561197960435530")

	if err != nil {
		fmt.Println(err)
		t.Errorf("GetPlayerSummaries request failed.")
	}

	for _, plsum := range psummarylist.PlayerSummary {
		if plsum.SteamID != "76561198092165728" && plsum.SteamID != "76561197960435530" {
			t.Errorf("Got Unexpected SteamId: %s, Expected: 76561198092165728 or 76561197960435530", plsum.SteamID)
		}
	}
}

func TestGetFriendList(t *testing.T) {
	dapi := NewApi(nil)
	dapi.SetApiKey("E09635A9F555CE8F0B0CCEECE8E40434")
	frdlist, err := dapi.GetFriendList("76561198092165728", "friend")
	if err != nil {
		t.Errorf("GetFriendList request failed.%s\n", err)
	}

	myoldfrd := FriendInfo{
		SteamID:      "76561198096441766",
		RelationShip: "friend",
		FriendSince:  1389535609,
	}

	foundflag := false
	for _, frd := range frdlist {
		if frd.SteamID == myoldfrd.SteamID && frd.RelationShip == myoldfrd.RelationShip && frd.FriendSince == myoldfrd.FriendSince {
			foundflag = true
		}
	}

	if !foundflag {
		t.Errorf("Got friends list:%+v\nExpected instance not found, %+v\n", frdlist, myoldfrd)
	}
}

func TestGetLeagueListing(t *testing.T) {
	dapi := NewApi(nil)
	dapi.SetApiKey("E09635A9F555CE8F0B0CCEECE8E40434")
	leagues, err := dapi.GetLeagueListing()
	if err != nil {
		t.Errorf("GetLeagueListing failed.%s\n", err)
	}
	log.Printf("leagues->\n%+v", leagues.Leagues)

	//Look for TI8:
	//		"Name":"#DOTA_Item_The_International_2018",
	//		"leagueid":9870,
	//		"description":"#DOTA_Item_Desc_The_International_2018",
	//		"tournament_url":"http://www.dota2.com/international/overview/",
	//		"itemdef":17428
	foundflag := false
	for _, league := range leagues.Leagues {
		if league.Name == "#DOTA_Item_The_International_2018" &&
			league.LeagueID == 9870 &&
			league.Description == "#DOTA_Item_Desc_The_International_2018" &&
			league.TournamentURL == "http://www.dota2.com/international/overview/" &&
			league.ItemDef == 17428 {
			foundflag = true
			break
		}
	}
	if !foundflag {
		t.Errorf("The International 2018 not found.\n")
	}
}

func TestGetLiveLeagueGames(t *testing.T) {
	dapi := NewApi(nil)
	dapi.SetApiKey("E09635A9F555CE8F0B0CCEECE8E40434")

	leagues, err := dapi.GetLiveLeagueGames()
	if err != nil {
		t.Errorf("GetLiveLeagueGames failed.%s\n", err)
	}

	if leagues.Status == 200 && len(leagues.Leagues) == 0 {
		t.Errorf("Got response(status=200) but no parse results.\n")
	}

}

func TestGetServerInfo(t *testing.T) {
	dapi := NewApi(nil)

	srverinfo, err := dapi.GetServerInfo()
	if err != nil {
		t.Errorf("GetServerInfo failed, %s\n", err)
	}
	tm := time.Now().Unix()
	if math.Abs(float64(srverinfo.ServerTime)-float64(tm)) > 60 { //usually the difference between local and server time should < 60s
		t.Errorf("Server time(%d) isn't consistent with localtime(%d).\n", srverinfo.ServerTime, tm)
	}
}
