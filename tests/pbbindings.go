package tests

import (
	proto "google.golang.org/protobuf/proto"
	amino "github.com/tendermint/go-amino"
	testspb "github.com/tendermint/go-amino/tests/pb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	time "time"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

func (goo EmptyStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.EmptyStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.EmptyStruct
			msg = pbov
			return
		}
		pbo = new(testspb.EmptyStruct)
	}
	msg = pbo
	return
}
func (goo *EmptyStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.EmptyStruct = msg.(*testspb.EmptyStruct)
	{
		if pbo != nil {
		}
	}
	return
}
func (_ EmptyStruct) GetTypeURL() (typeURL string) {
	return "/tests.EmptyStruct"
}
func (goo EmptyStruct) IsEmpty() (empty bool) {
	{
		empty = true
	}
	return
}
func (goo PrimitivesStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.PrimitivesStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.PrimitivesStruct
			msg = pbov
			return
		}
		pbo = new(testspb.PrimitivesStruct)
		{
			pbo.Int8 = int32(goo.Int8)
		}
		{
			pbo.Int16 = int32(goo.Int16)
		}
		{
			pbo.Int32 = goo.Int32
		}
		{
			pbo.Int32Fixed = goo.Int32Fixed
		}
		{
			pbo.Int64 = goo.Int64
		}
		{
			pbo.Int64Fixed = goo.Int64Fixed
		}
		{
			pbo.Int = int64(goo.Int)
		}
		{
			pbo.Byte = uint32(goo.Byte)
		}
		{
			pbo.Uint8 = uint32(goo.Uint8)
		}
		{
			pbo.Uint16 = uint32(goo.Uint16)
		}
		{
			pbo.Uint32 = goo.Uint32
		}
		{
			pbo.Uint32Fixed = goo.Uint32Fixed
		}
		{
			pbo.Uint64 = goo.Uint64
		}
		{
			pbo.Uint64Fixed = goo.Uint64Fixed
		}
		{
			pbo.Uint = uint64(goo.Uint)
		}
		{
			pbo.Str = goo.Str
		}
		{
			gool := len(goo.Bytes)
			if gool == 0 {
				pbo.Bytes = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Bytes[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.Bytes = pbos
			}
		}
		{
			if !amino.IsEmptyTime(goo.Time) {
				pbo.Time = timestamppb.New(goo.Time)
			}
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.Empty.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.Empty = pbom.(*testspb.EmptyStruct)
		}
	}
	msg = pbo
	return
}
func (goo *PrimitivesStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.PrimitivesStruct = msg.(*testspb.PrimitivesStruct)
	{
		if pbo != nil {
			{
				goo.Int8 = int8(pbo.Int8)
			}
			{
				goo.Int16 = int16(pbo.Int16)
			}
			{
				goo.Int32 = pbo.Int32
			}
			{
				goo.Int32Fixed = pbo.Int32Fixed
			}
			{
				goo.Int64 = pbo.Int64
			}
			{
				goo.Int64Fixed = pbo.Int64Fixed
			}
			{
				goo.Int = int(pbo.Int)
			}
			{
				goo.Byte = uint8(pbo.Byte)
			}
			{
				goo.Uint8 = uint8(pbo.Uint8)
			}
			{
				goo.Uint16 = uint16(pbo.Uint16)
			}
			{
				goo.Uint32 = pbo.Uint32
			}
			{
				goo.Uint32Fixed = pbo.Uint32Fixed
			}
			{
				goo.Uint64 = pbo.Uint64
			}
			{
				goo.Uint64Fixed = pbo.Uint64Fixed
			}
			{
				goo.Uint = uint(pbo.Uint)
			}
			{
				goo.Str = pbo.Str
			}
			{
				var pbol int = 0
				if pbo.Bytes != nil {
					pbol = len(pbo.Bytes)
				}
				if pbol == 0 {
					goo.Bytes = nil
				} else {
					var goos = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Bytes[i]
							{
								goos[i] = uint8(pboe)
							}
						}
					}
					goo.Bytes = goos
				}
			}
			{
				goo.Time = pbo.Time.AsTime()
			}
			{
				if pbo.Empty != nil {
					err = goo.Empty.FromPBMessage(cdc, pbo.Empty)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ PrimitivesStruct) GetTypeURL() (typeURL string) {
	return "/tests.PrimitivesStruct"
}
func (goo PrimitivesStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if goo.Int8 != 0 {
				return false
			}
		}
		{
			if goo.Int16 != 0 {
				return false
			}
		}
		{
			if goo.Int32 != 0 {
				return false
			}
		}
		{
			if goo.Int32Fixed != 0 {
				return false
			}
		}
		{
			if goo.Int64 != 0 {
				return false
			}
		}
		{
			if goo.Int64Fixed != 0 {
				return false
			}
		}
		{
			if goo.Int != 0 {
				return false
			}
		}
		{
			if goo.Byte != 0 {
				return false
			}
		}
		{
			if goo.Uint8 != 0 {
				return false
			}
		}
		{
			if goo.Uint16 != 0 {
				return false
			}
		}
		{
			if goo.Uint32 != 0 {
				return false
			}
		}
		{
			if goo.Uint32Fixed != 0 {
				return false
			}
		}
		{
			if goo.Uint64 != 0 {
				return false
			}
		}
		{
			if goo.Uint64Fixed != 0 {
				return false
			}
		}
		{
			if goo.Uint != 0 {
				return false
			}
		}
		{
			if goo.Str != "" {
				return false
			}
		}
		{
			if len(goo.Bytes) != 0 {
				return false
			}
		}
		{
			if !amino.IsEmptyTime(goo.Time) {
				return false
			}
		}
		{
			e := goo.Empty.IsEmpty()
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo ShortArraysStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.ShortArraysStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.ShortArraysStruct
			msg = pbov
			return
		}
		pbo = new(testspb.ShortArraysStruct)
		{
			gool := len(goo.TimeAr)
			if gool == 0 {
				pbo.TimeAr = nil
			} else {
				var pbos = make([]*timestamppb.Timestamp, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.TimeAr[i]
						{
							if !amino.IsEmptyTime(gooe) {
								pbos[i] = timestamppb.New(gooe)
							}
						}
					}
				}
				pbo.TimeAr = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo *ShortArraysStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.ShortArraysStruct = msg.(*testspb.ShortArraysStruct)
	{
		if pbo != nil {
			{
				var goos = [0]time.Time{}
				for i := 0; i < 0; i += 1 {
					{
						pboe := pbo.TimeAr[i]
						{
							goos[i] = pboe.AsTime()
						}
					}
				}
				goo.TimeAr = goos
			}
		}
	}
	return
}
func (_ ShortArraysStruct) GetTypeURL() (typeURL string) {
	return "/tests.ShortArraysStruct"
}
func (goo ShortArraysStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if len(goo.TimeAr) != 0 {
				return false
			}
		}
	}
	return
}
func (goo ArraysStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.ArraysStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.ArraysStruct
			msg = pbov
			return
		}
		pbo = new(testspb.ArraysStruct)
		{
			gool := len(goo.Int8Ar)
			if gool == 0 {
				pbo.Int8Ar = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int8Ar[i]
						{
							pbos[i] = int32(gooe)
						}
					}
				}
				pbo.Int8Ar = pbos
			}
		}
		{
			gool := len(goo.Int16Ar)
			if gool == 0 {
				pbo.Int16Ar = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int16Ar[i]
						{
							pbos[i] = int32(gooe)
						}
					}
				}
				pbo.Int16Ar = pbos
			}
		}
		{
			gool := len(goo.Int32Ar)
			if gool == 0 {
				pbo.Int32Ar = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32Ar[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Int32Ar = pbos
			}
		}
		{
			gool := len(goo.Int32FixedAr)
			if gool == 0 {
				pbo.Int32FixedAr = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32FixedAr[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Int32FixedAr = pbos
			}
		}
		{
			gool := len(goo.Int64Ar)
			if gool == 0 {
				pbo.Int64Ar = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64Ar[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Int64Ar = pbos
			}
		}
		{
			gool := len(goo.Int64FixedAr)
			if gool == 0 {
				pbo.Int64FixedAr = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64FixedAr[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Int64FixedAr = pbos
			}
		}
		{
			gool := len(goo.IntAr)
			if gool == 0 {
				pbo.IntAr = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.IntAr[i]
						{
							pbos[i] = int64(gooe)
						}
					}
				}
				pbo.IntAr = pbos
			}
		}
		{
			gool := len(goo.ByteAr)
			if gool == 0 {
				pbo.ByteAr = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.ByteAr[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.ByteAr = pbos
			}
		}
		{
			gool := len(goo.Uint8Ar)
			if gool == 0 {
				pbo.Uint8Ar = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint8Ar[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.Uint8Ar = pbos
			}
		}
		{
			gool := len(goo.Uint16Ar)
			if gool == 0 {
				pbo.Uint16Ar = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint16Ar[i]
						{
							pbos[i] = uint32(gooe)
						}
					}
				}
				pbo.Uint16Ar = pbos
			}
		}
		{
			gool := len(goo.Uint32Ar)
			if gool == 0 {
				pbo.Uint32Ar = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32Ar[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Uint32Ar = pbos
			}
		}
		{
			gool := len(goo.Uint32FixedAr)
			if gool == 0 {
				pbo.Uint32FixedAr = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32FixedAr[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Uint32FixedAr = pbos
			}
		}
		{
			gool := len(goo.Uint64Ar)
			if gool == 0 {
				pbo.Uint64Ar = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64Ar[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Uint64Ar = pbos
			}
		}
		{
			gool := len(goo.Uint64FixedAr)
			if gool == 0 {
				pbo.Uint64FixedAr = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64FixedAr[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Uint64FixedAr = pbos
			}
		}
		{
			gool := len(goo.UintAr)
			if gool == 0 {
				pbo.UintAr = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.UintAr[i]
						{
							pbos[i] = uint64(gooe)
						}
					}
				}
				pbo.UintAr = pbos
			}
		}
		{
			gool := len(goo.StrAr)
			if gool == 0 {
				pbo.StrAr = nil
			} else {
				var pbos = make([]string, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.StrAr[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.StrAr = pbos
			}
		}
		{
			gool := len(goo.BytesAr)
			if gool == 0 {
				pbo.BytesAr = nil
			} else {
				var pbos = make([][]byte, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.BytesAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint8, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = byte(gooe)
										}
									}
								}
								pbos[i] = pbos1
							}
						}
					}
				}
				pbo.BytesAr = pbos
			}
		}
		{
			gool := len(goo.TimeAr)
			if gool == 0 {
				pbo.TimeAr = nil
			} else {
				var pbos = make([]*timestamppb.Timestamp, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.TimeAr[i]
						{
							if !amino.IsEmptyTime(gooe) {
								pbos[i] = timestamppb.New(gooe)
							}
						}
					}
				}
				pbo.TimeAr = pbos
			}
		}
		{
			gool := len(goo.EmptyAr)
			if gool == 0 {
				pbo.EmptyAr = nil
			} else {
				var pbos = make([]*testspb.EmptyStruct, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.EmptyAr[i]
						{
							pbom := proto.Message(nil)
							pbom, err = gooe.ToPBMessage(cdc)
							if err != nil {
								return
							}
							pbos[i] = pbom.(*testspb.EmptyStruct)
						}
					}
				}
				pbo.EmptyAr = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo *ArraysStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.ArraysStruct = msg.(*testspb.ArraysStruct)
	{
		if pbo != nil {
			{
				var goos = [4]int8{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Int8Ar[i]
						{
							goos[i] = int8(pboe)
						}
					}
				}
				goo.Int8Ar = goos
			}
			{
				var goos = [4]int16{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Int16Ar[i]
						{
							goos[i] = int16(pboe)
						}
					}
				}
				goo.Int16Ar = goos
			}
			{
				var goos = [4]int32{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Int32Ar[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.Int32Ar = goos
			}
			{
				var goos = [4]int32{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Int32FixedAr[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.Int32FixedAr = goos
			}
			{
				var goos = [4]int64{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Int64Ar[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.Int64Ar = goos
			}
			{
				var goos = [4]int64{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Int64FixedAr[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.Int64FixedAr = goos
			}
			{
				var goos = [4]int{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.IntAr[i]
						{
							goos[i] = int(pboe)
						}
					}
				}
				goo.IntAr = goos
			}
			{
				var goos = [4]uint8{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.ByteAr[i]
						{
							goos[i] = uint8(pboe)
						}
					}
				}
				goo.ByteAr = goos
			}
			{
				var goos = [4]uint8{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Uint8Ar[i]
						{
							goos[i] = uint8(pboe)
						}
					}
				}
				goo.Uint8Ar = goos
			}
			{
				var goos = [4]uint16{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Uint16Ar[i]
						{
							goos[i] = uint16(pboe)
						}
					}
				}
				goo.Uint16Ar = goos
			}
			{
				var goos = [4]uint32{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Uint32Ar[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.Uint32Ar = goos
			}
			{
				var goos = [4]uint32{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Uint32FixedAr[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.Uint32FixedAr = goos
			}
			{
				var goos = [4]uint64{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Uint64Ar[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.Uint64Ar = goos
			}
			{
				var goos = [4]uint64{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.Uint64FixedAr[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.Uint64FixedAr = goos
			}
			{
				var goos = [4]uint{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.UintAr[i]
						{
							goos[i] = uint(pboe)
						}
					}
				}
				goo.UintAr = goos
			}
			{
				var goos = [4]string{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.StrAr[i]
						{
							goos[i] = pboe
						}
					}
				}
				goo.StrAr = goos
			}
			{
				var goos = [4][]uint8{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.BytesAr[i]
						{
							var pbol int = 0
							if pboe != nil {
								pbol = len(pboe)
							}
							if pbol == 0 {
								goos[i] = nil
							} else {
								var goos1 = make([]uint8, pbol)
								for i := 0; i < pbol; i += 1 {
									{
										pboe := pboe[i]
										{
											goos1[i] = uint8(pboe)
										}
									}
								}
								goos[i] = goos1
							}
						}
					}
				}
				goo.BytesAr = goos
			}
			{
				var goos = [4]time.Time{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.TimeAr[i]
						{
							goos[i] = pboe.AsTime()
						}
					}
				}
				goo.TimeAr = goos
			}
			{
				var goos = [4]EmptyStruct{}
				for i := 0; i < 4; i += 1 {
					{
						pboe := pbo.EmptyAr[i]
						{
							if pboe != nil {
								err = goos[i].FromPBMessage(cdc, pboe)
								if err != nil {
									return
								}
							}
						}
					}
				}
				goo.EmptyAr = goos
			}
		}
	}
	return
}
func (_ ArraysStruct) GetTypeURL() (typeURL string) {
	return "/tests.ArraysStruct"
}
func (goo ArraysStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if len(goo.Int8Ar) != 0 {
				return false
			}
		}
		{
			if len(goo.Int16Ar) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32Ar) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32FixedAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64Ar) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64FixedAr) != 0 {
				return false
			}
		}
		{
			if len(goo.IntAr) != 0 {
				return false
			}
		}
		{
			if len(goo.ByteAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint8Ar) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint16Ar) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32Ar) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32FixedAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64Ar) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64FixedAr) != 0 {
				return false
			}
		}
		{
			if len(goo.UintAr) != 0 {
				return false
			}
		}
		{
			if len(goo.StrAr) != 0 {
				return false
			}
		}
		{
			if len(goo.BytesAr) != 0 {
				return false
			}
		}
		{
			if len(goo.TimeAr) != 0 {
				return false
			}
		}
		{
			if len(goo.EmptyAr) != 0 {
				return false
			}
		}
	}
	return
}
func (goo ArraysArraysStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.ArraysArraysStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.ArraysArraysStruct
			msg = pbov
			return
		}
		pbo = new(testspb.ArraysArraysStruct)
		{
			gool := len(goo.Int8ArAr)
			if gool == 0 {
				pbo.Int8ArAr = nil
			} else {
				var pbos = make([]*testspb.Int8List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int8ArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = int32(gooe)
										}
									}
								}
								pbos[i] = &testspb.Int8List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int8ArAr = pbos
			}
		}
		{
			gool := len(goo.Int16ArAr)
			if gool == 0 {
				pbo.Int16ArAr = nil
			} else {
				var pbos = make([]*testspb.Int16List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int16ArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = int32(gooe)
										}
									}
								}
								pbos[i] = &testspb.Int16List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int16ArAr = pbos
			}
		}
		{
			gool := len(goo.Int32ArAr)
			if gool == 0 {
				pbo.Int32ArAr = nil
			} else {
				var pbos = make([]*testspb.Int32List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32ArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Int32List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int32ArAr = pbos
			}
		}
		{
			gool := len(goo.Int32FixedArAr)
			if gool == 0 {
				pbo.Int32FixedArAr = nil
			} else {
				var pbos = make([]*testspb.Fixed32Int32List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32FixedArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Fixed32Int32List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int32FixedArAr = pbos
			}
		}
		{
			gool := len(goo.Int64ArAr)
			if gool == 0 {
				pbo.Int64ArAr = nil
			} else {
				var pbos = make([]*testspb.Int64List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64ArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Int64List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int64ArAr = pbos
			}
		}
		{
			gool := len(goo.Int64FixedArAr)
			if gool == 0 {
				pbo.Int64FixedArAr = nil
			} else {
				var pbos = make([]*testspb.Fixed64Int64List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64FixedArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Fixed64Int64List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int64FixedArAr = pbos
			}
		}
		{
			gool := len(goo.IntArAr)
			if gool == 0 {
				pbo.IntArAr = nil
			} else {
				var pbos = make([]*testspb.IntList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.IntArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = int64(gooe)
										}
									}
								}
								pbos[i] = &testspb.IntList{Value: pbos1}
							}
						}
					}
				}
				pbo.IntArAr = pbos
			}
		}
		{
			gool := len(goo.ByteArAr)
			if gool == 0 {
				pbo.ByteArAr = nil
			} else {
				var pbos = make([][]byte, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.ByteArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint8, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = byte(gooe)
										}
									}
								}
								pbos[i] = pbos1
							}
						}
					}
				}
				pbo.ByteArAr = pbos
			}
		}
		{
			gool := len(goo.Uint8ArAr)
			if gool == 0 {
				pbo.Uint8ArAr = nil
			} else {
				var pbos = make([][]byte, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint8ArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint8, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = byte(gooe)
										}
									}
								}
								pbos[i] = pbos1
							}
						}
					}
				}
				pbo.Uint8ArAr = pbos
			}
		}
		{
			gool := len(goo.Uint16ArAr)
			if gool == 0 {
				pbo.Uint16ArAr = nil
			} else {
				var pbos = make([]*testspb.Uint16List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint16ArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = uint32(gooe)
										}
									}
								}
								pbos[i] = &testspb.Uint16List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint16ArAr = pbos
			}
		}
		{
			gool := len(goo.Uint32ArAr)
			if gool == 0 {
				pbo.Uint32ArAr = nil
			} else {
				var pbos = make([]*testspb.Uint32List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32ArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Uint32List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint32ArAr = pbos
			}
		}
		{
			gool := len(goo.Uint32FixedArAr)
			if gool == 0 {
				pbo.Uint32FixedArAr = nil
			} else {
				var pbos = make([]*testspb.Fixed32Uint32List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32FixedArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Fixed32Uint32List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint32FixedArAr = pbos
			}
		}
		{
			gool := len(goo.Uint64ArAr)
			if gool == 0 {
				pbo.Uint64ArAr = nil
			} else {
				var pbos = make([]*testspb.Uint64List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64ArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Uint64List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint64ArAr = pbos
			}
		}
		{
			gool := len(goo.Uint64FixedArAr)
			if gool == 0 {
				pbo.Uint64FixedArAr = nil
			} else {
				var pbos = make([]*testspb.Fixed64Uint64List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64FixedArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Fixed64Uint64List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint64FixedArAr = pbos
			}
		}
		{
			gool := len(goo.UintArAr)
			if gool == 0 {
				pbo.UintArAr = nil
			} else {
				var pbos = make([]*testspb.UintList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.UintArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = uint64(gooe)
										}
									}
								}
								pbos[i] = &testspb.UintList{Value: pbos1}
							}
						}
					}
				}
				pbo.UintArAr = pbos
			}
		}
		{
			gool := len(goo.StrArAr)
			if gool == 0 {
				pbo.StrArAr = nil
			} else {
				var pbos = make([]*testspb.StringList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.StrArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]string, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.StringList{Value: pbos1}
							}
						}
					}
				}
				pbo.StrArAr = pbos
			}
		}
		{
			gool := len(goo.BytesArAr)
			if gool == 0 {
				pbo.BytesArAr = nil
			} else {
				var pbos = make([]*testspb.BytesList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.BytesArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([][]byte, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											gool2 := len(gooe)
											if gool2 == 0 {
												pbos1[i] = nil
											} else {
												var pbos2 = make([]uint8, gool2)
												for i := 0; i < gool2; i += 1 {
													{
														gooe := gooe[i]
														{
															pbos2[i] = byte(gooe)
														}
													}
												}
												pbos1[i] = pbos2
											}
										}
									}
								}
								pbos[i] = &testspb.BytesList{Value: pbos1}
							}
						}
					}
				}
				pbo.BytesArAr = pbos
			}
		}
		{
			gool := len(goo.TimeArAr)
			if gool == 0 {
				pbo.TimeArAr = nil
			} else {
				var pbos = make([]*testspb.TimeList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.TimeArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]*timestamppb.Timestamp, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											if !amino.IsEmptyTime(gooe) {
												pbos1[i] = timestamppb.New(gooe)
											}
										}
									}
								}
								pbos[i] = &testspb.TimeList{Value: pbos1}
							}
						}
					}
				}
				pbo.TimeArAr = pbos
			}
		}
		{
			gool := len(goo.EmptyArAr)
			if gool == 0 {
				pbo.EmptyArAr = nil
			} else {
				var pbos = make([]*testspb.EmptyStructList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.EmptyArAr[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]*testspb.EmptyStruct, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbom := proto.Message(nil)
											pbom, err = gooe.ToPBMessage(cdc)
											if err != nil {
												return
											}
											pbos1[i] = pbom.(*testspb.EmptyStruct)
										}
									}
								}
								pbos[i] = &testspb.EmptyStructList{Value: pbos1}
							}
						}
					}
				}
				pbo.EmptyArAr = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo *ArraysArraysStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.ArraysArraysStruct = msg.(*testspb.ArraysArraysStruct)
	{
		if pbo != nil {
			{
				var goos = [2][2]int8{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Int8ArAr[i]
						{
							var goos1 = [2]int8{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = int8(pboe)
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Int8ArAr = goos
			}
			{
				var goos = [2][2]int16{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Int16ArAr[i]
						{
							var goos1 = [2]int16{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = int16(pboe)
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Int16ArAr = goos
			}
			{
				var goos = [2][2]int32{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Int32ArAr[i]
						{
							var goos1 = [2]int32{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Int32ArAr = goos
			}
			{
				var goos = [2][2]int32{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Int32FixedArAr[i]
						{
							var goos1 = [2]int32{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Int32FixedArAr = goos
			}
			{
				var goos = [2][2]int64{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Int64ArAr[i]
						{
							var goos1 = [2]int64{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Int64ArAr = goos
			}
			{
				var goos = [2][2]int64{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Int64FixedArAr[i]
						{
							var goos1 = [2]int64{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Int64FixedArAr = goos
			}
			{
				var goos = [2][2]int{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.IntArAr[i]
						{
							var goos1 = [2]int{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = int(pboe)
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.IntArAr = goos
			}
			{
				var goos = [2][2]uint8{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.ByteArAr[i]
						{
							var goos1 = [2]uint8{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe[i]
									{
										goos1[i] = uint8(pboe)
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.ByteArAr = goos
			}
			{
				var goos = [2][2]uint8{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Uint8ArAr[i]
						{
							var goos1 = [2]uint8{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe[i]
									{
										goos1[i] = uint8(pboe)
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Uint8ArAr = goos
			}
			{
				var goos = [2][2]uint16{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Uint16ArAr[i]
						{
							var goos1 = [2]uint16{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = uint16(pboe)
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Uint16ArAr = goos
			}
			{
				var goos = [2][2]uint32{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Uint32ArAr[i]
						{
							var goos1 = [2]uint32{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Uint32ArAr = goos
			}
			{
				var goos = [2][2]uint32{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Uint32FixedArAr[i]
						{
							var goos1 = [2]uint32{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Uint32FixedArAr = goos
			}
			{
				var goos = [2][2]uint64{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Uint64ArAr[i]
						{
							var goos1 = [2]uint64{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Uint64ArAr = goos
			}
			{
				var goos = [2][2]uint64{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.Uint64FixedArAr[i]
						{
							var goos1 = [2]uint64{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.Uint64FixedArAr = goos
			}
			{
				var goos = [2][2]uint{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.UintArAr[i]
						{
							var goos1 = [2]uint{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = uint(pboe)
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.UintArAr = goos
			}
			{
				var goos = [2][2]string{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.StrArAr[i]
						{
							var goos1 = [2]string{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.StrArAr = goos
			}
			{
				var goos = [2][2][]uint8{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.BytesArAr[i]
						{
							var goos1 = [2][]uint8{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										var pbol int = 0
										if pboe != nil {
											pbol = len(pboe)
										}
										if pbol == 0 {
											goos1[i] = nil
										} else {
											var goos2 = make([]uint8, pbol)
											for i := 0; i < pbol; i += 1 {
												{
													pboe := pboe[i]
													{
														goos2[i] = uint8(pboe)
													}
												}
											}
											goos1[i] = goos2
										}
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.BytesArAr = goos
			}
			{
				var goos = [2][2]time.Time{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.TimeArAr[i]
						{
							var goos1 = [2]time.Time{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										goos1[i] = pboe.AsTime()
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.TimeArAr = goos
			}
			{
				var goos = [2][2]EmptyStruct{}
				for i := 0; i < 2; i += 1 {
					{
						pboe := pbo.EmptyArAr[i]
						{
							var goos1 = [2]EmptyStruct{}
							for i := 0; i < 2; i += 1 {
								{
									pboe := pboe.Value[i]
									{
										if pboe != nil {
											err = goos1[i].FromPBMessage(cdc, pboe)
											if err != nil {
												return
											}
										}
									}
								}
							}
							goos[i] = goos1
						}
					}
				}
				goo.EmptyArAr = goos
			}
		}
	}
	return
}
func (_ ArraysArraysStruct) GetTypeURL() (typeURL string) {
	return "/tests.ArraysArraysStruct"
}
func (goo ArraysArraysStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if len(goo.Int8ArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Int16ArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32ArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32FixedArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64ArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64FixedArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.IntArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.ByteArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint8ArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint16ArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32ArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32FixedArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64ArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64FixedArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.UintArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.StrArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.BytesArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.TimeArAr) != 0 {
				return false
			}
		}
		{
			if len(goo.EmptyArAr) != 0 {
				return false
			}
		}
	}
	return
}
func (goo SlicesStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.SlicesStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.SlicesStruct
			msg = pbov
			return
		}
		pbo = new(testspb.SlicesStruct)
		{
			gool := len(goo.Int8Sl)
			if gool == 0 {
				pbo.Int8Sl = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int8Sl[i]
						{
							pbos[i] = int32(gooe)
						}
					}
				}
				pbo.Int8Sl = pbos
			}
		}
		{
			gool := len(goo.Int16Sl)
			if gool == 0 {
				pbo.Int16Sl = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int16Sl[i]
						{
							pbos[i] = int32(gooe)
						}
					}
				}
				pbo.Int16Sl = pbos
			}
		}
		{
			gool := len(goo.Int32Sl)
			if gool == 0 {
				pbo.Int32Sl = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32Sl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Int32Sl = pbos
			}
		}
		{
			gool := len(goo.Int32FixedSl)
			if gool == 0 {
				pbo.Int32FixedSl = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32FixedSl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Int32FixedSl = pbos
			}
		}
		{
			gool := len(goo.Int64Sl)
			if gool == 0 {
				pbo.Int64Sl = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64Sl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Int64Sl = pbos
			}
		}
		{
			gool := len(goo.Int64FixedSl)
			if gool == 0 {
				pbo.Int64FixedSl = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64FixedSl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Int64FixedSl = pbos
			}
		}
		{
			gool := len(goo.IntSl)
			if gool == 0 {
				pbo.IntSl = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.IntSl[i]
						{
							pbos[i] = int64(gooe)
						}
					}
				}
				pbo.IntSl = pbos
			}
		}
		{
			gool := len(goo.ByteSl)
			if gool == 0 {
				pbo.ByteSl = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.ByteSl[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.ByteSl = pbos
			}
		}
		{
			gool := len(goo.Uint8Sl)
			if gool == 0 {
				pbo.Uint8Sl = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint8Sl[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.Uint8Sl = pbos
			}
		}
		{
			gool := len(goo.Uint16Sl)
			if gool == 0 {
				pbo.Uint16Sl = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint16Sl[i]
						{
							pbos[i] = uint32(gooe)
						}
					}
				}
				pbo.Uint16Sl = pbos
			}
		}
		{
			gool := len(goo.Uint32Sl)
			if gool == 0 {
				pbo.Uint32Sl = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32Sl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Uint32Sl = pbos
			}
		}
		{
			gool := len(goo.Uint32FixedSl)
			if gool == 0 {
				pbo.Uint32FixedSl = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32FixedSl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Uint32FixedSl = pbos
			}
		}
		{
			gool := len(goo.Uint64Sl)
			if gool == 0 {
				pbo.Uint64Sl = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64Sl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Uint64Sl = pbos
			}
		}
		{
			gool := len(goo.Uint64FixedSl)
			if gool == 0 {
				pbo.Uint64FixedSl = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64FixedSl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.Uint64FixedSl = pbos
			}
		}
		{
			gool := len(goo.UintSl)
			if gool == 0 {
				pbo.UintSl = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.UintSl[i]
						{
							pbos[i] = uint64(gooe)
						}
					}
				}
				pbo.UintSl = pbos
			}
		}
		{
			gool := len(goo.StrSl)
			if gool == 0 {
				pbo.StrSl = nil
			} else {
				var pbos = make([]string, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.StrSl[i]
						{
							pbos[i] = gooe
						}
					}
				}
				pbo.StrSl = pbos
			}
		}
		{
			gool := len(goo.BytesSl)
			if gool == 0 {
				pbo.BytesSl = nil
			} else {
				var pbos = make([][]byte, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.BytesSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint8, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = byte(gooe)
										}
									}
								}
								pbos[i] = pbos1
							}
						}
					}
				}
				pbo.BytesSl = pbos
			}
		}
		{
			gool := len(goo.TimeSl)
			if gool == 0 {
				pbo.TimeSl = nil
			} else {
				var pbos = make([]*timestamppb.Timestamp, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.TimeSl[i]
						{
							if !amino.IsEmptyTime(gooe) {
								pbos[i] = timestamppb.New(gooe)
							}
						}
					}
				}
				pbo.TimeSl = pbos
			}
		}
		{
			gool := len(goo.EmptySl)
			if gool == 0 {
				pbo.EmptySl = nil
			} else {
				var pbos = make([]*testspb.EmptyStruct, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.EmptySl[i]
						{
							pbom := proto.Message(nil)
							pbom, err = gooe.ToPBMessage(cdc)
							if err != nil {
								return
							}
							pbos[i] = pbom.(*testspb.EmptyStruct)
						}
					}
				}
				pbo.EmptySl = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo *SlicesStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.SlicesStruct = msg.(*testspb.SlicesStruct)
	{
		if pbo != nil {
			{
				var pbol int = 0
				if pbo.Int8Sl != nil {
					pbol = len(pbo.Int8Sl)
				}
				if pbol == 0 {
					goo.Int8Sl = nil
				} else {
					var goos = make([]int8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int8Sl[i]
							{
								goos[i] = int8(pboe)
							}
						}
					}
					goo.Int8Sl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int16Sl != nil {
					pbol = len(pbo.Int16Sl)
				}
				if pbol == 0 {
					goo.Int16Sl = nil
				} else {
					var goos = make([]int16, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int16Sl[i]
							{
								goos[i] = int16(pboe)
							}
						}
					}
					goo.Int16Sl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int32Sl != nil {
					pbol = len(pbo.Int32Sl)
				}
				if pbol == 0 {
					goo.Int32Sl = nil
				} else {
					var goos = make([]int32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int32Sl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.Int32Sl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int32FixedSl != nil {
					pbol = len(pbo.Int32FixedSl)
				}
				if pbol == 0 {
					goo.Int32FixedSl = nil
				} else {
					var goos = make([]int32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int32FixedSl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.Int32FixedSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int64Sl != nil {
					pbol = len(pbo.Int64Sl)
				}
				if pbol == 0 {
					goo.Int64Sl = nil
				} else {
					var goos = make([]int64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int64Sl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.Int64Sl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int64FixedSl != nil {
					pbol = len(pbo.Int64FixedSl)
				}
				if pbol == 0 {
					goo.Int64FixedSl = nil
				} else {
					var goos = make([]int64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int64FixedSl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.Int64FixedSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.IntSl != nil {
					pbol = len(pbo.IntSl)
				}
				if pbol == 0 {
					goo.IntSl = nil
				} else {
					var goos = make([]int, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.IntSl[i]
							{
								goos[i] = int(pboe)
							}
						}
					}
					goo.IntSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.ByteSl != nil {
					pbol = len(pbo.ByteSl)
				}
				if pbol == 0 {
					goo.ByteSl = nil
				} else {
					var goos = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.ByteSl[i]
							{
								goos[i] = uint8(pboe)
							}
						}
					}
					goo.ByteSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint8Sl != nil {
					pbol = len(pbo.Uint8Sl)
				}
				if pbol == 0 {
					goo.Uint8Sl = nil
				} else {
					var goos = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint8Sl[i]
							{
								goos[i] = uint8(pboe)
							}
						}
					}
					goo.Uint8Sl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint16Sl != nil {
					pbol = len(pbo.Uint16Sl)
				}
				if pbol == 0 {
					goo.Uint16Sl = nil
				} else {
					var goos = make([]uint16, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint16Sl[i]
							{
								goos[i] = uint16(pboe)
							}
						}
					}
					goo.Uint16Sl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint32Sl != nil {
					pbol = len(pbo.Uint32Sl)
				}
				if pbol == 0 {
					goo.Uint32Sl = nil
				} else {
					var goos = make([]uint32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint32Sl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.Uint32Sl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint32FixedSl != nil {
					pbol = len(pbo.Uint32FixedSl)
				}
				if pbol == 0 {
					goo.Uint32FixedSl = nil
				} else {
					var goos = make([]uint32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint32FixedSl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.Uint32FixedSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint64Sl != nil {
					pbol = len(pbo.Uint64Sl)
				}
				if pbol == 0 {
					goo.Uint64Sl = nil
				} else {
					var goos = make([]uint64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint64Sl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.Uint64Sl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint64FixedSl != nil {
					pbol = len(pbo.Uint64FixedSl)
				}
				if pbol == 0 {
					goo.Uint64FixedSl = nil
				} else {
					var goos = make([]uint64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint64FixedSl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.Uint64FixedSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.UintSl != nil {
					pbol = len(pbo.UintSl)
				}
				if pbol == 0 {
					goo.UintSl = nil
				} else {
					var goos = make([]uint, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.UintSl[i]
							{
								goos[i] = uint(pboe)
							}
						}
					}
					goo.UintSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.StrSl != nil {
					pbol = len(pbo.StrSl)
				}
				if pbol == 0 {
					goo.StrSl = nil
				} else {
					var goos = make([]string, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.StrSl[i]
							{
								goos[i] = pboe
							}
						}
					}
					goo.StrSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.BytesSl != nil {
					pbol = len(pbo.BytesSl)
				}
				if pbol == 0 {
					goo.BytesSl = nil
				} else {
					var goos = make([][]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.BytesSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint8, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe[i]
											{
												goos1[i] = uint8(pboe)
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.BytesSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.TimeSl != nil {
					pbol = len(pbo.TimeSl)
				}
				if pbol == 0 {
					goo.TimeSl = nil
				} else {
					var goos = make([]time.Time, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.TimeSl[i]
							{
								goos[i] = pboe.AsTime()
							}
						}
					}
					goo.TimeSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.EmptySl != nil {
					pbol = len(pbo.EmptySl)
				}
				if pbol == 0 {
					goo.EmptySl = nil
				} else {
					var goos = make([]EmptyStruct, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.EmptySl[i]
							{
								if pboe != nil {
									err = goos[i].FromPBMessage(cdc, pboe)
									if err != nil {
										return
									}
								}
							}
						}
					}
					goo.EmptySl = goos
				}
			}
		}
	}
	return
}
func (_ SlicesStruct) GetTypeURL() (typeURL string) {
	return "/tests.SlicesStruct"
}
func (goo SlicesStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if len(goo.Int8Sl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int16Sl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32Sl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32FixedSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64Sl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64FixedSl) != 0 {
				return false
			}
		}
		{
			if len(goo.IntSl) != 0 {
				return false
			}
		}
		{
			if len(goo.ByteSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint8Sl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint16Sl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32Sl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32FixedSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64Sl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64FixedSl) != 0 {
				return false
			}
		}
		{
			if len(goo.UintSl) != 0 {
				return false
			}
		}
		{
			if len(goo.StrSl) != 0 {
				return false
			}
		}
		{
			if len(goo.BytesSl) != 0 {
				return false
			}
		}
		{
			if len(goo.TimeSl) != 0 {
				return false
			}
		}
		{
			if len(goo.EmptySl) != 0 {
				return false
			}
		}
	}
	return
}
func (goo SlicesSlicesStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.SlicesSlicesStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.SlicesSlicesStruct
			msg = pbov
			return
		}
		pbo = new(testspb.SlicesSlicesStruct)
		{
			gool := len(goo.Int8SlSl)
			if gool == 0 {
				pbo.Int8SlSl = nil
			} else {
				var pbos = make([]*testspb.Int8List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int8SlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = int32(gooe)
										}
									}
								}
								pbos[i] = &testspb.Int8List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int8SlSl = pbos
			}
		}
		{
			gool := len(goo.Int16SlSl)
			if gool == 0 {
				pbo.Int16SlSl = nil
			} else {
				var pbos = make([]*testspb.Int16List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int16SlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = int32(gooe)
										}
									}
								}
								pbos[i] = &testspb.Int16List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int16SlSl = pbos
			}
		}
		{
			gool := len(goo.Int32SlSl)
			if gool == 0 {
				pbo.Int32SlSl = nil
			} else {
				var pbos = make([]*testspb.Int32List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32SlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Int32List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int32SlSl = pbos
			}
		}
		{
			gool := len(goo.Int32FixedSlSl)
			if gool == 0 {
				pbo.Int32FixedSlSl = nil
			} else {
				var pbos = make([]*testspb.Fixed32Int32List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32FixedSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Fixed32Int32List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int32FixedSlSl = pbos
			}
		}
		{
			gool := len(goo.Int64SlSl)
			if gool == 0 {
				pbo.Int64SlSl = nil
			} else {
				var pbos = make([]*testspb.Int64List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64SlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Int64List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int64SlSl = pbos
			}
		}
		{
			gool := len(goo.Int64FixedSlSl)
			if gool == 0 {
				pbo.Int64FixedSlSl = nil
			} else {
				var pbos = make([]*testspb.Fixed64Int64List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64FixedSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Fixed64Int64List{Value: pbos1}
							}
						}
					}
				}
				pbo.Int64FixedSlSl = pbos
			}
		}
		{
			gool := len(goo.IntSlSl)
			if gool == 0 {
				pbo.IntSlSl = nil
			} else {
				var pbos = make([]*testspb.IntList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.IntSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]int64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = int64(gooe)
										}
									}
								}
								pbos[i] = &testspb.IntList{Value: pbos1}
							}
						}
					}
				}
				pbo.IntSlSl = pbos
			}
		}
		{
			gool := len(goo.ByteSlSl)
			if gool == 0 {
				pbo.ByteSlSl = nil
			} else {
				var pbos = make([][]byte, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.ByteSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint8, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = byte(gooe)
										}
									}
								}
								pbos[i] = pbos1
							}
						}
					}
				}
				pbo.ByteSlSl = pbos
			}
		}
		{
			gool := len(goo.Uint8SlSl)
			if gool == 0 {
				pbo.Uint8SlSl = nil
			} else {
				var pbos = make([][]byte, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint8SlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint8, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = byte(gooe)
										}
									}
								}
								pbos[i] = pbos1
							}
						}
					}
				}
				pbo.Uint8SlSl = pbos
			}
		}
		{
			gool := len(goo.Uint16SlSl)
			if gool == 0 {
				pbo.Uint16SlSl = nil
			} else {
				var pbos = make([]*testspb.Uint16List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint16SlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = uint32(gooe)
										}
									}
								}
								pbos[i] = &testspb.Uint16List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint16SlSl = pbos
			}
		}
		{
			gool := len(goo.Uint32SlSl)
			if gool == 0 {
				pbo.Uint32SlSl = nil
			} else {
				var pbos = make([]*testspb.Uint32List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32SlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Uint32List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint32SlSl = pbos
			}
		}
		{
			gool := len(goo.Uint32FixedSlSl)
			if gool == 0 {
				pbo.Uint32FixedSlSl = nil
			} else {
				var pbos = make([]*testspb.Fixed32Uint32List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32FixedSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint32, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Fixed32Uint32List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint32FixedSlSl = pbos
			}
		}
		{
			gool := len(goo.Uint64SlSl)
			if gool == 0 {
				pbo.Uint64SlSl = nil
			} else {
				var pbos = make([]*testspb.Uint64List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64SlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Uint64List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint64SlSl = pbos
			}
		}
		{
			gool := len(goo.Uint64FixedSlSl)
			if gool == 0 {
				pbo.Uint64FixedSlSl = nil
			} else {
				var pbos = make([]*testspb.Fixed64Uint64List, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64FixedSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.Fixed64Uint64List{Value: pbos1}
							}
						}
					}
				}
				pbo.Uint64FixedSlSl = pbos
			}
		}
		{
			gool := len(goo.UintSlSl)
			if gool == 0 {
				pbo.UintSlSl = nil
			} else {
				var pbos = make([]*testspb.UintList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.UintSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]uint64, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = uint64(gooe)
										}
									}
								}
								pbos[i] = &testspb.UintList{Value: pbos1}
							}
						}
					}
				}
				pbo.UintSlSl = pbos
			}
		}
		{
			gool := len(goo.StrSlSl)
			if gool == 0 {
				pbo.StrSlSl = nil
			} else {
				var pbos = make([]*testspb.StringList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.StrSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]string, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbos1[i] = gooe
										}
									}
								}
								pbos[i] = &testspb.StringList{Value: pbos1}
							}
						}
					}
				}
				pbo.StrSlSl = pbos
			}
		}
		{
			gool := len(goo.BytesSlSl)
			if gool == 0 {
				pbo.BytesSlSl = nil
			} else {
				var pbos = make([]*testspb.BytesList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.BytesSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([][]byte, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											gool2 := len(gooe)
											if gool2 == 0 {
												pbos1[i] = nil
											} else {
												var pbos2 = make([]uint8, gool2)
												for i := 0; i < gool2; i += 1 {
													{
														gooe := gooe[i]
														{
															pbos2[i] = byte(gooe)
														}
													}
												}
												pbos1[i] = pbos2
											}
										}
									}
								}
								pbos[i] = &testspb.BytesList{Value: pbos1}
							}
						}
					}
				}
				pbo.BytesSlSl = pbos
			}
		}
		{
			gool := len(goo.TimeSlSl)
			if gool == 0 {
				pbo.TimeSlSl = nil
			} else {
				var pbos = make([]*testspb.TimeList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.TimeSlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]*timestamppb.Timestamp, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											if !amino.IsEmptyTime(gooe) {
												pbos1[i] = timestamppb.New(gooe)
											}
										}
									}
								}
								pbos[i] = &testspb.TimeList{Value: pbos1}
							}
						}
					}
				}
				pbo.TimeSlSl = pbos
			}
		}
		{
			gool := len(goo.EmptySlSl)
			if gool == 0 {
				pbo.EmptySlSl = nil
			} else {
				var pbos = make([]*testspb.EmptyStructList, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.EmptySlSl[i]
						{
							gool1 := len(gooe)
							if gool1 == 0 {
								pbos[i] = nil
							} else {
								var pbos1 = make([]*testspb.EmptyStruct, gool1)
								for i := 0; i < gool1; i += 1 {
									{
										gooe := gooe[i]
										{
											pbom := proto.Message(nil)
											pbom, err = gooe.ToPBMessage(cdc)
											if err != nil {
												return
											}
											pbos1[i] = pbom.(*testspb.EmptyStruct)
										}
									}
								}
								pbos[i] = &testspb.EmptyStructList{Value: pbos1}
							}
						}
					}
				}
				pbo.EmptySlSl = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo *SlicesSlicesStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.SlicesSlicesStruct = msg.(*testspb.SlicesSlicesStruct)
	{
		if pbo != nil {
			{
				var pbol int = 0
				if pbo.Int8SlSl != nil {
					pbol = len(pbo.Int8SlSl)
				}
				if pbol == 0 {
					goo.Int8SlSl = nil
				} else {
					var goos = make([][]int8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int8SlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]int8, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = int8(pboe)
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Int8SlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int16SlSl != nil {
					pbol = len(pbo.Int16SlSl)
				}
				if pbol == 0 {
					goo.Int16SlSl = nil
				} else {
					var goos = make([][]int16, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int16SlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]int16, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = int16(pboe)
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Int16SlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int32SlSl != nil {
					pbol = len(pbo.Int32SlSl)
				}
				if pbol == 0 {
					goo.Int32SlSl = nil
				} else {
					var goos = make([][]int32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int32SlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]int32, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Int32SlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int32FixedSlSl != nil {
					pbol = len(pbo.Int32FixedSlSl)
				}
				if pbol == 0 {
					goo.Int32FixedSlSl = nil
				} else {
					var goos = make([][]int32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int32FixedSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]int32, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Int32FixedSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int64SlSl != nil {
					pbol = len(pbo.Int64SlSl)
				}
				if pbol == 0 {
					goo.Int64SlSl = nil
				} else {
					var goos = make([][]int64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int64SlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]int64, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Int64SlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int64FixedSlSl != nil {
					pbol = len(pbo.Int64FixedSlSl)
				}
				if pbol == 0 {
					goo.Int64FixedSlSl = nil
				} else {
					var goos = make([][]int64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int64FixedSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]int64, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Int64FixedSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.IntSlSl != nil {
					pbol = len(pbo.IntSlSl)
				}
				if pbol == 0 {
					goo.IntSlSl = nil
				} else {
					var goos = make([][]int, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.IntSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]int, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = int(pboe)
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.IntSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.ByteSlSl != nil {
					pbol = len(pbo.ByteSlSl)
				}
				if pbol == 0 {
					goo.ByteSlSl = nil
				} else {
					var goos = make([][]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.ByteSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint8, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe[i]
											{
												goos1[i] = uint8(pboe)
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.ByteSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint8SlSl != nil {
					pbol = len(pbo.Uint8SlSl)
				}
				if pbol == 0 {
					goo.Uint8SlSl = nil
				} else {
					var goos = make([][]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint8SlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint8, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe[i]
											{
												goos1[i] = uint8(pboe)
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Uint8SlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint16SlSl != nil {
					pbol = len(pbo.Uint16SlSl)
				}
				if pbol == 0 {
					goo.Uint16SlSl = nil
				} else {
					var goos = make([][]uint16, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint16SlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint16, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = uint16(pboe)
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Uint16SlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint32SlSl != nil {
					pbol = len(pbo.Uint32SlSl)
				}
				if pbol == 0 {
					goo.Uint32SlSl = nil
				} else {
					var goos = make([][]uint32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint32SlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint32, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Uint32SlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint32FixedSlSl != nil {
					pbol = len(pbo.Uint32FixedSlSl)
				}
				if pbol == 0 {
					goo.Uint32FixedSlSl = nil
				} else {
					var goos = make([][]uint32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint32FixedSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint32, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Uint32FixedSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint64SlSl != nil {
					pbol = len(pbo.Uint64SlSl)
				}
				if pbol == 0 {
					goo.Uint64SlSl = nil
				} else {
					var goos = make([][]uint64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint64SlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint64, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Uint64SlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint64FixedSlSl != nil {
					pbol = len(pbo.Uint64FixedSlSl)
				}
				if pbol == 0 {
					goo.Uint64FixedSlSl = nil
				} else {
					var goos = make([][]uint64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint64FixedSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint64, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.Uint64FixedSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.UintSlSl != nil {
					pbol = len(pbo.UintSlSl)
				}
				if pbol == 0 {
					goo.UintSlSl = nil
				} else {
					var goos = make([][]uint, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.UintSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]uint, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = uint(pboe)
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.UintSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.StrSlSl != nil {
					pbol = len(pbo.StrSlSl)
				}
				if pbol == 0 {
					goo.StrSlSl = nil
				} else {
					var goos = make([][]string, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.StrSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]string, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.StrSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.BytesSlSl != nil {
					pbol = len(pbo.BytesSlSl)
				}
				if pbol == 0 {
					goo.BytesSlSl = nil
				} else {
					var goos = make([][][]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.BytesSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([][]uint8, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												var pbol2 int = 0
												if pboe != nil {
													pbol2 = len(pboe)
												}
												if pbol2 == 0 {
													goos1[i] = nil
												} else {
													var goos2 = make([]uint8, pbol2)
													for i := 0; i < pbol2; i += 1 {
														{
															pboe := pboe[i]
															{
																goos2[i] = uint8(pboe)
															}
														}
													}
													goos1[i] = goos2
												}
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.BytesSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.TimeSlSl != nil {
					pbol = len(pbo.TimeSlSl)
				}
				if pbol == 0 {
					goo.TimeSlSl = nil
				} else {
					var goos = make([][]time.Time, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.TimeSlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]time.Time, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												goos1[i] = pboe.AsTime()
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.TimeSlSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.EmptySlSl != nil {
					pbol = len(pbo.EmptySlSl)
				}
				if pbol == 0 {
					goo.EmptySlSl = nil
				} else {
					var goos = make([][]EmptyStruct, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.EmptySlSl[i]
							{
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe.Value)
								}
								if pbol1 == 0 {
									goos[i] = nil
								} else {
									var goos1 = make([]EmptyStruct, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe.Value[i]
											{
												if pboe != nil {
													err = goos1[i].FromPBMessage(cdc, pboe)
													if err != nil {
														return
													}
												}
											}
										}
									}
									goos[i] = goos1
								}
							}
						}
					}
					goo.EmptySlSl = goos
				}
			}
		}
	}
	return
}
func (_ SlicesSlicesStruct) GetTypeURL() (typeURL string) {
	return "/tests.SlicesSlicesStruct"
}
func (goo SlicesSlicesStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if len(goo.Int8SlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int16SlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32SlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32FixedSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64SlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64FixedSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.IntSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.ByteSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint8SlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint16SlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32SlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32FixedSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64SlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64FixedSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.UintSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.StrSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.BytesSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.TimeSlSl) != 0 {
				return false
			}
		}
		{
			if len(goo.EmptySlSl) != 0 {
				return false
			}
		}
	}
	return
}
func (goo PointersStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.PointersStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.PointersStruct
			msg = pbov
			return
		}
		pbo = new(testspb.PointersStruct)
		{
			if goo.Int8Pt != nil {
				dgoo := *goo.Int8Pt
				pbo.Int8Pt = int32(dgoo)
			}
		}
		{
			if goo.Int16Pt != nil {
				dgoo := *goo.Int16Pt
				pbo.Int16Pt = int32(dgoo)
			}
		}
		{
			if goo.Int32Pt != nil {
				dgoo := *goo.Int32Pt
				pbo.Int32Pt = dgoo
			}
		}
		{
			if goo.Int32FixedPt != nil {
				dgoo := *goo.Int32FixedPt
				pbo.Int32FixedPt = dgoo
			}
		}
		{
			if goo.Int64Pt != nil {
				dgoo := *goo.Int64Pt
				pbo.Int64Pt = dgoo
			}
		}
		{
			if goo.Int64FixedPt != nil {
				dgoo := *goo.Int64FixedPt
				pbo.Int64FixedPt = dgoo
			}
		}
		{
			if goo.IntPt != nil {
				dgoo := *goo.IntPt
				pbo.IntPt = int64(dgoo)
			}
		}
		{
			if goo.BytePt != nil {
				dgoo := *goo.BytePt
				pbo.BytePt = uint32(dgoo)
			}
		}
		{
			if goo.Uint8Pt != nil {
				dgoo := *goo.Uint8Pt
				pbo.Uint8Pt = uint32(dgoo)
			}
		}
		{
			if goo.Uint16Pt != nil {
				dgoo := *goo.Uint16Pt
				pbo.Uint16Pt = uint32(dgoo)
			}
		}
		{
			if goo.Uint32Pt != nil {
				dgoo := *goo.Uint32Pt
				pbo.Uint32Pt = dgoo
			}
		}
		{
			if goo.Uint32FixedPt != nil {
				dgoo := *goo.Uint32FixedPt
				pbo.Uint32FixedPt = dgoo
			}
		}
		{
			if goo.Uint64Pt != nil {
				dgoo := *goo.Uint64Pt
				pbo.Uint64Pt = dgoo
			}
		}
		{
			if goo.Uint64FixedPt != nil {
				dgoo := *goo.Uint64FixedPt
				pbo.Uint64FixedPt = dgoo
			}
		}
		{
			if goo.UintPt != nil {
				dgoo := *goo.UintPt
				pbo.UintPt = uint64(dgoo)
			}
		}
		{
			if goo.StrPt != nil {
				dgoo := *goo.StrPt
				pbo.StrPt = dgoo
			}
		}
		{
			if goo.BytesPt != nil {
				dgoo := *goo.BytesPt
				gool := len(dgoo)
				if gool == 0 {
					pbo.BytesPt = nil
				} else {
					var pbos = make([]uint8, gool)
					for i := 0; i < gool; i += 1 {
						{
							gooe := dgoo[i]
							{
								pbos[i] = byte(gooe)
							}
						}
					}
					pbo.BytesPt = pbos
				}
			}
		}
		{
			if goo.TimePt != nil {
				dgoo := *goo.TimePt
				pbo.TimePt = timestamppb.New(dgoo)
			}
		}
		{
			if goo.EmptyPt != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.EmptyPt.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.EmptyPt = pbom.(*testspb.EmptyStruct)
				if pbo.EmptyPt == nil {
					pbo.EmptyPt = new(testspb.EmptyStruct)
				}
			}
		}
	}
	msg = pbo
	return
}
func (goo *PointersStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.PointersStruct = msg.(*testspb.PointersStruct)
	{
		if pbo != nil {
			{
				goo.Int8Pt = new(int8)
				*goo.Int8Pt = int8(pbo.Int8Pt)
			}
			{
				goo.Int16Pt = new(int16)
				*goo.Int16Pt = int16(pbo.Int16Pt)
			}
			{
				goo.Int32Pt = new(int32)
				*goo.Int32Pt = pbo.Int32Pt
			}
			{
				goo.Int32FixedPt = new(int32)
				*goo.Int32FixedPt = pbo.Int32FixedPt
			}
			{
				goo.Int64Pt = new(int64)
				*goo.Int64Pt = pbo.Int64Pt
			}
			{
				goo.Int64FixedPt = new(int64)
				*goo.Int64FixedPt = pbo.Int64FixedPt
			}
			{
				goo.IntPt = new(int)
				*goo.IntPt = int(pbo.IntPt)
			}
			{
				goo.BytePt = new(uint8)
				*goo.BytePt = uint8(pbo.BytePt)
			}
			{
				goo.Uint8Pt = new(uint8)
				*goo.Uint8Pt = uint8(pbo.Uint8Pt)
			}
			{
				goo.Uint16Pt = new(uint16)
				*goo.Uint16Pt = uint16(pbo.Uint16Pt)
			}
			{
				goo.Uint32Pt = new(uint32)
				*goo.Uint32Pt = pbo.Uint32Pt
			}
			{
				goo.Uint32FixedPt = new(uint32)
				*goo.Uint32FixedPt = pbo.Uint32FixedPt
			}
			{
				goo.Uint64Pt = new(uint64)
				*goo.Uint64Pt = pbo.Uint64Pt
			}
			{
				goo.Uint64FixedPt = new(uint64)
				*goo.Uint64FixedPt = pbo.Uint64FixedPt
			}
			{
				goo.UintPt = new(uint)
				*goo.UintPt = uint(pbo.UintPt)
			}
			{
				goo.StrPt = new(string)
				*goo.StrPt = pbo.StrPt
			}
			{
				goo.BytesPt = new([]uint8)
				var pbol int = 0
				if pbo.BytesPt != nil {
					pbol = len(pbo.BytesPt)
				}
				if pbol == 0 {
					*goo.BytesPt = nil
				} else {
					var goos = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.BytesPt[i]
							{
								goos[i] = uint8(pboe)
							}
						}
					}
					*goo.BytesPt = goos
				}
			}
			{
				goo.TimePt = new(time.Time)
				*goo.TimePt = pbo.TimePt.AsTime()
			}
			{
				if pbo.EmptyPt != nil {
					goo.EmptyPt = new(EmptyStruct)
					err = (*goo.EmptyPt).FromPBMessage(cdc, pbo.EmptyPt)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ PointersStruct) GetTypeURL() (typeURL string) {
	return "/tests.PointersStruct"
}
func (goo PointersStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if goo.Int8Pt != nil {
				dgoo := *goo.Int8Pt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Int16Pt != nil {
				dgoo := *goo.Int16Pt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Int32Pt != nil {
				dgoo := *goo.Int32Pt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Int32FixedPt != nil {
				dgoo := *goo.Int32FixedPt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Int64Pt != nil {
				dgoo := *goo.Int64Pt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Int64FixedPt != nil {
				dgoo := *goo.Int64FixedPt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.IntPt != nil {
				dgoo := *goo.IntPt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.BytePt != nil {
				dgoo := *goo.BytePt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Uint8Pt != nil {
				dgoo := *goo.Uint8Pt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Uint16Pt != nil {
				dgoo := *goo.Uint16Pt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Uint32Pt != nil {
				dgoo := *goo.Uint32Pt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Uint32FixedPt != nil {
				dgoo := *goo.Uint32FixedPt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Uint64Pt != nil {
				dgoo := *goo.Uint64Pt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.Uint64FixedPt != nil {
				dgoo := *goo.Uint64FixedPt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.UintPt != nil {
				dgoo := *goo.UintPt
				if dgoo != 0 {
					return false
				}
			}
		}
		{
			if goo.StrPt != nil {
				dgoo := *goo.StrPt
				if dgoo != "" {
					return false
				}
			}
		}
		{
			if goo.BytesPt != nil {
				dgoo := *goo.BytesPt
				if len(dgoo) != 0 {
					return false
				}
			}
		}
		{
			if goo.TimePt != nil {
				return false
			}
		}
		{
			if goo.EmptyPt != nil {
				return false
			}
		}
	}
	return
}
func (goo PointerSlicesStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.PointerSlicesStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.PointerSlicesStruct
			msg = pbov
			return
		}
		pbo = new(testspb.PointerSlicesStruct)
		{
			gool := len(goo.Int8PtSl)
			if gool == 0 {
				pbo.Int8PtSl = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int8PtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = int32(dgoo)
							}
						}
					}
				}
				pbo.Int8PtSl = pbos
			}
		}
		{
			gool := len(goo.Int16PtSl)
			if gool == 0 {
				pbo.Int16PtSl = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int16PtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = int32(dgoo)
							}
						}
					}
				}
				pbo.Int16PtSl = pbos
			}
		}
		{
			gool := len(goo.Int32PtSl)
			if gool == 0 {
				pbo.Int32PtSl = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32PtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.Int32PtSl = pbos
			}
		}
		{
			gool := len(goo.Int32FixedPtSl)
			if gool == 0 {
				pbo.Int32FixedPtSl = nil
			} else {
				var pbos = make([]int32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int32FixedPtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.Int32FixedPtSl = pbos
			}
		}
		{
			gool := len(goo.Int64PtSl)
			if gool == 0 {
				pbo.Int64PtSl = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64PtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.Int64PtSl = pbos
			}
		}
		{
			gool := len(goo.Int64FixedPtSl)
			if gool == 0 {
				pbo.Int64FixedPtSl = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Int64FixedPtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.Int64FixedPtSl = pbos
			}
		}
		{
			gool := len(goo.IntPtSl)
			if gool == 0 {
				pbo.IntPtSl = nil
			} else {
				var pbos = make([]int64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.IntPtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = int64(dgoo)
							}
						}
					}
				}
				pbo.IntPtSl = pbos
			}
		}
		{
			gool := len(goo.BytePtSl)
			if gool == 0 {
				pbo.BytePtSl = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.BytePtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = byte(dgoo)
							}
						}
					}
				}
				pbo.BytePtSl = pbos
			}
		}
		{
			gool := len(goo.Uint8PtSl)
			if gool == 0 {
				pbo.Uint8PtSl = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint8PtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = byte(dgoo)
							}
						}
					}
				}
				pbo.Uint8PtSl = pbos
			}
		}
		{
			gool := len(goo.Uint16PtSl)
			if gool == 0 {
				pbo.Uint16PtSl = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint16PtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = uint32(dgoo)
							}
						}
					}
				}
				pbo.Uint16PtSl = pbos
			}
		}
		{
			gool := len(goo.Uint32PtSl)
			if gool == 0 {
				pbo.Uint32PtSl = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32PtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.Uint32PtSl = pbos
			}
		}
		{
			gool := len(goo.Uint32FixedPtSl)
			if gool == 0 {
				pbo.Uint32FixedPtSl = nil
			} else {
				var pbos = make([]uint32, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint32FixedPtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.Uint32FixedPtSl = pbos
			}
		}
		{
			gool := len(goo.Uint64PtSl)
			if gool == 0 {
				pbo.Uint64PtSl = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64PtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.Uint64PtSl = pbos
			}
		}
		{
			gool := len(goo.Uint64FixedPtSl)
			if gool == 0 {
				pbo.Uint64FixedPtSl = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Uint64FixedPtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.Uint64FixedPtSl = pbos
			}
		}
		{
			gool := len(goo.UintPtSl)
			if gool == 0 {
				pbo.UintPtSl = nil
			} else {
				var pbos = make([]uint64, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.UintPtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = uint64(dgoo)
							}
						}
					}
				}
				pbo.UintPtSl = pbos
			}
		}
		{
			gool := len(goo.StrPtSl)
			if gool == 0 {
				pbo.StrPtSl = nil
			} else {
				var pbos = make([]string, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.StrPtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = dgoo
							}
						}
					}
				}
				pbo.StrPtSl = pbos
			}
		}
		{
			gool := len(goo.BytesPtSl)
			if gool == 0 {
				pbo.BytesPtSl = nil
			} else {
				var pbos = make([][]byte, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.BytesPtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								gool1 := len(dgoo)
								if gool1 == 0 {
									pbos[i] = nil
								} else {
									var pbos1 = make([]uint8, gool1)
									for i := 0; i < gool1; i += 1 {
										{
											gooe := dgoo[i]
											{
												pbos1[i] = byte(gooe)
											}
										}
									}
									pbos[i] = pbos1
								}
							}
						}
					}
				}
				pbo.BytesPtSl = pbos
			}
		}
		{
			gool := len(goo.TimePtSl)
			if gool == 0 {
				pbo.TimePtSl = nil
			} else {
				var pbos = make([]*timestamppb.Timestamp, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.TimePtSl[i]
						{
							if gooe != nil {
								dgoo := *gooe
								pbos[i] = timestamppb.New(dgoo)
							}
						}
					}
				}
				pbo.TimePtSl = pbos
			}
		}
		{
			gool := len(goo.EmptyPtSl)
			if gool == 0 {
				pbo.EmptyPtSl = nil
			} else {
				var pbos = make([]*testspb.EmptyStruct, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.EmptyPtSl[i]
						{
							if gooe != nil {
								pbom := proto.Message(nil)
								pbom, err = gooe.ToPBMessage(cdc)
								if err != nil {
									return
								}
								pbos[i] = pbom.(*testspb.EmptyStruct)
								if pbos[i] == nil {
									pbos[i] = new(testspb.EmptyStruct)
								}
							}
						}
					}
				}
				pbo.EmptyPtSl = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo *PointerSlicesStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.PointerSlicesStruct = msg.(*testspb.PointerSlicesStruct)
	{
		if pbo != nil {
			{
				var pbol int = 0
				if pbo.Int8PtSl != nil {
					pbol = len(pbo.Int8PtSl)
				}
				if pbol == 0 {
					goo.Int8PtSl = nil
				} else {
					var goos = make([]*int8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int8PtSl[i]
							{
								goos[i] = new(int8)
								*goos[i] = int8(pboe)
							}
						}
					}
					goo.Int8PtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int16PtSl != nil {
					pbol = len(pbo.Int16PtSl)
				}
				if pbol == 0 {
					goo.Int16PtSl = nil
				} else {
					var goos = make([]*int16, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int16PtSl[i]
							{
								goos[i] = new(int16)
								*goos[i] = int16(pboe)
							}
						}
					}
					goo.Int16PtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int32PtSl != nil {
					pbol = len(pbo.Int32PtSl)
				}
				if pbol == 0 {
					goo.Int32PtSl = nil
				} else {
					var goos = make([]*int32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int32PtSl[i]
							{
								goos[i] = new(int32)
								*goos[i] = pboe
							}
						}
					}
					goo.Int32PtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int32FixedPtSl != nil {
					pbol = len(pbo.Int32FixedPtSl)
				}
				if pbol == 0 {
					goo.Int32FixedPtSl = nil
				} else {
					var goos = make([]*int32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int32FixedPtSl[i]
							{
								goos[i] = new(int32)
								*goos[i] = pboe
							}
						}
					}
					goo.Int32FixedPtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int64PtSl != nil {
					pbol = len(pbo.Int64PtSl)
				}
				if pbol == 0 {
					goo.Int64PtSl = nil
				} else {
					var goos = make([]*int64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int64PtSl[i]
							{
								goos[i] = new(int64)
								*goos[i] = pboe
							}
						}
					}
					goo.Int64PtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Int64FixedPtSl != nil {
					pbol = len(pbo.Int64FixedPtSl)
				}
				if pbol == 0 {
					goo.Int64FixedPtSl = nil
				} else {
					var goos = make([]*int64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Int64FixedPtSl[i]
							{
								goos[i] = new(int64)
								*goos[i] = pboe
							}
						}
					}
					goo.Int64FixedPtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.IntPtSl != nil {
					pbol = len(pbo.IntPtSl)
				}
				if pbol == 0 {
					goo.IntPtSl = nil
				} else {
					var goos = make([]*int, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.IntPtSl[i]
							{
								goos[i] = new(int)
								*goos[i] = int(pboe)
							}
						}
					}
					goo.IntPtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.BytePtSl != nil {
					pbol = len(pbo.BytePtSl)
				}
				if pbol == 0 {
					goo.BytePtSl = nil
				} else {
					var goos = make([]*uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.BytePtSl[i]
							{
								goos[i] = new(uint8)
								*goos[i] = uint8(pboe)
							}
						}
					}
					goo.BytePtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint8PtSl != nil {
					pbol = len(pbo.Uint8PtSl)
				}
				if pbol == 0 {
					goo.Uint8PtSl = nil
				} else {
					var goos = make([]*uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint8PtSl[i]
							{
								goos[i] = new(uint8)
								*goos[i] = uint8(pboe)
							}
						}
					}
					goo.Uint8PtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint16PtSl != nil {
					pbol = len(pbo.Uint16PtSl)
				}
				if pbol == 0 {
					goo.Uint16PtSl = nil
				} else {
					var goos = make([]*uint16, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint16PtSl[i]
							{
								goos[i] = new(uint16)
								*goos[i] = uint16(pboe)
							}
						}
					}
					goo.Uint16PtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint32PtSl != nil {
					pbol = len(pbo.Uint32PtSl)
				}
				if pbol == 0 {
					goo.Uint32PtSl = nil
				} else {
					var goos = make([]*uint32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint32PtSl[i]
							{
								goos[i] = new(uint32)
								*goos[i] = pboe
							}
						}
					}
					goo.Uint32PtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint32FixedPtSl != nil {
					pbol = len(pbo.Uint32FixedPtSl)
				}
				if pbol == 0 {
					goo.Uint32FixedPtSl = nil
				} else {
					var goos = make([]*uint32, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint32FixedPtSl[i]
							{
								goos[i] = new(uint32)
								*goos[i] = pboe
							}
						}
					}
					goo.Uint32FixedPtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint64PtSl != nil {
					pbol = len(pbo.Uint64PtSl)
				}
				if pbol == 0 {
					goo.Uint64PtSl = nil
				} else {
					var goos = make([]*uint64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint64PtSl[i]
							{
								goos[i] = new(uint64)
								*goos[i] = pboe
							}
						}
					}
					goo.Uint64PtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.Uint64FixedPtSl != nil {
					pbol = len(pbo.Uint64FixedPtSl)
				}
				if pbol == 0 {
					goo.Uint64FixedPtSl = nil
				} else {
					var goos = make([]*uint64, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Uint64FixedPtSl[i]
							{
								goos[i] = new(uint64)
								*goos[i] = pboe
							}
						}
					}
					goo.Uint64FixedPtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.UintPtSl != nil {
					pbol = len(pbo.UintPtSl)
				}
				if pbol == 0 {
					goo.UintPtSl = nil
				} else {
					var goos = make([]*uint, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.UintPtSl[i]
							{
								goos[i] = new(uint)
								*goos[i] = uint(pboe)
							}
						}
					}
					goo.UintPtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.StrPtSl != nil {
					pbol = len(pbo.StrPtSl)
				}
				if pbol == 0 {
					goo.StrPtSl = nil
				} else {
					var goos = make([]*string, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.StrPtSl[i]
							{
								goos[i] = new(string)
								*goos[i] = pboe
							}
						}
					}
					goo.StrPtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.BytesPtSl != nil {
					pbol = len(pbo.BytesPtSl)
				}
				if pbol == 0 {
					goo.BytesPtSl = nil
				} else {
					var goos = make([]*[]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.BytesPtSl[i]
							{
								goos[i] = new([]uint8)
								var pbol1 int = 0
								if pboe != nil {
									pbol1 = len(pboe)
								}
								if pbol1 == 0 {
									*goos[i] = nil
								} else {
									var goos1 = make([]uint8, pbol1)
									for i := 0; i < pbol1; i += 1 {
										{
											pboe := pboe[i]
											{
												goos1[i] = uint8(pboe)
											}
										}
									}
									*goos[i] = goos1
								}
							}
						}
					}
					goo.BytesPtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.TimePtSl != nil {
					pbol = len(pbo.TimePtSl)
				}
				if pbol == 0 {
					goo.TimePtSl = nil
				} else {
					var goos = make([]*time.Time, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.TimePtSl[i]
							{
								goos[i] = new(time.Time)
								*goos[i] = pboe.AsTime()
							}
						}
					}
					goo.TimePtSl = goos
				}
			}
			{
				var pbol int = 0
				if pbo.EmptyPtSl != nil {
					pbol = len(pbo.EmptyPtSl)
				}
				if pbol == 0 {
					goo.EmptyPtSl = nil
				} else {
					var goos = make([]*EmptyStruct, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.EmptyPtSl[i]
							{
								if pboe != nil {
									goos[i] = new(EmptyStruct)
									err = (*goos[i]).FromPBMessage(cdc, pboe)
									if err != nil {
										return
									}
								}
							}
						}
					}
					goo.EmptyPtSl = goos
				}
			}
		}
	}
	return
}
func (_ PointerSlicesStruct) GetTypeURL() (typeURL string) {
	return "/tests.PointerSlicesStruct"
}
func (goo PointerSlicesStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if len(goo.Int8PtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int16PtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32PtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int32FixedPtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64PtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Int64FixedPtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.IntPtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.BytePtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint8PtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint16PtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32PtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint32FixedPtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64PtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.Uint64FixedPtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.UintPtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.StrPtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.BytesPtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.TimePtSl) != 0 {
				return false
			}
		}
		{
			if len(goo.EmptyPtSl) != 0 {
				return false
			}
		}
	}
	return
}
func (goo ComplexSt) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.ComplexSt
	{
		if goo.IsEmpty() {
			var pbov *testspb.ComplexSt
			msg = pbov
			return
		}
		pbo = new(testspb.ComplexSt)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.PrField.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.PrField = pbom.(*testspb.PrimitivesStruct)
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ArField.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ArField = pbom.(*testspb.ArraysStruct)
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.SlField.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.SlField = pbom.(*testspb.SlicesStruct)
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.PtField.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.PtField = pbom.(*testspb.PointersStruct)
		}
	}
	msg = pbo
	return
}
func (goo *ComplexSt) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.ComplexSt = msg.(*testspb.ComplexSt)
	{
		if pbo != nil {
			{
				if pbo.PrField != nil {
					err = goo.PrField.FromPBMessage(cdc, pbo.PrField)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.ArField != nil {
					err = goo.ArField.FromPBMessage(cdc, pbo.ArField)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.SlField != nil {
					err = goo.SlField.FromPBMessage(cdc, pbo.SlField)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.PtField != nil {
					err = goo.PtField.FromPBMessage(cdc, pbo.PtField)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ ComplexSt) GetTypeURL() (typeURL string) {
	return "/tests.ComplexSt"
}
func (goo ComplexSt) IsEmpty() (empty bool) {
	{
		empty = true
		{
			e := goo.PrField.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			e := goo.ArField.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			e := goo.SlField.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			e := goo.PtField.IsEmpty()
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo EmbeddedSt1) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.EmbeddedSt1
	{
		if goo.IsEmpty() {
			var pbov *testspb.EmbeddedSt1
			msg = pbov
			return
		}
		pbo = new(testspb.EmbeddedSt1)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.PrimitivesStruct.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.PrimitivesStruct = pbom.(*testspb.PrimitivesStruct)
		}
	}
	msg = pbo
	return
}
func (goo *EmbeddedSt1) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.EmbeddedSt1 = msg.(*testspb.EmbeddedSt1)
	{
		if pbo != nil {
			{
				if pbo.PrimitivesStruct != nil {
					err = goo.PrimitivesStruct.FromPBMessage(cdc, pbo.PrimitivesStruct)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ EmbeddedSt1) GetTypeURL() (typeURL string) {
	return "/tests.EmbeddedSt1"
}
func (goo EmbeddedSt1) IsEmpty() (empty bool) {
	{
		empty = true
		{
			e := goo.PrimitivesStruct.IsEmpty()
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo EmbeddedSt2) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.EmbeddedSt2
	{
		if goo.IsEmpty() {
			var pbov *testspb.EmbeddedSt2
			msg = pbov
			return
		}
		pbo = new(testspb.EmbeddedSt2)
		{
			pbom := proto.Message(nil)
			pbom, err = goo.PrimitivesStruct.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.PrimitivesStruct = pbom.(*testspb.PrimitivesStruct)
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ArraysStruct.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ArraysStruct = pbom.(*testspb.ArraysStruct)
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.SlicesStruct.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.SlicesStruct = pbom.(*testspb.SlicesStruct)
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.PointersStruct.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.PointersStruct = pbom.(*testspb.PointersStruct)
		}
	}
	msg = pbo
	return
}
func (goo *EmbeddedSt2) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.EmbeddedSt2 = msg.(*testspb.EmbeddedSt2)
	{
		if pbo != nil {
			{
				if pbo.PrimitivesStruct != nil {
					err = goo.PrimitivesStruct.FromPBMessage(cdc, pbo.PrimitivesStruct)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.ArraysStruct != nil {
					err = goo.ArraysStruct.FromPBMessage(cdc, pbo.ArraysStruct)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.SlicesStruct != nil {
					err = goo.SlicesStruct.FromPBMessage(cdc, pbo.SlicesStruct)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.PointersStruct != nil {
					err = goo.PointersStruct.FromPBMessage(cdc, pbo.PointersStruct)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ EmbeddedSt2) GetTypeURL() (typeURL string) {
	return "/tests.EmbeddedSt2"
}
func (goo EmbeddedSt2) IsEmpty() (empty bool) {
	{
		empty = true
		{
			e := goo.PrimitivesStruct.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			e := goo.ArraysStruct.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			e := goo.SlicesStruct.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			e := goo.PointersStruct.IsEmpty()
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo EmbeddedSt3) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.EmbeddedSt3
	{
		if goo.IsEmpty() {
			var pbov *testspb.EmbeddedSt3
			msg = pbov
			return
		}
		pbo = new(testspb.EmbeddedSt3)
		{
			if goo.PrimitivesStruct != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.PrimitivesStruct.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.PrimitivesStruct = pbom.(*testspb.PrimitivesStruct)
				if pbo.PrimitivesStruct == nil {
					pbo.PrimitivesStruct = new(testspb.PrimitivesStruct)
				}
			}
		}
		{
			if goo.ArraysStruct != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.ArraysStruct.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.ArraysStruct = pbom.(*testspb.ArraysStruct)
				if pbo.ArraysStruct == nil {
					pbo.ArraysStruct = new(testspb.ArraysStruct)
				}
			}
		}
		{
			if goo.SlicesStruct != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.SlicesStruct.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.SlicesStruct = pbom.(*testspb.SlicesStruct)
				if pbo.SlicesStruct == nil {
					pbo.SlicesStruct = new(testspb.SlicesStruct)
				}
			}
		}
		{
			if goo.PointersStruct != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.PointersStruct.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.PointersStruct = pbom.(*testspb.PointersStruct)
				if pbo.PointersStruct == nil {
					pbo.PointersStruct = new(testspb.PointersStruct)
				}
			}
		}
		{
			if goo.EmptyStruct != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.EmptyStruct.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.EmptyStruct = pbom.(*testspb.EmptyStruct)
				if pbo.EmptyStruct == nil {
					pbo.EmptyStruct = new(testspb.EmptyStruct)
				}
			}
		}
	}
	msg = pbo
	return
}
func (goo *EmbeddedSt3) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.EmbeddedSt3 = msg.(*testspb.EmbeddedSt3)
	{
		if pbo != nil {
			{
				if pbo.PrimitivesStruct != nil {
					goo.PrimitivesStruct = new(PrimitivesStruct)
					err = (*goo.PrimitivesStruct).FromPBMessage(cdc, pbo.PrimitivesStruct)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.ArraysStruct != nil {
					goo.ArraysStruct = new(ArraysStruct)
					err = (*goo.ArraysStruct).FromPBMessage(cdc, pbo.ArraysStruct)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.SlicesStruct != nil {
					goo.SlicesStruct = new(SlicesStruct)
					err = (*goo.SlicesStruct).FromPBMessage(cdc, pbo.SlicesStruct)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.PointersStruct != nil {
					goo.PointersStruct = new(PointersStruct)
					err = (*goo.PointersStruct).FromPBMessage(cdc, pbo.PointersStruct)
					if err != nil {
						return
					}
				}
			}
			{
				if pbo.EmptyStruct != nil {
					goo.EmptyStruct = new(EmptyStruct)
					err = (*goo.EmptyStruct).FromPBMessage(cdc, pbo.EmptyStruct)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ EmbeddedSt3) GetTypeURL() (typeURL string) {
	return "/tests.EmbeddedSt3"
}
func (goo EmbeddedSt3) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if goo.PrimitivesStruct != nil {
				return false
			}
		}
		{
			if goo.ArraysStruct != nil {
				return false
			}
		}
		{
			if goo.SlicesStruct != nil {
				return false
			}
		}
		{
			if goo.PointersStruct != nil {
				return false
			}
		}
		{
			if goo.EmptyStruct != nil {
				return false
			}
		}
	}
	return
}
func (goo EmbeddedSt4) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.EmbeddedSt4
	{
		if goo.IsEmpty() {
			var pbov *testspb.EmbeddedSt4
			msg = pbov
			return
		}
		pbo = new(testspb.EmbeddedSt4)
		{
			pbo.Foo1 = int64(goo.Foo1)
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.PrimitivesStruct.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.PrimitivesStruct = pbom.(*testspb.PrimitivesStruct)
		}
		{
			pbo.Foo2 = goo.Foo2
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.ArraysStructField.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.ArraysStructField = pbom.(*testspb.ArraysStruct)
		}
		{
			gool := len(goo.Foo3)
			if gool == 0 {
				pbo.Foo3 = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Foo3[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.Foo3 = pbos
			}
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.SlicesStruct.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.SlicesStruct = pbom.(*testspb.SlicesStruct)
		}
		{
			pbo.Foo4 = goo.Foo4
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.PointersStructField.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.PointersStructField = pbom.(*testspb.PointersStruct)
		}
		{
			pbo.Foo5 = uint64(goo.Foo5)
		}
	}
	msg = pbo
	return
}
func (goo *EmbeddedSt4) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.EmbeddedSt4 = msg.(*testspb.EmbeddedSt4)
	{
		if pbo != nil {
			{
				goo.Foo1 = int(pbo.Foo1)
			}
			{
				if pbo.PrimitivesStruct != nil {
					err = goo.PrimitivesStruct.FromPBMessage(cdc, pbo.PrimitivesStruct)
					if err != nil {
						return
					}
				}
			}
			{
				goo.Foo2 = pbo.Foo2
			}
			{
				if pbo.ArraysStructField != nil {
					err = goo.ArraysStructField.FromPBMessage(cdc, pbo.ArraysStructField)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Foo3 != nil {
					pbol = len(pbo.Foo3)
				}
				if pbol == 0 {
					goo.Foo3 = nil
				} else {
					var goos = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Foo3[i]
							{
								goos[i] = uint8(pboe)
							}
						}
					}
					goo.Foo3 = goos
				}
			}
			{
				if pbo.SlicesStruct != nil {
					err = goo.SlicesStruct.FromPBMessage(cdc, pbo.SlicesStruct)
					if err != nil {
						return
					}
				}
			}
			{
				goo.Foo4 = pbo.Foo4
			}
			{
				if pbo.PointersStructField != nil {
					err = goo.PointersStructField.FromPBMessage(cdc, pbo.PointersStructField)
					if err != nil {
						return
					}
				}
			}
			{
				goo.Foo5 = uint(pbo.Foo5)
			}
		}
	}
	return
}
func (_ EmbeddedSt4) GetTypeURL() (typeURL string) {
	return "/tests.EmbeddedSt4"
}
func (goo EmbeddedSt4) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if goo.Foo1 != 0 {
				return false
			}
		}
		{
			e := goo.PrimitivesStruct.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			if goo.Foo2 != "" {
				return false
			}
		}
		{
			e := goo.ArraysStructField.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			if len(goo.Foo3) != 0 {
				return false
			}
		}
		{
			e := goo.SlicesStruct.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			if goo.Foo4 != false {
				return false
			}
		}
		{
			e := goo.PointersStructField.IsEmpty()
			if e == false {
				return false
			}
		}
		{
			if goo.Foo5 != 0 {
				return false
			}
		}
	}
	return
}
func (goo EmbeddedSt5) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.EmbeddedSt5
	{
		if goo.IsEmpty() {
			var pbov *testspb.EmbeddedSt5
			msg = pbov
			return
		}
		pbo = new(testspb.EmbeddedSt5)
		{
			pbo.Foo1 = int64(goo.Foo1)
		}
		{
			if goo.PrimitivesStruct != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.PrimitivesStruct.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.PrimitivesStruct = pbom.(*testspb.PrimitivesStruct)
				if pbo.PrimitivesStruct == nil {
					pbo.PrimitivesStruct = new(testspb.PrimitivesStruct)
				}
			}
		}
		{
			pbo.Foo2 = goo.Foo2
		}
		{
			if goo.ArraysStructField != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.ArraysStructField.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.ArraysStructField = pbom.(*testspb.ArraysStruct)
				if pbo.ArraysStructField == nil {
					pbo.ArraysStructField = new(testspb.ArraysStruct)
				}
			}
		}
		{
			gool := len(goo.Foo3)
			if gool == 0 {
				pbo.Foo3 = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Foo3[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.Foo3 = pbos
			}
		}
		{
			if goo.SlicesStruct != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.SlicesStruct.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.SlicesStruct = pbom.(*testspb.SlicesStruct)
				if pbo.SlicesStruct == nil {
					pbo.SlicesStruct = new(testspb.SlicesStruct)
				}
			}
		}
		{
			pbo.Foo4 = goo.Foo4
		}
		{
			if goo.PointersStructField != nil {
				pbom := proto.Message(nil)
				pbom, err = goo.PointersStructField.ToPBMessage(cdc)
				if err != nil {
					return
				}
				pbo.PointersStructField = pbom.(*testspb.PointersStruct)
				if pbo.PointersStructField == nil {
					pbo.PointersStructField = new(testspb.PointersStruct)
				}
			}
		}
		{
			pbo.Foo5 = uint64(goo.Foo5)
		}
	}
	msg = pbo
	return
}
func (goo *EmbeddedSt5) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.EmbeddedSt5 = msg.(*testspb.EmbeddedSt5)
	{
		if pbo != nil {
			{
				goo.Foo1 = int(pbo.Foo1)
			}
			{
				if pbo.PrimitivesStruct != nil {
					goo.PrimitivesStruct = new(PrimitivesStruct)
					err = (*goo.PrimitivesStruct).FromPBMessage(cdc, pbo.PrimitivesStruct)
					if err != nil {
						return
					}
				}
			}
			{
				goo.Foo2 = pbo.Foo2
			}
			{
				if pbo.ArraysStructField != nil {
					goo.ArraysStructField = new(ArraysStruct)
					err = (*goo.ArraysStructField).FromPBMessage(cdc, pbo.ArraysStructField)
					if err != nil {
						return
					}
				}
			}
			{
				var pbol int = 0
				if pbo.Foo3 != nil {
					pbol = len(pbo.Foo3)
				}
				if pbol == 0 {
					goo.Foo3 = nil
				} else {
					var goos = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Foo3[i]
							{
								goos[i] = uint8(pboe)
							}
						}
					}
					goo.Foo3 = goos
				}
			}
			{
				if pbo.SlicesStruct != nil {
					goo.SlicesStruct = new(SlicesStruct)
					err = (*goo.SlicesStruct).FromPBMessage(cdc, pbo.SlicesStruct)
					if err != nil {
						return
					}
				}
			}
			{
				goo.Foo4 = pbo.Foo4
			}
			{
				if pbo.PointersStructField != nil {
					goo.PointersStructField = new(PointersStruct)
					err = (*goo.PointersStructField).FromPBMessage(cdc, pbo.PointersStructField)
					if err != nil {
						return
					}
				}
			}
			{
				goo.Foo5 = uint(pbo.Foo5)
			}
		}
	}
	return
}
func (_ EmbeddedSt5) GetTypeURL() (typeURL string) {
	return "/tests.EmbeddedSt5"
}
func (goo EmbeddedSt5) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if goo.Foo1 != 0 {
				return false
			}
		}
		{
			if goo.PrimitivesStruct != nil {
				return false
			}
		}
		{
			if goo.Foo2 != "" {
				return false
			}
		}
		{
			if goo.ArraysStructField != nil {
				return false
			}
		}
		{
			if len(goo.Foo3) != 0 {
				return false
			}
		}
		{
			if goo.SlicesStruct != nil {
				return false
			}
		}
		{
			if goo.Foo4 != false {
				return false
			}
		}
		{
			if goo.PointersStructField != nil {
				return false
			}
		}
		{
			if goo.Foo5 != 0 {
				return false
			}
		}
	}
	return
}
func (goo PrimitivesStructDef) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.PrimitivesStructDef
	{
		if goo.IsEmpty() {
			var pbov *testspb.PrimitivesStructDef
			msg = pbov
			return
		}
		pbo = new(testspb.PrimitivesStructDef)
		{
			pbo.Int8 = int32(goo.Int8)
		}
		{
			pbo.Int16 = int32(goo.Int16)
		}
		{
			pbo.Int32 = goo.Int32
		}
		{
			pbo.Int32Fixed = goo.Int32Fixed
		}
		{
			pbo.Int64 = goo.Int64
		}
		{
			pbo.Int64Fixed = goo.Int64Fixed
		}
		{
			pbo.Int = int64(goo.Int)
		}
		{
			pbo.Byte = uint32(goo.Byte)
		}
		{
			pbo.Uint8 = uint32(goo.Uint8)
		}
		{
			pbo.Uint16 = uint32(goo.Uint16)
		}
		{
			pbo.Uint32 = goo.Uint32
		}
		{
			pbo.Uint32Fixed = goo.Uint32Fixed
		}
		{
			pbo.Uint64 = goo.Uint64
		}
		{
			pbo.Uint64Fixed = goo.Uint64Fixed
		}
		{
			pbo.Uint = uint64(goo.Uint)
		}
		{
			pbo.Str = goo.Str
		}
		{
			gool := len(goo.Bytes)
			if gool == 0 {
				pbo.Bytes = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Bytes[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.Bytes = pbos
			}
		}
		{
			if !amino.IsEmptyTime(goo.Time) {
				pbo.Time = timestamppb.New(goo.Time)
			}
		}
		{
			pbom := proto.Message(nil)
			pbom, err = goo.Empty.ToPBMessage(cdc)
			if err != nil {
				return
			}
			pbo.Empty = pbom.(*testspb.EmptyStruct)
		}
	}
	msg = pbo
	return
}
func (goo *PrimitivesStructDef) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.PrimitivesStructDef = msg.(*testspb.PrimitivesStructDef)
	{
		if pbo != nil {
			{
				goo.Int8 = int8(pbo.Int8)
			}
			{
				goo.Int16 = int16(pbo.Int16)
			}
			{
				goo.Int32 = pbo.Int32
			}
			{
				goo.Int32Fixed = pbo.Int32Fixed
			}
			{
				goo.Int64 = pbo.Int64
			}
			{
				goo.Int64Fixed = pbo.Int64Fixed
			}
			{
				goo.Int = int(pbo.Int)
			}
			{
				goo.Byte = uint8(pbo.Byte)
			}
			{
				goo.Uint8 = uint8(pbo.Uint8)
			}
			{
				goo.Uint16 = uint16(pbo.Uint16)
			}
			{
				goo.Uint32 = pbo.Uint32
			}
			{
				goo.Uint32Fixed = pbo.Uint32Fixed
			}
			{
				goo.Uint64 = pbo.Uint64
			}
			{
				goo.Uint64Fixed = pbo.Uint64Fixed
			}
			{
				goo.Uint = uint(pbo.Uint)
			}
			{
				goo.Str = pbo.Str
			}
			{
				var pbol int = 0
				if pbo.Bytes != nil {
					pbol = len(pbo.Bytes)
				}
				if pbol == 0 {
					goo.Bytes = nil
				} else {
					var goos = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Bytes[i]
							{
								goos[i] = uint8(pboe)
							}
						}
					}
					goo.Bytes = goos
				}
			}
			{
				goo.Time = pbo.Time.AsTime()
			}
			{
				if pbo.Empty != nil {
					err = goo.Empty.FromPBMessage(cdc, pbo.Empty)
					if err != nil {
						return
					}
				}
			}
		}
	}
	return
}
func (_ PrimitivesStructDef) GetTypeURL() (typeURL string) {
	return "/tests.PrimitivesStructDef"
}
func (goo PrimitivesStructDef) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if goo.Int8 != 0 {
				return false
			}
		}
		{
			if goo.Int16 != 0 {
				return false
			}
		}
		{
			if goo.Int32 != 0 {
				return false
			}
		}
		{
			if goo.Int32Fixed != 0 {
				return false
			}
		}
		{
			if goo.Int64 != 0 {
				return false
			}
		}
		{
			if goo.Int64Fixed != 0 {
				return false
			}
		}
		{
			if goo.Int != 0 {
				return false
			}
		}
		{
			if goo.Byte != 0 {
				return false
			}
		}
		{
			if goo.Uint8 != 0 {
				return false
			}
		}
		{
			if goo.Uint16 != 0 {
				return false
			}
		}
		{
			if goo.Uint32 != 0 {
				return false
			}
		}
		{
			if goo.Uint32Fixed != 0 {
				return false
			}
		}
		{
			if goo.Uint64 != 0 {
				return false
			}
		}
		{
			if goo.Uint64Fixed != 0 {
				return false
			}
		}
		{
			if goo.Uint != 0 {
				return false
			}
		}
		{
			if goo.Str != "" {
				return false
			}
		}
		{
			if len(goo.Bytes) != 0 {
				return false
			}
		}
		{
			if !amino.IsEmptyTime(goo.Time) {
				return false
			}
		}
		{
			e := goo.Empty.IsEmpty()
			if e == false {
				return false
			}
		}
	}
	return
}
func (goo Concrete1) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.Concrete1
	{
		if goo.IsEmpty() {
			var pbov *testspb.Concrete1
			msg = pbov
			return
		}
		pbo = new(testspb.Concrete1)
	}
	msg = pbo
	return
}
func (goo *Concrete1) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.Concrete1 = msg.(*testspb.Concrete1)
	{
		if pbo != nil {
		}
	}
	return
}
func (_ Concrete1) GetTypeURL() (typeURL string) {
	return "/tests.Concrete1"
}
func (goo Concrete1) IsEmpty() (empty bool) {
	{
		empty = true
	}
	return
}
func (goo Concrete2) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.Concrete2
	{
		if goo.IsEmpty() {
			var pbov *testspb.Concrete2
			msg = pbov
			return
		}
		pbo = new(testspb.Concrete2)
	}
	msg = pbo
	return
}
func (goo *Concrete2) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.Concrete2 = msg.(*testspb.Concrete2)
	{
		if pbo != nil {
		}
	}
	return
}
func (_ Concrete2) GetTypeURL() (typeURL string) {
	return "/tests.Concrete2"
}
func (goo Concrete2) IsEmpty() (empty bool) {
	{
		empty = true
	}
	return
}
func (goo ConcreteWrappedBytes) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.ConcreteWrappedBytes
	{
		if goo.IsEmpty() {
			var pbov *testspb.ConcreteWrappedBytes
			msg = pbov
			return
		}
		pbo = new(testspb.ConcreteWrappedBytes)
		{
			gool := len(goo.Value)
			if gool == 0 {
				pbo.Value = nil
			} else {
				var pbos = make([]uint8, gool)
				for i := 0; i < gool; i += 1 {
					{
						gooe := goo.Value[i]
						{
							pbos[i] = byte(gooe)
						}
					}
				}
				pbo.Value = pbos
			}
		}
	}
	msg = pbo
	return
}
func (goo *ConcreteWrappedBytes) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.ConcreteWrappedBytes = msg.(*testspb.ConcreteWrappedBytes)
	{
		if pbo != nil {
			{
				var pbol int = 0
				if pbo.Value != nil {
					pbol = len(pbo.Value)
				}
				if pbol == 0 {
					goo.Value = nil
				} else {
					var goos = make([]uint8, pbol)
					for i := 0; i < pbol; i += 1 {
						{
							pboe := pbo.Value[i]
							{
								goos[i] = uint8(pboe)
							}
						}
					}
					goo.Value = goos
				}
			}
		}
	}
	return
}
func (_ ConcreteWrappedBytes) GetTypeURL() (typeURL string) {
	return "/tests.ConcreteWrappedBytes"
}
func (goo ConcreteWrappedBytes) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if len(goo.Value) != 0 {
				return false
			}
		}
	}
	return
}
func (goo InterfaceFieldsStruct) ToPBMessage(cdc *amino.Codec) (msg proto.Message, err error) {
	var pbo *testspb.InterfaceFieldsStruct
	{
		if goo.IsEmpty() {
			var pbov *testspb.InterfaceFieldsStruct
			msg = pbov
			return
		}
		pbo = new(testspb.InterfaceFieldsStruct)
		{
			if goo.F1 != nil {
				typeUrl := goo.F1.(amino.Object).GetTypeURL()
				bz := []byte(nil)
				bz, err = cdc.MarshalBinaryBare(goo.F1)
				if err != nil {
					return
				}
				pbo.F1 = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
		{
			if goo.F2 != nil {
				typeUrl := goo.F2.(amino.Object).GetTypeURL()
				bz := []byte(nil)
				bz, err = cdc.MarshalBinaryBare(goo.F2)
				if err != nil {
					return
				}
				pbo.F2 = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
		{
			if goo.F3 != nil {
				typeUrl := goo.F3.(amino.Object).GetTypeURL()
				bz := []byte(nil)
				bz, err = cdc.MarshalBinaryBare(goo.F3)
				if err != nil {
					return
				}
				pbo.F3 = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
		{
			if goo.F4 != nil {
				typeUrl := goo.F4.(amino.Object).GetTypeURL()
				bz := []byte(nil)
				bz, err = cdc.MarshalBinaryBare(goo.F4)
				if err != nil {
					return
				}
				pbo.F4 = &anypb.Any{TypeUrl: typeUrl, Value: bz}
			}
		}
	}
	msg = pbo
	return
}
func (goo *InterfaceFieldsStruct) FromPBMessage(cdc *amino.Codec, msg proto.Message) (err error) {
	var pbo *testspb.InterfaceFieldsStruct = msg.(*testspb.InterfaceFieldsStruct)
	{
		if pbo != nil {
			{
				typeUrl := pbo.F1.TypeUrl
				bz := pbo.F1.Value
				goop := &goo.F1
				err = cdc.UnmarshalBinaryAny(typeUrl, bz, goop)
				if err != nil {
					return
				}
			}
			{
				typeUrl := pbo.F2.TypeUrl
				bz := pbo.F2.Value
				goop := &goo.F2
				err = cdc.UnmarshalBinaryAny(typeUrl, bz, goop)
				if err != nil {
					return
				}
			}
			{
				typeUrl := pbo.F3.TypeUrl
				bz := pbo.F3.Value
				goop := &goo.F3
				err = cdc.UnmarshalBinaryAny(typeUrl, bz, goop)
				if err != nil {
					return
				}
			}
			{
				typeUrl := pbo.F4.TypeUrl
				bz := pbo.F4.Value
				goop := &goo.F4
				err = cdc.UnmarshalBinaryAny(typeUrl, bz, goop)
				if err != nil {
					return
				}
			}
		}
	}
	return
}
func (_ InterfaceFieldsStruct) GetTypeURL() (typeURL string) {
	return "/tests.InterfaceFieldsStruct"
}
func (goo InterfaceFieldsStruct) IsEmpty() (empty bool) {
	{
		empty = true
		{
			if goo.F1 != nil {
				return false
			}
		}
		{
			if goo.F2 != nil {
				return false
			}
		}
		{
			if goo.F3 != nil {
				return false
			}
		}
		{
			if goo.F4 != nil {
				return false
			}
		}
	}
	return
}
