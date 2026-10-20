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

type MediaIOUtil struct {
}

var MediaIOUtil_Instance *MediaIOUtil = NewMediaIOUtil()

func NewMediaIOUtil() *MediaIOUtil {
    this := &MediaIOUtil{}
    return this
}


func MediaIOUtil_GetInstance() *MediaIOUtil {
        return MediaIOUtil_Instance
}


func (this *MediaIOUtil) GetReaderFormatNames() []string {
        return ImageIO_GetReaderFormatNames()
}


func (this *MediaIOUtil) GetWriterFormatNames() []string {
        return ImageIO_GetWriterFormatNames()
}

