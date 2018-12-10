# A Dota2 SDK in Golang

a dota2 api for getting dota2 player/match information.(some steam api also implemented.)

## Get it ##

Use `go get -u github.com/Katsusan/go-dota2` to get or update it.

## Usage ##

### Quick start ###

Here is the example to get dota2 match history by match id.

```go
package main

import (
    "fmt"
    dota2api "github.com/Katsusan/go-dota2"
)

func main() {
    dapi := dota2api.NewApi(nil)
    dapi.SetApiKey("AAFB3717E64F8A3C51200A3F7F7988F8") // 这里替换为自己申请的apikey
    matchdetail, err :=  dapi.GetMatchDetails("4080856812") //以TI8决赛第五场为例
    if err != nil {
        //....
    }
    fmt.Println("天辉队伍：",matchdetail.RadiantName)  //PSG.LGD
    fmt.Println("夜魇队伍：",matchdetail.DireName)  //OG
    fmt.Println("比赛时长(s)：",matchdetail.Duration) //2188
    fmt.Println("一血时间(s)：",matchdetail.FirstBloodTime) //153
}

```

## Supported API ##
- GetMatchHistory(根据指定账号ID获取历史比赛)
  - [x] Status (状态码，意义未知)
  - [x] ResultNum (本次响应中的比赛数量)
  - [x] TotalNum  (单次查询的总比赛数)
  - [x] RemainNum (后续查询会返回的比赛数)
  - [x] Matches (简易比赛信息，参照MatchInfo)

- GetMatchDetails(根据指定比赛ID获取比赛详细信息)
    - [x] MatchID (比赛ID)
    - [x] MatchSeqNum (比赛记录的位置)
    - [x] RadiantWin (天辉赢为True，否则为False)
    - [x] PreGameDuration ()
    - [x] Duration (比赛时长)
    - [x] StartTime (比赛开始时间)
    - [x] FirstBloodTime (一血时间)
    - [x] HumanPlayers (人类玩家数量)
    - [x] LeagueID (联赛ID)
    - [x] PostiveVotes (点赞数)
    - [x] NegativeVotes (反对数)
    - [x] GameMode (比赛模式)
    - [x] LobbyType (比赛类型)
    - [x] RadiantCaptain (天辉队长账号ID)
    - [x] DireCaptain (夜魇队长账号ID)
    - [x] TowerStatusRadiant (天辉防御塔状况)
    - [x] TowerStatusDire (夜魇防御塔状况)
    - [x] BarracksStatusRadiant (天辉兵营状况)
    - [x] BarracksStatusDire (夜魇兵营状况)
    - [x] Cluster (比赛所在的服务器集群，用于获取录像)
    - [x] Engine
    - [x] PickBans (BP一览)
    - [x] Players (选手一览)
    - [x] RadiantTeamID (天辉队伍ID)
    - [x] DireTeamID (夜魇队伍ID)
    - [x] RadiantTeamComplete
    - [x] DireTeamComplete
    - [x] DireName (夜魇队伍名)
    - [x] RadiantName (天辉队伍名)
    - [x] Flags
    - [x] RadiantScore (天辉评分)
    - [x] DireScore (夜魇评分)
    - [x] DireLogo (夜魇队伍logo地址)
    - [x] RadiantLogo (天辉队伍logo地址)

- GetLeagueListing(获取联赛一览)
    - Leagues (联赛列表)
        - [x] Name (联赛名称)
        - [x] LeagueID (联赛ID) 
        - [x] Description (联赛描述)
        - [x] TournamentURL (联赛URL)
        - [x] ItemDef (联赛相关的物品ID)



- GetPlayerSummaries(获取选手信息一览)
    - PlayerSummary (选手概要)
        - [x] SteamID (选手steam ID)
        - [x] CommunityVisibilityState (社区个人信息是否对他人可见)
        - [x] ProfileState (个人档案状态)
        - [x] PersonaName (昵称)
        - [x] LastLogoff (自从上次下线steam后所经历的时间)
        - [x] ProfileURL (个人档案URL)
        - [x] Avatar (头像图片URL)
        - [x] AvatarMedium (头像(中)图片URL)
        - [x] AvatarFull (头像(大)图片URL)
        - [x] PersonaState (在线状态，0->Offline, 1->Online, 2->Busy, 3->Away, 4->Snooze, 5->Looking to trade, 6->Looking to play)
        - [x] PrimaryClanID ()
        - [x] TimeCreated (个人档案创建时间)
        - [x] PersonaStateFlags (个人状态flag)


- GetLiveLeagueGames(获取进行中的比赛信息)






## relevant structures ##

- MatchInfo(比赛信息)
    -  MatchID (比赛ID)
    -  StartTime (开始时间)
    -  LobbyType (比赛类型)
    -  RadiantTeamID (天辉队伍ID)
    -  DireTeamID (夜魇队伍ID)
    -  Player (选手信息)
        + AccountID (dota2账号ID)
        + PlayerSlot (选手游戏里的位置)
        + HeroID (英雄ID)


    

