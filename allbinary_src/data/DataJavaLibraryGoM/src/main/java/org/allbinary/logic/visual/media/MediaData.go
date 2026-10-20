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

type MediaData struct {
    Name string
    Type string
    Extension string
}

var MediaData_HashMap *HashMap = StdUtil_GetInstance().CreateHashMap()

func (this *MediaData) initFields() {
    
    
    
}


func NewMediaData(a_Name string) *MediaData {
    this := &MediaData{}
    this.initFields()
        this.Name = a_Name
        this.Extension = "." + this.Name
        this.Type = MediaTypeData_GetInstance().NOT_RESIZABLE_MEDIA
        MediaData_HashMap.Put(a_Name, this)
    return this
}


func (this *MediaData) GetName() string {
        return this.Name
}


func (this *MediaData) GetType() string {
        return this.Type
}


func (this *MediaData) GetExtension() string {
        return this.Extension
}


func (this *MediaData) IsConvertableTo(mediaData *MediaData) bool {
        if this == MediaDataFactory_GetInstance().GIF && mediaData == MediaDataFactory_GetInstance().JPG {
            return true
        }
        return false
}


func (this *MediaData) ToString() string {
        return this.GetName()
}

func (this *MediaData) String() string {
    return this.ToString()
}


// throws Exception
func MediaData_Get(a_Name string) *MediaData {
        var mediaData *MediaData = any(MediaData_HashMap.Get(a_Name)).(*MediaData); _ = mediaData
        if mediaData != nil {
            return mediaData
        } else {
            panic(errors.New(fmt.Sprint("No Such MediaData")))
        }
}


func MediaData_GetDefault() *MediaData {
        return MediaDataFactory_GetInstance().JPG
}


func MediaData_ToHashMap() *HashMap {
        return MediaData_HashMap
}

