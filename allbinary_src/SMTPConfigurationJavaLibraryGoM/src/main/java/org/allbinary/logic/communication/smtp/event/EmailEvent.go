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

type EmailEvent struct {
    EventObject
    EmailInfo *EmailInfo
    UserEmailEventNameData *UserEmailEventNameData
    Attempts int
}


func (this *EmailEvent) initFields() {
    
    
    
}


func NewEmailEvent(emailEventHandler any, userEmailEventNameData *UserEmailEventNameData, emailInfo *EmailInfo, attempts int) *EmailEvent {
    this := &EmailEvent{}
        this.EventObject = *NewEventObject(emailEventHandler)
        this.initFields()
        this.UserEmailEventNameData = userEmailEventNameData
        this.EmailInfo = emailInfo
    return this
}


func (this *EmailEvent) GetEmailInfo() *EmailInfo {
        return this.EmailInfo
}


func (this *EmailEvent) GetAttempts() int {
        return this.Attempts
}


func (this *EmailEvent) SetAttempts(attempts int) {
        this.Attempts = attempts
}


func (this *EmailEvent) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("EmailEvent Log: \n")
        stringBuffer.Append("Event Name: ")
        stringBuffer.Append(fmt.Sprint(this.UserEmailEventNameData))
        stringBuffer.Append(fmt.Sprint(this.GetEmailInfo()))
        stringBuffer.Append("\nNumber of Attempts: ")
        stringBuffer.Appendint(this.GetAttempts())
        return fmt.Sprint(stringBuffer)
}

func (this *EmailEvent) String() string {
    return this.ToString()
}

