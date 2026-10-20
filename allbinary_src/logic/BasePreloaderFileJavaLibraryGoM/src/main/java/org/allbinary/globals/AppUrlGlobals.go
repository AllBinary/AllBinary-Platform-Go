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

//implements
type AppUrlGlobals struct {
    UrlGlobalsInterface
    Path string
}

var AppUrlGlobals_NULL_APP_URL_GLOBALS *AppUrlGlobals = NewAppUrlGlobals()

func (this *AppUrlGlobals) initFields() {
    this.Path = StringUtil_GetInstance().EMPTY_STRING
}

func NewAppUrlGlobals() *AppUrlGlobals {
    this := &AppUrlGlobals{}
    this.UrlGlobalsInterface = *NewUrlGlobalsInterface()
    this.initFields()
    return this
}


func (this *AppUrlGlobals) IsTestingMode() bool {
        return false
}


func (this *AppUrlGlobals) GetWebappPath() string {
        return this.Path
}


// synchronized
func (this *AppUrlGlobals) SetWebappPath(path string) {
        this.Path = path
        if !(strings.HasSuffix(this.Path, "\\") || strings.HasSuffix(this.Path, "/")) {
            this.Path = this.Path + "\\"
        }
}


func (this *AppUrlGlobals) GetTestHtmlPath() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


// throws Exception
func (this *AppUrlGlobals) GetMainPath() string {
        return this.Path
}

