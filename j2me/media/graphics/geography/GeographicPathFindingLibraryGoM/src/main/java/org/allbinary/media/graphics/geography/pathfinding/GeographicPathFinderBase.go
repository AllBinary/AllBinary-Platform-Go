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

type GeographicPathFinderBase struct {
    CommonStrings *CommonStrings
}

// The Java class GeographicPathFinderBase has subclasses so its variables are this interface
type IGeographicPathFinderBase interface {
    GeographicPathFinderBase_() *GeographicPathFinderBase
    SearchTotalPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, totalPaths int) IBasicArrayList
    SearchTotalPathN(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, totalPaths int, multipassState *MultipassState) IBasicArrayList
}

func (this *GeographicPathFinderBase) GeographicPathFinderBase_() *GeographicPathFinderBase {
    return this
}

var GeographicPathFinderBase_NULL_GEOGRAPHIC_PATH_FINDER_BASE IGeographicPathFinderBase = NewGeographicPathFinderBase()

func (this *GeographicPathFinderBase) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}

func NewGeographicPathFinderBase() *GeographicPathFinderBase {
    this := &GeographicPathFinderBase{}
    this.initFields()
    return this
}


// throws Exception
func (this *GeographicPathFinderBase) SearchTotalPath(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, totalPaths int) IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *GeographicPathFinderBase) SearchTotalPathN(startPathFindingNodeList IBasicArrayList, endPathFindingNodeList IBasicArrayList, totalPaths int, multipassState *MultipassState) IBasicArrayList {
        panic(errors.New(fmt.Sprint()))
}

