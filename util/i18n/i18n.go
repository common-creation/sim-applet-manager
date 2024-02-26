package i18n

import (
	"encoding/json"

	"github.com/jeandeaual/go-locale"
	"github.com/k0kubun/pp/v3"
)

type I18n struct {
	t map[string]string
}

func NewI18n() *I18n {
	i := new(I18n)

	lang, _ := locale.GetLanguage()
	if lang == "ja" {
		b, err := langs.ReadFile("langs/ja.json")
		if err == nil {
			m := map[string]string{}
			err = json.Unmarshal(b, &m)
			if err == nil {
				i.t = m
			}
		} else {
			pp.Print(err)
		}
	}
	if i.t == nil {
		b, err := langs.ReadFile("langs/en.json")
		if err == nil {
			m := map[string]string{}
			err = json.Unmarshal(b, &m)
			if err != nil {
				// TOOD: エラーダイアログ
				panic(err)
			}
			i.t = m
		}
	}
	pp.Println("lang:", lang)
	pp.Println("mapping:", i.t)
	return i
}

func (i *I18n) T(key string) string {
	if v, ok := i.t[key]; ok {
		return v
	}
	return key
}
