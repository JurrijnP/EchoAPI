package eapi

import (
    "encoding/json"
    //"fmt"
    "io/ioutil"
    "math/rand"
    "reflect"
    "strconv"
    "strings"
    "time"
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
    
    var i   int    = 0
    var md  string = ""
    var mdb string = ""
    var mdk string = ""
    for _, arg := range args {
        md  = ""
        mdb = ""
        mdk = ""
        
        if strings.Contains(fmsg, "{{data[?"+strconv.Itoa(i)+"]}}") {
            if reflect.TypeOf(arg).Kind() == reflect.Slice {
                fmsg = strings.Replace(fmsg, "{{data[?"+strconv.Itoa(i)+"]}}", "{{data[..."+strconv.Itoa(i)+"]}}", -1)
            } else {
                fmsg = strings.Replace(fmsg, "{{data[?"+strconv.Itoa(i)+"]}}", "{{data["+strconv.Itoa(i)+"]}}", -1)
            }
            //fmt.Println(fmsg)
        }
        
        if strings.Contains(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}") {
            if strings.HasPrefix(strings.Split(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}")[1], "(") {
                md = strings.Split(strings.Split(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}(")[1], ")")[0]
                mdk = "("+md+")"
                for _,v := range md {
                    mdb = string(v) + mdb
                }
            } else {
                md = ""
                mdb = ""
                mdk = ""
            }
        } else {
            if strings.HasPrefix(strings.Split(fmsg, "{{data["+strconv.Itoa(i)+"]}}")[1], "(") {
                md = strings.Split(strings.Split(fmsg, "{{data["+strconv.Itoa(i)+"]}}(")[1], ")")[0]
                mdk = "("+md+")"
                for _,v := range md {
                    mdb = string(v) + mdb
                }
            } else {
                md = ""
                mdb = ""
                mdk = ""
            }
        }
        
        if strings.Contains(md, "```") {
            md = md[:strings.Index(md, "```")+3]+"\n"+md[strings.Index(md, "```")+3:]
            mdb = md[:strings.Index(md, "```")]+"\n"+md[strings.Index(md, "```"):]
        }
        
        //fmt.Println(md, mdk)
        
        switch v := arg.(type) {
        case bool:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatBool(v)+mdb, -1)
        case uint8:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatUint(uint64(v), 10)+mdb, -1)
        case uint16:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatUint(uint64(v), 10)+mdb, -1)
        case uint32:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatUint(uint64(v), 10)+mdb, -1)
        case uint64:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatUint(uint64(v), 10)+mdb, -1)
        case int8:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatInt(int64(v), 10)+mdb, -1)
        case int16:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatInt(int64(v), 10)+mdb, -1)
        case int32:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatInt(int64(v), 10)+mdb, -1)
        case int64:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatInt(int64(v), 10)+mdb, -1)
        case int:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatInt(int64(v), 10)+mdb, -1)
        case uint:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatUint(uint64(v), 10)+mdb, -1)
        case float32:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatFloat(float64(v), 'f', -1, 32)+mdb, -1)
        case float64:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatFloat(v, 'f', -1, 64)+mdb, -1)
        case complex128:
            var fr float64 = real(arg.(complex128))
            var fi float64 = imag(arg.(complex128))
            var sc string = strconv.FormatFloat(fr, 'f', -1, 64) + "+" + strconv.FormatFloat(fi, 'f', -1, 64) + "i"
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+sc+mdb, -1)
        case complex64:
            var fr float32 = real(arg.(complex64))
            var fi float32 = imag(arg.(complex64))
            var sc string = strconv.FormatFloat(float64(fr), 'f', -1, 32) + "+" + strconv.FormatFloat(float64(fi), 'f', -1, 32) + "i"
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+sc+mdb, -1)
        case string:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+v+mdb, -1)
        case []string:
            if strings.Contains(fmsg, "{{data[...") {
                //fmt.Println(md)
                if md != "" {
                    for vi := range v {
                        v[vi] = md+v[vi]+mdb
                    }
                }
                jv := strings.Join(v, ", ")
                fmsg = strings.Replace(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}"+mdk, jv, -1)
                goto next
            }
            for _, av := range arg.([]string) {
                if strings.HasPrefix(strings.Split(fmsg, "{{data["+strconv.Itoa(i)+"]}}")[1], "(") {
                    md = strings.Split(strings.Split(fmsg, "{{data["+strconv.Itoa(i)+"]}}(")[1], ")")[0]
                    mdk = "("+md+")"
                    for _,v := range md {
                        mdb = string(v) + mdb
                    }
                    
                    if strings.Contains(md, "```") {
                        md = md[:strings.Index(md, "```")+3]+"\n"+md[strings.Index(md, "```")+3:]
                        mdb = md[:strings.Index(md, "```")]+"\n"+md[strings.Index(md, "```"):]
                    }
                } else {
                    md = ""
                    mdb = ""
                    mdk = ""
                }
                
                fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", md+av+mdb, -1)
                i++
            }
        case []int:
            if strings.Contains(fmsg, "{{data[...") {
                var sjav []string
                for _, av := range arg.([]int) {
                    sjav = append(sjav, md+strconv.Itoa(av)+mdb)
                }
                jav := strings.Join(sjav, ", ")
                fmsg = strings.Replace(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}"+mdk, jav, -1)
                goto next
            }
            for _, av := range arg.([]int) {
                if strings.HasPrefix(strings.Split(fmsg, "{{data["+strconv.Itoa(i)+"]}}")[1], "(") {
                    md = strings.Split(strings.Split(fmsg, "{{data["+strconv.Itoa(i)+"]}}(")[1], ")")[0]
                    mdk = "("+md+")"
                    for _,v := range md {
                        mdb = string(v) + mdb
                    }
                    
                    if strings.Contains(md, "```") {
                        md = md[:strings.Index(md, "```")+3]+"\n"+md[strings.Index(md, "```")+3:]
                        mdb = md[:strings.Index(md, "```")]+"\n"+md[strings.Index(md, "```"):]
                    }
                } else {
                    md = ""
                    mdb = ""
                    mdk = ""
                }
                
                fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.Itoa(av)+mdb, -1)
                i++
            }
        case []float64:
            if strings.Contains(fmsg, "{{data[...") {
                var sjav []string
                for _, av := range arg.([]float64) {
                    sjav = append(sjav, md+strconv.FormatFloat(av, 'f', -1, 64)+mdb)
                }
                jav := strings.Join(sjav, ", ")
                fmsg = strings.Replace(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}"+mdk, jav, -1)
                goto next
            }
            for _, av := range arg.([]float64) {
                if strings.HasPrefix(strings.Split(fmsg, "{{data["+strconv.Itoa(i)+"]}}")[1], "(") {
                    md = strings.Split(strings.Split(fmsg, "{{data["+strconv.Itoa(i)+"]}}(")[1], ")")[0]
                    mdk = "("+md+")"
                    for _,v := range md {
                        mdb = string(v) + mdb
                    }
                    
                    if strings.Contains(md, "```") {
                        md = md[:strings.Index(md, "```")+3]+"\n"+md[strings.Index(md, "```")+3:]
                        mdb = md[:strings.Index(md, "```")]+"\n"+md[strings.Index(md, "```"):]
                    }
                } else {
                    md = ""
                    mdb = ""
                    mdk = ""
                }
                
                fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}"+mdk, md+strconv.FormatFloat(av, 'f', -1, 64)+mdb, -1)
                i++
            }
        }
        goto next
        next: i++
    }

    return fmsg
}