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

type PartnerIdentifierFileUtil struct {
    LogUtil *LogUtil
    FILE_NAME string
}

var PartnerIdentifierFileUtil_Instance *PartnerIdentifierFileUtil = NewPartnerIdentifierFileUtil()

func (this *PartnerIdentifierFileUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.FILE_NAME = "partner.txt"
}

func NewPartnerIdentifierFileUtil() *PartnerIdentifierFileUtil {
    this := &PartnerIdentifierFileUtil{}
    this.initFields()
    return this
}


func PartnerIdentifierFileUtil_GetInstance() *PartnerIdentifierFileUtil {
        return PartnerIdentifierFileUtil_Instance
}


func (this *PartnerIdentifierFileUtil) Get() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET, e)
            tryControl1, tryResult1 = 1, StringUtil_GetInstance().NULL_STRING; return
                }
            }()
            var resourceUtil *ResourceUtil = ResourceUtil_GetInstance(); _ = resourceUtil
            var partnerString string = StringUtil_GetInstance().EMPTY_STRING; _ = partnerString
            tryControl1, tryResult1 = 1, partnerString; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

