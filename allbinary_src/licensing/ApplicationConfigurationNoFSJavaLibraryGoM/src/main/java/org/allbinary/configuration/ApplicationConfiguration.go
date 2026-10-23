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
}

var ApplicationConfiguration_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *ApplicationConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Fullscreen = true
    this.ShowTitleBar = false
    this.ProgressBarView = false
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
            this.Read()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


// throws Exception
func (this *ApplicationConfiguration) Read() {
        this.SetFullscreen(true)
        this.SetProgressBarView(false)
        this.SetShowTitleBar(false)
        this.LogUtil.PutF("Read Configuration: " + this.ToString(), this, "read")
}


// throws Exception
func (this *ApplicationConfiguration) Write() {
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

