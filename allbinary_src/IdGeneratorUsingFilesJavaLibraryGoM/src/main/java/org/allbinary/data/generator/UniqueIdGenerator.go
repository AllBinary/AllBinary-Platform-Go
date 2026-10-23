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

type UniqueIdGenerator struct {
    LogUtil *LogUtil
    NewFile *AbFile
    Id int64
}

// Implements IdGeneratorInterface

// The Java class UniqueIdGenerator has subclasses so its variables are this interface
type IUniqueIdGenerator interface {
    UniqueIdGenerator_() *UniqueIdGenerator
    Initialize(value int)
    SetFile(filePathName string, name string)
    GetNext() string
}

func (this *UniqueIdGenerator) UniqueIdGenerator_() *UniqueIdGenerator {
    return this
}

var UniqueIdGenerator_EXT string = ".unq"

func (this *UniqueIdGenerator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.Id = 0
}


func NewUniqueIdGenerator() *UniqueIdGenerator {
    this := &UniqueIdGenerator{}
    this.initFields()
    return this
}


func (this *UniqueIdGenerator) Initialize(value int) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().IDLOGGING) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "initialize", e)
            }
                }
            }()
            this.NewFile.CreateNewFile()
            var idData *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(this.NewFile); _ = idData
            idData.WriteLong(value)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
// synchronized
func (this *UniqueIdGenerator) SetFile(filePathName string, name string) {
        this.NewFile = AbFile_CreateAbFile(filePathName)
}


// throws IOException
// synchronized
func (this *UniqueIdGenerator) GetNext() string {
        var idData *AbDataInputStream = nil; _ = idData
        var idOutData *AbDataOutputStream = nil; _ = idOutData
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().IDLOGGING) {
                var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
                this.LogUtil.Put(commonStrings.EXCEPTION, this, "getNext", e)
            }
            tryControl1, tryResult1 = 1, "Error"; return
                }
            StreamUtil_GetInstance().Close(idData)
            StreamUtil_GetInstance().Close(idOutData)
            }()
            var idFile IAbFileInputStream = NewAbFileInputStream(this.NewFile); _ = idFile
            idData = NewAbDataInputStream(idFile)
            this.Id = idData.ReadLong()
            var idOutFile *AbFileOutputStream = AbFileOutputStream_CreateFromAbFile(this.NewFile); _ = idOutFile
            idOutData = NewAbDataOutputStream(idOutFile)
            idOutData.WriteLong(this.Id + 1)
            var idLong int64 = int64(this.Id); _ = idLong
            tryControl1, tryResult1 = 1, fmt.Sprint(idLong); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

