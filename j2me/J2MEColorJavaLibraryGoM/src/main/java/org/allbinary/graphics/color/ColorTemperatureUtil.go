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

type ColorTemperatureUtil struct {
    BasicColorFactory *BasicColorFactory
}

var ColorTemperatureUtil_Instance *ColorTemperatureUtil = NewColorTemperatureUtil()

func (this *ColorTemperatureUtil) initFields() {
    this.BasicColorFactory = BasicColorFactory_GetInstance()
}

func NewColorTemperatureUtil() *ColorTemperatureUtil {
    this := &ColorTemperatureUtil{}
    this.initFields()
    return this
}


func ColorTemperatureUtil_GetInstance() *ColorTemperatureUtil {
        return ColorTemperatureUtil_Instance
}


func (this *ColorTemperatureUtil) GetBasicColor(thirdTemp int) *BasicColor {
        var basicColor *BasicColor = this.BasicColorFactory.RED; _ = basicColor
        if thirdTemp > 3700 {
            basicColor = this.BasicColorFactory.PURPLE
        } else if thirdTemp > 3200 {
            basicColor = this.BasicColorFactory.BLUE
        } else if thirdTemp > 2700 {
            basicColor = this.BasicColorFactory.PUCE
        } else if thirdTemp > 2200 {
            basicColor = this.BasicColorFactory.WHITE
        } else if thirdTemp > 1700 {
            basicColor = this.BasicColorFactory.GREEN
        } else if thirdTemp > 1200 {
            basicColor = this.BasicColorFactory.YELLOW
        } else if thirdTemp > 700 {
            basicColor = this.BasicColorFactory.ORANGE
        }
        return basicColor
}

