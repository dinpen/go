package net

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// 判断是否为合法的 IPv4 地址
func IsIPv4(ip string) bool {
	nip := net.ParseIP(ip)
	if nip != nil {
		nip = nip.To4()
		if nip != nil {
			return true
		}
	}
	return false
}

// 判断是否为合法的 IPv6 地址
func IsIPv6(ip string) bool {
	nip := net.ParseIP(ip)
	if nip != nil {
		nip = nip.To16()
		if nip != nil {
			return true
		}
	}
	return false
}

// 判断是否为合法的 CIDR
func IsCIDR(ip string) bool {
	if _, _, err := net.ParseCIDR(ip); err != nil {
		return false
	}
	return true
}

// 判断是否为合法的 IP 或 CIDR
func ValidateIPOrCIDR(ip string) bool {
	parseIP := net.ParseIP(ip)
	if parseIP == nil {
		_, _, parseCIDR := net.ParseCIDR(ip)
		if parseCIDR != nil {
			return false
		}
	}
	return true
}

// IP 转换为整数
func IPConvertToInt(ip string) (int64, error) {
	ips := strings.Split(ip, ".")
	if len(ips) != 4 {
		return 0, errors.New("invalid ip")
	}

	var res int
	for i, v := range ips {
		n, e := strconv.Atoi(v)
		if e != nil || n > 255 {
			return 0, errors.New("invalid ip")
		}
		res = res | i<<uint(8*(3-i))
	}
	return int64(res), nil
}
