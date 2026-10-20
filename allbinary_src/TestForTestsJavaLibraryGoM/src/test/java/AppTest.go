// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type AppTest struct {
}


func NewAppTest() *AppTest {
    this := &AppTest{}
    return this
}


func (this *AppTest) TestAddition() {
        Assertions_AssertEquals(2, 1 + 1)
}


func (this *AppTest) TestStringEquality() {
        var expected string = "Hello World!"; _ = expected
        var actual string = "Hello World!"; _ = actual
        Assertions_AssertEquals(expected, actual)
}


func (this *AppTest) TestBooleanAssertions() {
        Assertions_AssertTrue(true)
        Assertions_AssertFalse(false)
}

