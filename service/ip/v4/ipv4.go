package v4

import (
	"dtapps/dta/library/utils/gostring"
	_ "embed"
	"encoding/binary"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net"
)

var (
	header  []byte
	country []byte
	area    []byte
	v4ip    uint32
	offset  uint32
	start   uint32
	end     uint32
)

//go:embed qqwry.dat
var dat []byte

type Pointer struct {
	Offset   uint32
	ItemLen  uint32
	IndexLen uint32
}

// Result 返回
type Result struct {
	IP      string `json:"ip,omitempty"`      // 输入的ip地址
	Country string `json:"country,omitempty"` // 国家或地区
	Area    string `json:"area,omitempty"`    // 区域
}

func (q *Pointer) readData(length uint32) (rs []byte) {
	end := q.Offset + length
	dataNum := uint32(len(dat))
	if q.Offset > dataNum {
		return nil
	}

	if end > dataNum {
		end = dataNum
	}
	rs = dat[q.Offset:end]
	q.Offset = end
	return rs
}

func (q *Pointer) Find(ip string) (res Result) {
	res.IP = ip
	q.Offset = 0

	v4ip = binary.BigEndian.Uint32(net.ParseIP(ip).To4())
	offset = q.searchIndex(v4ip)
	q.Offset = offset + q.ItemLen

	enc := simplifiedchinese.GBK.NewDecoder()
	country, area = q.getAddr()

	res.Country, _ = enc.String(string(country))
	res.Country = gostring.SpaceAndLineBreak(res.Country)

	res.Area, _ = enc.String(string(area))

	// Delete CZ88.NET (防止不相关的信息产生干扰）
	if res.Area == " CZ88.NET" || res.Area == "" {
		res.Area = ""
	} else {
		res.Area = " " + res.Area
	}

	res.Area = gostring.SpaceAndLineBreak(res.Area)

	return
}

func (q *Pointer) getAddr() ([]byte, []byte) {
	mode := q.readData(1)[0]
	if mode == 0x01 {
		// [IP][0x01][国家和地区信息的绝对偏移地址]
		q.Offset = byteToUInt32(q.readData(3))
		return q.getAddr()
	}
	// [IP][0x02][信息的绝对偏移][...] or [IP][国家][...]
	_offset := q.Offset - 1
	c1 := q.readArea(_offset)
	if mode == 0x02 {
		q.Offset = 4 + _offset
	} else {
		q.Offset = _offset + uint32(1+len(c1))
	}
	c2 := q.readArea(q.Offset)
	return c1, c2
}

func (q *Pointer) readArea(offset uint32) []byte {
	q.Offset = offset
	mode := q.readData(1)[0]
	if mode == 0x01 || mode == 0x02 {
		return q.readArea(byteToUInt32(q.readData(3)))
	}
	q.Offset = offset
	return q.readString()
}

func (q *Pointer) readString() []byte {
	data := make([]byte, 0)
	for {
		buf := q.readData(1)
		if buf[0] == 0 {
			break
		}
		data = append(data, buf[0])
	}
	return data
}

func (q *Pointer) searchIndex(ip uint32) uint32 {
	q.ItemLen = 4
	q.IndexLen = 7
	header = dat[0:8]
	start = binary.LittleEndian.Uint32(header[:4])
	end = binary.LittleEndian.Uint32(header[4:])

	buf := make([]byte, q.IndexLen)

	for {
		mid := start + q.IndexLen*(((end-start)/q.IndexLen)>>1)
		buf = dat[mid : mid+q.IndexLen]
		_ip := binary.LittleEndian.Uint32(buf[:q.ItemLen])

		if end-start == q.IndexLen {
			if ip >= binary.LittleEndian.Uint32(dat[end:end+q.ItemLen]) {
				buf = dat[end : end+q.IndexLen]
			}
			return byteToUInt32(buf[q.ItemLen:])
		}

		if _ip > ip {
			end = mid
		} else if _ip < ip {
			start = mid
		} else if _ip == ip {
			return byteToUInt32(buf[q.ItemLen:])
		}
	}
}

func byteToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}
