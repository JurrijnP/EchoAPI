package eapi

import (
    "encoding/json"
    //"fmt"
    "io/ioutil"
    "math/rand"
    "reflect"
    "regexp"
    "strconv"
    "strings"
    "time"
)

var (
    patternDataKey   = regexp.MustCompile(`{{data(\[\.\.\.|\[\?|\[)([0-9]+)\](}}\((.+?)\)|}})`)
    patternCodeblock = regexp.MustCompile(`^([~*_]*)((\x60\x60\x60(.+?))|(\x60\x60\x60))([~*_]*)$`)
)

func (t *Translation) GetQuickResponse(cat, key string) (msg string) {
    msg, err := t.GetResponse(cat, key)
    if err != nil {
        return ""
    }
    return msg
}

func (t *Translation) GetResponse(cat, key string) (msg string, err error) {
    lb, err := ioutil.ReadFile("translation/"+t.Language+".json")
    if err != nil {
        return "", err
    }
    
    //fmt.Printf("\n\ncat: %v\nkey: %v", cat, key)
    
    switch cat {
    case "SubHelp":
        tsh := CategorySubHelp{}
        json.Unmarshal(lb, &tsh)
        if reflect.ValueOf(tsh).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(tsh).Field(0).FieldByName(key).String()
        }
    case "Errors":
        te := CategoryErrors{}
        json.Unmarshal(lb, &te)
        if reflect.ValueOf(te).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(te).Field(0).FieldByName(key).String()
        }
    case "Attachments":
        tsh := CategoryAttachments{}
        json.Unmarshal(lb, &tsh)
    case "General":
        tg := CategoryGeneral{}
        json.Unmarshal(lb, &tg)
        if reflect.ValueOf(tg).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(tg).Field(0).FieldByName(key).String()
        }
    case "Events":
        te := CategoryEvents{}
        json.Unmarshal(lb, &te)
        if reflect.ValueOf(te).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(te).Field(0).FieldByName(key).String()
        }
    case "Filters":
        tf := CategoryFilters{}
        json.Unmarshal(lb, &tf)
        if reflect.ValueOf(tf).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(tf).Field(0).FieldByName(key).String()
        }
    case "Commands_Discord":
        tcd := CategoryCommandsDiscord{}
        json.Unmarshal(lb, &tcd)
        if reflect.ValueOf(tcd).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(tcd).Field(0).FieldByName(key).String()
        }
    case "Commands_Echo":
        tce := CategoryCommandsEcho{}
        json.Unmarshal(lb, &tce)
        if reflect.ValueOf(tce).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(tce).Field(0).FieldByName(key).String()
        }
    case "Commands_Misc":
        tcm := CategoryCommandsMisc{}
        json.Unmarshal(lb, &tcm)
        if reflect.ValueOf(tcm).Field(0).FieldByName(key).IsValid() {
            if key == "EightBall" {
                rand.Seed(time.Now().Unix())
                msg = reflect.ValueOf(tcm).Field(0).FieldByName(key).Index(rand.Intn(reflect.ValueOf(tcm).Field(0).FieldByName(key).Len())).String()
            } else {
                msg = reflect.ValueOf(tcm).Field(0).FieldByName(key).String()
            }
        }
    case "Profiles":
        tp := CategoryProfiles{}
        json.Unmarshal(lb, &tp)
        if reflect.ValueOf(tp).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(tp).Field(0).FieldByName(key).String()
        }
    /*case "Attachments":
        tm.Category = Attachments{}
    case "Attachments":
        tm.Category = Attachments{}*/
    }
    //fmt.Printf("\nmsg: %v", msg)
    
    if strings.Contains(msg, "{{prefix}}") && t.prefix != "" {
        msg = strings.Replace(msg, "{{prefix}}", t.prefix, -1)
        return msg, err
    }
    if msg != "" {
        return msg, err
    }
    
    // This should NEVER!!!!! happen.
    //
    // This means an invalid request in the code was made.
    // Something which can only be fixed with a restart.
    return "Oh no... Something went very wrong...\nPlease DM <@167717172343209984> with this error + the message you send before this.", err
}

