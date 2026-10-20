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

type CommonButtons struct {
    STANDARD_BUTTON_SIZE int
    NORMAL_BUTTON *Rectangle
    LARGE_BUTTON *Rectangle
}

var CommonButtons_Instance *CommonButtons = NewCommonButtons()

func (this *CommonButtons) initFields() {
    
    
    
}


func CommonButtons_GetInstance() *CommonButtons {
        return CommonButtons_Instance
}


func NewCommonButtons() *CommonButtons {
    this := &CommonButtons{}
    this.initFields()
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var commonButtonSize int = 128; _ = commonButtonSize
        for commonButtonSize > 64 {
            var totalColumns int = displayInfo.GetLastWidth() / commonButtonSize; _ = totalColumns
            var totalRows int = displayInfo.GetLastHeight() / commonButtonSize; _ = totalRows
            var max_ int = totalColumns; _ = max_
            if totalRows > max_ {
                max_ = totalRows
            }
            if max_ > 9 {
                break
            }
            commonButtonSize = commonButtonSize >> 1
        }
        this.STANDARD_BUTTON_SIZE = commonButtonSize
        this.NORMAL_BUTTON = NewRectangle(PointFactory_GetInstance().ZERO_ZERO, this.STANDARD_BUTTON_SIZE, this.STANDARD_BUTTON_SIZE)
        this.LARGE_BUTTON = NewRectangle(PointFactory_GetInstance().ZERO_ZERO, this.STANDARD_BUTTON_SIZE << 1, this.STANDARD_BUTTON_SIZE << 1)
    return this
}

