package eapi

import (
    "bytes"
    "bufio"
    "io/ioutil"
)

func RandomNickname() string {
    bn, err := ioutil.ReadFile("nicknames.txt")
    if err != nil {
        return ""
    }

    var names []string
    scanner := bufio.NewScanner(bytes.NewReader(bn))
	for scanner.Scan() {
        names = append(names, scanner.Text())
    }
    
    if len(names) > 0 {
        return names[RandomNumber(0, len(names)-1)]
    }
    return ""
}