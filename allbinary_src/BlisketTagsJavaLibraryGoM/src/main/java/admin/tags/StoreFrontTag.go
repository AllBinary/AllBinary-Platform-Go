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

type StoreFrontTag struct {
    CustomTagSupport
    LogUtil *LogUtil
    Command string
    StoreName string
    PropertiesHashMap *HashMap
}


func (this *StoreFrontTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}


func NewStoreFrontTag() *StoreFrontTag {
    this := &StoreFrontTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
    return this
}


func (this *StoreFrontTag) SetCommand(command string) {
        this.Command = command
}


func (this *StoreFrontTag) SetStoreName(value string) {
        this.StoreName = value
}


// throws LicensingException
func (this *StoreFrontTag) GetCurrentLocation() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewStoreFrontHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("getCurrentLocation", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws LicensingException
func (this *StoreFrontTag) GetCurrentHomeLocation() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var anyType any = NewStoreFrontHelperFactory().GetInstance(this.PropertiesHashMap, this.PageContext); _ = anyType
            var method *Method = reflect.TypeOf(anyType).GetMethod("getCurrentHomeLocation", nil); _ = method
            var result string = any(method.Invoke(anyType, nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *StoreFrontTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
                }
            }()
            if this.Command != "" {
                this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
                this.PropertiesHashMap.Put(StoreFrontData_GetInstance().NAME, this.StoreName)
                if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.GETCURRENTLOCATION) == 0 {
                    this.PageContext.GetOut().Print(this.GetCurrentLocation())
                } else if strings.Compare(this.Command, org_Allbinary.Globals.GLOBALS2.GETCURRENTHOMELOCATION) == 0 {
                    this.PageContext.GetOut().Print(this.GetCurrentHomeLocation())
                }
            }
            tryControl1, tryResult1 = 1, CustomTagSupport_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

