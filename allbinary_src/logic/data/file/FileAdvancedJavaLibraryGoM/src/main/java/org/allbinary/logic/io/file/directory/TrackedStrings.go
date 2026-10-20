/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2026 AllBinary 
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

type TrackedStrings struct {
    TARGET_PATH_WINDOWS string
    TARGET_PATH_UNIX string
    HTML_TEMP string
    APPLICATION string
    APP string
    GIT_COMMAND string
    CHANGE_DIRECTORY_OPTION string
    LIST_FILES_COMMAND string
    ERROR_UNMATCH_OPTION string
    PATHSPEC_SEPARATOR string
    SEP_BY_NULL_CHAR_INSTEAD_OF_NEW_LINE string
}

var TrackedStrings_Instance *TrackedStrings = NewTrackedStrings()

func (this *TrackedStrings) initFields() {
    this.TARGET_PATH_WINDOWS = "\\target\\"
    this.TARGET_PATH_UNIX = "/target/"
    this.HTML_TEMP = "HTMLTemp"
    this.APPLICATION = "Application"
    this.APP = "\\app"
    this.GIT_COMMAND = "git"
    this.CHANGE_DIRECTORY_OPTION = "-C"
    this.LIST_FILES_COMMAND = "ls-files"
    this.ERROR_UNMATCH_OPTION = "--error-unmatch"
    this.PATHSPEC_SEPARATOR = "--"
    this.SEP_BY_NULL_CHAR_INSTEAD_OF_NEW_LINE = "-z"
}

func NewTrackedStrings() *TrackedStrings {
    this := &TrackedStrings{}
    this.initFields()
    return this
}


func TrackedStrings_GetInstance() *TrackedStrings {
        return TrackedStrings_Instance
}

