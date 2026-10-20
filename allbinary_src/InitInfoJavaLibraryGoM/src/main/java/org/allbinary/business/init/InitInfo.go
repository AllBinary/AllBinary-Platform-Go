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
    MainPath *AbPath
    TestHtmlPath *AbPath
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
func (this *InitInfo) Set() {
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


// Java overloaded method: Go needs a different name
// synchronized
func (this *InitInfo) Set(hashMap *HashMap) {
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


func (this *InitInfo) SetTestHtmlPath(value *AbPath) {
        this.TestHtmlPath = value
}


func (this *InitInfo) SetMainPath(value *AbPath) {
        this.MainPath = value
}


func (this *InitInfo) IsMainPathValid(abPath *AbPath) bool {
        return true
}


func (this *InitInfo) IsTestHtmlPathValid(abPath *AbPath) bool {
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

