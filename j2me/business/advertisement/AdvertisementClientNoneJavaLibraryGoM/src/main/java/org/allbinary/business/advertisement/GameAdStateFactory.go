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

type GameAdStateFactory struct {
    GameAdStateArray []*GameAdState
    GameAdState *GameAdState
}

// Implements GameAdStateFactoryInterface

var GameAdStateFactory_Instance *GameAdStateFactory = NewGameAdStateFactory()

func (this *GameAdStateFactory) initFields() {
    this.GameAdStateArray = []*GameAdState{NewGameAdState(NewAppNameLicensingAdConfiguration())}
    this.GameAdState = this.GameAdStateArray[0]
}

func NewGameAdStateFactory() *GameAdStateFactory {
    this := &GameAdStateFactory{}
    this.initFields()
    return this
}


func GameAdStateFactory_GetInstance() *GameAdStateFactory {
        return GameAdStateFactory_Instance
}


func (this *GameAdStateFactory) GetCurrentInstance() *GameAdState {
        return this.GameAdState
}


// throws Exception
func (this *GameAdStateFactory) GetInstanceForApp(softwareInformation *SoftwareInformation) IGameAdStateBase {
        this.GameAdState = this.GameAdStateArray[0]
        return this.GameAdStateArray[0]
}


func (this *GameAdStateFactory) IsEnabled() bool {
        return false
}

