package parser

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

//works only for mirea

func GetTable(urlid string, c chan *Page) {
	url := fmt.Sprintf("https://priem.mirea.ru/competitions_api/entrants?competitions[]=%v", urlid)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	applicants := gjson.Get(string(body), "data.0.app_count").Int()
	list := make(List, 0, applicants)
	for _, v := range gjson.Get(string(body), "data.0.entrants").Array() {
		list = append(list, Student{
			Id:         int(v.Get("spn").Int()),
			Priority:   int(v.Get("p").Int()),
			Acceptance: bool(v.Get("acc").Int() > 0),
			Sum:        int(v.Get("fm").Int()),
			IHP:        bool(v.Get("iHP").Int() > 0),
			IHPO:       bool(v.Get("iHPO").Int() > 0),
			// ihp osnovnoy vp
			//ihpo prohodnoy vp
			//
			//add hp and mp
		})
	}

	raw := gjson.Get(string(body), "data.0.updated_at").String()
	layout := "2006-01-02T15:04:05.000000Z"
	t, err := time.Parse(layout, raw)
	if err != nil {
		log.Fatal(err)
	}
	plan := gjson.Get(string(body), "data.0.plan").Int()
	ti1 := gjson.Get(string(body), "data.0.title").String() + "\n"
	ti1 += gjson.Get(string(body), "data.0.programSet_title").String()
	c <- &Page{
		Applicants: int(applicants),
		Time:       t,
		Planned:    int(plan),
		Title:      ti1,
		List:       list,
	}
}
