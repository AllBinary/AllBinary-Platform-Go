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

type FileItemUtil struct {
    LogUtil *LogUtil
    DEFAULT_CHARSET string
}

var FileItemUtil_Instance *FileItemUtil = NewFileItemUtil()

func (this *FileItemUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.DEFAULT_CHARSET = "ISO-8859-1"
}

func NewFileItemUtil() *FileItemUtil {
    this := &FileItemUtil{}
    this.initFields()
    return this
}


func FileItemUtil_GetInstance() *FileItemUtil {
        return FileItemUtil_Instance
}


// throws Exception
func (this *FileItemUtil) GetString(byteArray []int8) string {
        return NewString(byteArray, this.DEFAULT_CHARSET)
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *FileItemUtil) GetString(fileItem *FileItemStream) string {
        return NewString(this.GetBytes(fileItem), this.DEFAULT_CHARSET)
}


// throws Exception
func (this *FileItemUtil) GetBytes(fileItem *FileItemStream) []int8 {
        var streamUtil *StreamUtil = StreamUtil_GetInstance(); _ = streamUtil
        var inputStream *InputStream = nil; _ = inputStream
        var outputStream *strings.Builder = nil; _ = outputStream
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 []int8) {
            defer func() {
            streamUtil.Close(outputStream)
            streamUtil.Close(inputStream)
            }()
            inputStream = fileItem.OpenStream()
            outputStream = &strings.Builder{}
            if fileItem.IsFormField() {
                this.LogUtil.PutF("FileItemStream FieldName: " + fmt.Sprint(fileItem.GetFieldName()), this, "write()")
            } else {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Uploaded File FieldName: ")
                stringBuffer.Append(fileItem.GetFieldName())
                stringBuffer.Append(" name = ")
                stringBuffer.Append(fileItem.GetName())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "write()")
            }
            outputStream = any(streamUtil.Get(inputStream, outputStream, make([]int8, 16384))).(*strings.Builder)
            tryControl1, tryResult1 = 1, outputStream.ToByteArray(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *FileItemUtil) Write(fileItem *FileItem, file *AbFile) {
        fileItem.Write(AbFileNativeUtil_Get(file))
}

