package gen

import (
	"strconv"
	"strings"
)

func (t *Truck) ToAccessStr() string {
	return "obj.track" + strconv.Itoa(t.Num)
}

func (d *DeffInfo) ToTracksAccessStr() string {
	r := make([]string, 0, len(d.Tracks))

	for _, v := range d.Tracks {
		r = append(r, v.ToAccessStr())
	}

	return strings.Join(r, ",")
}

func (d *DeffInfo) ToTracksHeaderStr() string {
	r := make([]string, 0, len(d.Tracks))

	for _, v := range d.Tracks {
		r = append(r, v.Body)
	}

	return strings.Join(r, "\n")
}
