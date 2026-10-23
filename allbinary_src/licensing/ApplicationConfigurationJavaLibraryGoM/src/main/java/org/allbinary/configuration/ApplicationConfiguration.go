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

type ApplicationConfiguration struct {
    LogUtil *LogUtil
    Fullscreen bool
    ShowTitleBar bool
    ProgressBarView bool
    FILE string
}

var ApplicationConfiguration_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *ApplicationConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Fullscreen = true
    this.ShowTitleBar = false
    this.ProgressBarView = false
    this.FILE = "ReloadConfiguration.dat"
}


func ApplicationConfiguration_GetInstance() *ApplicationConfiguration {
        if ApplicationConfiguration_Instance == NullUtil_GetInstance().NULL_OBJECT {
            ApplicationConfiguration_Instance = NewApplicationConfiguration()
        }
        return any(ApplicationConfiguration_Instance).(*ApplicationConfiguration)
}


func NewApplicationConfiguration() *ApplicationConfiguration {
    this := &ApplicationConfiguration{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            if FileFactory_GetInstance().IsFile(this.FILE) {
                this.Read()
            } else {
                this.Write()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


// throws Exception
func (this *ApplicationConfiguration) Read() {
        var fileInputStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileInputStreamFactory
        var fileInputStream IInputStream = fileInputStreamFactory.GetFileInputStreamInstance(StringUtil_GetInstance().EMPTY_STRING, this.FILE); _ = fileInputStream
        var dataInputStream *AbDataInputStream = NewAbDataInputStream(fileInputStream); _ = dataInputStream
        var fullScreen int = dataInputStream.ReadInt(); _ = fullScreen
        if fullScreen == 0 {
            this.SetFullscreen(false)
        } else if fullScreen == 1 {
            this.SetFullscreen(true)
        } else {
            panic(errors.New(fmt.Sprint("Invalid FullScreen ActivityConfiguration")))
        }
        var progressBarView int = dataInputStream.ReadInt(); _ = progressBarView
        if progressBarView == 0 {
            this.SetProgressBarView(false)
        } else if progressBarView == 1 {
            this.SetProgressBarView(true)
        } else {
            panic(errors.New(fmt.Sprint("Invalid ProgressBarView ActivityConfiguration")))
        }
        var showTitleBar int = dataInputStream.ReadInt(); _ = showTitleBar
        if showTitleBar == 0 {
            this.SetShowTitleBar(false)
        } else if showTitleBar == 1 {
            this.SetShowTitleBar(true)
        } else {
            panic(errors.New(fmt.Sprint("Invalid ShowTitleBar ActivityConfiguration")))
        }
        this.LogUtil.PutF("Read Configuration: " + this.ToString(), this, "read")
}


// throws Exception
func (this *ApplicationConfiguration) Write() {
        var closeable *Closeable = NullCloseable_NULL_CLOSEABLE; _ = closeable
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            StreamUtil_GetInstance().Close(closeable)
            }()
            this.LogUtil.PutF("Write Configuration: " + this.ToString(), this, "write")
            var fileInputStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileInputStreamFactory
            var fileOutputStream IOutputStream = fileInputStreamFactory.GetFileOutputStreamInstance(StringUtil_GetInstance().EMPTY_STRING, this.FILE); _ = fileOutputStream
            var dataOutputStream *AbDataOutputStream = NewAbDataOutputStream(fileOutputStream); _ = dataOutputStream
            closeable = dataOutputStream
            if this.IsFullscreen() {
                dataOutputStream.WriteInt(1)
            } else {
                dataOutputStream.WriteInt(0)
            }
            if this.IsProgressBarView() {
                dataOutputStream.WriteInt(1)
            } else {
                dataOutputStream.WriteInt(0)
            }
            if this.IsShowTitleBar() {
                dataOutputStream.WriteInt(1)
            } else {
                dataOutputStream.WriteInt(0)
            }
            dataOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *ApplicationConfiguration) Update(gameFeature IFeature) {
        if gameFeature == MainFeatureFactory_GetInstance().FULL_SCREEN {
            var features *Features = Features_GetInstance(); _ = features
            if features.IsFeature(gameFeature) {
                if !this.IsFullscreen() {
                    this.SetFullscreen(true)
                    this.Write()
                }
            } else {
                if this.IsFullscreen() {
                    this.SetFullscreen(false)
                    this.Write()
                }
            }
        }
}


func (this *ApplicationConfiguration) SetFullscreen(fullscreen bool) {
        this.Fullscreen = fullscreen
}


func (this *ApplicationConfiguration) IsFullscreen() bool {
        return this.Fullscreen
}


func (this *ApplicationConfiguration) SetShowTitleBar(showTitleBar bool) {
        this.ShowTitleBar = showTitleBar
}


func (this *ApplicationConfiguration) IsShowTitleBar() bool {
        return this.ShowTitleBar
}


func (this *ApplicationConfiguration) SetProgressBarView(progressBarView bool) {
        this.ProgressBarView = progressBarView
}


func (this *ApplicationConfiguration) IsProgressBarView() bool {
        return this.ProgressBarView
}


func (this *ApplicationConfiguration) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("isFullscreen: ")
        stringBuffer.Appendboolean(this.IsFullscreen())
        stringBuffer.Append(" isProgressBarView: ")
        stringBuffer.Appendboolean(this.IsProgressBarView())
        stringBuffer.Append(" isShowTitleBar: ")
        stringBuffer.Appendboolean(this.IsShowTitleBar())
        return fmt.Sprint(stringBuffer)
}

func (this *ApplicationConfiguration) String() string {
    return this.ToString()
}

