package main

import (
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"
)

var bigEndian = binary.LittleEndian

type ClassReader struct {
	bytecode []byte
}

func NewClassReader(bytecode []byte) *ClassReader {
	return &ClassReader{bytecode: bytecode}
}

func (this *ClassReader) ReadUint32() uint32 {
	value := bigEndian.Uint32(this.bytecode[:4])
	this.bytecode = this.bytecode[4:]
	return value
}

func (this *ClassReader) ReadBytes(len int) []byte {
	bytes := this.bytecode[:len]
	this.bytecode = this.bytecode[len:]
	return bytes
}

func (this *ClassReader) ReadIp() string {
	bytes := this.ReadBytes(4)
	return fmt.Sprintf("%d.%d.%d.%d", bytes[0], bytes[1], bytes[2], bytes[3])
}

func (this *ClassReader) ReadPort() uint16 {
	value := this.ReadBytes(4)
	return binary.BigEndian.Uint16(value[0:2])
}

type (
	DWORD uint32
)

type TCP_CONNECTION_OFFLOAD_STATE int

const (
	TcpConnectionOffloadStateInHost     TCP_CONNECTION_OFFLOAD_STATE = 0
	TcpConnectionOffloadStateOffloading TCP_CONNECTION_OFFLOAD_STATE = 1
	TcpConnectionOffloadStateOffloaded  TCP_CONNECTION_OFFLOAD_STATE = 2
	TcpConnectionOffloadStateUploading  TCP_CONNECTION_OFFLOAD_STATE = 3
	TcpConnectionOffloadStateMax        TCP_CONNECTION_OFFLOAD_STATE = 4
)

type MIB_TCPROW2 struct {
	dwState        DWORD
	dwLocalAddr    DWORD
	dwLocalPort    DWORD
	dwRemoteAddr   DWORD
	dwRemotePort   DWORD
	dwOwningPid    DWORD
	dwOffloadState TCP_CONNECTION_OFFLOAD_STATE
}

type MIB_TCPTABLE2 struct {
	dwNumEntries DWORD
	table        []MIB_TCPROW2
}

func main() {
	call := syscall.NewLazyDLL("Iphlpapi.dll")
	getTcpTable2 := call.NewProc("GetTcpTable2")
	var n uint32 = 0
	table := &MIB_TCPTABLE2{}
	r1, _, _ := getTcpTable2.Call(uintptr(unsafe.Pointer(table)), uintptr(unsafe.Pointer(&n)), 1)
	// if r1 != syscall.ERROR_INSUFFICIENT_BUFFER {
	// something bad happened; use syscall.Error(r1) to diagnose
	fmt.Println(r1)
	// }
	b := make([]byte, n)
	r2, _, _ := getTcpTable2.Call(uintptr(unsafe.Pointer(&b[0])), uintptr(unsafe.Pointer(&n)), 1)
	if r2 != 0 {
		// something bad happened; use syscall.Error(r1) to diagnose
		fmt.Println(r2)
	}
	reader := NewClassReader(b)
	loopCount := reader.ReadUint32()
	fmt.Println("count = ", loopCount)
	fmt.Printf("\t%s\t%16s\t%16s\t\t\t%s\t%s\n", "State", "Local Address", "Foreign Address", "PID", "Off Load State")
	for i := uint32(0); i < loopCount; i++ {
		bb := reader.ReadBytes(28)
		// fmt.Println("bytes: ", bb)
		r := NewClassReader(bb)
		fmt.Printf("\t%-6d\t%s:%-16d\t%s:%-16d\t%d\t%d\n", r.ReadUint32(), r.ReadIp(), r.ReadPort(), r.ReadIp(), r.ReadPort(), r.ReadUint32(), r.ReadUint32())
	}
}
