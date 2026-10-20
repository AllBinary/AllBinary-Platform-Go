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

type Advertisements struct {
    BannerList IBasicArrayList
    PageList IBasicArrayList
    MyRandomFactory *MyRandomFactory
}


func (this *Advertisements) initFields() {
    this.BannerList = NewBasicArrayListD()
    this.PageList = NewBasicArrayListD()
    this.MyRandomFactory = MyRandomFactory_GetInstance()
}

func NewAdvertisements() *Advertisements {
    this := &Advertisements{}
    this.initFields()
    return this
}


func (this *Advertisements) AddPage(advertisementProcessorInterface IAdvertisementProcessorInterface) {
        if !this.PageList.Contains(advertisementProcessorInterface) {
            this.PageList.Add(advertisementProcessorInterface)
        }
}


func (this *Advertisements) StartRandomPage() {
        var size int = this.PageList.Size(); _ = size
        if size > 0 {
            var random int = this.MyRandomFactory.GetAbsoluteNextInt(size); _ = random
            var advertisementProcessorInterface IAdvertisementProcessorInterface = any(this.PageList.BasicArrayList_().ObjectArray[random]).(IAdvertisementProcessorInterface); _ = advertisementProcessorInterface
            advertisementProcessorInterface.Start()
        }
}


func (this *Advertisements) ProcessRandomPage() {
        if LockedUtil_GetInstance().IsLocked() {
            var size int = this.PageList.Size(); _ = size
            if size > 0 {
                var random int = this.MyRandomFactory.GetAbsoluteNextInt(size); _ = random
                var advertisementProcessorInterface IAdvertisementProcessorInterface = any(this.PageList.BasicArrayList_().ObjectArray[random]).(IAdvertisementProcessorInterface); _ = advertisementProcessorInterface
                advertisementProcessorInterface.Process()
            }
        }
}


func (this *Advertisements) StopAllPage() {
        for index := int(this.PageList.Size()); Java_PreDecrement(&index) >= 0;  {
            var advertisementProcessorInterface IAdvertisementProcessorInterface = any(this.PageList.BasicArrayList_().ObjectArray[index]).(IAdvertisementProcessorInterface); _ = advertisementProcessorInterface
            advertisementProcessorInterface.Stop()
        }
}


func (this *Advertisements) AddBanner(advertisementProcessorInterface IAdvertisementProcessorInterface) {
        if !this.BannerList.Contains(advertisementProcessorInterface) {
            this.BannerList.Add(advertisementProcessorInterface)
        }
}


func (this *Advertisements) StartAll() {
        for index := int(this.BannerList.Size()); Java_PreDecrement(&index) >= 0;  {
            var advertisementProcessorInterface IAdvertisementProcessorInterface = any(this.BannerList.BasicArrayList_().ObjectArray[index]).(IAdvertisementProcessorInterface); _ = advertisementProcessorInterface
            advertisementProcessorInterface.Start()
        }
}


func (this *Advertisements) StopAllBanner() {
        for index := int(this.BannerList.Size()); Java_PreDecrement(&index) >= 0;  {
            var advertisementProcessorInterface IAdvertisementProcessorInterface = any(this.BannerList.BasicArrayList_().ObjectArray[index]).(IAdvertisementProcessorInterface); _ = advertisementProcessorInterface
            advertisementProcessorInterface.Stop()
        }
}


func (this *Advertisements) GetBanner(index int) IAdvertisementProcessorInterface {
        var advertisementProcessorInterface IAdvertisementProcessorInterface = any(this.BannerList.BasicArrayList_().ObjectArray[index]).(IAdvertisementProcessorInterface); _ = advertisementProcessorInterface
        return advertisementProcessorInterface
}


func (this *Advertisements) StopAll() {
        this.StopAllBanner()
        this.StopAllPage()
}

