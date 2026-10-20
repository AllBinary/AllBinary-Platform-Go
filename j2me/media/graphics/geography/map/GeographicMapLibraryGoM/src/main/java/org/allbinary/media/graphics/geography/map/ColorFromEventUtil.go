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

type ColorFromEventUtil struct {
    YELLOW *BasicColor
    COLOR_INT int
}

var ColorFromEventUtil_Instance *ColorFromEventUtil = NewColorFromEventUtil()

func (this *ColorFromEventUtil) initFields() {
    this.YELLOW = BasicColorFactory_GetInstance().YELLOW
    this.COLOR_INT = this.YELLOW.IntValue()
}

func NewColorFromEventUtil() *ColorFromEventUtil {
    this := &ColorFromEventUtil{}
    this.initFields()
    return this
}


func ColorFromEventUtil_GetInstance() *ColorFromEventUtil {
        return ColorFromEventUtil_Instance
}


func (this *ColorFromEventUtil) GetForegroundColor(gameCanvasInterface IAllBinaryGameCanvas) *BasicColor {
        var color *BasicColor = this.YELLOW; _ = color
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(gameCanvasInterface.GetLayerManager()).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        if geographicMapInterface != nil {
            var foregroundBasicColor *BasicColor = geographicMapInterface.GetForegroundBasicColor(); _ = foregroundBasicColor
            color = foregroundBasicColor
        }
        return color
}

