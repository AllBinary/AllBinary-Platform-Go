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
    LogUtil *LogUtil
    RegistrationCode string
    NAME string
    FILE string
}

var RegistrationConfiguration_SINGLETON *RegistrationConfiguration = NewRegistrationConfiguration()

func (this *RegistrationConfiguration) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.RegistrationCode = "No Registration Code"
    this.NAME = "registrationid"
    this.FILE = "Registration.dat"
}


func NewRegistrationConfiguration() *RegistrationConfiguration {
    this := &RegistrationConfiguration{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            if FileFactory_GetInstance().IsFile(this.FILE) {
                this.Read()
            } else {
                this.Write()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func RegistrationConfiguration_GetInstance() *RegistrationConfiguration {
        return RegistrationConfiguration_SINGLETON
}


// throws Exception
func (this *RegistrationConfiguration) Read() {
        var fileInputStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileInputStreamFactory
        var fileInputStream IInputStream = fileInputStreamFactory.GetFileInputStreamInstance(StringUtil_GetInstance().EMPTY_STRING, this.FILE); _ = fileInputStream
        var dataInputStream IDataInputStream = NewDataInputStream(fileInputStream); _ = dataInputStream
        this.SetRegistrationCode(dataInputStream.ReadUTF())
        PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("Read Configuration: ").Append(this.ToString())), this, "read")
}


// throws Exception
func (this *RegistrationConfiguration) Write() {
        var dataOutputStream *AbDataOutputStream = nil; _ = dataOutputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            StreamUtil_GetInstance().Close(dataOutputStream)
            }()
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Write Configuration: ").Append(this.ToString())), this, "write")
            var fileInputStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileInputStreamFactory
            var fileOutputStream IOutputStream = fileInputStreamFactory.GetFileOutputStreamInstance(StringUtil_GetInstance().EMPTY_STRING, this.FILE); _ = fileOutputStream
            dataOutputStream = NewAbDataOutputStream(fileOutputStream)
            dataOutputStream.WriteUTF(this.GetRegistrationCode())
            dataOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
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

