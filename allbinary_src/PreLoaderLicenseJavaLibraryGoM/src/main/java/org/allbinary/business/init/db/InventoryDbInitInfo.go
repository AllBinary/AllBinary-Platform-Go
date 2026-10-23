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

type InventoryDbInitInfo struct {
    DbInitInfo
}

var InventoryDbInitInfo_INITFILE string = "inventorydbinitdata.dat"




// Java overloads of new InventoryDbInitInfo
func NewInventoryDbInitInfo(args ...any) *InventoryDbInitInfo {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewInventoryDbInitInfo_1()
        }
        if len(args) == 1 && Java_IsArg[bool](args[0], pass) {
            return NewInventoryDbInitInfo_2(Java_Arg[bool](args[0]))
        }
    }
    panic("No Java overload of new InventoryDbInitInfo for the arguments")
}
func NewInventoryDbInitInfo_1() *InventoryDbInitInfo {
    this := &InventoryDbInitInfo{}
        this.DbInitInfo = *NewDbInitInfo(InventoryDbInitInfo_INITFILE, true)
    return this
}


func NewInventoryDbInitInfo_2(read bool) *InventoryDbInitInfo {
    this := &InventoryDbInitInfo{}
        this.DbInitInfo = *NewDbInitInfo(InventoryDbInitInfo_INITFILE, read)
    return this
}

