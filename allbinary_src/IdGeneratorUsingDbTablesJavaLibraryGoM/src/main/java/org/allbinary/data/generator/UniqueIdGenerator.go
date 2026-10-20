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
    IdGeneratorEntity *IdGeneratorEntity
    Name string
}

// Implements IdGeneratorInterface

var UniqueIdGenerator_EXT string = ".unq"

func (this *UniqueIdGenerator) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewUniqueIdGenerator() *UniqueIdGenerator {
    this := &UniqueIdGenerator{}
    this.initFields()
        this.IdGeneratorEntity = any(IdGeneratorEntityFactory_GetInstance()).(*IdGeneratorEntity)
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
            var vector *BasicArrayList = NewBasicArrayListD(); _ = vector
            vector.Add(this.Name)
            vector.Add(fmt.Sprint(fmt.Sprint(value)))
            this.IdGeneratorEntity.Insert(vector)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
// synchronized
func (this *UniqueIdGenerator) SetFile(filePathName string, name string) {
        this.Name = name
}


// throws IOException
// synchronized
func (this *UniqueIdGenerator) GetNext() string {
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
            }()
            var idLong int64 = this.IdGeneratorEntity.Get(this.Name); _ = idLong
            var newValue int64 = fmt.Sprint(int64(idLong) + 1); _ = newValue
            this.IdGeneratorEntity.Update(this.Name, newValue)
            tryControl1, tryResult1 = 1, fmt.Sprint(idLong); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

