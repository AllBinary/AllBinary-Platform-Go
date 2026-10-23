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

type WeblisketFinder struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    SubDirectory *SubDirectory
}

var WeblisketFinder_Instance *WeblisketFinder = NewWeblisketFinder()
var WeblisketFinder_KEY string = "installerdata.dat"

func (this *WeblisketFinder) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.SubDirectory = SubDirectory_GetInstance()
}


func WeblisketFinder_GetInstance() *WeblisketFinder {
        return WeblisketFinder_Instance
}


func NewWeblisketFinder() *WeblisketFinder {
    this := &WeblisketFinder{}
    this.initFields()
    return this
}


func (this *WeblisketFinder) FindAll(path string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "findAll", e)
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "findAll", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            this.LogUtil.PutF(this.CommonStrings.START, this, "findAll")
            tryControl1, tryResult1 = 1, this.SubDirectory.Search(WeblisketFinder_KEY, AbFile_CreateAbFile(path)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *WeblisketFinder) GetInstallationPath(file *AbFile) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().PRELOADERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "getInstallationPath", e)
            }
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            var filePath string = file.GetPath(); _ = filePath
            var end int = strings.Index(filePath, string(WeblisketFinder_KEY)); _ = end
            if end < 0 {
                tryControl1, tryResult1 = 1, nil; return
            }
            tryControl1, tryResult1 = 1, filePath[0:end]; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func WeblisketFinder_Main(args []string) {
        var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
        var vector IBasicArrayList = WeblisketFinder_GetInstance().FindAll(stringUtil.EMPTY_STRING); _ = vector
        fmt.Println(vector.Size())
}

