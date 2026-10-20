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

//TWB - replace with file visitor version
type ImageFileFilter struct {
    FileFilter
}


func NewImageFileFilter() *ImageFileFilter {
    this := &ImageFileFilter{}
    this.FileFilter = *NewFileFilter()
    return this
}


func (this *ImageFileFilter) Accept(f *File) bool {
        if File_IsDirectory() {
            return true
        }
        var s string = f.GetName(); _ = s
        var i int = strings.LastIndex(s, string('.')); _ = i
        if i > 0 && i < len(s) - 1 {
            var ext string = strings.ToLower(s[i + 1:]); _ = ext
            if (ext == "jpg") || (ext == "jpeg") {
                return true
            }
        }
        return false
}


func (this *ImageFileFilter) GetDescription() string {
        return "Accepts .jpg and .jpeg files"
}

