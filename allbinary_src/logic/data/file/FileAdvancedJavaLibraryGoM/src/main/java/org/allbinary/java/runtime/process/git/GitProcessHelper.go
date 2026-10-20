/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2026 AllBinary 
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

//Do not use this instead just skip target, build, Application paths.
type GitProcessHelper struct {
    TrackedStrings *TrackedStrings
}

var GitProcessHelper_Instance *GitProcessHelper = NewGitProcessHelper()

func (this *GitProcessHelper) initFields() {
    this.TrackedStrings = TrackedStrings_GetInstance()
}


func GitProcessHelper_GetInstance() *GitProcessHelper {
        return GitProcessHelper_Instance
}


func NewGitProcessHelper() *GitProcessHelper {
    this := &GitProcessHelper{}
    this.initFields()
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var THIS_IS_SLOW string = "This is slow consider path exclusion instead"; _ = THIS_IS_SLOW
        logUtil.PutF(THIS_IS_SLOW, this, commonStrings.CONSTRUCTOR)
    return this
}


// throws IOException, InterruptedException
func (this *GitProcessHelper) TrackedFiles(rootAsString string, pathspec string) List[string] {
        return this.TrackedFiles(rootAsString, pathspec, ".java")
}


// throws IOException, InterruptedException
// Java overloaded method: Go needs a different name
func (this *GitProcessHelper) TrackedFiles(rootAsString string, pathspec string, suffix string) List[string] {
        var process *Process = NewProcessBuilder(this.TrackedStrings.GIT_COMMAND, this.TrackedStrings.CHANGE_DIRECTORY_OPTION, rootAsString, this.TrackedStrings.LIST_FILES_COMMAND, this.TrackedStrings.SEP_BY_NULL_CHAR_INSTEAD_OF_NEW_LINE, this.TrackedStrings.PATHSPEC_SEPARATOR, pathspec).RedirectError(ProcessBuilder_Redirect.INHERIT).Start(); _ = process
        var output []int8 = process.GetInputStream().ReadAllBytes(); _ = output
        var exitCode int = process.WaitFor(); _ = exitCode
        if exitCode != 0 {
            panic(errors.New(fmt.Sprint(fmt.Sprint(NewStringMaker().Append("git ls-files failed for ").Append(rootAsString).Append(" with exit code ").Appendint(exitCode)))))
        }
        var files List[string] = NewArrayList(); _ = files
        var start int = 0; _ = start
        var relativePath string; _ = relativePath
        for index := int(0); index < len(output); index++ {
            if int(output[index]) == 0 {
                relativePath = NewString(output, start, index - start, StandardCharsets_UTF_8)
                if strings.HasSuffix(relativePath, suffix) {
                    files.Add(relativePath)
                }
                start = index + 1
            }
        }
        return files
}


func (this *GitProcessHelper) IsTracked(file *AbFile) bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var nativeFile *File = AbFileNativeUtil_Get(file); _ = nativeFile
            var parentFile *File = nativeFile.GetParentFile(); _ = parentFile
            var process *Process = NewProcessBuilder(this.TrackedStrings.GIT_COMMAND, this.TrackedStrings.CHANGE_DIRECTORY_OPTION, parentFile.GetPath(), this.TrackedStrings.LIST_FILES_COMMAND, this.TrackedStrings.ERROR_UNMATCH_OPTION, this.TrackedStrings.PATHSPEC_SEPARATOR, nativeFile.GetName()).RedirectErrorStream(true).Start(); _ = process
            process.GetInputStream().ReadAllBytes()
            tryControl1, tryResult1 = 1, process.WaitFor() == 0; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

