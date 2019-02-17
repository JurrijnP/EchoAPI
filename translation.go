package eapi

import (
    "encoding/json"
    //"errors"
    "io/ioutil"
    //"fmt"
    "math/rand"
    "reflect"
    "strconv"
    "strings"
    "time"
)

func New(lang string) (t *Translation, err error) {
    
    t = &Translation {
        Language: "",
        UseCustom: false,
        CustomResponses: Custom{},
    }

    var tc TranslationConfig
    tc.load()
    
    if tc.languageIsValid(lang) {
        t.Language = lang
    } else {
        return nil, ErrLanguageNotSupported
    }
    return t, err
}

func (t *Translation) Prefix(prefix string) {
    t.prefix = prefix
}

func (t *Translation) Code(lang string) string {
    var tc TranslationConfig
    tc.load()
    
    if !tc.languageIsValid(lang) {
        return ""
    }
    
    for k, v := range tc.Definitions {
        for _, lv := range v {
            if strings.ToLower(lang) == lv {
                return k
            }
        }
    }
    return ""
}

func (t *Translation) FullFormat(lang string) string {
    var tc TranslationConfig
    tc.load()
    
    if !tc.languageIsValid(lang) {
        return ""
    }
    
    for _, v := range tc.Definitions {
        for _, lv := range v {
            if strings.ToLower(lang) == lv {
                return strings.Title(v[1])+"/"+strings.Title(v[0])
            }
        }
    }
    return ""
}

func (t *Translation) Formatted(lang string, mode int) string {
    var tc TranslationConfig
    tc.load()
    
    if !tc.languageIsValid(lang) {
        return ""
    }
    
    for _, v := range tc.Definitions {
        for _, lv := range v {
            if strings.ToLower(lang) == lv {
                switch mode {
                case 0:
                    return strings.Title(v[1])
                default:
                    return strings.Title(v[0])
                }
            }
        }
    }
    return ""
}

func (t *Translation) LanguageIsSupported(lang string) bool {
    var tc TranslationConfig
    tc.load()
    
    if !tc.languageIsValid(lang) {
        return false
    }
    
    for _, v := range tc.Supported {
        if strings.ToLower(t.Code(lang)) == v {
            return true
        }
    }
    return false
}

func (t *Translation) LanguageIsValid(lang string) bool {
    var tc TranslationConfig
    tc.load()
    return tc.languageIsValid(lang)
}

func (tc *TranslationConfig) load() {
    tcb, err := ioutil.ReadFile("translation/config.json")
    if err != nil {
        return
    }
    json.Unmarshal(tcb, &tc)
}

func (tc *TranslationConfig) languageIsValid(lang string) bool {
    for _, v := range tc.Definitions {
        for _, lv := range v {
            if strings.ToLower(lang) == lv {
                return true
            }
        }
    }
    return false
}

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
        tg := CategoryCommandsDiscord{}
        json.Unmarshal(lb, &tg)
        if reflect.ValueOf(tg).Field(0).FieldByName(key).IsValid() {
            msg = reflect.ValueOf(tg).Field(0).FieldByName(key).String()
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

    var i int = 0
    for _, arg := range args {
        switch v := arg.(type) {
        case bool:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatBool(v), -1)
        case uint8:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatUint(uint64(v), 10), -1)
        case uint16:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatUint(uint64(v), 10), -1)
        case uint32:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatUint(uint64(v), 10), -1)
        case uint64:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatUint(uint64(v), 10), -1)
        case int8:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatInt(int64(v), 10), -1)
        case int16:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatInt(int64(v), 10), -1)
        case int32:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatInt(int64(v), 10), -1)
        case int64:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatInt(int64(v), 10), -1)
        case int:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatInt(int64(v), 10), -1)
        case uint:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatUint(uint64(v), 10), -1)
        case float32:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatFloat(float64(v), 'f', -1, 32), -1)
        case float64:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatFloat(v, 'f', -1, 64), -1)
        case complex128:
            var fr float64 = real(arg.(complex128))
            var fi float64 = imag(arg.(complex128))
            var sc string = strconv.FormatFloat(fr, 'f', -1, 64) + "+" + strconv.FormatFloat(fi, 'f', -1, 64) + "i"
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", sc, -1)
        case complex64:
            var fr float32 = real(arg.(complex64))
            var fi float32 = imag(arg.(complex64))
            var sc string = strconv.FormatFloat(float64(fr), 'f', -1, 32) + "+" + strconv.FormatFloat(float64(fi), 'f', -1, 32) + "i"
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", sc, -1)
        case string:
            fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", v, -1)
        case []string:
            if strings.Contains(msg, "{{data[...") {
                jv := strings.Join(v, ", ")
                fmsg = strings.Replace(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}", jv, -1)
                i++
            }
            for _, av := range arg.([]string) {
                fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", av, -1)
                i++
            }
        case []int:
            if strings.Contains(msg, "{{data[...") {
                var sjav []string
                for _, av := range arg.([]int) {
                    sjav = append(sjav, strconv.Itoa(av))
                }
                jav := strings.Join(sjav, ", ")
                fmsg = strings.Replace(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}", jav, -1)
                i++
            }
            for _, av := range arg.([]int) {
                fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.Itoa(av), -1)
                i++
            }
        case []float64:
            if strings.Contains(msg, "{{data[...") {
                var sjav []string
                for _, av := range arg.([]float64) {
                    sjav = append(sjav, strconv.FormatFloat(av, 'f', -1, 64))
                }
                jav := strings.Join(sjav, ", ")
                fmsg = strings.Replace(fmsg, "{{data[..."+strconv.Itoa(i)+"]}}", jav, -1)
                i++
            }
            for _, av := range arg.([]float64) {
                fmsg = strings.Replace(fmsg, "{{data["+strconv.Itoa(i)+"]}}", strconv.FormatFloat(av, 'f', -1, 64), -1)
                i++
            }
        }
        i++
    }

    return fmsg
}