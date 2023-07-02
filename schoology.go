package gology

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

//this is going to look remarkably OO for go code
//because I can't think of a better way and it kinda
//works here

type SchoologySession struct {
	http    *http.Client
	rootUrl string
}

func (sc *SchoologySession) Init(sessiontoken, schoolhash string) {
	jar, err := cookiejar.New(nil)
	sessionCookie := &http.Cookie{
		Name:   "SESS" + schoolhash,
		Value:  sessiontoken,
		Path:   "/",
		Domain: sc.rootUrl,
	}
	var cookies []*http.Cookie
	cookies = append(cookies, sessionCookie)
	url, _ := url.Parse(sc.rootUrl)
	jar.SetCookies(url, cookies)
	if err != nil {
		panic(err)
	}
	sc.http = &http.Client{
		Jar: jar,
	}
}

func (sc *SchoologySession) GetNotifications() NotificationResponse {
	res, err := sc.http.Get(fmt.Sprintf("%v/iapi2/site-navigation/notifications", sc.rootUrl))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var nt NotificationResponse
	err = json.Unmarshal(b, &nt)
	if err != nil {
		panic(err)
	}
	return nt
}
