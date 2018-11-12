package dota2

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

//string : summary of steam API urls
const (
	BASE_URL                     string = "http://api.steampowered.com/"
	GET_MATCH_HISTORY            string = "IDOTA2Match_570/GetMatchHistory/v001/"
	GET_MATCH_HISTORY_BY_SEQ_NUM string = "IDOTA2Match_570/GetMatchHistoryBySequenceNum/v0001/"
	GET_MATCH_DETAILS            string = "IDOTA2Match_570/GetMatchDetails/v001/"
	GET_LEAGUE_LISTING           string = "IDOTA2Match_205790/GetLeagueListing/v0001/"
	GET_LIVE_LEAGUE_GAMES        string = "IDOTA2Match_570/GetLiveLeagueGames/v0001/"
	GET_TEAM_INFO_BY_TEAM_ID     string = "IDOTA2Match_570/GetTeamInfoByTeamID/v001/"
	GET_PLAYER_SUMMARIES         string = "ISteamUser/GetPlayerSummaries/v0002/"
	GET_FRIEND_LIST              string = "ISteamUser/GetFriendList/v0001/"
	GET_SERVER_INFO              string = "ISteamWebAPIUtil/GetServerInfo/v0001/"
	GET_HEROES                   string = "IEconDOTA2_570/GetHeroes/v0001/"
	GET_GAME_ITEMS               string = "IEconDOTA2_570/GetGameItems/v0001/"
	GET_TOURNAMENT_PRIZE_POOL    string = "IEconDOTA2_570/GetTournamentPrizePool/v1/"
	GET_TOP_LIVE_GAME            string = "IDOTA2Match_570/GetTopLiveGame/v1/"
	BASE_ITEMS_IMAGES_URL        string = "http://cdn.dota2.com/apps/dota2/images/items/"
	BASE_HERO_IMAGES_URL         string = "http://cdn.dota2.com/apps/dota2/images/heroes/"
)

/*
In order to use the Steam Web API, you have to request a key here: https://steamcommunity.com/dev/apikey. This key acts as your secret identifier when making requests to the API, so don't lose or share it.
Requests contain the following elements
Base URL
Usually https://api.steampowered.com but there's no guarantee that it won't change in the future.
Interface Name
Indicates which method group (interface) you want to use. Methods are usually grouped by similarities, but that doesn't seem to be a hard and fast rule.
Method Name
Indicates which method within the interface you want to use.
Method Version
Indicates which version of the method you want to use. Valve will occasionally (almost never) update the API and roll the version number on the method. This allows applications to continue working with older versions while new applications can take advantage of the new versions
Parameters
sometimes optional A one to many list of parameters to be passed to the selected method. Parameters are delimited by the & character.

Request URL format
https://{base_url}/{interface}/{method}/{version}?{parameters}

Sample request URL
http://api.steampowered.com/ISteamWebAPIUtil/GetSupportedAPIList/v1/

Sample request URL with parameters
http://api.steampowered.com/ISteamWebAPIUtil/GetSupportedAPIList/v1/?key=1234567890&steamid=000123000456
*/
var (
	URLMap = map[string]string{
		"GetMatchHistory":         BASE_URL + GET_MATCH_HISTORY,
		"GetMatchHistoryBySeqNum": BASE_URL + GET_MATCH_HISTORY_BY_SEQ_NUM,
		"GetMatchDetails":         BASE_URL + GET_MATCH_DETAILS,
		"GetLeagueListing":        BASE_URL + GET_LEAGUE_LISTING,
		"GetLiveLeagueGames":      BASE_URL + GET_LIVE_LEAGUE_GAMES,
		"GetTeamInfoByTeamId":     BASE_URL + GET_TEAM_INFO_BY_TEAM_ID,
		"GetPlayerSummaries":      BASE_URL + GET_PLAYER_SUMMARIES,
		"GetFriendList":           BASE_URL + GET_FRIEND_LIST,
		"GetServerInfo":           BASE_URL + GET_SERVER_INFO,
		"GetHeroes":               BASE_URL + GET_HEROES,
		"GetGameItems":            BASE_URL + GET_GAME_ITEMS,
		"GetTournamentPrizePool":  BASE_URL + GET_TOURNAMENT_PRIZE_POOL,
		"GetTopLiveGame":          BASE_URL + GET_TOP_LIVE_GAME,
		"BaseItemsImagesUrl":      BASE_ITEMS_IMAGES_URL,
		"BaseHeroImagesUrl":       BASE_HERO_IMAGES_URL,
	}
)

var (
	URLMapError = errors.New("Cannot find correspond URL in URLMap")
)

type Dota2api struct {
	apikey string
	client *http.Client
}

func NewApi(apiclient *http.Client) *Dota2api {
	if apiclient == nil {
		apiclient = http.DefaultClient
	}

	dapi := &Dota2api{
		apikey: "",
		client: apiclient,
	}

	return dapi
}

func (d *Dota2api) SetApiKey(apikey string) {
	d.apikey = apikey
}

//GetMatchHistory : get recent dota2 match history of player for user's dota2 account id(not steam id).
//example:
//	GetMatchHistory("123400001")
//return:
//	the detailed information of certain dota2 match.
func (d *Dota2api) GetMatchHistory(accountid string) (MatchHistory, error) {
	var mh MatchHistory
	url, found := URLMap["GetMatchHistory"]
	if !found {
		return mh, URLMapError
	}

	formurl := url + "?key=" + d.apikey + "&account_id=" + accountid
	bmatchhistory, err := d.RequestForURL(formurl)
	if err != nil {
		return mh, err
	}

	var mhwrap MatchHistoryWrapper
	err = json.Unmarshal(bmatchhistory, &mhwrap)
	if err != nil {
		return mh, err
	}
	mh = mhwrap.Result
	return mh, nil
}

