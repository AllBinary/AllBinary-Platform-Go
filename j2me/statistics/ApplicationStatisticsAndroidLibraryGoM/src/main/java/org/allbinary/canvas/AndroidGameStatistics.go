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

type AndroidGameStatistics struct {
    BaseGameStatistics
    View IBaseViewWrapper
    TotalOnDraws int64
    ANDROID_STRING_ARRAY []string
    TOTAL_ONDRAWS string
    ONDRAWS_RATE string
}

// The Java class AndroidGameStatistics has subclasses so its variables are this interface
type IAndroidGameStatistics interface {
    IBaseGameStatistics
    AndroidGameStatistics_() *AndroidGameStatistics
    InitView(view *View)
    Process()
    NextOnDraw()
    ToStringArray() []string
    ToString() string
}

func (this *AndroidGameStatistics) AndroidGameStatistics_() *AndroidGameStatistics {
    return this
}


func (this *AndroidGameStatistics) initFields() {
    this.View = BaseViewWrapper_NULL_VIEW_WRAPPER
    
    this.ANDROID_STRING_ARRAY = make([]string, 14)
    this.TOTAL_ONDRAWS = " Total onDraws: "
    this.ONDRAWS_RATE = " onDraws(/10) Sec: "
}


func NewAndroidGameStatistics() *AndroidGameStatistics {
    this := &AndroidGameStatistics{}
    this.BaseGameStatistics = *NewBaseGameStatistics()
    this.initFields()
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        for index := int(0); index < len(this.ANDROID_STRING_ARRAY); index++ {
            this.ANDROID_STRING_ARRAY[index] = stringUtil.EMPTY_STRING
        }
    return this
}


func (this *AndroidGameStatistics) InitView(view *View) {
        this.BaseGameStatistics.Init()
        this.View = NewViewWrapper(view)
        this.TotalOnDraws = 0
}


func (this *AndroidGameStatistics) Process() {
        this.View.PostInvalidate()
}


func (this *AndroidGameStatistics) NextOnDraw() {
        this.TotalOnDraws++
}


func (this *AndroidGameStatistics) ToStringArray() []string {
        var stringArray []string = this.BaseGameStatistics.ToStringArray(); _ = stringArray
        for index := int(0); index < len(stringArray); index++ {
            this.ANDROID_STRING_ARRAY[index] = stringArray[index]
        }
        var totalTime int64 = this.GetTimeDelayHelper().GetElapsed(this.GameTickTimeDelayHelper.StartTime); _ = totalTime
        totalTime = (totalTime / 10000)
        this.ANDROID_STRING_ARRAY[10] = this.TOTAL_ONDRAWS
        this.ANDROID_STRING_ARRAY[11] = fmt.Sprint(this.TotalOnDraws)
        this.ANDROID_STRING_ARRAY[12] = this.ONDRAWS_RATE
        this.ANDROID_STRING_ARRAY[13] = fmt.Sprint(this.TotalOnDraws / totalTime)
        return this.ANDROID_STRING_ARRAY
}


func (this *AndroidGameStatistics) ToString() string {
        var totalTime int64 = this.GetTimeDelayHelper().GetElapsed(this.GameTickTimeDelayHelper.StartTime); _ = totalTime
        totalTime = (totalTime / 1000)
        if totalTime > 0 {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.BaseGameStatistics.ToStringAt(totalTime))
            if this.TotalOnDraws > 0 {
                stringBuffer.Append(this.TOTAL_ONDRAWS)
                stringBuffer.Append(fmt.Sprint(this.TotalOnDraws))
                stringBuffer.Append(this.ONDRAWS_RATE)
                var drawsOverTime string = fmt.Sprint(this.TotalOnDraws / totalTime); _ = drawsOverTime
                stringBuffer.Append(drawsOverTime)
            }
            return fmt.Sprint(stringBuffer)
        } else {
            return this.NOT
        }
}

func (this *AndroidGameStatistics) String() string {
    return this.ToString()
}

