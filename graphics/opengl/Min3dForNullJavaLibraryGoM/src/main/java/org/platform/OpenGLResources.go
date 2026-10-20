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

type OpenGLResources struct {
}

var OpenGLResources_Instance *OpenGLResources = NewOpenGLResources()

func NewOpenGLResources() *OpenGLResources {
    this := &OpenGLResources{}
    return this
}


func OpenGLResources_GetInstance() *OpenGLResources {
        return OpenGLResources_Instance
}


func (this *OpenGLResources) Init() {
}


func (this *OpenGLResources) IsExcluded(name string) bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLResources) GetIdentifier(name string, defType string, defPackage string) int {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *OpenGLResources) OpenRawResource(id int) IInputStream {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *OpenGLResources) OpenRawResourceFromName(resourceName string) IInputStream {
        panic(errors.New(fmt.Sprint()))
}


// throws Exception
func (this *OpenGLResources) GetResourceAsStream(resourceName string) IInputStream {
        panic(errors.New(fmt.Sprint()))
}


func (this *OpenGLResources) GetPackageName() string {
        panic(errors.New(fmt.Sprint()))
}

