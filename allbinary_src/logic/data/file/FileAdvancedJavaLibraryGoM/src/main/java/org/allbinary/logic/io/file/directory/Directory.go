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

type Directory struct {
}

var Directory_Instance *Directory = NewDirectory()


func Directory_GetInstance() *Directory {
        return Directory_Instance
}


func NewDirectory() *Directory {
    this := &Directory{}
    return this
}


func (this *Directory) Create(directoryAbPath *AbPath) bool {
        return this.Create(directoryAbPath.ToFileSystemString())
}


// throws Exception
func (this *Directory) Remove(existingDirectoryAbPath *AbPath) {
        this.Remove(existingDirectoryAbPath.ToFileSystemString())
}


// Java overloaded method: Go needs a different name
func (this *Directory) Create(directory string) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            PreLogUtil_PutOE("Error Creating Directories: " + directory, this, "create", e)
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var directoryFile *AbFile = FileFactory_GetInstance().GetInstance(directory); _ = directoryFile
            if !directoryFile.IsDirectory() {
                PreLogUtil_Put("Creating Directories: " + directory, this, "create")
                if !directoryFile.Mkdirs() {
                    tryControl1, tryResult1 = 1, false; return
                }
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *Directory) Remove(existingDirectory string) {
        var existingDirectoryFile *AbFile = FileFactory_GetInstance().GetInstance(existingDirectory); _ = existingDirectoryFile
        if existingDirectoryFile.IsDirectory() {
            if len(existingDirectoryFile.List()) > 0 {
                panic(errors.New(fmt.Sprint("Did not remove category: " + existingDirectory + " because files exist")))
            } else if existingDirectoryFile.Delete() {
                PreLogUtil_Put("Successfully Removed: " + existingDirectory, "Directory", "remove")
            }
        } else {
            panic(errors.New(fmt.Sprint("Category Directory Is Missing:" + existingDirectory)))
        }
}


func (this *Directory) Search(fileFilter *FileFilter, file *AbFile) *BasicArrayList {
        return this.Search(fileFilter, file, false)
}


// Java overloaded method: Go needs a different name
func (this *Directory) Search(fileFilter *FileFilter, file *AbFile, isRecursiveSearch bool) *BasicArrayList {
        var fileList *BasicArrayList = NewBasicArrayListD(); _ = fileList
        var recursiveFileList *BasicArrayList; _ = recursiveFileList
        if file.IsDirectory() {
            var fileArray []any = file.ListFilesFileFilter(fileFilter); _ = fileArray
            if fileArray == nil {
                return fileList
            }
            var files []*AbFile = FileWrapperUtil_WrapFiles(fileArray); _ = files
            if files == nil {
                return fileList
            }
            var size int = len(files); _ = size
            for index := int(0); index < size; index++ {
                fileList.Add(files[index])
                if isRecursiveSearch {
                    recursiveFileList = this.Search(fileFilter, files[index], isRecursiveSearch)
                    fileList.AddAllList(recursiveFileList)
                }
            }
        }
        return fileList
}


// Java overloaded method: Go needs a different name
func (this *Directory) Search(file *AbFile) *BasicArrayList {
        return this.Search(file, false)
}


// Java overloaded method: Go needs a different name
func (this *Directory) Search(file *AbFile, isRecursiveSearch bool) *BasicArrayList {
        var fileList *BasicArrayList = NewBasicArrayListD(); _ = fileList
        var recursiveFileList *BasicArrayList; _ = recursiveFileList
        if file.IsDirectory() {
            var fileArray []any = file.ListFiles(); _ = fileArray
            if fileArray == nil {
                return fileList
            }
            var files []*AbFile = FileWrapperUtil_WrapFiles(fileArray); _ = files
            if files == nil {
                return fileList
            }
            for index := int(0); index < len(files); index++ {
                fileList.Add(files[index])
                if isRecursiveSearch {
                    recursiveFileList = this.Search(files[index], isRecursiveSearch)
                    fileList.AddAllList(recursiveFileList)
                }
            }
        }
        return fileList
}


// Java overloaded method: Go needs a different name
func (this *Directory) Search(searchValue string, file *AbFile) *BasicArrayList {
        return this.Search(searchValue, file, false)
}


// Java overloaded method: Go needs a different name
func (this *Directory) Search(searchValue string, file *AbFile, isRecursiveSearch bool) *BasicArrayList {
        var fileList *BasicArrayList = NewBasicArrayListD(); _ = fileList
        var recursiveFileList *BasicArrayList; _ = recursiveFileList
        if file.IsDirectory() {
            var fileArray []any = file.ListFiles(); _ = fileArray
            if fileArray == nil {
                return fileList
            }
            var files []*AbFile = FileWrapperUtil_WrapFiles(fileArray); _ = files
            if files == nil {
                return fileList
            }
            for index := int(0); index < len(files); index++ {
                if files[index].GetPath().IndexOf(searchValue) >= 0 {
                    fileList.Add(files[index])
                }
                if isRecursiveSearch {
                    recursiveFileList = this.Search(searchValue, files[index], isRecursiveSearch)
                    fileList.AddAllList(recursiveFileList)
                }
            }
        }
        return fileList
}


// Java overloaded method: Go needs a different name
func (this *Directory) Search(level int, file *AbFile) *BasicArrayList {
        return this.Search(level, file, false)
}


// Java overloaded method: Go needs a different name
func (this *Directory) Search(level int, file *AbFile, isRecursiveSearch bool) *BasicArrayList {
        var fileList *BasicArrayList = NewBasicArrayListD(); _ = fileList
        var recursiveFileList *BasicArrayList; _ = recursiveFileList
        if file.IsDirectory() {
            var fileArray []any = file.ListFiles(); _ = fileArray
            if fileArray == nil {
                return fileList
            }
            var files []*AbFile = FileWrapperUtil_WrapFiles(fileArray); _ = files
            if files == nil {
                return fileList
            }
            for index := int(0); index < len(files); index++ {
                fileList.Add(files[index])
                if level <= 0 {
                    return fileList
                }
                recursiveFileList = this.Search(level - 1, files[index], isRecursiveSearch)
                fileList.AddAllList(recursiveFileList)
            }
        }
        return fileList
}