func (t *Translation) GetActionResponse(filter, action string) (msg string) {
    lb, err := ioutil.ReadFile("translation/"+t.Language+".json")
    if err != nil {
        return ""
    }
    
    tfa := CategoryActions{}
    json.Unmarshal(lb, &tfa)
    
    switch filter {
    case "AntiLink":
        if reflect.ValueOf(tfa).Field(0).FieldByName("AntiLink_"+action).IsValid() {
            msg = reflect.ValueOf(tfa).Field(0).FieldByName("AntiLink_"+action).String()
        }
    case "WordFilter":
        if reflect.ValueOf(tfa).Field(0).FieldByName("WordFilter_"+action).IsValid() {
            msg = reflect.ValueOf(tfa).Field(0).FieldByName("WordFilter_"+action).String()
        }
    case "NameFilter":
        if reflect.ValueOf(tfa).Field(0).FieldByName("NameFilter_"+action).IsValid() {
            msg = reflect.ValueOf(tfa).Field(0).FieldByName("NameFilter_"+action).String()
        }
    case "Kick":
        if reflect.ValueOf(tfa).Field(0).FieldByName("Kick_"+action).IsValid() {
            msg = reflect.ValueOf(tfa).Field(0).FieldByName("Kick_"+action).String()
        }
    case "Ban":
        if reflect.ValueOf(tfa).Field(0).FieldByName("Ban_"+action).IsValid() {
            msg = reflect.ValueOf(tfa).Field(0).FieldByName("Ban_"+action).String()
        }
    }
    
    if strings.Contains(msg, "{{prefix}}") && t.prefix != "" {
        msg = strings.Replace(msg, "{{prefix}}", t.prefix, -1)
        return msg
    }
    if msg != "" {
        return msg
    }
    
    // Should never happen
    return "!!!ERROR!!! Please join https://discord.gg/YWBKbkD & mention JurrijnP/Proxy with this error."
}

func (t *Translation) SetResponseOptions(msg string, options []int) (fmsg string) {
    var patternOption = regexp.MustCompile(`{{option\[([0-9]+?)\]}}\((((.|[\n])*?)\|((.|[\n])*?))\)`)
    
    if !patternOption.MatchString(msg) {
        return msg
    } else if len(options) == 0 {
        return msg
    }
    
    if len(patternOption.FindAllString(msg, -1)) < len(options) {
        options = options[0:len(patternOption.FindAllString(msg, -1))]
    }
    
    fmsg = msg
    
    for oi := range options {
        // TESTING
        // fmt.Println(strings.Split(patternOption.ReplaceAllString(patternOption.FindAllString(msg, -1)[oi], "$2"), "|"))
        
        option := strings.Split(patternOption.ReplaceAllString(patternOption.FindAllString(msg, -1)[oi], "$2"), "|")
        if len(option) < options[oi] || options[oi] < 0 {
            fmsg = strings.Replace(fmsg, patternOption.FindAllString(msg, -1)[oi], option[0], 1)
        } else {
            fmsg = strings.Replace(fmsg, patternOption.FindAllString(msg, -1)[oi], option[options[oi]], 1)
        }
    }
    return fmsg
}

func (tmd *TranslationMDData) setEmpty() {
    tmd.Start = ""
    tmd.End   = ""
}

func (tmd *TranslationMDData) setMarkdown(dk string) {
    tmd.Start = patternDataKey.ReplaceAllString(dk, "$4")
    if patternCodeblock.MatchString(tmd.Start) && patternCodeblock.ReplaceAllString(tmd.Start, "$3") != "" {
        md2 := patternCodeblock.ReplaceAllString(tmd.Start, "$1```$6")
        for _,v := range md2 {
            tmd.End = string(v)+tmd.End
        }
    } else {
        for _,v := range tmd.Start {
            tmd.End = string(v)+tmd.End
        }
    }
    
    if strings.Contains(tmd.Start, "```") {
        tmd.End = tmd.End[:strings.Index(tmd.End, "```")]+"\n"+tmd.End[strings.Index(tmd.End, "```"):]
        mdc := patternCodeblock.ReplaceAllString(tmd.Start, "$2")
        tmd.Start = tmd.Start[:strings.Index(tmd.Start, mdc)+len(mdc)]+"\n"+tmd.Start[strings.Index(tmd.Start, mdc)+len(mdc):]
    }
}

