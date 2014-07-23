package main

import (
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"net/http"
	"strings"
)

type Tweet struct {
	Contributors interface{} `json:"contributors"`
	Coordinates  interface{} `json:"coordinates"`
	CreatedAt    string      `json:"created_at"`
	Entities     struct {
		Hashtags     []interface{} `json:"hashtags"`
		Symbols      []interface{} `json:"symbols"`
		Urls         []interface{} `json:"urls"`
		UserMentions []interface{} `json:"user_mentions"`
	} `json:"entities"`
	FavoriteCount        float64     `json:"favorite_count"`
	Favorited            bool        `json:"favorited"`
	FilterLevel          string      `json:"filter_level"`
	Geo                  interface{} `json:"geo"`
	ID                   float64     `json:"id"`
	IdStr                string      `json:"id_str"`
	InReplyToScreenName  string      `json:"in_reply_to_screen_name"`
	InReplyToStatusID    float64     `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr string      `json:"in_reply_to_status_id_str"`
	InReplyToUserID      float64     `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string      `json:"in_reply_to_user_id_str"`
	Lang                 string      `json:"lang"`
	Place                interface{} `json:"place"`
	RetweetCount         float64     `json:"retweet_count"`
	Retweeted            bool        `json:"retweeted"`
	Source               string      `json:"source"`
	Text                 string      `json:"text"`
	Truncated            bool        `json:"truncated"`
	User                 TwitterUser `json:"user"`
}

type TwitterUser struct {
	ContributorsEnabled            bool        `json:"contributors_enabled"`
	CreatedAt                      string      `json:"created_at"`
	DefaultProfile                 bool        `json:"default_profile"`
	DefaultProfileImage            bool        `json:"default_profile_image"`
	Description                    string      `json:"description"`
	FavouritesCount                float64     `json:"favourites_count"`
	FollowRequestSent              interface{} `json:"follow_request_sent"`
	FollowersCount                 float64     `json:"followers_count"`
	Following                      interface{} `json:"following"`
	FriendsCount                   float64     `json:"friends_count"`
	GeoEnabled                     bool        `json:"geo_enabled"`
	ID                             float64     `json:"id"`
	IdStr                          string      `json:"id_str"`
	IsTranslationEnabled           bool        `json:"is_translation_enabled"`
	IsTranslator                   bool        `json:"is_translator"`
	Lang                           string      `json:"lang"`
	ListedCount                    float64     `json:"listed_count"`
	Location                       string      `json:"location"`
	Name                           string      `json:"name"`
	Notifications                  interface{} `json:"notifications"`
	ProfileBackgroundColor         string      `json:"profile_background_color"`
	ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string      `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileBannerURL               string      `json:"profile_banner_url"`
	ProfileImageURL                string      `json:"profile_image_url"`
	ProfileImageUrlHttps           string      `json:"profile_image_url_https"`
	ProfileLinkColor               string      `json:"profile_link_color"`
	ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string      `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
	Protected                      bool        `json:"protected"`
	ScreenName                     string      `json:"screen_name"`
	StatusesCount                  float64     `json:"statuses_count"`
	TimeZone                       string      `json:"time_zone"`
	URL                            string      `json:"url"`
	UtcOffset                      float64     `json:"utc_offset"`
	Verified                       bool        `json:"verified"`
}

type FollowList struct {
	NextCursor        float64       `json:"next_cursor"`
	NextCursorStr     string        `json:"next_cursor_str"`
	PreviousCursor    float64       `json:"previous_cursor"`
	PreviousCursorStr string        `json:"previous_cursor_str"`
	Users             []TwitterUser `json:"users"`
}

func GetFollowers(cursor string, logindata oauth.AccessToken, c *oauth.Consumer) (Flist FollowList) {
	var response *http.Response

	defer func() {
		if Flist.NextCursorStr == "" {
			Flist.NextCursorStr = "0"
		}
	}()
	var e error
	if cursor != "0" {
		response, e = c.Get(
			"https://api.twitter.com/1.1/friends/list.json",
			map[string]string{
				"count":  "200",
				"cursor": cursor,
			},
			&logindata)
	} else {
		response, e = c.Get(
			"https://api.twitter.com/1.1/friends/list.json",
			map[string]string{
				"count": "200",
			},
			&logindata)
	}

	if e != nil {
		fmt.Println(e)
		return Flist
	}

	b, e := ioutil.ReadAll(response.Body)

	if e != nil {
		fmt.Println("Could not read json for followers")
		return Flist
	}

	e = json.Unmarshal(b, &Flist)
	if e != nil {
		fmt.Println("Could not decode json for followers")
	}

	return Flist
}

func ProduceNameList(logindata oauth.AccessToken, c *oauth.Consumer, TM map[string]Tweet) []string {
	Chunks := make([]string, 0)
	Flist := GetFollowers("0", logindata, c)
	Chunks = MakeUserList(Flist, Chunks, TM)

	for Flist.NextCursorStr != "0" {
		Flist = GetFollowers(Flist.NextCursorStr, logindata, c)
		Chunks = MakeUserList(Flist, Chunks, TM)
	}

	return Chunks
}

func MakeUserList(Flist FollowList, input []string, TM map[string]Tweet) []string {
	RunningList := ""
	for c, v := range Flist.Users {
		RunningList = RunningList + " " + v.ScreenName
		T := Tweet{}
		T.User = v
		TM[strings.ToLower(v.ScreenName)] = T
		if c%50 == 0 {
			input = append(input, RunningList)
			RunningList = ""
		}
	}
	input = append(input, RunningList)
	return input
}

type RemovePacket struct {
	Delete struct {
		Status struct {
			ID        float64 `json:"id"`
			IdStr     string  `json:"id_str"`
			UserID    float64 `json:"user_id"`
			UserIdStr string  `json:"user_id_str"`
		} `json:"status"`
	} `json:"delete"`
}
