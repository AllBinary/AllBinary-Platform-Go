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

type InitInfo struct {
    TESTING string
    MAINPATH string
    TESTHTMLPATH string
    Testing string
    MainPath IAbPath
    TestHtmlPath IAbPath
    HasRead bool
}

var InitInfo_Instance *InitInfo = NewInitInfo()

func (this *InitInfo) initFields() {
    this.TESTING = "TESTING"
    this.MAINPATH = "MAINPATH"
    this.TESTHTMLPATH = "TESTHTMLPATH"
    this.Testing = nil
    this.MainPath = nil
    this.TestHtmlPath = nil
    this.HasRead = false
}


func InitInfo_GetInstance() *InitInfo {
        return InitInfo_Instance
}


func NewInitInfo() *InitInfo {
    this := &InitInfo{}
    this.initFields()
    return this
}


// throws InitException
// synchronized
func (this *InitInfo) Set_1() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE("error", "InitInfo", "set()", e)
            }
                }
            }()
            var initInfoEntity *InitInfoEntity = NewInitInfoEntity(); _ = initInfoEntity
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
                PreLogUtil_Put("created entity", "InitInfo", "set()")
            }
            if !initInfoEntity.Is() {
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
                    PreLogUtil_Put("adding", "InitInfo", "set()")
                }
                initInfoEntity.Add()
            } else {
                if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADER) {
                    PreLogUtil_Put("update", "InitInfo", "set()")
                }
                initInfoEntity.Update()
            }
            this.HasRead = false
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of InitInfo.set
func (this *InitInfo) Set(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            this.Set_1()
            return
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            this.Set_2(Java_Arg[*HashMap](args[0]))
            return
        }
    }
    panic("No Java overload of InitInfo.set for the arguments")
}


// synchronized
func (this *InitInfo) Set_2(hashMap *HashMap) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE("error", "InitInfo", "set()", e)
            }
                }
            }()
            this.Testing = any(hashMap.Get(this.TESTING)).(string)
            this.MainPath = NewAbPath(any(hashMap.Get(this.MAINPATH)).(string), StringUtil_GetInstance().EMPTY_STRING)
            this.TestHtmlPath = NewAbPath(any(hashMap.Get(this.TESTHTMLPATH)).(string), StringUtil_GetInstance().EMPTY_STRING)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *InitInfo) Get() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE("error", "InitInfo", "set()", e)
            }
            tryControl1 = 1; return
                }
            }()
            this.Testing = nil
            this.MainPath = nil
            this.TestHtmlPath = nil
            var initInfoEntity *InitInfoEntity = NewInitInfoEntity(); _ = initInfoEntity
            initInfoEntity.Get()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// synchronized
func (this *InitInfo) SetHasRead(value bool) {
        this.HasRead = value
}


// synchronized
func (this *InitInfo) UpdateIfNeeded() {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE("error", "InitInfo", "updateIfNeeded()", e)
            }
            tryControl1 = 1; return
                }
            }()
            if !this.HasRead {
                this.HasRead = true
                this.Get()
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *InitInfo) IsTesting() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if LogConfigTypes_LOGGING.Contains(LogConfigTypeFactory_GetInstance().PRELOADERERROR) {
                PreLogUtil_PutOE("error", "InitInfo", "updateIfNeeded()", e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            this.UpdateIfNeeded()
            tryControl1, tryResult1 = 1, BooleanUtil_GetInstance().GetFromString(this.Testing); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InitInfo) GetTesting() string {
        this.UpdateIfNeeded()
        return this.Testing
}


func (this *InitInfo) GetTestHtmlPath() string {
        this.UpdateIfNeeded()
        if this.TestHtmlPath != nil {
            return fmt.Sprint(this.TestHtmlPath)
        } else {
            return nil
        }
}


func (this *InitInfo) GetMainPath() string {
        this.UpdateIfNeeded()
        if this.MainPath != nil {
            return fmt.Sprint(this.MainPath)
        } else {
            return nil
        }
}


func (this *InitInfo) SetTesting(value string) {
        this.Testing = value
}


func (this *InitInfo) SetTestHtmlPath(value IAbPath) {
        this.TestHtmlPath = value
}


func (this *InitInfo) SetMainPath(value IAbPath) {
        this.MainPath = value
}


func (this *InitInfo) IsMainPathValid(abPath IAbPath) bool {
        return true
}


func (this *InitInfo) IsTestHtmlPathValid(abPath IAbPath) bool {
        return true
}


// throws Exception
func (this *InitInfo) IsTestingValid(testing string) bool {
        return BooleanUtil_GetInstance().IsStringBoolean(testing)
}


func (this *InitInfo) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(this.TESTING, this.GetTesting())
        hashMap.Put(this.MAINPATH, this.GetMainPath())
        hashMap.Put(this.TESTHTMLPATH, this.GetTestHtmlPath())
        return hashMap
}

