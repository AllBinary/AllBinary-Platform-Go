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

type AbFileItem struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Name string
    FieldName string
    ByteArray []int8
}

// Implements *FileItem


func (this *AbFileItem) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
}


func NewAbFileItem(name string, fieldName string, byteArray []int8) *AbFileItem {
    this := &AbFileItem{}
    this.initFields()
        this.Name = name
        this.FieldName = fieldName
        this.ByteArray = byteArray
    return this
}


// throws IOException
func (this *AbFileItem) GetInputStream() IInputStream {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *AbFileItem) GetContentType() string {
        return nil
}


func (this *AbFileItem) GetName() string {
        return this.Name
}


func (this *AbFileItem) IsInMemory() bool {
        return true
}


func (this *AbFileItem) GetSize() int64 {
        return int64(len(this.ByteArray))
}


func (this *AbFileItem) Get() []int8 {
        return this.ByteArray
}


// throws UnsupportedEncodingException
func (this *AbFileItem) GetString_1(encoding string) string {
        return NewString(this.ByteArray, encoding)
}

// Java overloads of AbFileItem.getString
func (this *AbFileItem) GetString(args ...any) string {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetString_1(Java_Arg[string](args[0]))
        }
        if len(args) == 0 {
            return this.GetString_2()
        }
    }
    panic("No Java overload of AbFileItem.getString for the arguments")
}


func (this *AbFileItem) GetString_2() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().HTTPERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getString", e)
            }
            tryControl1, tryResult1 = 1, StringUtil_GetInstance().EMPTY_STRING; return
                }
            }()
            tryControl1, tryResult1 = 1, FileItemUtil_GetInstance().GetString(this.ByteArray); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *AbFileItem) Write_1(file *File) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}

// Java overloads of AbFileItem.write
func (this *AbFileItem) Write(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*File](args[0], pass) {
            this.Write_1(Java_Arg[*File](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[*AbFile](args[0], pass) {
            this.Write_2(Java_Arg[*AbFile](args[0]))
            return
        }
    }
    panic("No Java overload of AbFileItem.write for the arguments")
}


// throws Exception
func (this *AbFileItem) Write_2(file *AbFile) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *AbFileItem) Delete() {
        ForcedLogUtil_Log(this.CommonStrings.NOT_IMPLEMENTED, "delete()")
}


func (this *AbFileItem) GetFieldName() string {
        return this.FieldName
}


func (this *AbFileItem) SetFieldName(name string) {
        this.FieldName = name
}


func (this *AbFileItem) IsFormField() bool {
        if StringValidationUtil_GetInstance().IsEmpty(this.FieldName) {
            return false
        } else {
            return true
        }
}


func (this *AbFileItem) SetFormField(state bool) {
}


// throws IOException
func (this *AbFileItem) GetOutputStream() IOutputStream {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *AbFileItem) GetHeaders() *FileItemHeaders {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *AbFileItem) SetHeaders(fih *FileItemHeaders) {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}

