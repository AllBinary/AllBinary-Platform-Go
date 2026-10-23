/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
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

type BaseGameStatistics struct {
    TimeDelayHelper ITimeDelayHelper
    UpdateDelayHelper ITimeDelayHelper
    TotalRefreshes int64
    TotalFrames int64
    GameTickTimeDelayHelper *GameTickTimeDelayHelper
    DEFAULT_SCALE_FACTOR int
    NOT string
    EMPTY_STRING string
    STRING_ARRAY []string
    SPACE_CHAR rune
    PLUS_CHAR rune
    X_CHAR rune
    CHAR_ARRAY [][]rune
    PrimitiveLongSingleton *PrimitiveLongSingleton
}

// The Java class BaseGameStatistics has subclasses so its variables are this interface
type IBaseGameStatistics interface {
    BaseGameStatistics_() *BaseGameStatistics
    Init()
    Add(string_ string)
    Process()
    NextFrame()
    NextRefresh()
    GetRefreshRate() int16
    GetTimeDelayHelper() ITimeDelayHelper
    To2DCharArray() [][]rune
    ToStringArray() []string
    ToStringAt(totalTime int64) string
    ToString() string
}

func (this *BaseGameStatistics) BaseGameStatistics_() *BaseGameStatistics {
    return this
}


func (this *BaseGameStatistics) initFields() {
    this.TimeDelayHelper = NewTimeDelayHelper(Integer_MAX_VALUE)
    this.UpdateDelayHelper = NewTimeDelayHelper(2000)
    
    
    this.GameTickTimeDelayHelper = GameTickTimeDelayHelperFactory_GetInstance()
    this.DEFAULT_SCALE_FACTOR = ScaleFactorFactory_GetInstance().DEFAULT_SCALE_FACTOR
    this.NOT = "Not enough Time to Calculate"
    this.EMPTY_STRING = StringUtil_GetInstance().EMPTY_STRING
    this.STRING_ARRAY = []string{" Total Time: ", this.EMPTY_STRING, " Total Frames: ", this.EMPTY_STRING, " Frames/10 Sec: ", this.EMPTY_STRING, " Total Paints: ", this.EMPTY_STRING, " Paints/10 Sec: ", this.EMPTY_STRING}
    this.SPACE_CHAR = ' '
    this.PLUS_CHAR = '+'
    this.X_CHAR = 'X'
    this.CHAR_ARRAY = [][]rune{[]rune{'F', 'r', 'a', 'm', 'e', 's', '/', '1', '0', ' ', 'S', 'e', 'c', ':'}, []rune{'X', 'X'}, []rune{'P', 'a', 'i', 'n', 't', 's', '/', '1', '0', ' ', 'S', 'e', 'c', ':'}, []rune{'X', 'X'}}
    this.PrimitiveLongSingleton = PrimitiveLongSingleton_GetInstance()
}


func NewBaseGameStatistics() *BaseGameStatistics {
    this := &BaseGameStatistics{}
    this.initFields()
    return this
}


func (this *BaseGameStatistics) Init() {
        this.TimeDelayHelper.SetStartTimeTNT()
        this.TotalRefreshes = 0
        this.TotalFrames = 0
}


func (this *BaseGameStatistics) Add(string_ string) {
}


func (this *BaseGameStatistics) Process() {
}


func (this *BaseGameStatistics) NextFrame() {
        this.TotalFrames++
}


func (this *BaseGameStatistics) NextRefresh() {
        this.TotalRefreshes++
}


func (this *BaseGameStatistics) GetRefreshRate() int16 {
        var elapsed int64 = this.TimeDelayHelper.GetElapsed(this.GameTickTimeDelayHelper.StartTime); _ = elapsed
        if elapsed > 1 {
            var time int16 = int16((this.TotalRefreshes / (elapsed >> this.DEFAULT_SCALE_FACTOR))); _ = time
            return time
        } else {
            return 0
        }
}


func (this *BaseGameStatistics) GetTimeDelayHelper() ITimeDelayHelper {
        return this.TimeDelayHelper
}


