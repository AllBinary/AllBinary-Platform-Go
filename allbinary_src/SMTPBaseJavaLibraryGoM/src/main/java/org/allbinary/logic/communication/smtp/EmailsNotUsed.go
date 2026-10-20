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

type EmailsNotUsed struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    EmailInfoVector *BasicArrayList
}

// Implements DomNodeInterface


func (this *EmailsNotUsed) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewEmailsNotUsed() *EmailsNotUsed {
    this := &EmailsNotUsed{}
    this.initFields()
        this.EmailInfoVector = NewBasicArrayListD()
    return this
}


// throws Exception
func (this *EmailsNotUsed) ToXmlNode(document *Document) *Node {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *Node) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGINGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "toXmlNode()", e)
            }
            var error_ string = "Error Processing Order.  Please contact us if error persists.<p/>"; _ = error_
            tryControl1, tryResult1 = 1, BasicErrorNodeUtil_Get(document, error_); return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENT) {
                this.LogUtil.PutF(this.CommonStrings.START, this, "toXmlNode")
            }
            var node *Node = document.CreateElement(EmailData_NAME); _ = node
            var size int = this.EmailInfoVector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var email *Email = any(this.EmailInfoVector.Get(index)).(*Email); _ = email
                var emailNode *Node = email.ToXmlNode(document); _ = emailNode
                if emailNode != nil {
                    node.AppendChild(emailNode)
                }
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PAYMENT) {
                this.LogUtil.PutF(this.CommonStrings.END, this, "toXmlNode()")
            }
            tryControl1, tryResult1 = 1, node; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

