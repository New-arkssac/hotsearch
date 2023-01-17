package target

import "fmt"

type WeiBuData1 struct {
	Data struct {
		Search []struct {
			Keyword string `json:"keyword"`
		} `json:"search"`
		Intelligence []struct {
			Keyword     string `json:"keyword"`
			Name        string `json:"name"`
			HoleGradeZh string `json:"holeGradeZh"`
			HoleClassZh string `json:"holeClassZh"`
		} `json:"intelligence"`
	} `json:"data"`
}

type WeiBuData2 struct {
	Data struct {
		Items []struct {
			Name string `json:"name"`
		} `json:"items"`
	} `json:"data"`
}

type WeiBu struct {
	name   string
	header map[string]string
	urls   map[TargetData]string
}

func (w *WeiBu) New() {
	var urls = make(map[TargetData]string)
	header := make(map[string]string)
	header["referer"] = "https://x.threatbook.com/"
	header["user-agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

	urls[new(WeiBuData1)] = "https://x.threatbook.com/v5/node/vnext/searchAndIntelRecommend"
	urls[new(WeiBuData2)] = "https://x.threatbook.com/v5/node/alltopics?page=1"
	w.name = "weibu"
	w.header = header
	w.urls = urls
}

func (b *WeiBu) Urls() map[TargetData]string {
	return b.urls
}

func (b *WeiBu) Name() string {
	return b.name
}

func (b *WeiBu) Header() map[string]string {
	return b.header
}

func (w *WeiBuData1) Decode() []string {
	var code []string

	for _, list := range w.Data.Search {
		code = append(code, list.Keyword)
	}

	for _, list := range w.Data.Intelligence {
		code = append(code, fmt.Sprintf("%s %s %s %s", list.Name, list.Keyword, list.HoleGradeZh, list.HoleClassZh))
	}

	return code
}

func (w WeiBuData2) Decode() []string {
	var code []string

	for _, list := range w.Data.Items {
		code = append(code, list.Name)
	}

	return code
}
