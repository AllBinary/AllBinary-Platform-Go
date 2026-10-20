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

type AcceptableResponseGenerator struct {
    LogUtil *LogUtil
}

var AcceptableResponseGenerator_Instance *AcceptableResponseGenerator = NewAcceptableResponseGenerator()

func (this *AcceptableResponseGenerator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func AcceptableResponseGenerator_GetInstance() *AcceptableResponseGenerator {
        return AcceptableResponseGenerator_Instance
}


func NewAcceptableResponseGenerator() *AcceptableResponseGenerator {
    this := &AcceptableResponseGenerator{}
    this.initFields()
    return this
}


// throws Exception
func (this *AcceptableResponseGenerator) Get(httpServletRequest *HttpServletRequest) string {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPERROR) {
                this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.GET, e)
            }
            panic(e)
                }
            }()
            var acceptableResponseUtil *AcceptableResponseUtil = AcceptableResponseUtil_GetInstance(); _ = acceptableResponseUtil
            var acceptable string = httpServletRequest.GetHeader("accept"); _ = acceptable
            var result string = acceptableResponseUtil.GetTagName(0); _ = result
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTP) {
                this.LogUtil.PutF("Request Type: " + acceptable, this, commonStrings.GET)
            }
            if acceptable != "" {
                var size int = acceptableResponseUtil.Size(); _ = size
                for index := int(0); index < size; index++ {
                    if strings.Compare(acceptable, acceptableResponseUtil.Get(index)) == 0 {
                        result = acceptableResponseUtil.GetTagName(index)
                    }
                }
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTP) {
                this.LogUtil.PutF("Response Type: " + result, this, commonStrings.GET)
            }
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

