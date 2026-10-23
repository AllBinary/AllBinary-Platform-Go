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


func (this *FileListFetcher) GetFiles_1(path string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var files IBasicArrayList = this.SubDirectory.Search(AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}

// Java overloads of FileListFetcher.getFiles
func (this *FileListFetcher) GetFiles(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.GetFiles_1(Java_Arg[string](args[0]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[[]string](args[1], pass) {
            return this.GetFiles_2(Java_Arg[string](args[0]), Java_Arg[[]string](args[1]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.GetFiles_3(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) && Java_IsArg[string](args[2], pass) {
            return this.GetFiles_4(Java_Arg[string](args[0]), Java_Arg[string](args[1]), Java_Arg[string](args[2]))
        }
    }
    panic("No Java overload of FileListFetcher.getFiles for the arguments")
}


func (this *FileListFetcher) GetFiles_2(path string, includeExtensions []string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            var size int = len(includeExtensions); _ = size
            for index := int(0); index < size; index++ {
                includeExtensionBasicArrayList.Add(includeExtensions[index])
            }
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionBooleanFileVisitor(includeExtensionBasicArrayList)); _ = visitorFileFilter
            var files IBasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


func (this *FileListFetcher) GetFiles_3(path string, includeExtension string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            includeExtensionBasicArrayList.Add(includeExtension)
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionBooleanFileVisitor(includeExtensionBasicArrayList)); _ = visitorFileFilter
            var files IBasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


func (this *FileListFetcher) GetFiles_4(path string, pathIncludes string, includeExtension string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            includeExtensionBasicArrayList.Add(includeExtension)
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionBooleanFileVisitor(includeExtensionBasicArrayList, pathIncludes)); _ = visitorFileFilter
            var files IBasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


func (this *FileListFetcher) GetTrackedFiles_1(path string, includeExtensions []string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            var size int = len(includeExtensions); _ = size
            for index := int(0); index < size; index++ {
                includeExtensionBasicArrayList.Add(includeExtensions[index])
            }
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor(includeExtensionBasicArrayList)); _ = visitorFileFilter
            var files IBasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}

// Java overloads of FileListFetcher.getTrackedFiles
func (this *FileListFetcher) GetTrackedFiles(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[[]string](args[1], pass) {
            return this.GetTrackedFiles_1(Java_Arg[string](args[0]), Java_Arg[[]string](args[1]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return this.GetTrackedFiles_2(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
        }
    }
    panic("No Java overload of FileListFetcher.getTrackedFiles for the arguments")
}


func (this *FileListFetcher) GetTrackedFiles_2(path string, includeExtension string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var includeExtensionBasicArrayList IBasicArrayList = NewBasicArrayListD(); _ = includeExtensionBasicArrayList
            includeExtensionBasicArrayList.Add(includeExtension)
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor(includeExtensionBasicArrayList)); _ = visitorFileFilter
            var files IBasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}


func (this *FileListFetcher) GetDirectories(path string) IBasicArrayList {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IBasicArrayList) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            fmt.Println(e)
            fmt.Println("Error: " + fmt.Sprint(e) + "\nMsg: " + e.Error())
                }
            }()
            var visitorFileFilter *VisitorFileFilter = NewVisitorFileFilter(NewDirectoryBooleanFileVisitor()); _ = visitorFileFilter
            var files IBasicArrayList = this.SubDirectory.Search(visitorFileFilter, AbFile_CreateAbFile(path)); _ = files
            tryControl1, tryResult1 = 1, files; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        return nil
}

