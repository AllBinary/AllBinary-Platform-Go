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

//Warning you must have sql root access
type DynamicInitDb struct {
    InitDb
}





// Java overloads of new DynamicInitDb
func NewDynamicInitDb(args ...any) *DynamicInitDb {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[AbeClientInformationInterface](args[0], pass) {
            return NewDynamicInitDb_1(Java_Arg[AbeClientInformationInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[AbeClientInformationInterface](args[0], pass) && Java_IsArg[DatabaseConnectionInfoInterface](args[1], pass) {
            return NewDynamicInitDb_2(Java_Arg[AbeClientInformationInterface](args[0]), Java_Arg[DatabaseConnectionInfoInterface](args[1]))
        }
    }
    panic("No Java overload of new DynamicInitDb for the arguments")
}
func NewDynamicInitDb_1(abeClientInformation AbeClientInformationInterface) *DynamicInitDb {
    this := &DynamicInitDb{}
    this.InitDb = *NewInitDb()
        this.SetHelper(InitDbHelperFactory_GetInstance(abeClientInformation))
    return this
}


func NewDynamicInitDb_2(abeClientInformation AbeClientInformationInterface, databaseConnectionInfoInterface DatabaseConnectionInfoInterface) *DynamicInitDb {
    this := &DynamicInitDb{}
    this.InitDb = *NewInitDb()
        this.SetHelper(InitDbHelperFactory_GetInstance(abeClientInformation, databaseConnectionInfoInterface))
    return this
}

