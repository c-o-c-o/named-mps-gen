package gen

import (
	"regexp"
	"strconv"
)

func GetDeffInfo(body string) *DeffInfo {
	r, _ := regexp.Compile("(--track([0-9]):.*?)(\n|$)")
	ls := r.FindAllStringSubmatch(body, -1)

	if ls == nil {
		return nil
	}

	tracks := make([]Truck, len(ls), 4)

	for i, v := range ls {
		tracks[i].Body = v[1]
		tracks[i].Num, _ = strconv.Atoi(v[2])
	}

	return &DeffInfo{
		DeffBody: body,
		Tracks:   tracks,
	}
}
