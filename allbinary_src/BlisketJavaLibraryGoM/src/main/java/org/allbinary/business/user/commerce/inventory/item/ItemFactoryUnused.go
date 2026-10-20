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

type ItemFactoryUnused struct {
}



func NewItemFactoryUnused() *ItemFactoryUnused {
    this := &ItemFactoryUnused{}
    return this
}


// throws MoneyException
func ItemFactoryUnused_GetInstance_1(hashMap *HashMap) TableMappingInterface {
        return any(NewItem(hashMap)).(TableMappingInterface)
}

// Java overloads of ItemFactoryUnused.getInstance
func ItemFactoryUnused_GetInstance(args ...any) TableMappingInterface {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return ItemFactoryUnused_GetInstance_1(Java_Arg[*HashMap](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HttpServletRequest](args[0], pass) {
            return ItemFactoryUnused_GetInstance_2(Java_Arg[*HttpServletRequest](args[0]))
        }
    }
    panic("No Java overload of ItemFactoryUnused.getInstance for the arguments")
}


// throws Exception, MoneyException
func ItemFactoryUnused_GetInstance_2(request *HttpServletRequest) TableMappingInterface {
        return any(NewItem(NewRequestParams(request).ToHashMap())).(TableMappingInterface)
}

