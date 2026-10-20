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

type LOGGING struct {
    Name string
}

var LOGGING_Logging IBasicArrayList = NewBasicArrayListD()
var LOGGING_GRAPHICSPAINT *LOGGING = NewLOGGING("Graphics Paint")
var LOGGING_GRAPHICSCREATION *LOGGING = NewLOGGING("Graphics Creation")
var LOGGING_GRAPHICSDELETION *LOGGING = NewLOGGING("Graphics Deletion")
var LOGGING_LOADINGERROR *LOGGING = NewLOGGING("Loading Error")
var LOGGING_MENUEVENT *LOGGING = NewLOGGING("Menu Action Event")
var LOGGING_MOUSEEVENT *LOGGING = NewLOGGING("Mouse Action Event")
var LOGGING_KEYEVENT *LOGGING = NewLOGGING("Keyboard Action Event")
var LOGGING_ACTIONEVENT *LOGGING = NewLOGGING("Action Event")
var LOGGING_ADJUSTMENTEVENT *LOGGING = NewLOGGING("Adjustment Action Event")
var LOGGING_COMPONENTEVENT *LOGGING = NewLOGGING("Component Action Event")
var LOGGING_CONTAINEREVENT *LOGGING = NewLOGGING("Container Action Event")
var LOGGING_FOCUSEVENT *LOGGING = NewLOGGING("Focus Action Event")
var LOGGING_HIERARCHYEVENT *LOGGING = NewLOGGING("Hierarchy Action Event")
var LOGGING_INPUTEVENT *LOGGING = NewLOGGING("Input Action Event")
var LOGGING_INVOCATIONEVENT *LOGGING = NewLOGGING("Invocation Action Event")
var LOGGING_ITEMEVENT *LOGGING = NewLOGGING("Item Action Event")
var LOGGING_PAINTEVENT *LOGGING = NewLOGGING("Paint Action Event")
var LOGGING_TEXTEVENT *LOGGING = NewLOGGING("Text Action Event")
var LOGGING_WINDOWEVENT *LOGGING = NewLOGGING("Window Action Event")

func (this *LOGGING) initFields() {
    
}


func NewLOGGING(value string) *LOGGING {
    this := &LOGGING{}
    this.initFields()
        this.Name = value
        LOGGING_Logging.Add(value)
    return this
}


func LOGGING_Contains(logType *LOGGING) bool {
        return LOGGING_Logging.Contains(fmt.Sprint(logType))
}


func (this *LOGGING) ToString() string {
        return this.Name
}

func (this *LOGGING) String() string {
    return this.ToString()
}

