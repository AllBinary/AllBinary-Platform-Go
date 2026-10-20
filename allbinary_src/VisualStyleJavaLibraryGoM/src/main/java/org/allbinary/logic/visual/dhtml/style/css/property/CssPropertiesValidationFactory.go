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

type CssPropertiesValidationFactory struct {
}



func NewCssPropertiesValidationFactory() *CssPropertiesValidationFactory {
    this := &CssPropertiesValidationFactory{}
    return this
}


// throws Exception
func CssPropertiesValidationFactory_GetInstance(nodeList *NodeList) *BasicArrayList {
        var propertiesVector *BasicArrayList = NewBasicArrayListD(); _ = propertiesVector
        for index := int(0); index < nodeList.GetLength(); index++ {
            var node *Node = nodeList.Item(index); _ = node
            if node.GetNodeName().CompareTo(CssPropertyData_GetInstance().NAME) == 0 {
                var cssPropertyValidationInterface ValidationInterface = any(CssPropertyValidationFactory_GetInstance(node)).(ValidationInterface); _ = cssPropertyValidationInterface
                propertiesVector.Add(cssPropertyValidationInterface)
            }
        }
        return propertiesVector
}


// Java overloaded method: Go needs a different name
func CssPropertiesValidationFactory_GetInstance(hashMap *HashMap) *BasicArrayList {
        var styles *BasicArrayList = NewBasicArrayListD(); _ = styles
        return styles
}

