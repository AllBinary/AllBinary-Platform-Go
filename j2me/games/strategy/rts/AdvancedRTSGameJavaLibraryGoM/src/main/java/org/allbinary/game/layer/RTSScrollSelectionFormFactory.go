/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type RTSScrollSelectionFormFactory struct {
}


func NewRTSScrollSelectionFormFactory() *RTSScrollSelectionFormFactory {
    this := &RTSScrollSelectionFormFactory{}
    return this
}


// throws Exception
func RTSScrollSelectionFormFactory_GetInstance(formLabel string, items []IABCustomItem) IScrollSelectionForm {
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        var touchButtonLocationHelper *TouchButtonLocationHelper = NewTouchButtonLocationHelper(); _ = touchButtonLocationHelper
        var displayInfo *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfo
        var menuX int = CommonButtons_GetInstance().STANDARD_BUTTON_SIZE + touchButtonLocationHelper.GetColumnsRemainderHalf(); _ = menuX
        var width int = displayInfo.GetLastWidth() - (menuX * 2); _ = width
        var menuRectangle *Rectangle = NewRectangle(PointFactory_GetInstance().CreateXY(menuX, 16), width, CommonButtons_GetInstance().STANDARD_BUTTON_SIZE); _ = menuRectangle
        return MultipleScrollSelectionHorizontalForm_CreateForm(formLabel, items, menuRectangle, FormTypeFactory_GetInstance().HORIZONTAL_FORM, 0, basicColorFactory.BLACK, basicColorFactory.WHITE)
}

