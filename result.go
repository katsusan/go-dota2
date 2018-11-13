package dota2

const (
	PICKBANCOUNT = 22
	PLAYERCOUNT  = 10

	LOBBYTYPE_INVALID              = -1
	LOBBYTYPE_PUBLIC_MATCHMAKING   = 0
	LOBBYTYPE_PRACTICE             = 1
	LOBBYTYPE_TOURNAMENT           = 2
	LOBBYTYPE_TUTORIAL             = 3
	LOBBYTYPE_COOPERATIION_WITH_AI = 4
	LOBBYTYPE_TEAM_MATCH           = 5
	LOBBYTYPE_SOLO_QUEUE           = 6
	LOBBYTYPE_RANKED_MATCHMAKING   = 7
	LOBBYTYPE_SOLO_MID             = 8

	SERIESTYPE_NONSERIES = 0
	SERIESTYPE_BESTOF3   = 1
	SERIESTYPE_BESTOF5   = 2

	LEAGUETIER_AMATEUR      = 1
	LEAGUETIER_PROFESSIONAL = 2
	LEAGUETIER_PREMIER      = 3

	GAMEMODE_UNKNOWN                = 0
	GAMEMODE_ALL_PICK               = 1
	GAMEMODE_CAPTAIN_MODE           = 2
	GAMEMODE_RANDOM_DRAFT           = 3
	GAMEMODE_SINGLE_DRAFT           = 4
	GAMEMODE_ALL_RADNOM             = 5
	GAMEMODE_INTRO                  = 6
	GAMEMODE_DIRETIDE               = 7
	GAMEMODE_REVERSE_CAPTAIN_MODE   = 8
	GAMEMODE_THE_GREEVILING         = 9
	GAMEMODE_TUTORIAL               = 10
	GAMEMODE_MID_ONLY               = 11
	GAMEMODE_LEAST_PLAYED           = 12
	GAMEMODE_NEW_PLAYER_POOL        = 13
	GAMEMODE_COMPENDIUM_MATCHMAKING = 14
	GAMEMODE_CUSTOM                 = 15
	GAMEMODE_CAPTAIN_DRAFT          = 16
	GAMEMODE_BALANCED_DRAFT         = 17
	GAMEMODE_ABILITY_DRAFT          = 18
	GAMEMODE_EVENT                  = 19
	GAMEMODE_ALL_RANDOM_DEATH_MATCH = 20
	GAMEMODE_SOLO_MID_1V1           = 21
	GAMEMODE_RANKED_ALL_PICK        = 22

	LEAVERSTATUS_NONE                     = 0
	LEAVERSTATUS_DISCONNECTED             = 1
	LEAVERSTATUS_DISCONNECTED_TOO_LONG    = 2
	LEAVERSTATUS_ABANDONED                = 3
	LEAVERSTATUS_AFK                      = 4
	LEAVERSTATUS_NEVER_CONNECTED          = 5
	LEAVERSTATUS_NEVER_CONNECTED_TOO_LONG = 6
)

type MatchHistoryWrapper struct {
	Result MatchHistory `json:"result"`
}

type MatchHistory struct {
	Status    int         `json:"status"`
	ResultNum int         `json:"num_results"`       //Number of matches within a single response
	TotalNum  int         `json:"total_results"`     //Total number of matches for this query
	RemainNum int         `json:"results_remaining"` //Number of matches remaining to be retrieved with subsequent API calls
	Matches   []MatchInfo `json:"matches"`           //Brief Information of a match
}

type MatchInfo struct {
	MatchID       int64 `json:"match_id"`        //Unique match ID
	MatchSeqNum   int64 `json:"match_seq_num"`   //Number indicating position in which this match was recorded
	StartTime     int32 `json:"start_time"`      //Unix timestamp of beginning of match
	LobbyType     int   `json:"lobby_type"`      //See const LOBBYTYPE_xx
	RadiantTeamID int   `json:"radiant_team_id"` //Unique Team ID
	DireTeamID    int   `json:"dire_team_id"`    //Unique Team ID
	Player        []struct {
		AccountID  int `json:"account_id"`  //Unique account ID
		PlayerSlot int `json:"player_slot"` //Player's position within the team
		HeroID     int `json:"hero_id"`     //Unique hero ID
	} `json:"players"`
}

type MatchDetailWrapper struct {
	Result MatchDetail `json: "result"`
}

