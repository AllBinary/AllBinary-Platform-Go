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

type SmallDelete struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    FileName string
    String string
}


func (this *SmallDelete) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
}


func NewSmallDelete(fileName string) *SmallDelete {
    this := &SmallDelete{}
    this.initFields()
        this.FileName = fileName
        this.String = AbFileSystem_GetInstance().ReadAsString(fileName)
    return this
}


func (this *SmallDelete) DeleteAtStart(start string, end string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().IDLOGGING) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "deleteAtStart", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var beginIndex int; _ = beginIndex
            var endIndex int; _ = endIndex
            beginIndex = strings.Index(this.String, string(start))
            endIndex = strings.Index(this.String, string(end))
            if endIndex > beginIndex && beginIndex >= 0 && endIndex >= 0 && beginIndex < 10 && endIndex < 525 {
                var text string = this.String[endIndex:len(this.String)]; _ = text
                var idFile *FileOutputStream = NewFileOutputStream(this.FileName); _ = idFile
                var idOutData IDataOutputStream = NewDataOutputStream(idFile); _ = idOutData
                idOutData.WriteBytes(text)
                tryControl1, tryResult1 = 1, true; return
            }
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *SmallDelete) DeleteBetween(start string, end string, newText string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().IDLOGGING) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "deleteAtStart", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var beginIndex int; _ = beginIndex
            var endIndex int; _ = endIndex
            beginIndex = strings.Index(this.String, string(start))
            endIndex = strings.Index(this.String, string(end))
            if endIndex > beginIndex && beginIndex >= 0 && endIndex >= 0 && beginIndex < 50 && endIndex < 425 {
                var newStart string = this.String[0:beginIndex]; _ = newStart
                var text string = this.String[endIndex + len(end):len(this.String)]; _ = text
                var idFile *FileOutputStream = NewFileOutputStream(this.FileName); _ = idFile
                var idOutData IDataOutputStream = NewDataOutputStream(idFile); _ = idOutData
                idOutData.WriteBytes(newStart + newText + text)
                tryControl1, tryResult1 = 1, true; return
            }
            tryControl1, tryResult1 = 1, false; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

