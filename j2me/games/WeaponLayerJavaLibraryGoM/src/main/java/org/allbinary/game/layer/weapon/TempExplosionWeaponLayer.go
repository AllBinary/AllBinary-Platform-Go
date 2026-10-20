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

type TempExplosionWeaponLayer struct {
    SimpleWeaponLayer
    WeaponProperties *WeaponProperties
    TimeDelayHelper ITimeDelayHelper
}


func (this *TempExplosionWeaponLayer) initFields() {
    this.WeaponProperties = NewWeaponProperties(-1, -1, 0, 1000, int16(0))
    this.TimeDelayHelper = NewTimeDelayHelper(0)
}


func NewTempExplosionWeaponLayer(name string, movement IMovement, animationInterface IAnimation, rectangle *Rectangle, viewPosition IViewPositionBase, timeDelay int) *TempExplosionWeaponLayer {
    this := &TempExplosionWeaponLayer{}
        this.SimpleWeaponLayer = *NewSimpleWeaponLayer(name, RemoteInfo_REMOTE_INFO, -1, movement, animationInterface, SimpleWeaponLayer_CreateDestroyed(), rectangle, viewPosition)
        this.initFields()
        this.TimeDelayHelper.TimeDelayHelper_().Delay = timeDelay
    return this
}


// throws Exception
func (this *TempExplosionWeaponLayer) ProcessTick(allBinaryLayerManager IAllBinaryLayerManager) {
        if this.TimeDelayHelper.IsTime(GameTickTimeDelayHelperFactory_GetInstance().StartTime) {
            this.GetCollidableInferface().Collide(this, this)
            this.TotalDamage = this.GetInitDamage() + 1
        }
        this.SimpleWeaponLayer.ProcessTick(allBinaryLayerManager)
}

