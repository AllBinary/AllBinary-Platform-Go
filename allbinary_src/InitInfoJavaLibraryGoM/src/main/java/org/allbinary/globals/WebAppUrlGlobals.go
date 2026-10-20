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

type WebAppUrlGlobals struct {
    UrlGlobalsInterface
    Path string
    IsWebappPathSet bool
}


func (this *WebAppUrlGlobals) initFields() {
    
    this.IsWebappPathSet = false
}

func NewWebAppUrlGlobals() *WebAppUrlGlobals {
    this := &WebAppUrlGlobals{}
    this.UrlGlobalsInterface = *NewUrlGlobalsInterface()
    this.initFields()
    return this
}


func (this *WebAppUrlGlobals) IsTestingMode() bool {
        return InitInfo_GetInstance().IsTesting()
}


func (this *WebAppUrlGlobals) GetWebappPath() string {
        return this.Path
}


// synchronized
func (this *WebAppUrlGlobals) SetWebappPath(path string) {
        path = path
        this.IsWebappPathSet = true
}


func (this *WebAppUrlGlobals) GetTestHtmlPath() string {
        return InitInfo_GetInstance().GetTestHtmlPath()
}


// throws Exception
func (this *WebAppUrlGlobals) GetMainPath() string {
        var mainPath string = InitInfo_GetInstance().GetMainPath(); _ = mainPath
        if !StringValidationUtil_GetInstance().IsEmpty(mainPath) {
            return mainPath
        } else if this.IsWebappPathSet {
            return this.GetWebappPath()
        } else {
            panic(errors.New(fmt.Sprint("Webapp Path is not set")))
        }
}

