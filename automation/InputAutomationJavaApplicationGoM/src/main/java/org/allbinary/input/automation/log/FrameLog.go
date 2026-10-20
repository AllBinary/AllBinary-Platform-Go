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

type FrameLog struct {
    Frame int64
    InfoStringBuffer *StringMaker
    ActionsStringBuffer *StringMaker
}

var FrameLog_FrameBytes []int8 = Java_GetBytes("Frame: ")
var FrameLog_Info []int8 = Java_GetBytes("Info: \n")
var FrameLog_Actions []int8 = Java_GetBytes("Actions: \n")

func (this *FrameLog) initFields() {
    
    this.InfoStringBuffer = NewStringMaker()
    this.ActionsStringBuffer = NewStringMaker()
}


func NewFrameLog(frame int64) *FrameLog {
    this := &FrameLog{}
    this.initFields()
        this.Frame = frame
    return this
}


func (this *FrameLog) AddInfo(infoString string) {
        this.InfoStringBuffer.Append(infoString)
}


func (this *FrameLog) AddActions(vectorOfStrings IBasicArrayList) {
        var size int = vectorOfStrings.Size(); _ = size
        for index := int(0); index < size; index++ {
            this.AddAction(any(vectorOfStrings.Get(index)).(string))
        }
}


func (this *FrameLog) AddAction(action string) {
        this.ActionsStringBuffer.Append(action)
        this.ActionsStringBuffer.Append(CommonSeps_GetInstance().NEW_LINE)
}


// throws Exception
func (this *FrameLog) Write() {
        var filePathStringBuffer *StringMaker = NewStringMaker(); _ = filePathStringBuffer
        filePathStringBuffer.Append(ImageOutputData_SAVE_PATH)
        filePathStringBuffer.Append(LongUtil_FillIn(fmt.Sprint(this.Frame)))
        filePathStringBuffer.Append(".txt")
        var filePath string = fmt.Sprint(filePathStringBuffer); _ = filePath
        var fileOutputStream *FileOutputStream = NewFileOutputStream(NewFile(filePath)); _ = fileOutputStream
        fileOutputStream.Write(FrameLog_FrameBytes)
        fileOutputStream.Write(Java_GetBytes(fmt.Sprint(this.GetFrame())))
        fileOutputStream.Write('\n')
        fileOutputStream.Write(FrameLog_Info)
        fileOutputStream.Write(Java_GetBytes(fmt.Sprint(this.InfoStringBuffer)))
        fileOutputStream.Write('\n')
        fileOutputStream.Write(FrameLog_Actions)
        fileOutputStream.Write(Java_GetBytes(fmt.Sprint(this.ActionsStringBuffer)))
        fileOutputStream.Close()
}


func (this *FrameLog) GetFrame() int64 {
        return this.Frame
}

