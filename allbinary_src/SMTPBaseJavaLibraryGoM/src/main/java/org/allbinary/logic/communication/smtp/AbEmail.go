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

type AbEmail struct {
    Email *Email
}

// Implements EmailInterface


func (this *AbEmail) initFields() {
    
}


func NewAbEmail(server string, from string, to string, subject string, textBody string, htmlAttachment string, contentBase string) *AbEmail {
    this := &AbEmail{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
            this.Email = NewEmail(from, to, subject, server, textBody, htmlAttachment, contentBase, true)
        } else {
            this.Email = NewEmail(from, to, subject, server, textBody, htmlAttachment, contentBase)
        }
    return this
}


func (this *AbEmail) GetMimeMessage() *MimeMessage {
        return this.Email.GetMimeMessage()
}


func (this *AbEmail) GetDebugInfo() string {
        return this.Email.GetDebugInfo()
}


// throws Exception
func (this *AbEmail) Log() string {
        return this.Email.Log()
}


// throws Exception
func (this *AbEmail) ToHashMap() *HashMap {
        return this.Email.ToHashMap()
}


// throws Exception
func (this *AbEmail) ToXmlNode(document *Document) *Node {
        return this.Email.ToXmlNode(document)
}


func (this *AbEmail) GetEmail() *Email {
        return any(this.Email).(*Email)
}


// throws Exception
func (this *AbEmail) ToXmlDoc() *Document {
        return this.Email.ToXmlDoc()
}

