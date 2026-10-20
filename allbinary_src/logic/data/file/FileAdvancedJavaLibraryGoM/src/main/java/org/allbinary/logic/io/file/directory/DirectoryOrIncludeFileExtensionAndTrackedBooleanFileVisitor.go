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

type DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor struct {
    IncludeFileExtensionsBooleanFileVisitor
    TrackedStrings *TrackedStrings
    GitProcessHelper *GitProcessHelper
    IncludesString string
}


func (this *DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor) initFields() {
    this.TrackedStrings = TrackedStrings_GetInstance()
    this.GitProcessHelper = GitProcessHelper_GetInstance()
    
}


func NewDirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor(filterStringBasicArrayList *BasicArrayList) *DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor {
    this := &DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor{}
        this.IncludeFileExtensionsBooleanFileVisitor = *NewIncludeFileExtensionsBooleanFileVisitor(filterStringBasicArrayList)
        this.initFields()
        this.IncludesString = nil
    return this
}


// Java overloaded constructor: Go needs a different name
func NewDirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor(filterStringBasicArrayList *BasicArrayList, includesString string) *DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor {
    this := &DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor{}
        this.IncludeFileExtensionsBooleanFileVisitor = *NewIncludeFileExtensionsBooleanFileVisitor(filterStringBasicArrayList)
        this.initFields()
        this.IncludesString = includesString
    return this
}


func (this *DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor) Visit(file *AbFile) bool {
        if file.IsDirectory() {
            return true
        }
        if this.IncludesString == "" || file.GetAbsolutePath().IndexOf(this.IncludesString) >= 0 {
            return this.IncludeFileExtensionsBooleanFileVisitor.Visit(file)
        } else {
            return false
        }
}


// Java overloaded method: Go needs a different name
func (this *DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor) Visit(file *AbFile, fileNameString string) bool {
        var filePath string = file.GetPath(); _ = filePath
        if this.IncludeFileExtensionsBooleanFileVisitor.Visit(file, fileNameString) {
            if !(strings.Contains(filePath, this.TrackedStrings.APPLICATION) || strings.Contains(filePath, this.TrackedStrings.APP) || strings.Contains(filePath, this.TrackedStrings.HTML_TEMP)) {
                if !(strings.Contains(filePath, this.TrackedStrings.TARGET_PATH_WINDOWS) || strings.Contains(filePath, this.TrackedStrings.TARGET_PATH_UNIX)) {
                    if this.GitProcessHelper.IsTracked(file) {
                        return true
                    } else {
                    }
                } else {
                }
            } else {
            }
        }
        return false
}


func DirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor_Main(args []string) {
        var commonFileStrings *CommonDataFileStrings = CommonDataFileStrings_GetInstance(); _ = commonFileStrings
        var list *BasicArrayList = NewBasicArrayListD(); _ = list
        list.Add(commonFileStrings.JAVA)
        var filePath string = "G:\\mnt\\bc\\mydev\\games\\ZeptoRacer\\platform\\j2se\\ZeptoRacerStaticPathsJ2SEM\\target\\dependency\\org\\allbinary\\logic\\system\\security\\licensing\\ZeptoRacerPCClientInformation.java"; _ = filePath
        var result bool = NewDirectoryOrIncludeFileExtensionAndTrackedBooleanFileVisitor(list).Visit(AbFile_CreateAbFileFromRawPath(filePath)); _ = result
        fmt.Println(filePath + " is Tracked: " + fmt.Sprint(result))
}

