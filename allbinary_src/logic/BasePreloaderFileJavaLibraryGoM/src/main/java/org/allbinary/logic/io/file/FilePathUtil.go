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

//Used to correct user input for paths
type FilePathUtil struct {
}



func NewFilePathUtil() *FilePathUtil {
    this := &FilePathUtil{}
    return this
}


func FilePathUtil_Adjust(path string) string {
        var abPathData *AbPathData = AbPathData_GetInstance(); _ = abPathData
        var stringValidationUtil *StringValidationUtil = StringValidationUtil_GetInstance(); _ = stringValidationUtil
        var filePathData *FilePathData = FilePathData_GetInstance(); _ = filePathData
        if stringValidationUtil.IsEmpty(path) {
            return filePathData.PATH_START
        }
        path = strings.ReplaceAll(path, string(abPathData.SEPARATORCHAR), string(filePathData.SEPARATORCHAR))
        return path
}

