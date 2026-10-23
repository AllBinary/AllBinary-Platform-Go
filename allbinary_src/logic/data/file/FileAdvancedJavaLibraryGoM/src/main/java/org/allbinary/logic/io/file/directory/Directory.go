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


func (this *Directory) Create_1(directoryAbPath IAbPath) bool {
        return this.Create(directoryAbPath.ToFileSystemString())
}

// Java overloads of Directory.create
func (this *Directory) Create(args ...any) bool {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[IAbPath](args[0], pass) {
            return this.Create_1(Java_Arg[IAbPath](args[0]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return this.Create_2(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of Directory.create for the arguments")
}


// throws Exception
func (this *Directory) Remove_1(existingDirectoryAbPath IAbPath) {
        this.Remove(existingDirectoryAbPath.ToFileSystemString())
}

// Java overloads of Directory.remove
func (this *Directory) Remove(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[IAbPath](args[0], pass) {
            this.Remove_1(Java_Arg[IAbPath](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            this.Remove_2(Java_Arg[string](args[0]))
            return
        }
    }
    panic("No Java overload of Directory.remove for the arguments")
}


func (this *Directory) Create_2(directory string) bool {
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
func (this *Directory) Remove_2(existingDirectory string) {
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


func (this *Directory) Search_1(fileFilter *FileFilter, file *AbFile) IBasicArrayList {
        return this.Search(fileFilter, file, false)
}

// Java overloads of Directory.search
func (this *Directory) Search(args ...any) IBasicArrayList {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*FileFilter](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            return this.Search_1(Java_Arg[*FileFilter](args[0]), Java_Arg[*AbFile](args[1]))
        }
        if len(args) == 3 && Java_IsArg[*FileFilter](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) && Java_IsArg[bool](args[2], pass) {
            return this.Search_2(Java_Arg[*FileFilter](args[0]), Java_Arg[*AbFile](args[1]), Java_Arg[bool](args[2]))
        }
        if len(args) == 1 && Java_IsArg[*AbFile](args[0], pass) {
            return this.Search_3(Java_Arg[*AbFile](args[0]))
        }
        if len(args) == 2 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[bool](args[1], pass) {
            return this.Search_4(Java_Arg[*AbFile](args[0]), Java_Arg[bool](args[1]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            return this.Search_5(Java_Arg[string](args[0]), Java_Arg[*AbFile](args[1]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) && Java_IsArg[bool](args[2], pass) {
            return this.Search_6(Java_Arg[string](args[0]), Java_Arg[*AbFile](args[1]), Java_Arg[bool](args[2]))
        }
        if len(args) == 2 && Java_IsArg[int](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            return this.Search_7(Java_Arg[int](args[0]), Java_Arg[*AbFile](args[1]))
        }
        if len(args) == 3 && Java_IsArg[int](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) && Java_IsArg[bool](args[2], pass) {
            return this.Search_8(Java_Arg[int](args[0]), Java_Arg[*AbFile](args[1]), Java_Arg[bool](args[2]))
        }
    }
    panic("No Java overload of Directory.search for the arguments")
}


func (this *Directory) Search_2(fileFilter *FileFilter, file *AbFile, isRecursiveSearch bool) IBasicArrayList {
        var fileList IBasicArrayList = NewBasicArrayListD(); _ = fileList
        var recursiveFileList IBasicArrayList; _ = recursiveFileList
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


func (this *Directory) Search_3(file *AbFile) IBasicArrayList {
        return this.Search(file, false)
}


func (this *Directory) Search_4(file *AbFile, isRecursiveSearch bool) IBasicArrayList {
        var fileList IBasicArrayList = NewBasicArrayListD(); _ = fileList
        var recursiveFileList IBasicArrayList; _ = recursiveFileList
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


func (this *Directory) Search_5(searchValue string, file *AbFile) IBasicArrayList {
        return this.Search(searchValue, file, false)
}


func (this *Directory) Search_6(searchValue string, file *AbFile, isRecursiveSearch bool) IBasicArrayList {
        var fileList IBasicArrayList = NewBasicArrayListD(); _ = fileList
        var recursiveFileList IBasicArrayList; _ = recursiveFileList
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


func (this *Directory) Search_7(level int, file *AbFile) IBasicArrayList {
        return this.Search(level, file, false)
}


func (this *Directory) Search_8(level int, file *AbFile, isRecursiveSearch bool) IBasicArrayList {
        var fileList IBasicArrayList = NewBasicArrayListD(); _ = fileList
        var recursiveFileList IBasicArrayList; _ = recursiveFileList
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

