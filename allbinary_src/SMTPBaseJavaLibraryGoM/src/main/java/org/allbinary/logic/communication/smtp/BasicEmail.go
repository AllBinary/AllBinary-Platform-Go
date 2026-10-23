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

type BasicEmail struct {
    Email *AbEmail
}

// Implements EmailInterface


func (this *BasicEmail) initFields() {
    
}


func NewBasicEmail(emailInfo *EmailInfo, to string) *BasicEmail {
    this := &BasicEmail{}
    this.initFields()
        var basicEmailInfo IBasicEmailInfo = emailInfo.GetBasicEmailInfo(); _ = basicEmailInfo
        this.Email = NewAbEmail(basicEmailInfo.GetEmailServerConfigurationInterface().GetSmtpServer(), basicEmailInfo.GetEmailServerConfigurationInterface().GetAccountName(), to, basicEmailInfo.GetSubject(), basicEmailInfo.GetTextBody(), emailInfo.GetHtmlAttachment(), emailInfo.GetContentBase())
    return this
}


func (this *BasicEmail) GetEmail() *Email {
        return any(this.Email.GetEmail()).(*Email)
}


func (this *BasicEmail) GetMimeMessage() *MimeMessage {
        return this.Email.GetMimeMessage()
}


func (this *BasicEmail) GetDebugInfo() string {
        return this.Email.GetDebugInfo()
}


// throws Exception
func (this *BasicEmail) Log() string {
        return this.Email.Log()
}


// throws Exception
func (this *BasicEmail) ToHashMap() *HashMap {
        return this.Email.ToHashMap()
}


// throws Exception
func (this *BasicEmail) ToXmlNode(document *Document) *Node {
        return this.Email.ToXmlNode(document)
}


// throws Exception
func (this *BasicEmail) ToXmlDoc() *Document {
        return this.Email.ToXmlDoc()
}

