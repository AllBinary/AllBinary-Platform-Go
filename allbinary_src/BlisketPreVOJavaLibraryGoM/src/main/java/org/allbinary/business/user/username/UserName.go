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

type UserName struct {
    LogUtil *LogUtil
    UserName string
}

var UserName_Instance *UserName = NewUserName()

func (this *UserName) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func UserName_GetInstance() *UserName {
        return UserName_Instance
}




// Java overloads of new UserName
func NewUserName(args ...any) *UserName {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewUserName_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewUserName_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new UserName for the arguments")
}
func NewUserName_1() *UserName {
    this := &UserName{}
    this.initFields()
    return this
}


func NewUserName_2(hashMap *HashMap) *UserName {
    this := &UserName{}
    this.initFields()
        this.UserName = any(hashMap.Get(UserData_USERNAME)).(string)
    return this
}


func (this *UserName) IsValid_1() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to validate form", this, "isUserNameValid()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            tryControl1, tryResult1 = 1, this.IsValid(this.UserName); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

// Java overloads of UserName.isValid
func (this *UserName) IsValid(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.IsValid_1()
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.IsValid_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of UserName.isValid for the arguments")
}


func (this *UserName) Get() string {
        return this.UserName
}


func (this *UserName) IsValid_2(aUserName string) bool {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put("Failed to validate form", this, commonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var booleanFactory *BooleanFactory = BooleanFactory_GetInstance(); _ = booleanFactory
            var valid bool = booleanFactory.TRUE; _ = valid
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VALIDATION) {
                this.LogUtil.PutF("UserName: " + aUserName, this, commonStrings.IS_VALID)
            }
            if !StringValidationUtil_GetInstance().IsValidRequired(aUserName, 5, UserData_MAXLEN) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VALIDATION) {
                    this.LogUtil.PutF("UserName is invalid", this, commonStrings.IS_VALID)
                }
                valid = booleanFactory.FALSE
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func UserName_GetValidationInfo(aUserName string) string {
        if !StringValidationUtil_GetInstance().IsValidRequired(aUserName, 5, UserData_MAXLEN) {
            return "Please enter a User Name with more than 5 characters.<br/>"
        }
        return StringUtil_GetInstance().EMPTY_STRING
}

