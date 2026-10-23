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

//data/init/views/TestStore/template/type/genericTemplateObjectConfig.xml
type FileUtil struct {
    LogUtil *LogUtil
    Directory *Directory
    LogConfigTypeFactory *LogConfigTypeFactory
    CommonStrings *CommonStrings
    StreamUtil *StreamUtil
    WRITE_LABEL string
    DATA_LABEL string
    WRITE_METHOD string
}

var FileUtil_Instance *FileUtil = NewFileUtil()

func (this *FileUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.Directory = Directory_GetInstance()
    this.LogConfigTypeFactory = LogConfigTypeFactory_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.StreamUtil = StreamUtil_GetInstance()
    this.WRITE_LABEL = "Write file: "
    this.DATA_LABEL = " data: "
    this.WRITE_METHOD = "write"
}


func FileUtil_GetInstance() *FileUtil {
        return FileUtil_Instance
}


func NewFileUtil() *FileUtil {
    this := &FileUtil{}
    this.initFields()
    return this
}


// throws Exception
func (this *FileUtil) Write_1(inputStream IInputStream, file *AbFile) {
        var dataOutputStream *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(file); _ = dataOutputStream
        this.Write(inputStream, dataOutputStream)
}

// Java overloads of FileUtil.write
func (this *FileUtil) Write(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[IInputStream](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            this.Write_1(Java_Arg[IInputStream](args[0]), Java_Arg[*AbFile](args[1]))
            return
        }
        if len(args) == 2 && Java_IsArg[IInputStream](args[0], pass) && Java_IsArg[*AbDataOutputStream](args[1], pass) {
            this.Write_2(Java_Arg[IInputStream](args[0]), Java_Arg[*AbDataOutputStream](args[1]))
            return
        }
        if len(args) == 3 && Java_IsArg[IInputStream](args[0], pass) && Java_IsArg[*AbDataOutputStream](args[1], pass) && Java_IsArg[[]int8](args[2], pass) {
            this.Write_3(Java_Arg[IInputStream](args[0]), Java_Arg[*AbDataOutputStream](args[1]), Java_Arg[[]int8](args[2]))
            return
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            this.Write_4(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
            return
        }
    }
    panic("No Java overload of FileUtil.write for the arguments")
}


// throws Exception
func (this *FileUtil) Write_2(inputStream IInputStream, dataOutputStream *AbDataOutputStream) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            StreamUtil_GetInstance().Close(dataOutputStream)
            }()
            dataOutputStream = any(StreamUtil_GetInstance().Get(inputStream, dataOutputStream, make([]int8, 16384))).(*AbDataOutputStream)
            dataOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *FileUtil) Write_3(inputStream IInputStream, dataOutputStream *AbDataOutputStream, buffer []int8) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            StreamUtil_GetInstance().Close(dataOutputStream)
            }()
            dataOutputStream = any(StreamUtil_GetInstance().Get(inputStream, dataOutputStream, buffer)).(*AbDataOutputStream)
            dataOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *FileUtil) Copy_1(dataInputStream IAbFileInputStream, dataOutputStream *AbDataOutputStream) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            StreamUtil_GetInstance().Close(dataOutputStream)
            StreamUtil_GetInstance().Close(dataInputStream)
            }()
            dataOutputStream = any(StreamUtil_GetInstance().Get(dataInputStream, dataOutputStream, make([]int8, 16384))).(*AbDataOutputStream)
            dataOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
}

// Java overloads of FileUtil.copy
func (this *FileUtil) Copy(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[IAbFileInputStream](args[0], pass) && Java_IsArg[*AbDataOutputStream](args[1], pass) {
            return this.Copy_1(Java_Arg[IAbFileInputStream](args[0]), Java_Arg[*AbDataOutputStream](args[1]))
        }
        if len(args) == 2 && Java_IsArg[IAbPath](args[0], pass) && Java_IsArg[IAbPath](args[1], pass) {
            return this.Copy_2(Java_Arg[IAbPath](args[0]), Java_Arg[IAbPath](args[1]))
        }
    }
    panic("No Java overload of FileUtil.copy for the arguments")
}


