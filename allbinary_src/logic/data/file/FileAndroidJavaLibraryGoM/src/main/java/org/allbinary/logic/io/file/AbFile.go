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

type AbFile struct {
    File *File
}

// Implements AbFileInterface

var AbFile_NULL_FILE *AbFile = AbFile_CreateAbFileFromRawPath(StringUtil_GetInstance().EMPTY_STRING)

func (this *AbFile) initFields() {
    
}


// throws Exception
func AbFile_CreateAbFileWithChild(file *AbFile, childPathName string) *AbFile {
        return NewAbFile(NewFile(file.GetFile(), childPathName))
}


// throws Exception
func AbFile_CreateAbFile(filePath string) *AbFile {
        return NewAbFile(NewFile(NewAbPath(filePath, StringUtil_GetInstance().EMPTY_STRING).ToFileSystemString()))
}


// throws Exception
func AbFile_CreateAbFilePathAndName(filePath string, fileName string) *AbFile {
        return NewAbFile(NewFile(NewAbPath(filePath, StringUtil_GetInstance().EMPTY_STRING).ToFileSystemString(), fileName))
}


func AbFile_CreateAbFileFromAbPath(abPath IAbPath) *AbFile {
        return NewAbFile(NewFile(abPath.ToFileSystemString()))
}


func AbFile_CreateAbFileFromRawPath(filePath string) *AbFile {
        return NewAbFile(NewFile(filePath))
}


func NewAbFile(file *File) *AbFile {
    this := &AbFile{}
    this.initFields()
        this.File = file
    return this
}


func (this *AbFile) GetFile() *File {
        return this.File
}


func (this *AbFile) GetName() string {
        return this.File.GetName()
}


func (this *AbFile) GetParent() string {
        return this.File.GetParent()
}


func (this *AbFile) GetParentFile() *File {
        return this.File.GetParentFile()
}


func (this *AbFile) GetPath() string {
        return this.File.GetPath()
}


func (this *AbFile) IsAbsolute() bool {
        return this.File.IsAbsolute()
}


func (this *AbFile) GetAbsolutePath() string {
        return this.File.GetAbsolutePath()
}


func (this *AbFile) GetAbsoluteFile() *File {
        return this.File.GetAbsoluteFile()
}


// throws IOException
func (this *AbFile) GetCanonicalPath() string {
        return this.File.GetCanonicalPath()
}


// throws IOException
func (this *AbFile) GetCanonicalFile() *File {
        return this.File.GetCanonicalFile()
}


func (this *AbFile) ToURI() *URI {
        return this.File.ToURI()
}


func (this *AbFile) CanRead() bool {
        return this.File.CanRead()
}


func (this *AbFile) CanWrite() bool {
        return this.File.CanWrite()
}


func (this *AbFile) Exists() bool {
        return File_Exists()
}


func (this *AbFile) IsDirectory() bool {
        return File_IsDirectory()
}


func (this *AbFile) IsFile() bool {
        return this.File.IsFile()
}


func (this *AbFile) IsHidden() bool {
        return this.File.IsHidden()
}


func (this *AbFile) LastModified() int64 {
        return this.File.LastModified()
}


func (this *AbFile) Length() int64 {
        return int64(this.File.Length())
}


// throws IOException
func (this *AbFile) CreateNewFile() bool {
        return this.File.CreateNewFile()
}


// throws IOException
func (this *AbFile) Delete() bool {
        return File_Delete()
}


func (this *AbFile) DeleteOnExit() {
        this.File.DeleteOnExit()
}


func (this *AbFile) List() []string {
        return File_List()
}


func (this *AbFile) ListWithFilter(filter *FilenameFilter) []string {
        return File_List(filter)
}


func (this *AbFile) ListFiles() []any {
        return any(this.File.ListFiles()).([]any)
}


func (this *AbFile) ListFilesFileNameFilter(filter *FilenameFilter) []any {
        return any(this.File.ListFiles(filter)).([]any)
}


func (this *AbFile) ListFilesFileFilter(filter *FileFilter) []any {
        return any(this.File.ListFiles(filter)).([]any)
}


func (this *AbFile) Mkdir() bool {
        return this.File.Mkdir()
}


func (this *AbFile) Mkdirs() bool {
        return this.File.Mkdirs()
}


func (this *AbFile) RenameTo(dest *AbFile) bool {
        return this.File.RenameTo(dest.GetFile())
}


func (this *AbFile) SetLastModified(time int64) bool {
        return this.File.SetLastModified(time)
}


func (this *AbFile) SetReadOnly() bool {
        return this.File.SetReadOnly()
}


func (this *AbFile) CompareTo(pathname *File) int {
        return this.File.CompareTo(pathname)
}


func (this *AbFile) Equals(obj any) bool {
        return TsUtil_GetInstance().EqualsNotstring(this.File, obj)
}


func (this *AbFile) HashCode() int {
        return Java_HashCode(this.File)
}


func (this *AbFile) ToString() string {
        return fmt.Sprint(this.File)
}

func (this *AbFile) String() string {
    return this.ToString()
}

