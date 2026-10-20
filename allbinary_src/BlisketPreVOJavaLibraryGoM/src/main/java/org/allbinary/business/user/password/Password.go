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

type Password struct {
    LogUtil *LogUtil
    Password string
}


func (this *Password) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewPassword(password string) *Password {
    this := &Password{}
    this.initFields()
        this.Password = password
    return this
}


func (this *Password) Set(value string) {
        this.Password = value
}


func (this *Password) Get() string {
        return this.Password
}


func (this *Password) IsValid() bool {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VALIDATIONERROR) {
                this.LogUtil.Put("Failed to validate form", "Password", commonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var valid bool = true; _ = valid
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VALIDATION) {
                this.LogUtil.PutF("Password: " + this.Password, this, commonStrings.IS_VALID)
            }
            if !StringValidationUtil_GetInstance().IsValidRequired(this.Password, 6, UserData_MAXLEN) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VALIDATION) {
                    this.LogUtil.PutF("Password is invalid", this, commonStrings.IS_VALID)
                }
                valid = false
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *Password) GetValidationInfo() string {
        if !StringValidationUtil_GetInstance().IsValidRequired(this.Password, 6, UserData_MAXLEN) {
            return "Please enter a Password with more than 6 characters.<br />"
        }
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        return stringUtil.EMPTY_STRING
}


func (this *Password) ToVector(secret string) IBasicArrayList {
        this.Password = StringUtil_GetInstance().GetNonNull(this.Password)
        var random int = NewRandom().NextInt(SuperCrypt_KEYMAX); _ = random
        var vector IBasicArrayList = NewBasicArrayListD(); _ = vector
        vector.Add(fmt.Sprint(int(random)))
        vector.Add(secret)
        vector.Add(NewSuperCrypt(random).Encrypt(this.Password))
        return vector
}


func (this *Password) ToHashMap(secret string) *HashMap {
        this.Password = StringUtil_GetInstance().GetNonNull(this.Password)
        var values *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = values
        var random int = NewRandom().NextInt(SuperCrypt_KEYMAX); _ = random
        values.Put(EntryData_GetInstance().ENCRYPTION, fmt.Sprint(int(random)))
        values.Put(UserData_SECRET, secret)
        values.Put(UserData_PASSWORD, NewSuperCrypt(random).Encrypt(this.Password))
        return values
}