// throws Exception
func (this *FileUtil) FixPath(file *AbFile, path IAbPath, realPath IAbPath, cloud string) IAbPath {
        var beginIndex int = strings.Index(file.GetPath(), string(cloud)) + len(cloud); _ = beginIndex
        var endIndex int = strings.Index(file.GetPath(), string(file.GetName())); _ = endIndex
        var fixedPathString string = fmt.Sprint(realPath.ToFileSystemString()) + file.GetPath()[beginIndex:endIndex]; _ = fixedPathString
        var fixedPath IAbPath = NewAbPath(fixedPathString, StringUtil_GetInstance().EMPTY_STRING); _ = fixedPath
        this.Directory.Create(fixedPath)
        var string_ string = fixedPath.ToFileSystemString(); _ = string_
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
            this.LogUtil.PutF(string_, this, "fixPath")
        }
        var outPath IAbPath = NewAbPath(string_, file.GetName()); _ = outPath
        return outPath
}


// throws Exception
func (this *FileUtil) CopyToCloud_1(file *AbFile, path IAbPath, realPath IAbPath, cloud string) {
        this.CopyToCloud(file, path, realPath, cloud, false, false)
}

// Java overloads of FileUtil.copyToCloud
func (this *FileUtil) CopyToCloud(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 4 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[IAbPath](args[1], pass) && Java_IsArg[IAbPath](args[2], pass) && Java_IsArg[string](args[3], pass) {
            this.CopyToCloud_1(Java_Arg[*AbFile](args[0]), Java_Arg[IAbPath](args[1]), Java_Arg[IAbPath](args[2]), Java_Arg[string](args[3]))
            return
        }
        if len(args) == 6 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[IAbPath](args[1], pass) && Java_IsArg[IAbPath](args[2], pass) && Java_IsArg[string](args[3], pass) && Java_IsArg[bool](args[4], pass) && Java_IsArg[bool](args[5], pass) {
            this.CopyToCloud_2(Java_Arg[*AbFile](args[0]), Java_Arg[IAbPath](args[1]), Java_Arg[IAbPath](args[2]), Java_Arg[string](args[3]), Java_Arg[bool](args[4]), Java_Arg[bool](args[5]))
            return
        }
        if len(args) == 2 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            this.CopyToCloud_3(Java_Arg[*AbFile](args[0]), Java_Arg[*AbFile](args[1]))
            return
        }
        if len(args) == 4 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) && Java_IsArg[bool](args[2], pass) && Java_IsArg[bool](args[3], pass) {
            this.CopyToCloud_4(Java_Arg[*AbFile](args[0]), Java_Arg[*AbFile](args[1]), Java_Arg[bool](args[2]), Java_Arg[bool](args[3]))
            return
        }
    }
    panic("No Java overload of FileUtil.copyToCloud for the arguments")
}


// throws Exception
func (this *FileUtil) CopyToCloud_2(file *AbFile, path IAbPath, realPath IAbPath, cloud string, overwriteNewer bool, overwriteAll bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILEERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Error Copying File File: ")
                stringBuffer.Append(fmt.Sprint(file))
                stringBuffer.Append(" to cloud.")
                this.LogUtil.Put(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyToCloud", e)
            }
            panic(e)
                }
            }()
            var outPath IAbPath = this.FixPath(file, path, realPath, cloud); _ = outPath
            var outFile *AbFile = AbFile_CreateAbFileFromAbPath(outPath); _ = outFile
            if !this.CopyPrepare(file, outFile, overwriteNewer, overwriteAll) {
                tryControl1 = 1; return
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                    var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                    stringBuffer.Append("Out File: ")
                    stringBuffer.Append(outFile.GetPath())
                    stringBuffer.Append(" In File: ")
                    stringBuffer.Append(file.GetPath())
                    stringBuffer.Append(" to cloud.")
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyToCloud")
                }
            }
            var dataOutputStream *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(outFile); _ = dataOutputStream
            var fileInputStream IAbFileInputStream = NewAbFileLocalInputStream(file); _ = fileInputStream
            this.Copy(fileInputStream, dataOutputStream)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *FileUtil) CopyToCloud_3(file *AbFile, outFile *AbFile) {
        this.CopyToCloud(file, outFile, false, false)
}


