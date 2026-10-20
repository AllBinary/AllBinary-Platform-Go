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

type FileListFetcher struct {
    SubDirectory *SubDirectory
}

var FileListFetcher_Instance *FileListFetcher = NewFileListFetcher()

func (this *FileListFetcher) initFields() {
    this.SubDirectory = SubDirectory_GetInstance()
}

func NewFileListFetcher() *FileListFetcher {
    this := &FileListFetcher{}
    this.initFields()
    return this
}


func FileListFetcher_GetInstance() *FileListFetcher {
        return FileListFetcher_Instance
}


func (this *FileListFetcher) GetFiles(path string) *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var files *BasicArrayList = this.SubDirectory.Search(AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


// Java overloaded method: Go needs a different name
func (this *FileListFetcher) GetFiles(path string, includeExtensions []string) *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList *BasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            var size int = len(includeExtensions); _ = size
            for index := int(0); index < size; index++ {
                includeExtensionBasicArrayList.Add(includeExtensions[index])
            }
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionBooleanFileVisitor(includeExtensionBasicArrayList)); _ = visitorFileFilter
            var files *BasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


// Java overloaded method: Go needs a different name
func (this *FileListFetcher) GetFiles(path string, includeExtension string) *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList *BasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            includeExtensionBasicArrayList.Add(includeExtension)
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionBooleanFileVisitor(includeExtensionBasicArrayList)); _ = visitorFileFilter
            var files *BasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


// Java overloaded method: Go needs a different name
func (this *FileListFetcher) GetFiles(path string, pathIncludes string, includeExtension string) *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList *BasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            includeExtensionBasicArrayList.Add(includeExtension)
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionBooleanFileVisitor(includeExtensionBasicArrayList, pathIncludes)); _ = visitorFileFilter
            var files *BasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


func (this *FileListFetcher) GetTrackedFiles(path string, includeExtensions []string) *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList *BasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            var size int = len(includeExtensions); _ = size
            for index := int(0); index < size; index++ {
                includeExtensionBasicArrayList.Add(includeExtensions[index])
            }
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor(includeExtensionBasicArrayList)); _ = visitorFileFilter
            var files *BasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


// Java overloaded method: Go needs a different name
func (this *FileListFetcher) GetTrackedFiles(path string, includeExtension string) *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList *BasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            includeExtensionBasicArrayList.Add(includeExtension)
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor(includeExtensionBasicArrayList)); _ = visitorFileFilter
            var files *BasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


func (this *FileListFetcher) GetDirectories(path string) *BasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 *BasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryBooleanFileVisitor()); _ = visitorFileFilter
            var files *BasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}

