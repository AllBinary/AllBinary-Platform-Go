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

type XmlRpcClient struct {
    Url *URL
}

// Implements XmlRpcHandler

// The Java class XmlRpcClient has subclasses so its variables are this interface
type IXmlRpcClient interface {
    XmlRpcClient_() *XmlRpcClient
    GetURL() *URL
    SetBasicAuthentication(user string, password string)
    Execute(args ...any) any
}

func (this *XmlRpcClient) XmlRpcClient_() *XmlRpcClient {
    return this
}


func (this *XmlRpcClient) initFields() {
    
}


func NewXmlRpcClient(url string) *XmlRpcClient {
    this := &XmlRpcClient{}
    this.initFields()
        this.Url = NewURL(url)
    return this
}


func (this *XmlRpcClient) GetURL() *URL {
        return this.Url
}


func (this *XmlRpcClient) SetBasicAuthentication(user string, password string) {
}


// throws XmlRpcException, IOException
func (this *XmlRpcClient) Execute_1(method string, params IBasicArrayList) any {
        return NullUtil_GetInstance().NULL_OBJECT
}

// Java overloads of XmlRpcClient.execute
func (this *XmlRpcClient) Execute(args ...any) any {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[IBasicArrayList](args[1], pass) {
            return this.Execute_1(Java_Arg[string](args[0]), Java_Arg[IBasicArrayList](args[1]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[IBasicArrayList](args[1], pass) && Java_IsArg[CryptInterface](args[2], pass) {
            return this.Execute_2(Java_Arg[string](args[0]), Java_Arg[IBasicArrayList](args[1]), Java_Arg[CryptInterface](args[2]))
        }
    }
    panic("No Java overload of XmlRpcClient.execute for the arguments")
}


// throws XmlRpcException, IOException
func (this *XmlRpcClient) Execute_2(method string, params IBasicArrayList, cryptInterface CryptInterface) any {
        return NullUtil_GetInstance().NULL_OBJECT
}

