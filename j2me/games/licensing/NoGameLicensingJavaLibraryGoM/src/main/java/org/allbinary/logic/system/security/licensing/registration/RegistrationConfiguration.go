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

type RegistrationConfiguration struct {
    RegistrationCode string
    NAME string
}

var RegistrationConfiguration_SINGLETON *RegistrationConfiguration = NewRegistrationConfiguration()

func (this *RegistrationConfiguration) initFields() {
    this.RegistrationCode = "No Registration Code"
    this.NAME = "registrationid"
}


func NewRegistrationConfiguration() *RegistrationConfiguration {
    this := &RegistrationConfiguration{}
    this.initFields()
    return this
}


func RegistrationConfiguration_GetInstance() *RegistrationConfiguration {
        return RegistrationConfiguration_SINGLETON
}


// throws Exception
func (this *RegistrationConfiguration) Read() {
}


// throws Exception
func (this *RegistrationConfiguration) Write() {
}


func (this *RegistrationConfiguration) ToString() string {
        return fmt.Sprint(NewStringMaker().Append("Registration Code: ").Append(this.GetRegistrationCode()))
}

func (this *RegistrationConfiguration) String() string {
    return this.ToString()
}


func (this *RegistrationConfiguration) SetRegistrationCode(registrationCode string) {
        this.RegistrationCode = registrationCode
}


func (this *RegistrationConfiguration) GetRegistrationCode() string {
        return this.RegistrationCode
}

