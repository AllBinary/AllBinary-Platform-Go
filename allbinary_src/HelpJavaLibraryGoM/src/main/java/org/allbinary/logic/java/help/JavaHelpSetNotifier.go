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

type JavaHelpSetNotifier struct {
}

var JavaHelpSetNotifier_Vector *BasicArrayList = NewBasicArrayListD()


func NewJavaHelpSetNotifier() *JavaHelpSetNotifier {
    this := &JavaHelpSetNotifier{}
    return this
}


func JavaHelpSetNotifier_Get() *BasicArrayList {
        return JavaHelpSetNotifier_Vector
}


func JavaHelpSetNotifier_IsNotified(helpSet *HelpSet) bool {
        var isNotified bool = false; _ = isNotified
        var url *URL = helpSet.GetHelpSetURL(); _ = url
        var size int = JavaHelpSetNotifier_Vector.Size(); _ = size
        var urlArray []any = JavaHelpSetNotifier_Vector.ToArray(); _ = urlArray
        for i := int(0); i < size; i++ {
            var nextURL *URL = any(urlArray[i]).(*URL); _ = nextURL
            if strings.Compare(fmt.Sprint(url), fmt.Sprint(nextURL)) == 0 {
                isNotified = true
            }
        }
        JavaHelpSetNotifier_Vector.Add(url)
        return isNotified
}

