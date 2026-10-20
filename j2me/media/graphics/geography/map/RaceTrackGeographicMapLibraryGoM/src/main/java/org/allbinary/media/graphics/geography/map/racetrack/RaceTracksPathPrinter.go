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

type RaceTracksPathPrinter struct {
}

// The Java class RaceTracksPathPrinter has subclasses so its variables are this interface
type IRaceTracksPathPrinter interface {
    RaceTracksPathPrinter_() *RaceTracksPathPrinter
    PrintPath(list IBasicArrayList)
    PrintPaths(list IBasicArrayList)
}

func (this *RaceTracksPathPrinter) RaceTracksPathPrinter_() *RaceTracksPathPrinter {
    return this
}


func NewRaceTracksPathPrinter() *RaceTracksPathPrinter {
    this := &RaceTracksPathPrinter{}
    return this
}


func (this *RaceTracksPathPrinter) PrintPath(list IBasicArrayList) {
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var anyType any = list.Get(index); _ = anyType
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append("pathList.add(cellPositionFactory.getInstance").Append(fmt.Sprint(anyType)).Append(");")), this, "printPath")
        }
}


func (this *RaceTracksPathPrinter) PrintPaths(list IBasicArrayList) {
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            fmt.Println(fmt.Sprint(NewStringMaker().Append(" Path ").Appendint(index).Append(CommonLabels_GetInstance().COLON_SEP)))
            this.PrintPath(any(list.Get(index)).(IBasicArrayList))
        }
}

