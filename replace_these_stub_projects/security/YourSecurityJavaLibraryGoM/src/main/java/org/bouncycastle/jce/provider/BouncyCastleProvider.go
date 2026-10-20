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

type BouncyCastleProvider struct {
    Provider
}

var BouncyCastleProvider_Info string = "BouncyCastle Security Provider v1.41"
var BouncyCastleProvider_PROVIDER_NAME string = "BC"


func NewBouncyCastleProvider() *BouncyCastleProvider {
    this := &BouncyCastleProvider{}
        this.Provider = *NewProvider(BouncyCastleProvider_PROVIDER_NAME, 1.41, BouncyCastleProvider_Info)
    return this
}


func (this *BouncyCastleProvider) SetParameter(parameterName string, parameter any) {
}

