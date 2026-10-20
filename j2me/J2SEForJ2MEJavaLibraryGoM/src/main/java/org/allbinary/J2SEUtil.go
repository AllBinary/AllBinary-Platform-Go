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

//J2SEForJ2ME
type J2SEUtil struct {
}

var J2SEUtil_Instance *J2SEUtil = NewJ2SEUtil()

func NewJ2SEUtil() *J2SEUtil {
    this := &J2SEUtil{}
    return this
}


func J2SEUtil_GetInstance() *J2SEUtil {
        return J2SEUtil_Instance
}


func (this *J2SEUtil) GetHashMapAsArray(hashMap *HashMap) []any {
        var objectArray []any = hashMap.KeySet().ToArray(); _ = objectArray
        return objectArray
}


func (this *J2SEUtil) GetMapAsArray(map_ Map[any, any]) []any {
        var objectArray []any = map_.KeySet().ToArray(); _ = objectArray
        return objectArray
}


func (this *J2SEUtil) GetSetAsArray(set Set[string]) []any {
        var objectArray []any = set.ToArray(); _ = objectArray
        return objectArray
}

