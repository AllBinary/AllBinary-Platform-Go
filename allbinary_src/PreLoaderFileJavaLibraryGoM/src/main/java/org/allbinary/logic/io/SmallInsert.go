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

type SmallInsert struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    FileName string
    String string
}


func (this *SmallInsert) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewSmallInsert(fileName string) *SmallInsert {
    this := &SmallInsert{}
    this.initFields()
        this.FileName = fileName
        this.String = AbFileSystem_GetInstance().ReadAsString(fileName)
    return this
}


func (this *SmallInsert) AtBeginning(text string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().IDLOGGING) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "insertAtBeginning", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var idFile *FileOutputStream = NewFileOutputStream(this.FileName); _ = idFile
            var idOutData IDataOutputStream = NewDataOutputStream(idFile); _ = idOutData
            idOutData.WriteBytes(text + this.String)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *SmallInsert) After(text string, after string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().IDLOGGING) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "after", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var index int = strings.Index(this.String, string(after)); _ = index
            if index < 0 {
                fmt.Println("No such start: " + after + " in: " + this.String)
                tryControl1, tryResult1 = 1, false; return
            }
            index = index + len(after)
            var start string = this.String[0:index]; _ = start
            var end string = this.String[index + 1:len(this.String)]; _ = end
            var idFile *FileOutputStream = NewFileOutputStream(this.FileName); _ = idFile
            var idOutData IDataOutputStream = NewDataOutputStream(idFile); _ = idOutData
            idOutData.WriteBytes(start + text + end)
            idOutData.Close()
            idFile.Close()
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *SmallInsert) AtEnd(text string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().IDLOGGING) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "insertAtEnd", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var idFile *FileOutputStream = NewFileOutputStream(this.FileName); _ = idFile
            var idOutData IDataOutputStream = NewDataOutputStream(idFile); _ = idOutData
            idOutData.WriteBytes(this.String + text)
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

