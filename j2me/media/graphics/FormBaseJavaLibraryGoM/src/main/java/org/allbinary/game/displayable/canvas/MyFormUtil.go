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

type MyFormUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    PopupMenuRectangle *Rectangle
}

var MyFormUtil_Instance *MyFormUtil = NewMyFormUtil()

func (this *MyFormUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.PopupMenuRectangle = RectangleFactory_SINGLETON
}

func NewMyFormUtil() *MyFormUtil {
    this := &MyFormUtil{}
    this.initFields()
    return this
}


func MyFormUtil_GetInstance() *MyFormUtil {
        return MyFormUtil_Instance
}


func (this *MyFormUtil) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var features *Features = Features_GetInstance(); _ = features
        var isOpenGL bool = features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL); _ = isOpenGL
        var TOTAL_CHAR_WIDTH int = func() int { if SWTJOGLProcessor_GetInstance().IsJOGL() && isOpenGL { return 2 }; return 3 }(); _ = TOTAL_CHAR_WIDTH
        var width int = MyFontProcessor_DefaultStringWidth(font, TOTAL_CHAR_WIDTH); _ = width
        var pointFactory *PointFactory = PointFactory_GetInstance(); _ = pointFactory
        if displayInfo.GetLastHeight() < 320 {
            this.PopupMenuRectangle = NewRectangle(pointFactory.CreateXY(0, 25), width, (font.GetHeight() * 4) + 2)
        } else {
            this.PopupMenuRectangle = NewRectangle(pointFactory.CreateXY(0, displayInfo.GetLastHalfHeight() - 70), width, (font.GetHeight() * 5))
        }
}


func (this *MyFormUtil) GetPopupMenuRectangle() *Rectangle {
        return this.PopupMenuRectangle
}

