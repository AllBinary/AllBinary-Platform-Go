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

type AuthenticationHelperFactory struct {
    TagHelperFactory
    LogUtil *LogUtil
}


func (this *AuthenticationHelperFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewAuthenticationHelperFactory() *AuthenticationHelperFactory {
    this := &AuthenticationHelperFactory{}
    this.TagHelperFactory = *NewTagHelperFactory()
    this.initFields()
    return this
}


// throws LicensingException
func (this *AuthenticationHelperFactory) GetInstance_1(hashMap *HashMap, pageContext *PageContext) any {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.GET_INSTANCE, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, NewAuthenticationHelper(hashMap, pageContext); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of AuthenticationHelperFactory.getInstance
func (this *AuthenticationHelperFactory) GetInstance(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*PageContext](args[1], pass) {
            return this.GetInstance_1(Java_Arg[*HashMap](args[0]), Java_Arg[*PageContext](args[1]))
        }
        if len(args) == 2 && Java_IsArg[*HashMap](args[0], pass) && Java_IsArg[*HttpServletRequest](args[1], pass) {
            return this.GetInstance_2(Java_Arg[*HashMap](args[0]), Java_Arg[*HttpServletRequest](args[1]))
        }
    }
    panic("No Java overload of AuthenticationHelperFactory.getInstance for the arguments")
}


// throws LicensingException
func (this *AuthenticationHelperFactory) GetInstance_2(hashMap *HashMap, httpServletRequest *HttpServletRequest) any {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 any) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERFACTORYERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.GET_INSTANCE, e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, NewAuthenticationHelper(hashMap, httpServletRequest); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

