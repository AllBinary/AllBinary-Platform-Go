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

type AbFileInputStream struct {
    FileInputStream
}

// The Java class AbFileInputStream has subclasses so its variables are this interface
type IAbFileInputStream interface {
    AbFileInputStream_() *AbFileInputStream
}

func (this *AbFileInputStream) AbFileInputStream_() *AbFileInputStream {
    return this
}





// Java overloads of new AbFileInputStream
func NewAbFileInputStream(args ...any) *AbFileInputStream {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewAbFileInputStream_1(Java_Arg[string](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*AbFile](args[0], pass) {
            return NewAbFileInputStream_2(Java_Arg[*AbFile](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*File](args[0], pass) {
            return NewAbFileInputStream_3(Java_Arg[*File](args[0]))
        }
    }
    panic("No Java overload of new AbFileInputStream for the arguments")
}
func NewAbFileInputStream_1(name string) *AbFileInputStream {
    this := &AbFileInputStream{}
        this.FileInputStream = *NewFileInputStream(name)
    return this
}


func NewAbFileInputStream_2(file *AbFile) *AbFileInputStream {
    this := &AbFileInputStream{}
        this.FileInputStream = *NewFileInputStream(AbFileNativeUtil_Get(file))
    return this
}


func NewAbFileInputStream_3(file *File) *AbFileInputStream {
    this := &AbFileInputStream{}
        this.FileInputStream = *NewFileInputStream(file)
    return this
}

