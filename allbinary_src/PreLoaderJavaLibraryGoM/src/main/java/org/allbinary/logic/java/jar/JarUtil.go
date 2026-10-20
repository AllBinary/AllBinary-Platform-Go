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

type JarUtil struct {
    JarURL *URL
}


func (this *JarUtil) initFields() {
    
}


func NewJarUtil(jarURL *URL) *JarUtil {
    this := &JarUtil{}
    this.initFields()
        this.JarURL = jarURL
    return this
}


func (this *JarUtil) Show() {
        var cache []int8 = make([]int8, 1024); _ = cache
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    ex := Java_ToError(recovered)
                    _ = ex
            fmt.Println(ex)
            fmt.Println(ex)
                }
            }()
            var conn *URLConnection = this.JarURL.OpenConnection(); _ = conn
            var jis *JarInputStream = NewJarInputStream(conn.GetInputStream()); _ = jis
            for true {
                var entry *JarEntry = jis.GetNextJarEntry(); _ = entry
                if entry != nil {
                    if !entry.IsDirectory() {
                        var offset int = 0; _ = offset
                        var i int = 0; _ = i
                        for (func() int { i = jis.Read(cache, offset, len(cache) - offset); return i }()) != -1 {
                            offset += i
                            if offset >= len(cache) {
                                var newcache []int8 = make([]int8, len(cache) + 1024); _ = newcache
                                copy(newcache[0:0 + len(cache)], cache[0:0 + len(cache)])
                                cache = newcache
                            }
                        }
                        var tmp []int8 = make([]int8, offset); _ = tmp
                        copy(tmp[0:0 + offset], cache[0:0 + offset])
                        fmt.Println(entry.GetName())
                    }
                } else {
                    break
                }
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


// throws Exception
func JarUtil_Main(args []string) {
        NewJarUtil(NewURL("file:///G:/mnt/bc/mydev/working/j2me/MiniSpaceWars/MiniSpaceWars/dist/MiniSpaceWars.jar")).Show()
}

