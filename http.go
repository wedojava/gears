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
	// var reTitle = regexp.MustCompile(`(?m)<title(.*?){0,1}>(?P<title>.*?)</title>`)
	var reTitle = regexp.MustCompile(`(?m)<meta name="twitter:title" content="(?P<title>.*?)"`)
	title := reTitle.FindStringSubmatch(rawBody)[1]

	return title
}
