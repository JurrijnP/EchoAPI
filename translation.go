package eapi

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "strings"
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

func (tc *TranslationConfig) load() {
    tcb, err := ioutil.ReadFile("translation/config.json")
    if err != nil {
        fmt.Println(err)
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

func (tc *TranslationConfig) langCode(lang string) string {
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

func (tc *TranslationConfig) langFormatted(lang string, format int) string {
    if !tc.languageIsValid(lang) {
        return ""
    }
    
    for _, v := range tc.Definitions {
        for _, lv := range v {
            if strings.ToLower(lang) == lv {
                switch format {
                case 0:
                    return strings.Title(v[1])
                case 2:
                    return strings.Title(v[1])+"/"+strings.Title(v[0])
                default:
                    return strings.Title(v[0])
                }
            }
        }
    }
    return ""
}

func (t *Translation) GetLangCode(lang string) string {
    var tc TranslationConfig
    tc.load()
    return tc.langCode(lang)
}

func (t *Translation) LangFormatted(lang string, format int) string {
    var tc TranslationConfig
    tc.load()

    return tc.langFormatted(lang, format)
}

func (t *Translation) LanguageIsSupported(lang string) bool {
    var tc TranslationConfig
    tc.load()
    
    if !tc.languageIsValid(lang) {
        return false
    }
    
    for _, v := range tc.Supported {
        if strings.ToLower(tc.langCode(lang)) == v {
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