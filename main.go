package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	userAgent := generateUserAgent("windows", "chrome")
	fmt.Println(userAgent)
}

func generateUserAgent(os, browser string) string {
	osChoices := map[string][]string{
		"windows": {"Windows NT 6.1", "Windows NT 6.2", "Windows NT 10.0"},
		"mac":     {"Macintosh; Intel Mac OS X 10_15_7", "Macintosh; Intel Mac OS X 10_14_6"},
		"linux":   {"X11; Linux x86_64", "X11; Ubuntu; Linux i686"},
	}

	browserChoices := map[string]string{
		"chrome":  "Mozilla/5.0 ({os}) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/{chrome_version} Safari/537.36",
		"firefox": "Mozilla/5.0 ({os}; rv:{firefox_version}) Gecko/20100101 Firefox/{firefox_version}",
		"safari":  "Mozilla/5.0 ({os}) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/{safari_version} Safari/605.1.15",
	}

	if os == "all" {
		os = randomChoiceString(osChoices)
	}

	if browser == "all" {
		browser = randomChoiceString(browserChoices)
	}

	osChoice := randomChoice(osChoices[os])

	if browser == "chrome" {
		chromeVersion := strconv.Itoa(rand.Intn(27)+64) + ".0." + strconv.Itoa(rand.Intn(9000)+1000) + ".0"
		return replacePlaceholders(browserChoices["chrome"], osChoice, chromeVersion)
	}

	if browser == "firefox" {
		firefoxVersion := strconv.Itoa(rand.Intn(41)+50) + ".0"
		return replacePlaceholders(browserChoices["firefox"], osChoice, firefoxVersion)
	}

	if browser == "safari" {
		safariVersion := strconv.Itoa(rand.Intn(4)+12) + ".0." + strconv.Itoa(rand.Intn(7)+1)
		return replacePlaceholders(browserChoices["safari"], osChoice, safariVersion)
	}

	return ""
}

func randomChoice(arr []string) string {
	return arr[rand.Intn(len(arr))]
}

func randomChoiceString(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys[rand.Intn(len(keys))]
}

func replacePlaceholders(template, os, version string) string {
	template = replacePlaceholder(template, "{os}", os)
	template = replacePlaceholder(template, "{chrome_version}", version)
	template = replacePlaceholder(template, "{firefox_version}", version)
	template = replacePlaceholder(template, "{safari_version}", version)
	return template
}

func replacePlaceholder(s, placeholder, value string) string {
	return strings.Replace(s, placeholder, value, -1)
}
