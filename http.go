package gears

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/pkg/errors"
)

// HttpGetBody request url to get Html Body, if get error occur it'll try n times.
func HttpGetBody(url string, n int) (string, error) {
	raw, err := http.Get(url)
	for err != nil && n > 0 {
		raw, err = http.Get(url)
		time.Sleep(time.Minute * 1)
		n--
	}
	if err != nil {
		return "", errors.Wrapf(err, "\n[-] gears.HttpGetBody()>Get() try times, but error occur still!\n[-] ")
	}
	rawBody, err := ioutil.ReadAll(raw.Body)
	defer raw.Body.Close()
	if err != nil {
		return "", errors.Wrap(err, "\n[-] gears.HttpGetBody()>ReadAll() Error!\n[-] ")
	}
	if raw.StatusCode != 200 {
		return "", errors.Wrap(err, "\n[-] gears.HttpGetBody()>Get() Error! Message: Cannot open the url.\n[-] ")
	}
	return string(rawBody), nil
}

// HttpGetTitleViaTwitterJS get post title via twitter share javascripts' json data
func HttpGetTitleViaTwitterJS(rawBody string) string {
	var a = regexp.MustCompile(`(?m)<meta name="twitter:title" content="(?P<title>.*?)"`)
	return a.FindStringSubmatch(rawBody)[1]
}

// HttpGetSiteViaTwitterJS get post site via twitter share javascripts' json data
func HttpGetSiteViaTwitterJS(rawBody string) string {
	var a = regexp.MustCompile(`(?m)<meta name="twitter:site" content="(?P<site>.*?)"`)
	return a.FindStringSubmatch(rawBody)[1]
}

func HttpGetDateViaMeta(rawBody string) string {
	var a = regexp.MustCompile(`(?m)<meta name="parsely-pub-date" content="(?P<date>.*?)".*?/>`)
	return a.FindStringSubmatch(rawBody)[1]
}

func HttpGetDomain(url string) string {
	var a = regexp.MustCompile(`(?m)https?://(\w+.\w+.\w+)/`)
	return a.FindStringSubmatch(url)[1]
}