//GetMatchDetails will get match details by match id
func (d *Dota2api) GetMatchDetails(matchid string) (MatchDetail, error) {
	var mdetail MatchDetail
	url, found := URLMap["GetMatchDetails"]
	if !found {
		return mdetail, URLMapError
	}

	formurl := url + "?key=" + d.apikey + "&match_id=" + matchid
	bmatchdetail, err := d.RequestForURL(formurl)
	if err != nil {
		return mdetail, err
	}

	var mdetailwrp MatchDetailWrapper
	err = json.Unmarshal(bmatchdetail, &mdetailwrp)
	if err != nil {
		return mdetail, err
	}
	mdetail = mdetailwrp.Result
	return mdetail, nil
}

//GetLeagueListing will get a list of leagues which can be viewed within DotaTV.
func (d *Dota2api) GetLeagueListing() (LeagueList, error) {
	var leagues LeagueList

	url, found := URLMap["GetLeagueListing"]
	if !found {
		return leagues, URLMapError
	}

	formurl := url + "?key=" + d.apikey
	log.Printf("formurl=%s\n", formurl)
	bleagues, err := d.RequestForURL(formurl)
	if err != nil {
		return leagues, err
	}
	//log.Printf("bytes->\n%s\n", bleagues)

	var leaguelistwrapper LeagueListWrapper
	err = json.Unmarshal(bleagues, &leaguelistwrapper)
	if err != nil {
		return leagues, err
	}

	leagues = leaguelistwrapper.League

	return leagues, nil

}

//GetPlayerSummaries will get basic profile information for 64-bit Steam IDs.
//API itself supports list of commma-delimated steam ids. eg: "71210000,7456222"
//example:
//	GetPlayerSummaries("71210000,7456222")
//return:
//	list of playersummary
func (d *Dota2api) GetPlayerSummaries(steamids string) (PlayerSummaryList, error) {
	var plsummarylist PlayerSummaryList
	url, found := URLMap["GetPlayerSummaries"]
	if !found {
		return plsummarylist, URLMapError
	}

	formurl := url + "?key=" + d.apikey + "&steamids=" + steamids
	bplayersummary, err := d.RequestForURL(formurl)
	if err != nil {
		return plsummarylist, err
	}

	var playersmrwrp PlayerSummaryWrapper
	err = json.Unmarshal(bplayersummary, &playersmrwrp)
	if err != nil {
		return plsummarylist, err
	}
	plsummarylist = playersmrwrp.Response
	return plsummarylist, err
}

//GetFriendList returns the friend list of any Steam user, only if Steam Community profile visibility is set to "Public".
//Nothing will be returned if the profile is private.
//example:
//	GetFriendList("76561198092165728", "friend")
//return:
//	slice of struct FriendInfo
func (d *Dota2api) GetFriendList(steamid string, relationship string) ([]FriendInfo, error) {
	var friendlist []FriendInfo
	url, found := URLMap["GetFriendList"]
	if !found {
		return friendlist, URLMapError
	}

	formurl := url + "?key=" + d.apikey + "&steamid=" + steamid + "&relationship=" + relationship
	bfriendlist, err := d.RequestForURL(formurl)
	if err != nil {
		return friendlist, err
	}

	var frdlistwrap FriendListWrapper
	err = json.Unmarshal(bfriendlist, &frdlistwrap)
	if err != nil {
		return friendlist, err

	}
	friendlist = frdlistwrap.FriendList.Friends

	return friendlist, err
}

//GetServerInfo will return WebAPI Server's time info.
func (d *Dota2api) GetServerInfo() (ServerInfo, error) {
	var srvinfo ServerInfo
	url, found := URLMap["GetServerInfo"]
	if !found {
		return srvinfo, URLMapError
	}

	bsrvinfo, err := d.RequestForURL(url)
	if err != nil {
		return srvinfo, err
	}

	err = json.Unmarshal(bsrvinfo, &srvinfo)
	if err != nil {
		return srvinfo, err
	}

	return srvinfo, nil
}

//GetLiveLeagueGames will return list of the detailed and real-time statistics of the games which are being played.
func (d *Dota2api) GetLiveLeagueGames() (LeagueGames, error) {
	var (
		leaguegameswarp LeagueGamesWrapper
		leaguegames     LeagueGames
	)
	url, found := URLMap["GetLiveLeagueGames"]
	if !found {
		return leaguegames, URLMapError
	}

	formurl := url + "?key=" + d.apikey
	log.Println("formurl->", formurl)
	bleagues, err := d.RequestForURL(formurl)
	if err != nil {
		return leaguegames, err
	}
	log.Printf("bleages->\n%s", bleagues)

	err = json.Unmarshal(bleagues, &leaguegameswarp)
	if err != nil {
		return leaguegames, err
	}

	leaguegames = leaguegameswarp.LgGames
	return leaguegames, nil

}

//RequestForURL will send http request to url and return the result with []byte
func (d *Dota2api) RequestForURL(url string) ([]byte, error) {
	var bresp []byte
	resp, err := d.client.Get(url)
	if err != nil {
		return bresp, err
	}
	defer resp.Body.Close()

	bresp, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return bresp, err
	}

	return bresp, nil
}
