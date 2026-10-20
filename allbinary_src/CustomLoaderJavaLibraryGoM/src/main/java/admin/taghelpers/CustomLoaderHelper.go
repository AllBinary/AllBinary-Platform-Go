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

type CustomLoaderHelper struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}


func (this *CustomLoaderHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}




// Java overloads of new CustomLoaderHelper
func NewCustomLoaderHelper(args ...any) *CustomLoaderHelper {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*PageContext](args[1], pass) {
            return NewCustomLoaderHelper_1(Java_Arg[*HashMap](args[0]), Java_Arg[*PageContext](args[1]))
        }
        if len(args) == 0 {
            return NewCustomLoaderHelper_2()
        }
    }
    panic("No Java overload of new CustomLoaderHelper for the arguments")
}
func NewCustomLoaderHelper_1(hashMap *HashMap, pageContext *PageContext) *CustomLoaderHelper {
    this := &CustomLoaderHelper{}
    this.initFields()
    return this
}


func NewCustomLoaderHelper_2() *CustomLoaderHelper {
    this := &CustomLoaderHelper{}
    this.initFields()
    return this
}


// throws Exception
func (this *CustomLoaderHelper) GetWebappPath() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, org_Allbinary.Globals.URLGLOBALS.GetWebappPath(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *CustomLoaderHelper) SetWebappPath(path string) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "setWebappPath()", e)
            }
                }
            }()
            var urlGlobals *AppUrlGlobals = NewAppUrlGlobals(); _ = urlGlobals
            urlGlobals.SetWebappPath(path)
            URLGLOBALS.Init(urlGlobals)
            return
        }(); tryControl1 == 1 {
            return
        }
}