type MatchDetail struct {
	MatchID               int64           `json:"match_id"`      //Unique match ID
	MatchSeqNum           int64           `json:"match_seq_num"` //Number indicating position in which this match was recorded
	RadiantWin            bool            `json:"radiant_win"`   //Win status of game,True for Radiant Win, False for Dire Win
	PreGameDuration       int             `json:"pre_game_duration"`
	Duration              int             `json:"duration"`                //Elapsed match time in seconds
	StartTime             int64           `json:"start_time"`              //Unix timestamp for beginning of match
	FirstBloodTime        int             `json:"first_blood_time"`        //Time elapsed in seconds since first blood of the match
	HumanPlayers          int             `json:"human_players"`           //Number of human players in the match
	LeagueID              int             `json:"leagueid"`                //Unique league ID
	PostiveVotes          int             `json:"positive_votes"`          //Number of positive/thumbs up votes
	NegativeVotes         int             `json:"negative_votes"`          //Number of negative/thumbs down votes
	GameMode              int             `json:"game_mode"`               //match mode, see consts GAMEMODE_xx, eg: 3 -> Random Draft
	LobbyType             int             `json:"lobby_type"`              //match type, see lobbies.json, eg: 7	-> Ranked matchmaking,天梯匹配
	RadiantCaptain        int64           `json:"radiant_captain"`         //Account ID for Radiant Captain
	DireCaptain           int64           `json:"dire_captain"`            //Account ID for Dire Captain
	TowerStatusRadiant    int             `json:"tower_status_radiant"`    //Status of Radiant Towers
	TowerStatusDire       int             `json:"tower_status_dire"`       //Status of Dire Towers
	BarracksStatusRadiant int             `json:"barracks_status_radiant"` //Status of Radiant barracks
	BarracksStatusDire    int             `json:"barracks_status_dire"`    //Status of Dire barracks
	Cluster               int             `json:"cluster"`                 //The server cluster the match was played on, used in retrieving replays
	Engine                int             `json:"engine"`
	PickBans              PickBanItem     `json:"picks_bans"`
	Players               PlayerStatistic `json:"players"`
	RadiantTeamID         int             `json:"radiant_team_id"`       //Radiant Team's unique ID
	DireTeamID            int             `json:"dire_team_id"`          //Dire Team's unique ID
	RadiantTeamComplete   int             `json:"radiant_team_complete"` //unknown field...
	DireTeamComplete      int             `json:"dire_team_complete"`
	DireName              string          `json:"dire_name"`     //Name of Dire Team
	RadiantName           string          `json:"radiant_name"`  //Name of Radiant Team
	Flags                 int             `json:"flags"`         //unknown field...
	RadiantScore          int             `json:"radiant_score"` //Match score of Radiant team
	DireScore             int             `json:"dire_score"`
	DireLogo              int64           `json:"dire_logo"`
	RadiantLogo           int64           `json:"radiant_logo"`
}

//PickBanItem is the data type which describes the pick/ban result.
//total bans -> 6 + 6, total picks -> 5 + 5, so usually 22 pick/ban results were got.
//map format: map[is_pick:false hero_id:67 team:0 order:0]
//				is_pick -> false means ban, true means pick
//				hero_id -> unique hero id
//				team	-> 0 means Radiant, 1 means Dire, 2 means Broadcaster, 3+ are not assigned
//				order	-> from 0 to 21, show the order of overall sequence of pick/ban
type PickBanItem [PICKBANCOUNT]map[string]interface{}

type PlayerStatistic [PLAYERCOUNT]map[string]interface{}

type LeagueListWrapper struct {
	League LeagueList `json:"result"`
}

type LeagueList struct {
	Leagues []struct {
		Name          string `json:"name"`           //Name of the league
		LeagueID      int    `json:"leagueid"`       //Unique league ID
		Description   string `json:"description"`    //Description of the league
		TournamentURL string `json:"tournament_url"` //League website information
		ItemDef       int    `json:"itemdef"`        //ID for an item associated with the tournament
	} `json:"leagues"`
}

type LeagueGamesWrapper struct {
	LgGames LeagueGames `json:"result"`
}

type LeagueGames struct {
	Leagues []LeagueGame `json:"games"`
	Status  uint16       `json:"status"`
}

type LeagueGame struct {
	Players []struct {
		AccountID uint64 `json:"account_id"`
		Name      string `json:"name"`
		HeroID    uint16 `json:"hero_id"`
		Team      uint8  `json:"team"`
	} `json:"players"`

	RadiantTeam struct {
		TeamName string `json:"team_name"`
		TeamID   uint64 `json:"team_id"`
		TeamLogo uint64 `json:"team_logo"`
		Complete bool   `json:"complete"`
	} `json:"radiant_team"`

	DireTeam struct {
		TeamName string `json:"team_name"`
		TeamID   uint64 `json:"team_id"`
		TeamLogo uint64 `json:"team_logo"`
		Complete bool   `json:"complete"`
	} `json:"dire_team"`

	LobbyID           uint64 `json:"lobby_id"`
	MatchID           uint64 `json:"match_id"`
	Spectators        uint32 `json:"spectators"`
	LeagueID          uint64 `json:"league_id"`
	LeagueNodeID      uint64 `json:"league_node_id"`
	StreamDelaySec    uint   `json:"stream_delay_s"`
	RadiantSeriesWins uint   `json:"radiant_series_wins"`
	DireSeriesWins    uint   `json:"dire_series_wins"`
	SeriesType        uint   `json:"series_type"`
	ScoreBoard        struct {
		Duration           float64       `json:"duration"`
		RoshanRespawnTimer int           `json:"roshan_respawn_timer"`
		Radiant            TeamStatistic `json:"radiant"`
		Dire               TeamStatistic `json:"dire"`
	} `json:"scoreboard"`
}

