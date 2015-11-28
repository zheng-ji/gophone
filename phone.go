/*
* Copyright (c) 2015, zheng-ji.info
* */
package gophone

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
	"strings"
)

const (
	INT_LEN            = 4
	CHAR_LEN           = 1
	HEAD_LENGTH        = 8
	PHONE_INDEX_LENGTH = 9
	CHUNK              = 100
	PHONE_DAT          = "phone.dat"
)

type PhoneRecord struct {
	PhoneNum string
	Province string
	City     string
	ZipCode  string
	AreaZone string
	CardType string
}

var content []byte

func init() {
	_, fulleFilename, _, _ := runtime.Caller(0)
	var err error
	content, err = ioutil.ReadFile(path.Join(path.Dir(fulleFilename), PHONE_DAT))
	if err != nil {
		panic(err)
	}
}

func Display() {
	fmt.Println(getVersion())
	fmt.Println(getTotalRecord())
	fmt.Println(getFirstRecordOffset())
}

func (pr PhoneRecord) String() string {
	_str := fmt.Sprintf("PhoneNum: %s\nAreaZone: %s\nCardType: %s\nCity: %s\nZipCode: %s\nProvince: %s\n", pr.PhoneNum, pr.AreaZone, pr.CardType, pr.City, pr.ZipCode, pr.Province)
	return _str
}

func getVersion() string {
	return string(content[0:INT_LEN])
}

func getTotalRecord() int32 {
	total := (int32(len(content)) - getFirstRecordOffset()) / PHONE_INDEX_LENGTH
	return total
}

func getFirstRecordOffset() int32 {
	var offset int32
	buffer := bytes.NewBuffer(content[INT_LEN : INT_LEN*2])
	binary.Read(buffer, binary.LittleEndian, &offset)
	return offset
}

func getIndexRecord(offset int32) (phone_prefix int32, record_offset int32, card_type byte) {
	buffer := bytes.NewBuffer(content[offset : offset+INT_LEN])
	binary.Read(buffer, binary.LittleEndian, &phone_prefix)
	buffer = bytes.NewBuffer(content[offset+INT_LEN : offset+INT_LEN*2])
	binary.Read(buffer, binary.LittleEndian, &record_offset)
	buffer = bytes.NewBuffer(content[offset+INT_LEN*2 : offset+INT_LEN*2+CHAR_LEN])
	binary.Read(buffer, binary.LittleEndian, &card_type)
	return
}

func getOpCompany(cardtype byte) string {
	var card_str = ""
	switch cardtype {
	case '1':
		card_str = "移动"
	case '2':
		card_str = "联通"
	case '3':
		card_str = "电信"
	case '4':
		card_str = "电信虚拟运营商"
	case '5':
		card_str = "联通虚拟运营商"
	default:
		card_str = "移动虚拟运营商"
	}
	return card_str
}

// BinarySearch
func Find(phone_num string) (pr *PhoneRecord, err error) {
	err = nil
	if len(phone_num) < 7 || len(phone_num) > 11 {
		return nil, errors.New("illegal phone length")
	}

	var left int32 = 0
	phone_seven_int, _ := strconv.ParseInt(phone_num[0:7], 10, 32)
	phone_seven_int32 := int32(phone_seven_int)
	total_len := int32(len(content))
	right := getTotalRecord()
	firstPhoneRecordOffset := getFirstRecordOffset()
	for {
		if left > right {
			break
		}
		mid := (left + right) / 2
		current_offset := firstPhoneRecordOffset + mid*PHONE_INDEX_LENGTH

		if current_offset >= total_len {
			break
		}
		cur_phone, record_offset, card_type := getIndexRecord(current_offset)
		if cur_phone > phone_seven_int32 {
			right = mid - 1
		} else if cur_phone < phone_seven_int32 {
			left = mid + 1
		} else {
			s := record_offset
			e := record_offset + int32(strings.Index(string(content[record_offset:record_offset+CHUNK]), "\000"))
			record_content := string(content[s:e])
			_tmp := strings.Split(record_content, "|")
			card_str := getOpCompany(card_type)
			pr = &PhoneRecord{
				PhoneNum: phone_num,
				Province: _tmp[0],
				City:     _tmp[1],
				ZipCode:  _tmp[2],
				AreaZone: _tmp[3],
				CardType: card_str,
			}
			return
		}
	}
	return nil, errors.New("num not found")
}
