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

type LicensingException struct {
    Exception
}

var LicensingException_SerialVersionUID int64 = 1




// Java overloads of new LicensingException
func NewLicensingException(args ...any) *LicensingException {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewLicensingException_1()
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewLicensingException_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new LicensingException for the arguments")
}
func NewLicensingException_1() *LicensingException {
    this := &LicensingException{}
    this.Exception = *NewException()
    return this
}


func NewLicensingException_2(msg string) *LicensingException {
    this := &LicensingException{}
        this.Exception = *NewException(msg)
    return this
}

