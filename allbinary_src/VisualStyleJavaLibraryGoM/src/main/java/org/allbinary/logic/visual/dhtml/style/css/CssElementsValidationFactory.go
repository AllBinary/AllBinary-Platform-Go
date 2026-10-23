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

type CssElementsValidationFactory struct {
}



func NewCssElementsValidationFactory() *CssElementsValidationFactory {
    this := &CssElementsValidationFactory{}
    return this
}


// throws Exception
func CssElementsValidationFactory_GetInstance_1(cssElementStyleNodeList IBasicArrayList) IBasicArrayList {
        var styles IBasicArrayList = NewBasicArrayListD(); _ = styles
        var size int = cssElementStyleNodeList.Size(); _ = size
        var cssElementStyleNode *Node; _ = cssElementStyleNode
        var cssValidationInterface ValidationInterface; _ = cssValidationInterface
        for index := int(0); index < size; index++ {
            cssElementStyleNode = any(cssElementStyleNodeList.Get(index)).(*Node)
            cssValidationInterface = any(CssElementValidationFactory_GetInstance(cssElementStyleNode)).(ValidationInterface)
            styles.Add(cssValidationInterface)
        }
        return styles
}

// Java overloads of CssElementsValidationFactory.getInstance
func CssElementsValidationFactory_GetInstance(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[IBasicArrayList](args[0], pass) {
            return CssElementsValidationFactory_GetInstance_1(Java_Arg[IBasicArrayList](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return CssElementsValidationFactory_GetInstance_2(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of CssElementsValidationFactory.getInstance for the arguments")
}


func CssElementsValidationFactory_GetInstance_2(hashMap *HashMap) IBasicArrayList {
        var styles IBasicArrayList = NewBasicArrayListD(); _ = styles
        return styles
}