// throws Exception
func (this *FileUtil) CopyToCloud_4(file *AbFile, outFile *AbFile, overwriteNewer bool, overwriteAll bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILEERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Error Copying File File: ")
                stringBuffer.Append(fmt.Sprint(file))
                stringBuffer.Append(" to cloud.")
                this.LogUtil.Put(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyToCloud", e)
            }
            panic(e)
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if !this.CopyPrepare(file, outFile, overwriteNewer, overwriteAll) {
                tryControl1 = 1; return
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append("Out File: ")
                    stringBuffer.Append(outFile.GetPath())
                    stringBuffer.Append(" In File: ")
                    stringBuffer.Append(file.GetPath())
                    stringBuffer.Append(" to cloud.")
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyToCloud")
                }
            }
            var dataOutputStream *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(outFile); _ = dataOutputStream
            var fileInputStream IAbFileInputStream = NewAbFileLocalInputStream(file); _ = fileInputStream
            this.Copy(fileInputStream, dataOutputStream)
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func (this *FileUtil) CopyPrepare(fromFile *AbFile, toFile *AbFile, overwriteNewer bool, overwriteAll bool) bool {
        if toFile.Exists() {
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if overwriteAll {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append("Overwriting File: ")
                    stringBuffer.Append(toFile.GetPath())
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyFile")
                }
                toFile.Delete()
                toFile.CreateNewFile()
            } else if toFile.LastModified() < fromFile.LastModified() && overwriteNewer {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append("Newer by: ")
                    stringBuffer.Appendlong(fromFile.LastModified() - toFile.LastModified())
                    stringBuffer.Append("ms ")
                    stringBuffer.Append("Copying File: ")
                    stringBuffer.Append(toFile.GetPath())
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyFile")
                }
                toFile.Delete()
                toFile.CreateNewFile()
            } else {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                    stringBuffer.Delete(0, stringBuffer.Length())
                    stringBuffer.Append("File Already Exists")
                    if overwriteNewer {
                        stringBuffer.Append(" And Is Not Older")
                    } else {
                        stringBuffer.Append(" And Not A In Overwrite Mode")
                    }
                    stringBuffer.Append(CommonLabels_GetInstance().COLON_SEP)
                    stringBuffer.Append(toFile.GetPath())
                    this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyFile")
                }
                return false
            }
        } else {
            toFile.CreateNewFile()
        }
        return true
}


// throws Exception
func (this *FileUtil) CopyFile_1(fromFile *AbFile, toFile *AbFile) {
        this.CopyFile(fromFile, toFile, false, false)
}

// Java overloads of FileUtil.copyFile
func (this *FileUtil) CopyFile(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) {
            this.CopyFile_1(Java_Arg[*AbFile](args[0]), Java_Arg[*AbFile](args[1]))
            return
        }
        if len(args) == 4 && Java_IsArg[*AbFile](args[0], pass) && Java_IsArg[*AbFile](args[1], pass) && Java_IsArg[bool](args[2], pass) && Java_IsArg[bool](args[3], pass) {
            this.CopyFile_2(Java_Arg[*AbFile](args[0]), Java_Arg[*AbFile](args[1]), Java_Arg[bool](args[2]), Java_Arg[bool](args[3]))
            return
        }
    }
    panic("No Java overload of FileUtil.copyFile for the arguments")
}