// FillResponseData is used to fill data into a message.
// Any base type can be used for 'args'.
// Slices can be processed in 2 ways.
func (t *Translation) FillResponseData(msg string, args ...interface{}) (fmsg string) {
    if args == nil {
		return msg
    }
    if !strings.Contains(msg, "{{data[") {
        return msg
    }
    fmsg = msg
    //fmt.Println(fmsg)
    
    var DataKey string = ""

Argloop:
    for _, arg := range args {
        if len(patternDataKey.FindAllString(fmsg, -1)) == 0 {
            return
        } else {
            DataKey = patternDataKey.FindAllString(fmsg, -1)[0]
            if patternDataKey.ReplaceAllString(DataKey, "$1") == "[?" {
                if reflect.TypeOf(arg).Kind() == reflect.Slice {
                    fmsg = strings.Replace(fmsg, DataKey, patternDataKey.ReplaceAllString(DataKey, "{{data[...$2]$3"), -1)
                    DataKey = patternDataKey.ReplaceAllString(DataKey, "{{data[...$2]$3")
                } else {
                    fmsg = strings.Replace(fmsg, DataKey, patternDataKey.ReplaceAllString(DataKey, "{{data[$2]$3"), -1)
                    DataKey = patternDataKey.ReplaceAllString(DataKey, "{{data[$2]$3")
                }
            }
        }
        
        if patternDataKey.ReplaceAllString(DataKey, "$3") != "}}" {
            t.markdown.setMarkdown(DataKey)
        } else {
            t.markdown.setEmpty()
        }
        
        switch v := arg.(type) {
        case bool:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatBool(v)                        +t.markdown.End, -1)
        case uint8:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatUint(uint64(v), 10)            +t.markdown.End, -1)
        case uint16:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatUint(uint64(v), 10)            +t.markdown.End, -1)
        case uint32:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatUint(uint64(v), 10)            +t.markdown.End, -1)
        case uint64:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatUint(v, 10)                    +t.markdown.End, -1)
        case int8:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatInt(int64(v), 10)              +t.markdown.End, -1)
        case int16:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatInt(int64(v), 10)              +t.markdown.End, -1)
        case int32:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatInt(int64(v), 10)              +t.markdown.End, -1)
        case int64:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatInt(v, 10)                     +t.markdown.End, -1)
        case int:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatInt(int64(v), 10)              +t.markdown.End, -1)
        case uint:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatUint(uint64(v), 10)            +t.markdown.End, -1)
        case float32:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatFloat(float64(v), 'f', -1, 32) +t.markdown.End, -1)
        case float64:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatFloat(v, 'f', -1, 64)          +t.markdown.End, -1)
        case complex128:
            var fr float64 = real(arg.(complex128))
            var fi float64 = imag(arg.(complex128))
            var sc string = strconv.FormatFloat(fr, 'f', -1, 64) + "+" + strconv.FormatFloat(fi, 'f', -1, 64) + "i"
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ sc +t.markdown.End, -1)
        case complex64:
            var fr float32 = real(arg.(complex64))
            var fi float32 = imag(arg.(complex64))
            var sc string = strconv.FormatFloat(float64(fr), 'f', -1, 32) + "+" + strconv.FormatFloat(float64(fi), 'f', -1, 32) + "i"
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ sc +t.markdown.End, -1)
        case string:
            fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ v  +t.markdown.End, -1)
        case []string:
            if patternDataKey.ReplaceAllString(DataKey, "$1") == "[..." {
                if t.markdown.Start != "" {
                    for vi := range v {
                        v[vi] = t.markdown.Start+ v[vi] +t.markdown.End
                    }
                }
                fmsg = strings.Replace(fmsg, DataKey, strings.Join(v, ", "), -1)
                break Argloop
            } else {
                for _, av := range arg.([]string) {
                    if len(patternDataKey.FindAllString(fmsg, -1)) == 0 {
                        return
                    } else {
                        DataKey = patternDataKey.FindAllString(fmsg, -1)[0]
                        if patternDataKey.ReplaceAllString(DataKey, "$3") != "}}" {
                            t.markdown.setMarkdown(DataKey)
                        } else {
                            t.markdown.setEmpty()
                        }
                        
                        fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ av +t.markdown.End, -1)
                    }
                }
            }
        case []int:
            if strings.Contains(fmsg, "{{data[...") {
                var sjav []string
                for _, av := range arg.([]int) {
                    sjav = append(sjav, t.markdown.Start + strconv.FormatInt(int64(av), 10) + t.markdown.Start)
                }
                fmsg = strings.Replace(fmsg, DataKey, strings.Join(sjav, ", "), -1)
                break Argloop
            } else {
                for _, av := range arg.([]int) {
                    if len(patternDataKey.FindAllString(fmsg, -1)) == 0 {
                        return
                    } else {
                        DataKey = patternDataKey.FindAllString(fmsg, -1)[0]
                        if patternDataKey.ReplaceAllString(DataKey, "$3") != "}}" {
                            t.markdown.setMarkdown(DataKey)
                        } else {
                            t.markdown.setEmpty()
                        }
                        
                        fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatInt(int64(av), 10) +t.markdown.End, -1)
                    }
                }
            }
        case []float64:
            if strings.Contains(fmsg, "{{data[...") {
                var sjav []string
                for _, av := range arg.([]float64) {
                    sjav = append(sjav, t.markdown.Start + strconv.FormatFloat(av, 'f', -1, 64) + t.markdown.End)
                }
                fmsg = strings.Replace(fmsg, DataKey, strings.Join(sjav, ", "), -1)
                break Argloop
            } else {
                for _, av := range arg.([]float64) {
                    if len(patternDataKey.FindAllString(fmsg, -1)) == 0 {
                        return
                    } else {
                        DataKey = patternDataKey.FindAllString(fmsg, -1)[0]
                        if patternDataKey.ReplaceAllString(DataKey, "$3") != "}}" {
                            t.markdown.setMarkdown(DataKey)
                        } else {
                            t.markdown.setEmpty()
                        }
                        
                        fmsg = strings.Replace(fmsg, DataKey, t.markdown.Start+ strconv.FormatFloat(av, 'f', -1, 64) +t.markdown.End, -1)
                    }
                }
            }
        }
    }

    return fmsg
}