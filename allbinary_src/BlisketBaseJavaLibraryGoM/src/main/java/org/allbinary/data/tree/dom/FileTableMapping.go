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

type FileTableMapping struct {
    File *AbFile
}

// Implements TableMappingInterface


func (this *FileTableMapping) initFields() {
    
}


func NewFileTableMapping(file *AbFile) *FileTableMapping {
    this := &FileTableMapping{}
    this.initFields()
        this.File = file
    return this
}


// throws Exception
func (this *FileTableMapping) GetKey() any {
        return any(this.File.GetPath()).(any)
}


// throws Exception
func (this *FileTableMapping) ToVector() IBasicArrayList {
        var filePropertyVector IBasicArrayList = NewBasicArrayListD(); _ = filePropertyVector
        filePropertyVector.Add(this.File.GetAbsolutePath())
        filePropertyVector.Add(this.File.GetCanonicalPath())
        filePropertyVector.Add(this.File.GetName())
        filePropertyVector.Add(this.File.GetParent())
        filePropertyVector.Add(this.File.GetPath())
        return filePropertyVector
}


// throws Exception
func (this *FileTableMapping) ToHashMap() *HashMap {
        var filePropertyHashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = filePropertyHashMap
        var pathData *AbPathData = AbPathData_GetInstance(); _ = pathData
        var pathUtil *PathUtil = PathUtil_GetInstance(); _ = pathUtil
        var rootFileName string = pathData.GetNameFromPath(pathUtil.GetWithoutExtension(this.File.GetPath())); _ = rootFileName
        var fileData *FileData = FileData_GetInstance(); _ = fileData
        filePropertyHashMap.Put(fileData.ROOT_NAME, rootFileName)
        filePropertyHashMap.Put(fileData.ISFILE, fmt.Sprint(fmt.Sprint(this.File.IsFile())))
        filePropertyHashMap.Put(fileData.ISDIRECTORY, fmt.Sprint(fmt.Sprint(this.File.IsDirectory())))
        filePropertyHashMap.Put(fileData.ISHIDDEN, fmt.Sprint(fmt.Sprint(this.File.IsHidden())))
        filePropertyHashMap.Put(fileData.ISABSOLUTE, fmt.Sprint(fmt.Sprint(this.File.IsAbsolute())))
        filePropertyHashMap.Put(fileData.ABSOLUTE_PATH, this.File.GetAbsolutePath())
        filePropertyHashMap.Put(fileData.CANONICAL_PATH, this.File.GetCanonicalPath())
        filePropertyHashMap.Put(fileData.NAME, this.File.GetName())
        filePropertyHashMap.Put(fileData.PARENT, this.File.GetParent())
        filePropertyHashMap.Put(fileData.PATH, this.File.GetPath())
        return filePropertyHashMap
}

