package manager

import (
	"encoding/json"
	"testing"
	"time"

	tunasync "github.com/tuna/tunasync/internal"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStatus(t *testing.T) {
	Convey("status json ser-de should work", t, func() {
		tz := "Asia/Tokyo"
		loc, err := time.LoadLocation(tz)
		So(err, ShouldBeNil)
		t := time.Date(2016, time.April, 16, 23, 8, 10, 0, loc)
		m := webMirrorStatus{
			Name:         "tunalinux",
			Status:       tunasync.Success,
			LastUpdate:   textTime{t},
			LastUpdateTs: stampTime{t},
			Size:         "5GB",
			Upstream:     "rsync://mirrors.tuna.tsinghua.edu.cn/tunalinux/",
		}

		b, err := json.Marshal(m)
		So(err, ShouldBeNil)
		//fmt.Println(string(b))
		var m2 webMirrorStatus
		err = json.Unmarshal(b, &m2)
		So(err, ShouldBeNil)
		// fmt.Printf("%#v", m2)
		So(m2.Name, ShouldEqual, m.Name)
		So(m2.Status, ShouldEqual, m.Status)
		So(m2.LastUpdate.Unix(), ShouldEqual, m.LastUpdate.Unix())
		So(m2.LastUpdateTs.Unix(), ShouldEqual, m.LastUpdate.Unix())
		So(m2.LastUpdate.UnixNano(), ShouldEqual, m.LastUpdate.UnixNano())
		So(m2.LastUpdateTs.UnixNano(), ShouldEqual, m.LastUpdate.UnixNano())
		So(m2.Size, ShouldEqual, m.Size)
		So(m2.Upstream, ShouldEqual, m.Upstream)
	})
}
