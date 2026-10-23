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

type BasicEmailInfo struct {
    LogUtil *LogUtil
    Subject string
    TextBody string
    EmailServerConfigurationInterface EmailServerConfigurationInterface
}

// The Java class BasicEmailInfo has subclasses so its variables are this interface
type IBasicEmailInfo interface {
    BasicEmailInfo_() *BasicEmailInfo
    GetSubject() string
    GetTextBody() string
    GetEmailServerConfigurationInterface() EmailServerConfigurationInterface
    SetEmailServerConfigurationInterface(emailServerConfigurationInterface EmailServerConfigurationInterface)
    ToString() string
    Init()
}

func (this *BasicEmailInfo) BasicEmailInfo_() *BasicEmailInfo {
    return this
}


func (this *BasicEmailInfo) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewBasicEmailInfo(subject string, textBody string) *BasicEmailInfo {
    this := &BasicEmailInfo{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.START, this, commonStrings.CONSTRUCTOR)
        }
        this.Subject = subject
        this.TextBody = textBody
    return this
}


func (this *BasicEmailInfo) GetSubject() string {
        return this.Subject
}


func (this *BasicEmailInfo) GetTextBody() string {
        return this.TextBody
}


func (this *BasicEmailInfo) GetEmailServerConfigurationInterface() EmailServerConfigurationInterface {
        return this.EmailServerConfigurationInterface
}


func (this *BasicEmailInfo) SetEmailServerConfigurationInterface(emailServerConfigurationInterface EmailServerConfigurationInterface) {
        this.EmailServerConfigurationInterface = emailServerConfigurationInterface
}


func (this *BasicEmailInfo) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(fmt.Sprint(this.GetEmailServerConfigurationInterface()))
        stringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
        stringBuffer.Append("Subject: \n")
        stringBuffer.Append(this.GetSubject())
        stringBuffer.Append("\nText Body: \n")
        stringBuffer.Append(this.GetTextBody())
        return fmt.Sprint(stringBuffer)
}

func (this *BasicEmailInfo) String() string {
    return this.ToString()
}


// throws Exception
func (this *BasicEmailInfo) Init() {
        panic(errors.New(fmt.Sprint()))
}

