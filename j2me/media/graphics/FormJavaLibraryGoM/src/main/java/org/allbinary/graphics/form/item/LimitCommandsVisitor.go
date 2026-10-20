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

type LimitCommandsVisitor struct {
    Visitor
    IsNotSmallDisplay bool
}


func (this *LimitCommandsVisitor) initFields() {
    
}


func NewLimitCommandsVisitor() *LimitCommandsVisitor {
    this := &LimitCommandsVisitor{}
    this.Visitor = *NewVisitor()
    this.initFields()
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        this.IsNotSmallDisplay = (displayInfo.IsPortrait() && displayInfo.GetLastHeight() > ScreenInfo_GetInstance().SMALL_WIDTH) || (!displayInfo.IsPortrait() && displayInfo.GetLastHeight() > ScreenInfo_GetInstance().SMALL_HEIGHT)
    return this
}


func (this *LimitCommandsVisitor) Visit(anyType any) any {
        var command ICommand = any(anyType).(ICommand); _ = command
        var booleanFactory *BooleanFactory = BooleanFactory_GetInstance(); _ = booleanFactory
        if this.IsNotSmallDisplay || command.GetPriority() == 1 || command == GameCommandsFactory_GetInstance().EXIT_COMMAND || command == GameCommandsFactory_GetInstance().EXIT_WITHOUT_PROGRESS_COMMAND {
            return booleanFactory.TRUE
        } else {
            return booleanFactory.FALSE
        }
}