func (this *BaseGameStatistics) To2DCharArray() [][]rune {
        var totalTime int64 = this.TimeDelayHelper.GetElapsed(this.GameTickTimeDelayHelper.StartTime); _ = totalTime
        totalTime = (totalTime >> this.DEFAULT_SCALE_FACTOR)
        if totalTime > 0 && this.UpdateDelayHelper.IsTime(this.GameTickTimeDelayHelper.StartTime) {
            var framesPerSec int = int((this.TotalFrames / totalTime)); _ = framesPerSec
            if framesPerSec < 10 {
                this.CHAR_ARRAY[1][0] = this.SPACE_CHAR
                this.CHAR_ARRAY[1][1] = this.PrimitiveLongSingleton.NUMBER_CHAR_ARRAY[framesPerSec]
            } else if framesPerSec < 100 {
                var tens int = framesPerSec / 10; _ = tens
                var removeTens int = tens * 10; _ = removeTens
                this.CHAR_ARRAY[1][0] = this.PrimitiveLongSingleton.NUMBER_CHAR_ARRAY[tens]
                this.CHAR_ARRAY[1][1] = this.PrimitiveLongSingleton.NUMBER_CHAR_ARRAY[framesPerSec - removeTens]
            } else {
                this.CHAR_ARRAY[1][0] = this.PLUS_CHAR
                this.CHAR_ARRAY[1][1] = this.PLUS_CHAR
            }
            var refreshesPerSec int = int((this.TotalRefreshes / totalTime)); _ = refreshesPerSec
            if refreshesPerSec < 10 {
                this.CHAR_ARRAY[3][0] = this.X_CHAR
                this.CHAR_ARRAY[3][1] = this.PrimitiveLongSingleton.NUMBER_CHAR_ARRAY[refreshesPerSec]
            } else if refreshesPerSec < 100 {
                var tens int = refreshesPerSec / 10; _ = tens
                var removeTens int = tens * 10; _ = removeTens
                this.CHAR_ARRAY[3][0] = this.PrimitiveLongSingleton.NUMBER_CHAR_ARRAY[refreshesPerSec / 10]
                this.CHAR_ARRAY[3][1] = this.PrimitiveLongSingleton.NUMBER_CHAR_ARRAY[refreshesPerSec - removeTens]
            } else {
                this.CHAR_ARRAY[3][0] = this.PLUS_CHAR
                this.CHAR_ARRAY[3][1] = this.PLUS_CHAR
            }
        }
        return this.CHAR_ARRAY
}


func (this *BaseGameStatistics) ToStringArray() []string {
        var totalTime int64 = this.TimeDelayHelper.GetElapsed(this.GameTickTimeDelayHelper.StartTime); _ = totalTime
        totalTime = (totalTime / 10000)
        if totalTime > 0 {
            this.STRING_ARRAY[1] = fmt.Sprint(totalTime)
            this.STRING_ARRAY[3] = fmt.Sprint(this.TotalFrames)
            this.STRING_ARRAY[5] = fmt.Sprint(this.TotalFrames / totalTime)
            this.STRING_ARRAY[7] = fmt.Sprint(this.TotalRefreshes)
            this.STRING_ARRAY[9] = fmt.Sprint(this.TotalRefreshes / totalTime)
        } else {
            var string_ string = this.EMPTY_STRING; _ = string_
            this.STRING_ARRAY[1] = string_
            this.STRING_ARRAY[3] = string_
            this.STRING_ARRAY[5] = string_
            this.STRING_ARRAY[7] = string_
            this.STRING_ARRAY[9] = string_
        }
        return this.STRING_ARRAY
}


func (this *BaseGameStatistics) ToStringAt(totalTime int64) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.STRING_ARRAY[0])
        stringBuffer.Appendlong(totalTime)
        stringBuffer.Append(this.STRING_ARRAY[2])
        stringBuffer.Appendlong(this.TotalFrames)
        stringBuffer.Append(this.STRING_ARRAY[4])
        stringBuffer.Appendlong(this.TotalFrames / totalTime)
        if this.TotalRefreshes > 0 {
            stringBuffer.Append(this.STRING_ARRAY[6])
            stringBuffer.Appendlong(this.TotalRefreshes)
            stringBuffer.Append(this.STRING_ARRAY[8])
            stringBuffer.Appendlong(this.TotalRefreshes / totalTime)
            stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        }
        return fmt.Sprint(stringBuffer)
}


func (this *BaseGameStatistics) ToString() string {
        var totalTime int64 = this.TimeDelayHelper.GetElapsed(this.GameTickTimeDelayHelper.StartTime); _ = totalTime
        totalTime = (totalTime / 1000)
        if totalTime > 0 {
            return this.ToStringAt(totalTime)
        } else {
            return this.NOT
        }
}

func (this *BaseGameStatistics) String() string {
    return this.ToString()
}