type TeamStatistic struct {
	Score         uint16 `json:"score"`
	TowerState    int64  `json:"tower_state"`
	BarracksState int32  `json:"barracks_state"`
	Picks         []struct {
		HeroID uint16 `json:"hero_id"`
	} `json:"picks"`
	Bans []struct {
		HeroID uint16 `json:"hero_id"`
	} `json:"bans"`
	Players []struct {
		PlayerSlot       uint8   `json:"player_slot"`
		AccountID        uint64  `json:"account_id"`
		HeroID           uint16  `json:"hero_id"`
		Kills            uint16  `json:"kills"`
		Death            uint16  `json:"death"`
		Assists          uint16  `json:"assists"`
		LastHits         uint16  `json:"last_hits"`
		Denies           uint16  `json:"denies"`
		Gold             uint32  `json:"gold"`
		Level            uint16  `json:"level"`
		GoldPerMin       uint16  `json:"gold_per_min"`
		XpPerMin         uint16  `json:"xp_per_min"`
		UltimateState    uint8   `json:"ultimate_state"`
		UltimateCoolDown uint8   `json:"ultimate_cooldown"`
		Item0            uint16  `json:"item0"`
		Item1            uint16  `json:"item1"`
		Item2            uint16  `json:"item2"`
		Item3            uint16  `json:"item3"`
		Item4            uint16  `json:"item4"`
		Item5            uint16  `json:"item5"`
		RespawnTimer     uint16  `json:"respawn_timer"`
		PositionX        float32 `json:"position_x"`
		PositionY        float32 `json:"position_y"`
		NetWorth         uint32  `json:"net_worth"`
	} `json:"players"`
	/*Abilities []struct {  //由于Valve返回不规范的json格式，此处会返回0-5个重复的Abilities键，暂且不解析
		AbilityID    uint16 `json:"ability_id"`
		AbilityLevel uint8  `json:"ability_level"`
	} `json:"abilities"`*/
}

type PlayerSummaryWrapper struct {
	Response PlayerSummaryList `json:"response"`
}

type PlayerSummaryList struct {
	PlayerSummary []PlayerSummary `json:"players"`
}

type PlayerSummary struct {
	SteamID                  string `json:"steamid"`                  //Unique Steam ID
	CommunityVisibilityState int    `json:"communityvisibilitystate"` //1->Private, 2->Friends only, 3->Friends of friends, 4->Users only, 5->Public
	ProfileState             int    `json:"profilestate"`             //unknown
	PersonaName              string `json:"personaname"`              //Equivalent of Steam username
	LastLogoff               int32  `json:"lastlogoff"`               //Unix timestamp since last time logged out of steam
	ProfileURL               string `json:"profileurl"`               //Steam profile URL
	Avatar                   string `json:"avatar"`                   //32x32 avatar image
	AvatarMedium             string `json:"avatarmedium"`             //64x64 avatar image
	AvatarFull               string `json:"avatarfull"`               //184x184 avatar image
	PersonaState             int    `json:"personastate"`             //0->Offline, 1->Online, 2->Busy, 3->Away, 4->Snooze, 5->Looking to trade, 6->Looking to play
	PrimaryClanID            string `json:"primaryclanid"`            // 64-bit unique clan identifier
	TimeCreated              int32  `json:"timecreated"`              // Unix timestamp of profile creation time
	PersonaStateFlags        int    `json:"personastateflags"`        //unknown
}

type FriendListWrapper struct {
	FriendList struct {
		Friends []FriendInfo `json:"friends"`
	} `json:"friendslist"`
}

type FriendInfo struct {
	SteamID      string `json:"steamid"`      //64 bit Steam ID of the friend
	RelationShip string `json:"relationship"` //Relationship qualifier
	FriendSince  int32  `json:"friend_since"` //Unix timestamp of the time when the relationship was created
}

type ServerInfo struct {
	ServerTime       int64  `json:"servertime"`       // Unix timestamp of WebAPI server.
	ServerTimeString string `json:"servertimestring"` //time string of WebAPI server.
}
