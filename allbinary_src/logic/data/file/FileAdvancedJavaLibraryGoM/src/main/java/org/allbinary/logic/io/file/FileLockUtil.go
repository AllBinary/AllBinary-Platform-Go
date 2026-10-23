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
func (this *FileLockUtil) GetAll(vector IBasicArrayList, isReturnOnFailure bool) IBasicArrayList {
        var fileLockVector IBasicArrayList = NewBasicArrayListD(); _ = fileLockVector
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
func (this *FileLockUtil) GetAllPossible(vector IBasicArrayList) IBasicArrayList {
        return this.GetAll(vector, false)
}


// throws Exception
func (this *FileLockUtil) GetAllOrNone(vector IBasicArrayList) IBasicArrayList {
        var fileLockVector IBasicArrayList = this.GetAll(vector, true); _ = fileLockVector
        if vector.Size() != fileLockVector.Size() {
            return NewBasicArrayListD()
        } else {
            return fileLockVector
        }
}


// throws Exception
func (this *FileLockUtil) GetLock_1(file *AbFile) *FileLock {
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

// Java overloads of FileLockUtil.getLock
func (this *FileLockUtil) GetLock(args ...any) *FileLock {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*AbFile](args[0], pass) {
            return this.GetLock_1(Java_Arg[*AbFile](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*AbFileOutputStream](args[0], pass) {
            return this.GetLock_2(Java_Arg[*AbFileOutputStream](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*FileChannel](args[0], pass) {
            return this.GetLock_3(Java_Arg[*FileChannel](args[0]))
        }
    }
    panic("No Java overload of FileLockUtil.getLock for the arguments")
}


// throws Exception
func (this *FileLockUtil) GetLock_2(fileOutputStream *AbFileOutputStream) *FileLock {
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
func (this *FileLockUtil) GetLock_3(fileChannel *FileChannel) *FileLock {
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

