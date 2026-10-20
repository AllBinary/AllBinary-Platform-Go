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

type BasicHighScoresFactory struct {
    HighScoresBase
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    HighScoresArray []IHighScores
    TOP string
    SCORES string
    PERSONAL_HIGH_SCORES string
    FETCH string
}


func (this *BasicHighScoresFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    this.HighScoresArray = make([]IHighScores, 1)
    this.TOP = "Top"
    this.SCORES = "Scores"
    this.PERSONAL_HIGH_SCORES = "Personal Top Scores"
    this.FETCH = "fetchHighScores"
}


func NewBasicHighScoresFactory(abeClientInformation AbeClientInformationInterface, softwareInformation *SoftwareInformation) *BasicHighScoresFactory {
    this := &BasicHighScoresFactory{}
    this.HighScoresBase = *NewHighScoresBase()
    this.initFields()
        this.AbeClientInformation = abeClientInformation
    return this
}


func (this *BasicHighScoresFactory) FetchHighScores(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener) {
        this.LogUtil.PutF("Getting Local HighScores", this, this.FETCH)
        this.FetchHighScoresPreload(gameInfo, highScoresResultsListener, true)
}


func (this *BasicHighScoresFactory) FetchHighScoresPreload(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener, preload bool) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.FETCH, e)
                }
            }()
            this.HighScoresArray[0] = RecordStoreHighScores_GetInstance(this.AbeClientInformation, gameInfo, this.TOP, this.PERSONAL_HIGH_SCORES, this.SCORES, NewScoreComparator(true))
            highScoresResultsListener.SetHighScoresArray(this.HighScoresArray)
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *BasicHighScoresFactory) CreateHighScoresHelper() HighScoresHelperBaseInterface {
        return NewHighScoresHelper2()
}


func BasicHighScoresFactory_Loaded(index2 int) bool {
        if index2 >= 0 {
            return true
        }
        return false
}

