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

//ActualPlatform
type GameStatisticsFactory struct {
    AndroidGameStatistics
    StringBuffer *StringMaker
}

var GameStatisticsFactory_SINGLETON IAndroidGameStatistics = NewGameStatisticsFactory()

func (this *GameStatisticsFactory) initFields() {
    this.StringBuffer = NewStringMaker()
}

func NewGameStatisticsFactory() *GameStatisticsFactory {
    this := &GameStatisticsFactory{}
    this.AndroidGameStatistics = *NewAndroidGameStatistics()
    this.initFields()
    return this
}


func GameStatisticsFactory_GetInstance() IAndroidGameStatistics {
        return GameStatisticsFactory_SINGLETON
}


func (this *GameStatisticsFactory) Add(string_ string) {
        if this.StringBuffer.Length() > 12000 {
            this.StringBuffer.Delete(0, this.StringBuffer.Length())
            this.StringBuffer.Append("Old Stats Cleared")
        }
        this.StringBuffer.Append(string_)
}


func (this *GameStatisticsFactory) ToString() string {
        return fmt.Sprint(NewStringMaker().Append(this.AndroidGameStatistics.ToString()).Append(fmt.Sprint(this.StringBuffer)))
}

func (this *GameStatisticsFactory) String() string {
    return this.ToString()
}

