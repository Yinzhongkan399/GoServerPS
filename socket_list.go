//go:build linux
// +build linux

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ListAll 导出：返回与原 Python ListAll 等价的 JSON 字符串（以及可能的错误）
func ListAll() (string, error) {
	curTime := float64(time.Now().UnixNano()) / 1e9

	total := make(map[string][][]interface{})

	if err := getDevInfo(total, curTime); err != nil {
		return "", err
	}
	if err := getTcpInfo(total, curTime); err != nil {
		return "", err
	}
	if err := getUdpInfo(total, curTime); err != nil {
		return "", err
	}
	if err := getRawInfo(total, curTime); err != nil {
		return "", err
	}
	if err := getIcmpInfo(total, curTime); err != nil {
		return "", err
	}

	b, err := json.Marshal(total)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

/* --------- 非导出辅助函数（与 Python 对应） --------- */

func tranV4IntoP(input string) string {
	parts := strings.Split(input, ":")
	if len(parts) < 2 {
		return input
	}
	iphex := parts[0]
	if len(iphex) < 8 {
		// 防御性处理
		iphex = fmt.Sprintf("%08s", iphex)
	}
	b1, _ := strconv.ParseInt(iphex[6:8], 16, 64)
	b2, _ := strconv.ParseInt(iphex[4:6], 16, 64)
	b3, _ := strconv.ParseInt(iphex[2:4], 16, 64)
	b4, _ := strconv.ParseInt(iphex[0:2], 16, 64)
	port, _ := strconv.ParseInt(parts[1], 16, 64)
	return fmt.Sprintf("%d.%d.%d.%d:%d", b1, b2, b3, b4, port)
}

func tranV6IntoP(input string) string {
	parts := strings.Split(input, ":")
	if len(parts) < 2 {
		return input
	}
	iphex := parts[0]
	// ensure length 32
	if len(iphex) < 32 {
		iphex = fmt.Sprintf("%032s", iphex)
	}
	groups := make([]string, 8)
	for i := 0; i < 8; i++ {
		groups[i] = strings.ToLower(iphex[i*4 : i*4+4])
	}
	port, _ := strconv.ParseInt(parts[1], 16, 64)
	return strings.Join(groups, ":") + ":" + strconv.FormatInt(port, 10)
}

func tranStateIntoStr(s int64) string {
	switch s {
	case 1:
		return "01(ESTABLISHED)"
	case 2:
		return "02(SYN_SENT)"
	case 3:
		return "03(SYN_RECV)"
	case 4:
		return "04(FIN_WAIT1)"
	case 5:
		return "05(FIN_WAIT2)"
	case 6:
		return "06(TIME_WAIT)"
	case 7:
		return "07(CLOSE)"
	case 8:
		return "08(CLOSE_WAIT)"
	case 9:
		return "09(LAST_ACK)"
	case 10:
		return "0A(LISTEN)"
	case 11:
		return "0B(CLOSING)"
	default:
		return fmt.Sprintf("%02X(UNDEFINED)", s)
	}
}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func getTcpInfo(total map[string][][]interface{}, curTime float64) error {
	if lines, err := readLines("/proc/net/tcp"); err == nil {
		data := [][]interface{}{}
		for i := 1; i < len(lines); i++ {
			fields := strings.Fields(lines[i])
			if len(fields) < 4 {
				continue
			}
			sl := strings.TrimSuffix(fields[0], ":")
			local := tranV4IntoP(fields[1])
			remote := tranV4IntoP(fields[2])
			stateVal, _ := strconv.ParseInt(fields[3], 16, 64)
			state := tranStateIntoStr(stateVal)
			data = append(data, []interface{}{curTime, sl, local, remote, state})
		}
		total["tcpipv4"] = data
	}

	if lines, err := readLines("/proc/net/tcp6"); err == nil {
		data := [][]interface{}{}
		for i := 1; i < len(lines); i++ {
			fields := strings.Fields(lines[i])
			if len(fields) < 4 {
				continue
			}
			sl := strings.TrimSuffix(fields[0], ":")
			local := tranV6IntoP(fields[1])
			remote := tranV6IntoP(fields[2])
			stateVal, _ := strconv.ParseInt(fields[3], 16, 64)
			state := tranStateIntoStr(stateVal)
			data = append(data, []interface{}{curTime, sl, local, remote, state})
		}
		total["tcpipv6"] = data
	}
	return nil
}

func getUdpInfo(total map[string][][]interface{}, curTime float64) error {
	if lines, err := readLines("/proc/net/udp"); err == nil {
		data := [][]interface{}{}
		for i := 1; i < len(lines); i++ {
			fields := strings.Fields(lines[i])
			if len(fields) < 4 {
				continue
			}
			sl := strings.TrimSuffix(fields[0], ":")
			local := tranV4IntoP(fields[1])
			remote := tranV4IntoP(fields[2])
			stateVal, _ := strconv.ParseInt(fields[3], 16, 64)
			state := tranStateIntoStr(stateVal)
			data = append(data, []interface{}{curTime, sl, local, remote, state})
		}
		total["udpipv4"] = data
	}

	if lines, err := readLines("/proc/net/udp6"); err == nil {
		data := [][]interface{}{}
		for i := 1; i < len(lines); i++ {
			fields := strings.Fields(lines[i])
			if len(fields) < 4 {
				continue
			}
			sl := strings.TrimSuffix(fields[0], ":")
			local := tranV6IntoP(fields[1])
			remote := tranV6IntoP(fields[2])
			stateVal, _ := strconv.ParseInt(fields[3], 16, 64)
			state := tranStateIntoStr(stateVal)
			data = append(data, []interface{}{curTime, sl, local, remote, state})
		}
		total["udpipv6"] = data
	}
	return nil
}

func getIcmpInfo(total map[string][][]interface{}, curTime float64) error {
	if lines, err := readLines("/proc/net/icmp"); err == nil {
		data := [][]interface{}{}
		for i := 1; i < len(lines); i++ {
			fields := strings.Fields(lines[i])
			if len(fields) < 4 {
				continue
			}
			sl := strings.TrimSuffix(fields[0], ":")
			local := tranV4IntoP(fields[1])
			remote := tranV4IntoP(fields[2])
			stateVal, _ := strconv.ParseInt(fields[3], 16, 64)
			state := tranStateIntoStr(stateVal)
			data = append(data, []interface{}{curTime, sl, local, remote, state})
		}
		total["icmpipv4"] = data
	}

	if lines, err := readLines("/proc/net/icmp6"); err == nil {
		data := [][]interface{}{}
		for i := 1; i < len(lines); i++ {
			fields := strings.Fields(lines[i])
			if len(fields) < 4 {
				continue
			}
			sl := strings.TrimSuffix(fields[0], ":")
			local := tranV6IntoP(fields[1])
			remote := tranV6IntoP(fields[2])
			stateVal, _ := strconv.ParseInt(fields[3], 16, 64)
			state := tranStateIntoStr(stateVal)
			data = append(data, []interface{}{curTime, sl, local, remote, state})
		}
		total["icmpipv6"] = data
	}
	return nil
}

func getRawInfo(total map[string][][]interface{}, curTime float64) error {
	if lines, err := readLines("/proc/net/raw"); err == nil {
		data := [][]interface{}{}
		for i := 1; i < len(lines); i++ {
			fields := strings.Fields(lines[i])
			if len(fields) < 4 {
				continue
			}
			sl := strings.TrimSuffix(fields[0], ":")
			local := tranV4IntoP(fields[1])
			remote := tranV4IntoP(fields[2])
			stateVal, _ := strconv.ParseInt(fields[3], 16, 64)
			state := tranStateIntoStr(stateVal)
			data = append(data, []interface{}{curTime, sl, local, remote, state})
		}
		total["rawipv4"] = data
	}

	if lines, err := readLines("/proc/net/raw6"); err == nil {
		data := [][]interface{}{}
		for i := 1; i < len(lines); i++ {
			fields := strings.Fields(lines[i])
			if len(fields) < 4 {
				continue
			}
			sl := strings.TrimSuffix(fields[0], ":")
			local := tranV6IntoP(fields[1])
			remote := tranV6IntoP(fields[2])
			stateVal, _ := strconv.ParseInt(fields[3], 16, 64)
			state := tranStateIntoStr(stateVal)
			data = append(data, []interface{}{curTime, sl, local, remote, state})
		}
		total["rawipv6"] = data
	}
	return nil
}

func getDevInfo(total map[string][][]interface{}, curTime float64) error {
	lines, err := readLines("/proc/net/dev")
	if err != nil {
		return err
	}
	data := [][]interface{}{}
	for i := 2; i < len(lines); i++ { // skip first two header lines
		fields := strings.Fields(lines[i])
		if len(fields) == 0 {
			continue
		}
		ifname := strings.TrimSuffix(fields[0], ":")
		data = append(data, []interface{}{curTime, ifname})
	}
	total["dev"] = data
	return nil
}
