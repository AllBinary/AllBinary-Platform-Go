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

type MouseStrings struct {
    MOUSE_PRESSED string
    MOUSE_RELEASED string
    MOUSE_CLICKED string
    MOUSE_DRAGGED string
    MOUSE_MOVED string
    MOUSE_PRESSED_LABEL string
    MOUSE_RELEASED_LABEL string
    MOUSE_CLICKED_LABEL string
    MOUSE_DRAGGED_LABEL string
    MOUSE_MOVED_LABEL string
}

var MouseStrings_Instance *MouseStrings = NewMouseStrings()

func (this *MouseStrings) initFields() {
    this.MOUSE_PRESSED = "mousePressed"
    this.MOUSE_RELEASED = "mouseReleased"
    this.MOUSE_CLICKED = "mouseClicked"
    this.MOUSE_DRAGGED = "mouseDragged"
    this.MOUSE_MOVED = "mouseMoved"
    this.MOUSE_PRESSED_LABEL = "MousePressed" + fmt.Sprint(CommonSeps_GetInstance().COLON_SEP)
    this.MOUSE_RELEASED_LABEL = "MouseReleased" + fmt.Sprint(CommonSeps_GetInstance().COLON_SEP)
    this.MOUSE_CLICKED_LABEL = "MouseClicked" + fmt.Sprint(CommonSeps_GetInstance().COLON_SEP)
    this.MOUSE_DRAGGED_LABEL = "MouseDragged" + fmt.Sprint(CommonSeps_GetInstance().COLON_SEP)
    this.MOUSE_MOVED_LABEL = "MouseMoved" + fmt.Sprint(CommonSeps_GetInstance().COLON_SEP)
}

func NewMouseStrings() *MouseStrings {
    this := &MouseStrings{}
    this.initFields()
    return this
}


func MouseStrings_GetInstance() *MouseStrings {
        return MouseStrings_Instance
}

