/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type CanvasStrings struct {
    CONSTRUCTOR string
    PAUSE string
    UN_PAUSE string
    SET_FULL_SCREEN_MODE string
    SIZE_CHANGED string
    SCALED_IMAGES string
    PAINT string
    ON_DISPLAY_CHANGE_EVENT string
    UPDATE_MEASUREMENT string
    FD_WIDTH string
    FD_HEIGHT string
    EMPTY_CHILD_NAME_LIST IBasicArrayList
    ABOUT string
    OPTIONS string
    FONT_ATLAS string
    OBJ_MODEL string
    MD2_MODEL string
    _PNG string
}

var CanvasStrings_Instance *CanvasStrings = NewCanvasStrings()

func (this *CanvasStrings) initFields() {
    this.CONSTRUCTOR = "MyCanvas::MyCanvas"
    this.PAUSE = "MyCanvas::pause"
    this.UN_PAUSE = "MyCanvas::unPause"
    this.SET_FULL_SCREEN_MODE = "setFullScreenMode"
    this.SIZE_CHANGED = "sizeChanged"
    this.SCALED_IMAGES = "scaledImages"
    this.PAINT = "paint"
    this.ON_DISPLAY_CHANGE_EVENT = "onDisplayChangeEvent"
    this.UPDATE_MEASUREMENT = "updateMeasurement"
    this.FD_WIDTH = " font/display width: "
    this.FD_HEIGHT = " font/display height: "
    this.EMPTY_CHILD_NAME_LIST = BasicArrayListUtil_GetInstance().GetImmutableInstance()
    this.ABOUT = "About"
    this.OPTIONS = "Options"
    this.FONT_ATLAS = "/font.png"
    this.OBJ_MODEL = "_obj"
    this.MD2_MODEL = "_md2"
    this._PNG = ".png"
}

func NewCanvasStrings() *CanvasStrings {
    this := &CanvasStrings{}
    this.initFields()
    return this
}


func CanvasStrings_GetInstance() *CanvasStrings {
        return CanvasStrings_Instance
}

