package gears

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/pkg/errors"
)

func HttpGetBody(url string) (string, error) {
	// Request the HTML page
	raw, err := http.Get(url)
	if err != nil {
		return "", errors.Wrapf(err, "[-] gears.HttpGetBody()>Get() Error!")
	}
	rawBody, err := ioutil.ReadAll(raw.Body)
	defer raw.Body.Close()
	if err != nil {
		return "", errors.Wrap(err, "[-] gears.HttpGetBody()>ReadAll() Error!")
	}
	if raw.StatusCode != 200 {
		return "", errors.Wrap(err, "[-] gears.HttpGetBody()>Get() Error! Message: Cannot open the url.")
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
