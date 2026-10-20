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

type ProfilingRefreshHelper struct {
    AndroidGameStatistics
    FirstTime bool
    BestFrameProcessingTime int
    WorstFrameProcessingTime int
    FrameProcessingTimeElapsed int64
}

var ProfilingRefreshHelper_Instance *ProfilingRefreshHelper = NewProfilingRefreshHelper()

func (this *ProfilingRefreshHelper) initFields() {
    this.FirstTime = true
    
    
    
}


func ProfilingRefreshHelper_GetInstance() IAndroidGameStatistics {
        return ProfilingRefreshHelper_Instance
}


func NewProfilingRefreshHelper() *ProfilingRefreshHelper {
    this := &ProfilingRefreshHelper{}
    this.AndroidGameStatistics = *NewAndroidGameStatistics()
    this.initFields()
    return this
}


func (this *ProfilingRefreshHelper) InitView(view *View) {
        this.AndroidGameStatistics.InitView(view)
        this.FrameProcessingTimeElapsed = this.GetTimeDelayHelper().GetStartTime()
        this.BestFrameProcessingTime = Integer_MAX_VALUE
        this.WorstFrameProcessingTime = 0
}


func (this *ProfilingRefreshHelper) NextFrame() {
        this.FrameProcessingTimeElapsed = time.Now().UnixMilli() - this.FrameProcessingTimeElapsed
        if this.FirstTime != true {
            if this.FrameProcessingTimeElapsed > int64(this.WorstFrameProcessingTime) {
                this.WorstFrameProcessingTime = int(this.FrameProcessingTimeElapsed)
            }
            if this.FrameProcessingTimeElapsed < int64(this.BestFrameProcessingTime) {
                this.BestFrameProcessingTime = int(this.FrameProcessingTimeElapsed)
            }
        } else {
            this.FirstTime = false
        }
        this.AndroidGameStatistics.NextFrame()
}


func (this *ProfilingRefreshHelper) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(this.AndroidGameStatistics.ToString()).Append(" Worst: ").Appendint(this.WorstFrameProcessingTime).Append(" Best: ").Appendint(this.BestFrameProcessingTime))
}

func (this *ProfilingRefreshHelper) String() string {
    return this.ToString()
}

