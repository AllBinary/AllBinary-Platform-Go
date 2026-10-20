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

type EmailInfo struct {
    LogUtil *LogUtil
    BasicEmailInfo *BasicEmailInfo
    HtmlAttachment string
    ContentBase string
}


func (this *EmailInfo) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewEmailInfo(basicEmailInfo *BasicEmailInfo) *EmailInfo {
    this := &EmailInfo{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.START, this, commonStrings.CONSTRUCTOR)
        }
        this.Init(basicEmailInfo, nil, nil)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewEmailInfo(basicEmailInfo *BasicEmailInfo, htmlAttachment string, contentBase string) *EmailInfo {
    this := &EmailInfo{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.START, this, commonStrings.CONSTRUCTOR)
        }
        this.Init(basicEmailInfo, nil, nil)
    return this
}


func (this *EmailInfo) Init(basicEmailInfo *BasicEmailInfo, htmlAttachment string, contentBase string) {
        this.BasicEmailInfo = basicEmailInfo
        this.HtmlAttachment = htmlAttachment
        this.ContentBase = contentBase
}


func (this *EmailInfo) GetBasicEmailInfo() *BasicEmailInfo {
        return this.BasicEmailInfo
}


func (this *EmailInfo) GetHtmlAttachment() string {
        return this.HtmlAttachment
}


func (this *EmailInfo) GetContentBase() string {
        return this.ContentBase
}


func (this *EmailInfo) ToString() string {
        return "Email Info: \n " + fmt.Sprint(this.GetBasicEmailInfo())
}

func (this *EmailInfo) String() string {
    return this.ToString()
}