// throws Exception
func (this *FileUtil) CopyFile_2(fromFile *AbFile, toFile *AbFile, overwriteNewer bool, overwriteAll bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILEERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Error Copying File fromFile: ")
                stringBuffer.Append(fmt.Sprint(fromFile))
                stringBuffer.Append(" to: ")
                stringBuffer.Append(fmt.Sprint(toFile))
                this.LogUtil.Put(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyFile", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Copying ")
                stringBuffer.Appendlong(int64(fromFile.Length()))
                stringBuffer.Append(" bytes from File: ")
                stringBuffer.Append(fromFile.GetPath())
                stringBuffer.Append(" to: ")
                stringBuffer.Append(toFile.GetPath())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyFile")
            }
            if !this.CopyPrepare(fromFile, toFile, overwriteNewer, overwriteAll) {
                tryControl1 = 1; return
            }
            var dataOutputStream *AbDataOutputStream = DataOutputStreamFactory_GetInstance().GetInstanceForAbFile(toFile); _ = dataOutputStream
            var fileInputStream IAbFileInputStream = NewAbFileInputStream(fromFile); _ = fileInputStream
            this.Copy(fileInputStream, dataOutputStream)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func FileUtil_GetNewDirectory(fromFile *AbFile) string {
        var newDirectory string = fromFile.GetPath(); _ = newDirectory
        var separatorChar string = FilePathData_GetInstance().PATH_START; _ = separatorChar
        if AbIOSystem_GetInstance().IsType("com.vobject.appengine.java.io") {
            separatorChar = AbPathData_GetInstance().SEPARATOR
        }
        var lastIndex int = strings.LastIndex(newDirectory, string(separatorChar)); _ = lastIndex
        if lastIndex >= len(newDirectory) {
            newDirectory = newDirectory[0:len(newDirectory) - 1]
            lastIndex = strings.LastIndex(newDirectory, string(separatorChar))
        }
        return newDirectory[lastIndex:len(newDirectory)]
}


// throws Exception
func (this *FileUtil) CopyDirectoryPortion(fromDirectoryAbPath IAbPath, toDirectoryAbPath IAbPath, overwriteNewer bool, overwriteAll bool, current int, total int) {
        var file *AbFile = AbFile_CreateAbFileFromAbPath(fromDirectoryAbPath); _ = file
        if !file.IsDirectory() {
            panic(errors.New(fmt.Sprint("Not a directory: " + fmt.Sprint(file.GetPath()))))
        }
        var fileList IBasicArrayList = this.Directory.Search(file, true); _ = fileList
        var size int = fileList.Size(); _ = size
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Searched: ")
        stringBuffer.Append(file.GetPath())
        stringBuffer.Append(" Total: ")
        stringBuffer.Appendint(size)
        var portion int = size / total + 1; _ = portion
        var start int = portion * current; _ = start
        var end int = start + portion; _ = end
        if end > size {
            end = size
        }
        stringBuffer.Append(" Section: ")
        stringBuffer.Appendint(start)
        stringBuffer.Append(" - ")
        stringBuffer.Appendint(end)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
            this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copySomeFilesToDirectory()")
        }
        var nextFile *AbFile; _ = nextFile
        for index := int(start); index < end; index++ {
            nextFile = any(fileList.Get(index)).(*AbFile)
            if nextFile.IsDirectory() {
            } else {
                var path string = NewAbPath(nextFile.GetPath(), StringUtil_GetInstance().EMPTY_STRING).GetPath(); _ = path
                var beginIndex int = len(fromDirectoryAbPath.GetPath()); _ = beginIndex
                if path.IndexOf(URLGLOBALS.GetWebappPath()) < 0 {
                    beginIndex -= len(NewAbPath(URLGLOBALS.GetWebappPath(), StringUtil_GetInstance().EMPTY_STRING).GetPath())
                }
                var newPath string = fmt.Sprint(toDirectoryAbPath.ToFileSystemString()) + path.Substring(beginIndex); _ = newPath
                var toFile *AbFile = AbFile_CreateAbFileFromAbPath(NewAbPath(newPath, StringUtil_GetInstance().EMPTY_STRING)); _ = toFile
                this.CopyFile(nextFile, toFile, overwriteNewer, overwriteAll)
            }
        }
}


