/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type CustomMapGeneratorBaseFactory struct {
}

// The Java class CustomMapGeneratorBaseFactory has subclasses so its variables are this interface
type ICustomMapGeneratorBaseFactory interface {
    CustomMapGeneratorBaseFactory_() *CustomMapGeneratorBaseFactory
    Create(raceTrackGeographicMap *RaceTrackGeographicMap) ICustomMapGeneratorBase
}

func (this *CustomMapGeneratorBaseFactory) CustomMapGeneratorBaseFactory_() *CustomMapGeneratorBaseFactory {
    return this
}

var CustomMapGeneratorBaseFactory_Instance ICustomMapGeneratorBaseFactory = NewCustomMapGeneratorBaseFactory()


func CustomMapGeneratorBaseFactory_GetInstance() ICustomMapGeneratorBaseFactory {
        return CustomMapGeneratorBaseFactory_Instance
}


func NewCustomMapGeneratorBaseFactory() *CustomMapGeneratorBaseFactory {
    this := &CustomMapGeneratorBaseFactory{}
    return this
}


// throws Exception
func (this *CustomMapGeneratorBaseFactory) Create(raceTrackGeographicMap *RaceTrackGeographicMap) ICustomMapGeneratorBase {
        return NewCustomMapGeneratorBase()
}

