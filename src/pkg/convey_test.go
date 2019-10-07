package pak

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//GoConvey
//	单元测试包
//	run: go test -v
//	测试覆盖率:
//		go get code.google.com/p/go.tools/cmd/cover
//		goconvey

//待测试代码
func Division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("被除数不能为 0")
	}
	return a / b, nil
}

func Test_Add(t *testing.T) {
	Convey("测试描述", t, func() {
		Convey("被除数为 0", func() {
			_, e = Division(10, 0)
			So(e, ShouldNotBeNil)
		})

		Convey("被除数不为 0", func() {
			n, e = Division(10, 2)
			So(e, ShouldBeNil)
			So(n, ShouldEqual, 5)
		})
	})
}
