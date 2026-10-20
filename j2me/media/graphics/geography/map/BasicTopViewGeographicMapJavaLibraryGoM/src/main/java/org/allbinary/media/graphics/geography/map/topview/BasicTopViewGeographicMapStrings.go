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

type BasicTopViewGeographicMapStrings struct {
    DEFAULT string
    OTHER string
    WALL string
    OFF_MAP string
    FLOOR string
    DOOR string
    STAIRS_UP string
    STAIRS_DOWN string
}

var BasicTopViewGeographicMapStrings_Instance *BasicTopViewGeographicMapStrings = NewBasicTopViewGeographicMapStrings()

func (this *BasicTopViewGeographicMapStrings) initFields() {
    this.DEFAULT = "Default"
    this.OTHER = "Other"
    this.WALL = "Wall"
    this.OFF_MAP = "OffMap"
    this.FLOOR = "Floor"
    this.DOOR = "Door"
    this.STAIRS_UP = "StairsUp"
    this.STAIRS_DOWN = "StairsDown"
}

func NewBasicTopViewGeographicMapStrings() *BasicTopViewGeographicMapStrings {
    this := &BasicTopViewGeographicMapStrings{}
    this.initFields()
    return this
}


func BasicTopViewGeographicMapStrings_GetInstance() *BasicTopViewGeographicMapStrings {
        return BasicTopViewGeographicMapStrings_Instance
}

