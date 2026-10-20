/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type TechRTSFormInput struct {
    RTSFormInput
    EVENT IAllBinaryEventObject
    NoMoneyGameNotificationEvent *GameNotificationEvent
}


func (this *TechRTSFormInput) initFields() {
    this.EVENT = NewAllBinaryEventObject(this)
    
}


func NewTechRTSFormInput(groupInterface []*Group) *TechRTSFormInput {
    this := &TechRTSFormInput{}
        this.RTSFormInput = *NewRTSFormInput(groupInterface)
        this.initFields()
        var basicColorFactory *BasicColorFactory = BasicColorFactory_GetInstance(); _ = basicColorFactory
        this.NoMoneyGameNotificationEvent = NewGameNotificationEvent(this, RTSGameStrings_GetInstance().NO_MONEY, SmallIntegerSingletonFactory_GetInstance().GetAt(2), basicColorFactory.WHITE, BooleanFactory_GetInstance().FALSE)
    return this
}


// throws Exception
func (this *TechRTSFormInput) SetAllBinaryGameLayerManager(allBinaryGameLayerManager IAllBinaryGameLayerManager) {
        this.RTSFormInput.SetAllBinaryGameLayerManager(allBinaryGameLayerManager)
        var geographicMapCompositeInterface GeographicMapCompositeInterface = any(allBinaryGameLayerManager).(GeographicMapCompositeInterface); _ = geographicMapCompositeInterface
        var geographicMapInterface IBasicGeographicMap = geographicMapCompositeInterface.GetGeographicMapInterface()[0]; _ = geographicMapInterface
        this.NoMoneyGameNotificationEvent.SetBasicColorP(geographicMapInterface.GetForegroundBasicColor())
}


// throws Exception
func (this *TechRTSFormInput) ProcessGameSpecific(associatedRtsLayer ICollidableDestroyableDamageableLayer, rtsPlayerLayerInterface RTSPlayerLayerInterface, layerManager IAllBinaryLayerManager, item IABCustomItem, itemIndex int) {
        this.RTSFormInput.Process(layerManager)
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        var technologyRTSInterfaceImageItem *TechnologyRTSInterfaceImageItem = any(item).(*TechnologyRTSInterfaceImageItem); _ = technologyRTSInterfaceImageItem
        var rtsInterface RTSInterface = technologyRTSInterfaceImageItem.GetRtsInterface(); _ = rtsInterface
        this.LogUtil.PutF(fmt.Sprint(stringMaker.Append("isUpgradeable: ").Appendboolean(rtsInterface.IsUpgradeable())), this, commonStrings.PROCESS)
        if rtsInterface.IsUpgradeable() {
            var cost int = rtsInterface.GetUpgradeCost(); _ = cost
            var capital ICapital = rtsPlayerLayerInterface.GetCapital(); _ = capital
            stringMaker.Delete(0, stringMaker.Length())
            this.LogUtil.PutF(fmt.Sprint(stringMaker.Appendint(cost).Append("<=").Appendint(capital.GetTotalMoney())), this, commonStrings.PROCESS)
            if cost <= capital.GetTotalMoney() {
                rtsPlayerLayerInterface.Add(BuildingSound_GetInstance())
                capital.RemoveMoney(cost)
                rtsInterface.Upgrade()
                technologyRTSInterfaceImageItem.Update()
                TechEventHandler_GetInstance().FireEvent(this.EVENT)
            } else {
                rtsPlayerLayerInterface.Add(ErrorSound_GetInstance())
                if !rtsPlayerLayerInterface.ImplmentsArtificialIntelligenceCompositeInterface() {
                    GameNotificationEventHandler_GetInstance().FireEvent(this.NoMoneyGameNotificationEvent)
                }
            }
        }
}

