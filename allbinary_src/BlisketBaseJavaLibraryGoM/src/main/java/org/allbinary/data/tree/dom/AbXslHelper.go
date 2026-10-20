/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2022 AllBinary
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

type AbXslHelper struct {
}

var AbXslHelper_Instance *AbXslHelper = NewAbXslHelper()

func NewAbXslHelper() *AbXslHelper {
    this := &AbXslHelper{}
    return this
}


func AbXslHelper_GetInstance() *AbXslHelper {
        return AbXslHelper_Instance
}


// throws Exception
func (this *AbXslHelper) Translate_1(xslFile *AbFile, xmlFile *AbFile) string {
        return this.Translate(xslFile.GetAbsolutePath(), xmlFile.GetAbsolutePath())
}

// Java overloads of AbXslHelper.translate
func (this *AbXslHelper) Translate(args ...any) string {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            return this.Translate_1(Java_Arg[*AbFile](args[0]), Java_Arg[*AbFile](args[1]))
        }
        if len(args) == 2 && Java_IsArg[IAbPath](args[0], pass) && Java_IsArg[IAbPath](args[1], pass) {
            return this.Translate_2(Java_Arg[IAbPath](args[0]), Java_Arg[IAbPath](args[1]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.Translate_3(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of AbXslHelper.translate for the arguments")
}


// throws Exception
func (this *AbXslHelper) Translate_2(xslPath IAbPath, xmlPath IAbPath) string {
        return this.Translate(xslPath.GetPath(), xmlPath.GetPath())
}


// throws Exception
func (this *AbXslHelper) Translate_3(xsltFilePath string, xmlFilePath string) string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            panic(e)
                }
            }()
            var outputStream *strings.Builder = &strings.Builder{}; _ = outputStream
            tryControl1, tryResult1 = 1, fmt.Sprint(XslHelper_GetInstance().Translate(NewStreamSource(xsltFilePath), NewStreamSource(xmlFilePath), NewStreamResult(outputStream)).(IOutputStream)); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

