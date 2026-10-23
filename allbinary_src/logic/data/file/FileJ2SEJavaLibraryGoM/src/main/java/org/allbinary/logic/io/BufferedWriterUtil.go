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

type BufferedWriterUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    CommonLabels *CommonLabels
    REMOVING_OLD string
}

var BufferedWriterUtil_Instance *BufferedWriterUtil = NewBufferedWriterUtil()

func (this *BufferedWriterUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.CommonLabels = CommonLabels_GetInstance()
    this.REMOVING_OLD = "Remove old for overwriting: "
}

func NewBufferedWriterUtil() *BufferedWriterUtil {
    this := &BufferedWriterUtil{}
    this.initFields()
    return this
}


func BufferedWriterUtil_GetInstance() *BufferedWriterUtil {
        return BufferedWriterUtil_Instance
}


// throws Exception
func (this *BufferedWriterUtil) Overwrite_1(path string, data string) {
        var abFile *AbFile = AbFile_CreateAbFile(path); _ = abFile
        if abFile.Exists() {
            this.LogUtil.PutF(this.REMOVING_OLD + path, this, this.CommonStrings.CREATE)
            abFile.Delete()
        } else {
            var name string = AbPathData_GetInstance().RemoveNameFromPath(path); _ = name
            var abFileDirectory *AbFile = AbFile_CreateAbFile(name); _ = abFileDirectory
            if abFileDirectory.Exists() {
            } else {
                var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
                this.LogUtil.PutF(fmt.Sprint(stringBuilder.Append(this.CommonStrings.CREATE).Append(this.CommonLabels.COLON_SEP).Append(name)), this, this.CommonStrings.CREATE)
                var result bool = abFileDirectory.Mkdirs(); _ = result
                stringBuilder.Delete(0, stringBuilder.Length())
                this.LogUtil.PutF(fmt.Sprint(stringBuilder.Append(this.CommonLabels.RESULT_).Appendboolean(result)), this, this.CommonStrings.CREATE)
            }
        }
        this.Write(abFile, data)
}

// Java overloads of BufferedWriterUtil.overwrite
func (this *BufferedWriterUtil) Overwrite(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            this.Overwrite_1(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
            return
        }
        if len(args) == 2 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[string](args[1], pass) {
            this.Overwrite_2(Java_Arg[*AbFile](args[0]), Java_Arg[string](args[1]))
            return
        }
    }
    panic("No Java overload of BufferedWriterUtil.overwrite for the arguments")
}


// throws Exception
func (this *BufferedWriterUtil) Overwrite_2(abFile *AbFile, data string) {
        if abFile.Exists() {
            this.LogUtil.PutF(this.REMOVING_OLD + fmt.Sprint(abFile.GetPath()), this, this.CommonStrings.CREATE)
            abFile.Delete()
        } else {
            var name string = AbPathData_GetInstance().RemoveNameFromPath(abFile.GetAbsolutePath()); _ = name
            var abFileDirectory *AbFile = AbFile_CreateAbFile(name); _ = abFileDirectory
            if abFileDirectory.Exists() {
            } else {
                var stringBuilder *StringMaker = NewStringMaker(); _ = stringBuilder
                this.LogUtil.PutF(fmt.Sprint(stringBuilder.Append(this.CommonStrings.CREATE).Append(this.CommonLabels.COLON_SEP).Append(name)), this, this.CommonStrings.CREATE)
                var result bool = abFileDirectory.Mkdirs(); _ = result
                stringBuilder.Delete(0, stringBuilder.Length())
                this.LogUtil.PutF(fmt.Sprint(stringBuilder.Append(this.CommonLabels.RESULT_).Appendboolean(result)), this, this.CommonStrings.CREATE)
            }
        }
        this.Write(abFile, data)
}


// throws Exception
func (this *BufferedWriterUtil) Write(abFile *AbFile, data string) {
        var fileOut *BufferedWriter = NewBufferedWriter(NewFileWriter(AbFileNativeUtil_Get(abFile))); _ = fileOut
        fileOut.Write(data, 0, len(data))
        fileOut.NewLine()
        fileOut.Flush()
}

