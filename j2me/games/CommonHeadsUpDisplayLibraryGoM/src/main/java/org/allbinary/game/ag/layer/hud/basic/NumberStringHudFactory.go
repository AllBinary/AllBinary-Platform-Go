/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

package ag

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
    . "org/allbinary"
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
var _ = NullUtil_GetInstance

type NumberStringHudFactory struct {
    NULL_NUMBER_STRING_HUD INumberStringHud
}

var NumberStringHudFactory_Instance any = NullUtil_GetInstance()

func (this *NumberStringHudFactory) initFields() {
    this.NULL_NUMBER_STRING_HUD = this.CreateHud()
}

func NewNumberStringHudFactory() *NumberStringHudFactory {
    this := &NumberStringHudFactory{}
    this.initFields()
    return this
}


func NumberStringHudFactory_GetInstance() *NumberStringHudFactory {
        if NumberStringHudFactory_Instance == NullUtil_GetInstance() {
            NumberStringHudFactory_Instance = NewNumberStringHudFactory()
        }
        return any(NumberStringHudFactory_Instance).(*NumberStringHudFactory)
}


func (this *NumberStringHudFactory) CreateHud() INumberStringHud {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 INumberStringHud) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            logUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
            panic(errors.New(fmt.Sprint()))
                }
            }()
            tryControl1, tryResult1 = 1, NewNumberStringHud(StringUtil_GetInstance().EMPTY_STRING, 9, BasicHudFactory_GetInstance().ABSOLUTE, 1, 0, BasicColorFactory_GetInstance().NULL_COLOR); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

