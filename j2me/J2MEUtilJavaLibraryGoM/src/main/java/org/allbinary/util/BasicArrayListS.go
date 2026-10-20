/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type BasicArrayListS struct {
    BasicArrayList
}

// The Java class BasicArrayListS has subclasses so its variables are this interface
type IBasicArrayListS interface {
    IBasicArrayList
    BasicArrayListS_() *BasicArrayListS
}

func (this *BasicArrayListS) BasicArrayListS_() *BasicArrayListS {
    return this
}



func NewBasicArrayListS(size int) *BasicArrayListS {
    this := &BasicArrayListS{}
        this.BasicArrayList = *NewBasicArrayList(NullUtil_GetInstance().NULL_OBJECT_ARRAY)
        if size < 0 {
            panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("Init Size Exception: ").Appendint(size)))))
        }
        this.ObjectArray = make([]any, size)
    return this
}
