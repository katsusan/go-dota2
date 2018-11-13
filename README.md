# A Dota2 SDK in Golang

a dota2 api for getting dota2 player/match information.(some steam api also integrated.)

## Get it ##

Use `go get -u github.com/Katsusan/go-dota2` to get or update it.

## Usage ##

### Quick start ###

Here is the example to get dota2 match history by account id.

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
    fmt.Println("天辉队伍：",matchdetail.DireName)
    fmt.Println("夜魇队伍：",matchdetail.RadiantName)
    fmt.Println("比赛时长(s)：",matchdetail.Duration)
    fmt.Println("一血时间(s)：",matchdetail.FirstBloodTime)
}

```

