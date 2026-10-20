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

type FileLockUtil struct {
    LogUtil *LogUtil
}

var FileLockUtil_Instance *FileLockUtil = NewFileLockUtil()

func (this *FileLockUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func FileLockUtil_GetInstance() *FileLockUtil {
        return FileLockUtil_Instance
}


func NewFileLockUtil() *FileLockUtil {
    this := &FileLockUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *FileLockUtil) GetAll(vector *BasicArrayList, isReturnOnFailure bool) *BasicArrayList {
        var fileLockVector *BasicArrayList = NewBasicArrayListD(); _ = fileLockVector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var file *AbFile = any(vector.Get(index)).(*AbFile); _ = file
            var fileLock *FileLock = this.GetLock(file); _ = fileLock
            if fileLock != nil {
                this.LogUtil.PutF("File Lock Obtained: " + fmt.Sprint(file.GetAbsolutePath()), this, "getAll")
                fileLockVector.Add(fileLock)
            } else if isReturnOnFailure {
                this.LogUtil.PutF("Total Locks Obtained: " + fmt.Sprint(fileLockVector.Size()), this, "getAll")
                return fileLockVector
            }
        }
        this.LogUtil.PutF("Total Locks Obtained: " + fmt.Sprint(fileLockVector.Size()), this, "getAll")
        return fileLockVector
}


// throws Exception
func (this *FileLockUtil) GetAllPossible(vector *BasicArrayList) *BasicArrayList {
        return this.GetAll(vector, false)
}


// throws Exception
func (this *FileLockUtil) GetAllOrNone(vector *BasicArrayList) *BasicArrayList {
        var fileLockVector *BasicArrayList = this.GetAll(vector, true); _ = fileLockVector
        if vector.Size() != fileLockVector.Size() {
            return NewBasicArrayListD()
        } else {
            return fileLockVector
        }
}


// throws Exception
func (this *FileLockUtil) GetLock(file *AbFile) *FileLock {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *FileLock) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Exception returns null", this, "getLock", e)
            tryControl1, tryResult1 = 1, nil; return
                }
            }()
            tryControl1, tryResult1 = 1, this.GetLock(AbFileOutputStream_CreateFromAbFileAppend(file, true)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *FileLockUtil) GetLock(fileOutputStream *AbFileOutputStream) *FileLock {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *FileLock) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Exception returns null", this, "getLock", e)
            tryControl1, tryResult1 = 1, nil; return
                }
            this.LogUtil.PutF("Finally - Closing FileOutputStream", this, "getLock")
            StreamUtil_GetInstance().Close(fileOutputStream)
            }()
            var fileLock *FileLock = this.GetLock(fileOutputStream.GetChannel()); _ = fileLock
            tryControl1, tryResult1 = 1, fileLock; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *FileLockUtil) GetLock(fileChannel *FileChannel) *FileLock {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *FileLock) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put("Exception returns null", this, "getLock", e)
            tryControl1, tryResult1 = 1, nil; return
                }
            this.LogUtil.PutF("Finally - Closing FileChannel", this, "getLock")
            fileChannel.Close()
            }()
            var fileLock *FileLock = fileChannel.TryLock(); _ = fileLock
            tryControl1, tryResult1 = 1, fileLock; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