// throws Exception
func (this *FileUtil) CopyDirectory(fromFile *AbFile, to *AbFile) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILEERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Error Copying Directory fromFile: ")
                stringBuffer.Append(fmt.Sprint(fromFile))
                stringBuffer.Append(" to: ")
                stringBuffer.Append(fmt.Sprint(to))
                this.LogUtil.Put(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyDirectory", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Copying Directory from: ")
                stringBuffer.Append(fromFile.GetPath())
                stringBuffer.Append(" to: ")
                stringBuffer.Append(to.GetPath())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyDirectory")
            }
            var newDirectory string = FileUtil_GetNewDirectory(fromFile); _ = newDirectory
            var newDirectoryAbPath IAbPath = NewAbPath(fmt.Sprint(to.GetPath() + AbPathData_GetInstance().SEPARATOR) + newDirectory, StringUtil_GetInstance().EMPTY_STRING); _ = newDirectoryAbPath
            if !this.Directory.Create(newDirectoryAbPath) {
                panic(errors.New(fmt.Sprint("Failed to create directory: " + fmt.Sprint(newDirectoryAbPath))))
            }
            var fileArray []*AbFile = FileWrapperUtil_WrapFiles(fromFile.ListFiles()); _ = fileArray
            var size int = len(fileArray); _ = size
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Copying ")
                stringBuffer.Appendint(size)
                stringBuffer.Append(" files from: ")
                stringBuffer.Append(fromFile.GetPath())
                stringBuffer.Append(" to: ")
                stringBuffer.Append(to.GetPath())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), "copyDirectory")
            }
            var file *AbFile; _ = file
            for index := int(0); index < size; index++ {
                file = fileArray[index]
                if file.IsFile() {
                    this.CopyFile(file, AbFile_CreateAbFilePathAndName(fmt.Sprint(newDirectoryAbPath), file.GetName()))
                } else if file.IsDirectory() {
                    this.CopyDirectory(file, AbFile_CreateAbFileFromAbPath(newDirectoryAbPath))
                } else {
                    panic(errors.New(fmt.Sprint("File Copy Error")))
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
// synchronized
func (this *FileUtil) Copy_2(fromAbPath IAbPath, to IAbPath) bool {
        var COPY string = "copy"; _ = COPY
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILEERROR) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Error Copying fromAbPath: ")
                stringBuffer.Append(fmt.Sprint(fromAbPath))
                stringBuffer.Append(" to: ")
                stringBuffer.Append(fmt.Sprint(to))
                this.LogUtil.Put(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), COPY, e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("Copying AbPaths from: ")
                stringBuffer.Append(fmt.Sprint(fromAbPath))
                stringBuffer.Append(" to: ")
                stringBuffer.Append(fmt.Sprint(to))
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), COPY)
            }
            if fromAbPath == nil {
                panic(errors.New(fmt.Sprint("Cannot Copy From A Null Location")))
            }
            if to == nil {
                panic(errors.New(fmt.Sprint("Cannot Copy To A Null Location")))
            }
            var fromLocationFile *AbFile = AbFile_CreateAbFileFromAbPath(fromAbPath); _ = fromLocationFile
            var toLocationFile *AbFile = AbFile_CreateAbFileFromAbPath(to); _ = toLocationFile
            if fromLocationFile.IsFile() {
                if toLocationFile.IsDirectory() {
                    var file *AbFile = AbFile_CreateAbFileWithChild(toLocationFile, fromLocationFile.GetName()); _ = file
                    this.CopyFile(fromLocationFile, file)
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                        stringBuffer.Append("Copied file=")
                        stringBuffer.Append(fromLocationFile.GetName())
                        stringBuffer.Append(" from: ")
                        stringBuffer.Append(fromLocationFile.GetPath())
                        stringBuffer.Append(" to: ")
                        stringBuffer.Append(toLocationFile.GetPath())
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), COPY)
                    }
                } else {
                    this.CopyFile(fromLocationFile, toLocationFile)
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                        stringBuffer.Append("Copied file with new name ")
                        stringBuffer.Append(" from: ")
                        stringBuffer.Append(fromLocationFile.GetPath())
                        stringBuffer.Append(" to: ")
                        stringBuffer.Append(toLocationFile.GetPath())
                        this.LogUtil.PutF(fmt.Sprint(stringBuffer), FileUtil_GetInstance(), COPY)
                    }
                }
            } else if fromLocationFile.IsDirectory() && toLocationFile.IsDirectory() {
                if rune(fromAbPath.ToFileSystemString()[len(fromAbPath.ToFileSystemString()) - 1]) == java_Io.File.SeparatorChar {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                        this.LogUtil.PutF("Copying subdirectories", FileUtil_GetInstance(), COPY)
                    }
                    var fileArray []*AbFile = FileWrapperUtil_WrapFiles(fromLocationFile.ListFiles()); _ = fileArray
                    for index := int(0); index < len(fileArray); index++ {
                        var file *AbFile = fileArray[index]; _ = file
                        if file.IsFile() {
                            var aFile *AbFile = AbFile_CreateAbFileWithChild(toLocationFile, file.GetName()); _ = aFile
                            this.CopyFile(file, aFile)
                        } else if file.IsDirectory() {
                            this.CopyDirectory(file, toLocationFile)
                        } else {
                            panic(errors.New(fmt.Sprint("File Copy Error")))
                        }
                    }
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                        this.LogUtil.PutF("Copied subdirectories", FileUtil_GetInstance(), COPY)
                    }
                } else {
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                        this.LogUtil.PutF("Copyinhg directory", FileUtil_GetInstance(), COPY)
                    }
                    this.CopyDirectory(fromLocationFile, toLocationFile)
                    if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(this.LogConfigTypeFactory.FILE) {
                        this.LogUtil.PutF("Copied directory", FileUtil_GetInstance(), COPY)
                    }
                }
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *FileUtil) ShouldSkip(file *AbFile, skipFiles []string) bool {
        var size int = len(skipFiles); _ = size
        var skipFile string; _ = skipFile
        for index := int(0); index < size; index++ {
            skipFile = skipFiles[index]
            if strings.Compare(file.GetPath(), skipFile) == 0 {
                return true
            }
        }
        return false
}


// throws Exception
func (this *FileUtil) Write_4(filePath string, string_ string) {
        var dataOutputStream *AbDataOutputStream = nil; _ = dataOutputStream
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
            StreamUtil_GetInstance().Close(dataOutputStream)
            }()
            PreLogUtil_Put(fmt.Sprint(NewStringMaker().Append(this.WRITE_LABEL).Append(filePath).Append(this.DATA_LABEL).Append(string_)), this, this.WRITE_METHOD)
            var fileInputStreamFactory *FileStreamFactory = FileStreamFactory_GetInstance(); _ = fileInputStreamFactory
            var fileOutputStream IOutputStream = fileInputStreamFactory.GetFileOutputStreamInstance(StringUtil_GetInstance().EMPTY_STRING, filePath); _ = fileOutputStream
            dataOutputStream = NewAbDataOutputStream(fileOutputStream)
            dataOutputStream.Write(Java_GetBytes(string_), 0, len(string_))
            dataOutputStream.Flush()
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *FileUtil) LoadFileAsList(file *AbFile, max_ int, byteArray1 []int8) IBasicArrayList {
        return FileUtil2_GetInstance().LoadFileAsList(file, max_, byteArray1)
}

