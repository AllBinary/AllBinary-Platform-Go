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

//ActualPlatform
type ExceptionUtil struct {
    PRETEND_EXCEPTION error
    NONE string
}

var ExceptionUtil_Instance *ExceptionUtil = NewExceptionUtil()

func (this *ExceptionUtil) initFields() {
    this.PRETEND_EXCEPTION = errors.New(fmt.Sprint("Not Really An Exception"))
    this.NONE = "No Stack Trace"
}


func ExceptionUtil_GetInstance() *ExceptionUtil {
        return ExceptionUtil_Instance
}


func NewExceptionUtil() *ExceptionUtil {
    this := &ExceptionUtil{}
    this.initFields()
    return this
}


func (this *ExceptionUtil) GetStackTrace(e error) string {
        var bs *strings.Builder = &strings.Builder{}; _ = bs
        fmt.Println(e)
        var output string = bs.String(); _ = output
        if output != "" {
            return output
        }
        return this.NONE
}

