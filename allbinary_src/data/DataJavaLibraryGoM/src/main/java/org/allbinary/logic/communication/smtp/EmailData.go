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

type EmailData struct {
}

var EmailData_SERVER string = "EMAIL_SERVER"
var EmailData_NAME string = "EMAIL_NAME"
var EmailData_FROM string = "EMAIL_FROM"
var EmailData_TO string = "EMAIL_TO"
var EmailData_CC string = "EMAIL_CC"
var EmailData_BCC string = "EMAIL_BCC"
var EmailData_SUBJECT string = "EMAIL_SUBJECT"
var EmailData_CONTENT string = "EMAIL_CONTENT"


func NewEmailData() *EmailData {
    this := &EmailData{}
    return this
}

