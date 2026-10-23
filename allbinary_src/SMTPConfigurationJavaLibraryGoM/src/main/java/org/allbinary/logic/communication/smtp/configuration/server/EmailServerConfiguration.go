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

type EmailServerConfiguration struct {
    LogUtil *LogUtil
    AccountName string
    Password string
    SmtpServer string
}

// Implements EmailServerConfigurationInterface


func (this *EmailServerConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
}




// Java overloads of new EmailServerConfiguration
func NewEmailServerConfiguration(args ...any) *EmailServerConfiguration {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewEmailServerConfiguration_1()
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewEmailServerConfiguration_2(Java_Arg[*HashMap](args[0]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return NewEmailServerConfiguration_3(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
    }
    panic("No Java overload of new EmailServerConfiguration for the arguments")
}
func NewEmailServerConfiguration_1() *EmailServerConfiguration {
    this := &EmailServerConfiguration{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.START, this, commonStrings.CONSTRUCTOR)
        }
        this.SetSmtpServer(StringUtil_GetInstance().EMPTY_STRING)
    return this
}


func NewEmailServerConfiguration_2(hashMap *HashMap) *EmailServerConfiguration {
    this := &EmailServerConfiguration{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.START, this, "EmailServerConfiguration(HashMap)")
        }
        this.SetAccountName(any(hashMap.Get(EmailServerConfigurationData_ACCOUNT)).(string))
        this.SetPassword(any(hashMap.Get(EmailServerConfigurationData_PASSWORD)).(string))
        this.SetSmtpServer(any(hashMap.Get(EmailServerConfigurationData_SERVER)).(string))
        this.Log()
    return this
}


func NewEmailServerConfiguration_3(account string, password string, server string) *EmailServerConfiguration {
    this := &EmailServerConfiguration{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.START, this, "EmailServerConfiguration(HashMap)")
        }
        this.SetAccountName(account)
        this.SetPassword(password)
        this.SetSmtpServer(server)
        this.Log()
    return this
}


func (this *EmailServerConfiguration) SetSmtpServer(value string) {
        this.SmtpServer = value
}


func (this *EmailServerConfiguration) GetSmtpServer() string {
        return this.SmtpServer
}


func (this *EmailServerConfiguration) GetAccountName() string {
        return this.AccountName
}


func (this *EmailServerConfiguration) SetAccountName(accountName string) {
        this.AccountName = accountName
}


func (this *EmailServerConfiguration) GetPassword() string {
        return this.Password
}


func (this *EmailServerConfiguration) SetPassword(password string) {
        this.Password = password
}


func (this *EmailServerConfiguration) Log() {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().EMAILLOGGING) {
            this.LogUtil.PutF(this.ToString(), this, "EmailServerConfiguration")
        }
}


func (this *EmailServerConfiguration) ToString() string {
        return "Email Server Configuration:\n" + "Smtp Server: \n" + fmt.Sprint(this.GetSmtpServer()) + "\nAccount Name: " + fmt.Sprint(this.GetAccountName()) + "\nPassword: " + fmt.Sprint(this.GetPassword())
}

func (this *EmailServerConfiguration) String() string {
    return this.ToString()
}

