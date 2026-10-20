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

type DateRangeOrderHistoryTag struct {
    StoreValidationTransformTag
    LogUtil *LogUtil
    Status string
}


func (this *DateRangeOrderHistoryTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewDateRangeOrderHistoryTag() *DateRangeOrderHistoryTag {
    this := &DateRangeOrderHistoryTag{}
        this.StoreValidationTransformTag = *NewStoreValidationTransformTag()
        this.initFields()
    return this
}


func (this *DateRangeOrderHistoryTag) SetStatus(value string) {
        this.Status = value
}


// throws JspTagException
func (this *DateRangeOrderHistoryTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, StoreValidationTransformTag_SKIP_BODY; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                this.LogUtil.PutF("Tag Start", this, this.TagStrings.DO_START_TAG)
            }
            this.SetName("Range Order History View")
            this.SetObjectFile("views.admin.orderhistory.DateRangeOrderHistoryView")
            if this.GetCommand() != nil {
                if strings.Compare(this.GetCommand(), org_Allbinary.Globals.GLOBALS2.VIEW) == 0 {
                } else {
                    panic(errors.New(fmt.Sprint("No Such View Command: " + fmt.Sprint(this.GetCommand()))))
                }
                tryControl1, tryResult1 = 1, this.StoreValidationTransformTag.DoStartTag(); return
            }
            panic(errors.New(fmt.Sprint("Command Null")))
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

