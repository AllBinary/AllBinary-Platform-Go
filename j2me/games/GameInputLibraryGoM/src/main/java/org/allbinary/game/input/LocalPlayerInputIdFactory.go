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

type LocalPlayerInputIdFactory struct {
    LogUtil *LogUtil
    MAX_NUMBER_OF_PLAYERS int
    MAX_FAST_DEVICEID int
    DeviceIdToplayerId []int
    PlayerIdToDeviceId []int
    PlayersInPlay []bool
    TotalDevicesMapped int
}

var LocalPlayerInputIdFactory_Instance *LocalPlayerInputIdFactory = NewLocalPlayerInputIdFactory()

func (this *LocalPlayerInputIdFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.MAX_NUMBER_OF_PLAYERS = 6
    this.MAX_FAST_DEVICEID = 65
    this.DeviceIdToplayerId = make([]int, this.MAX_FAST_DEVICEID)
    this.PlayerIdToDeviceId = make([]int, this.MAX_NUMBER_OF_PLAYERS)
    this.PlayersInPlay = make([]bool, this.MAX_NUMBER_OF_PLAYERS)
    this.TotalDevicesMapped = 0
}


func NewLocalPlayerInputIdFactory() *LocalPlayerInputIdFactory {
    this := &LocalPlayerInputIdFactory{}
    this.initFields()
        for index := int(len(this.DeviceIdToplayerId) - 1); index >= 0; index-- {
            this.DeviceIdToplayerId[index] = -1
        }
    return this
}


func LocalPlayerInputIdFactory_GetInstance() *LocalPlayerInputIdFactory {
        return LocalPlayerInputIdFactory_Instance
}


func (this *LocalPlayerInputIdFactory) GetDeviceId(playerInputId int) int {
        return this.PlayerIdToDeviceId[playerInputId]
}


func (this *LocalPlayerInputIdFactory) GetPlayerForDevice(deviceId int) int {
        var index int; _ = index
        if deviceId < this.MAX_FAST_DEVICEID {
            var playerInputId int = this.DeviceIdToplayerId[deviceId]; _ = playerInputId
            if playerInputId != -1 {
                return playerInputId
            }
        } else {
            for index = this.TotalDevicesMapped - 1; index >= 0; index-- {
                if this.PlayerIdToDeviceId[index] == deviceId {
                    return index
                }
            }
        }
        index = this.TotalDevicesMapped
        if deviceId < this.MAX_FAST_DEVICEID {
            this.DeviceIdToplayerId[deviceId] = index
        } else {
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Added DeviceId: ").Appendint(deviceId).Append(" beyond fast Id list")), this, "getPlayerForDevice")
        }
        this.PlayerIdToDeviceId[index] = deviceId
        AnalogLocationInputFactory_GetInstance().AddPlayer(this.TotalDevicesMapped)
        this.TotalDevicesMapped++
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Added DeviceId: ").Appendint(deviceId).Append(" at: ").Appendint(index).Append(" Total: ").Appendint(this.TotalDevicesMapped)), this, "getPlayerForDevice")
        return index
}


func (this *LocalPlayerInputIdFactory) GetTotalDevicesInPlay() int {
        var total int = 0; _ = total
        for index := int(this.TotalDevicesMapped - 1); index >= 0; index-- {
            if this.PlayersInPlay[index] == true {
                total++
            }
        }
        return total
}


func (this *LocalPlayerInputIdFactory) IsPlayerInPlay(playerInputId int) bool {
        return this.PlayersInPlay[playerInputId]
}


func (this *LocalPlayerInputIdFactory) SetPlayerInPlay(playerInputId int) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Setting PlayerInPlay with playerInputId: ").Appendint(playerInputId)), this, "setPlayerInPlay")
        this.PlayersInPlay[playerInputId] = true
}


func (this *LocalPlayerInputIdFactory) SetPlayerOutOfPlay(playerInputId int) {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Setting PlayerOutOfPlay with playerInputId: ").Appendint(playerInputId)), this, "setPlayerOutOfPlay")
        this.PlayersInPlay[playerInputId] = false
}


func (this *LocalPlayerInputIdFactory) GetPlayersInPlay() []bool {
        return this.PlayersInPlay
}

