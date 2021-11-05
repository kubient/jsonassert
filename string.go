package jsonassert

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func (a *Asserter) checkString(path, act, exp string) {
	a.tt.Helper()

	isExpRegEx, err := isRegEx(exp)
	if err != nil {
		a.tt.Errorf("expected string check for regex error '%s', path: '%s' , exp reg ex: '%s'", err, path, exp)
		return
	}
	if isExpRegEx {

		regex := getReqExPattern(exp)
		if regex == "" {
			a.tt.Errorf("can't get reg ex string from '%s', path: '%s'", exp, path)
			return
		}

		matched, err := regexp.MatchString(regex, act)
		if err != nil {
			a.tt.Errorf("error on matching: '%v' with: '%v' in path: %v", exp, act, path)
			return
		}

		if !matched {
			a.tt.Errorf("does not match by pattern: '%v' with: '%v' path: %v'", exp, act, path)
		}

	} else {
		if act != exp {
			if len(exp+act) < 50 {
				a.tt.Errorf("expected string at '%s' to be '%s' but was '%s'", path, exp, act)
			} else {
				a.tt.Errorf("expected string at '%s' to be\n'%s'\nbut was\n'%s'", path, exp, act)
			}
		}
	}
}

func extractString(s string) (string, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "", fmt.Errorf("cannot parse nothing as string")
	}
	if s[0] != '"' {
		return "", fmt.Errorf("cannot parse '%s' as string", s)
	}
	var str string
	err := json.Unmarshal([]byte(s), &str)
	return str, err
}

const regExField = `^<<<(.+)>>>$`

func isRegEx(str string) (bool, error) {
	return regexp.MatchString(regExField, str)
}

func getReqExPattern(exp string) string {

	r := regexp.MustCompile(regExField)
	match := r.FindStringSubmatch(exp)
	if match != nil {
		return match[1]
	}
	return ""
}
