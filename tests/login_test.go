package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "beego/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"net/url"
	"strings"
	"strconv"
	"beego/bootstrap"
	"encoding/json"
	"beego/response"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator), ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	bootstrap.AppConfig()
}

func TestLoginByCredentials(t *testing.T) {

	data := url.Values{}
	data.Add("email", "xx@xx.com")
	data.Add("password", "xxxx")

	r, _ := http.NewRequest("POST", "http://xxx/v1/login", strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	//r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	//r.Header.Add("Accept-Encoding", "gzip, deflate")
	//r.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	//r.Header.Add("Connection", "keep-alive")
	//r.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
	//r.ParseForm()
	//r.PostForm.Add("email", "xx@xx.com")
	//r.PostForm.Add("password", "xx")

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestLoginByCredentials", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		var result response.JsonResult
		err := json.Unmarshal(w.Body.Bytes(), &result)

		Convey("The Result Should be Json Format", func() {
			So(err, ShouldBeNil)
		})

		Convey("The Result Should Contain Token And expiredAt Data", func() {
			So(result.Data, ShouldContainKey, "token")
			So(result.Data, ShouldContainKey, "expiredAt")
		})
	})
}