package eapi

import (
    "fmt"
    
    "bytes"
    "bufio"
    "errors"
    "io/ioutil"
    "sort"
    "strings"
)

func RandomNickname() string {
    bn, err := ioutil.ReadFile("recourses/nicknames.txt")
    if err != nil {
        fmt.Println(err)
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

func (p *Pages) Length(l int) {
    if l > 1950 {
        p.length = 1950
    } else if l < 100 {
        p.length = 100
    } else {
        p.length = l
    }
}

func (p *Pages) StringToPages(msg string) {
    for len(msg) > p.length {
        p.Pages = append(p.Pages, msg[0:p.length])
        msg = msg[p.length:len(msg)]
    }
    p.Pages = append(p.Pages, msg)
}

func (p *Pages) SliceToPages(Data []string, Format string) error {

    if !strings.Contains(Format, "{{Value}}") {
        return errors.New("Function `SliceToPages`, Expected Format to contain '{Value}'.")
    }

    if Data == nil || len(Data) == 0 {
        return errors.New("Function `SliceToPages`, Data is empty.")
    }

    p.Pages = append(p.Pages, "")
    sort.Strings(Data)

    for di := range Data {
        if len(p.Pages[(len(p.Pages) - 1)]) >= p.length {
            p.Pages = append(p.Pages, "")
        }
        
        p.Pages[(len(p.Pages) - 1)] += strings.Replace(Format, "{{Value}}", Data[di], -1)
    }

    return nil
}

func (p *Pages) MapToPages(Data map[string]interface{}, Format string) error {
    //  
    //  Function that is used for commands like 'viewauto', 'profiles', 'viewdbs', 'view -db'
    //  'Length' is the character limit for the codeblock.
    //  'Format' is the format of how a line should look like. (eg. "Trigger: {Key}" or "{Key}: {Value}")
    
    Keys  := []string{}
    
    if (!strings.Contains(Format, "{{Key}}") && !strings.Contains(Format, "{{Value}}")) {
        return errors.New("Function `MapToPages`: Expected Format to contain either \"{{Key}}\" or \"{{Value}}\" or both.")
    }
    
    if Data == nil || len(Data) == 0 {
        return errors.New("Function `MapToPages`: Data is empty.")
    }
    
    for k, _ := range Data {
        Keys = append(Keys, strings.Replace(k, "`", "\\`", -1))
    }
    
    p.Pages = append(p.Pages, "")
    sort.Strings(Keys)
    
    for _, k := range Keys {
        if len(p.Pages[(len(p.Pages) - 1)]) > p.length {
            p.Pages = append(p.Pages, "")
        }
        
        p.Pages[(len(p.Pages) - 1)] += strings.Replace(strings.Replace(Format, "{{Value}}", fmt.Sprintf("%v", Data[k]), -1), "{{Key}}", k, -1)
    }
    return nil
}