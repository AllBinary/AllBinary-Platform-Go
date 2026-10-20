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

type URIResolverStrings struct {
    IMPORT_URL string
    ATTEMPT string
    BASE string
    NEW_PATH string
    NOTE string
    URL_GLOBAL string
    REQUIRED_EXTENSION string
    RESOLVE string
}

var URIResolverStrings_Instance *URIResolverStrings = NewURIResolverStrings()

func (this *URIResolverStrings) initFields() {
    this.IMPORT_URL = "/{import url}"
    this.ATTEMPT = "attempt to use xsl:import: href="
    this.BASE = "\nBase= "
    this.NEW_PATH = "\nNew path= "
    this.NOTE = "\nNote: "
    this.URL_GLOBAL = " is a urlglobal"
    this.REQUIRED_EXTENSION = "\nRequired Extension: "
    this.RESOLVE = "resolve"
}

func NewURIResolverStrings() *URIResolverStrings {
    this := &URIResolverStrings{}
    this.initFields()
    return this
}


func URIResolverStrings_GetInstance() *URIResolverStrings {
        return URIResolverStrings_Instance
}

