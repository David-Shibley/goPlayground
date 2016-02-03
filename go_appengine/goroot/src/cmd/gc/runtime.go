// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// NOTE: If you change this file you must run "./mkbuiltin"
// to update builtin.c.  This is not done automatically
// to avoid depending on having a working compiler binary.

// +build ignore

package PACKAGE

// emitted by compiler, not referred to by go programs

func newobject(typ *byte) *any
func panicindex()
func panicslice()
func panicdivide()
func throwreturn()
func throwinit()
func panicwrap(string, string, string)

func gopanic(interface{})
func gorecover(*int32) interface{}

func printbool(bool)
func printfloat(float64)
func printint(int64)
func printhex(uint64)
func printuint(uint64)
func printcomplex(complex128)
func printstring(string)
func printpointer(any)
func printiface(any)
func printeface(any)
func printslice(any)
func printnl()
func printsp()

func concatstring2(string, string) string
func concatstring3(string, string, string) string
func concatstring4(string, string, string, string) string
func concatstring5(string, string, string, string, string) string
func concatstrings([]string) string

func cmpstring(string, string) int
func eqstring(string, string) bool
func intstring(int64) string
func slicebytetostring([]byte) string
func slicebytetostringtmp([]byte) string
func slicerunetostring([]rune) string
func stringtoslicebyte(string) []byte
func stringtoslicerune(string) []rune
func stringiter(string, int) int
func stringiter2(string, int) (retk int, retv rune)
func slicecopy(to any, fr any, wid uintptr) int
func slicestringcopy(to any, fr any) int

// interface conversions
func typ2Itab(typ *byte, typ2 *byte, cache **byte) (ret *byte)
func convI2E(elem any) (ret any)
func convI2I(typ *byte, elem any) (ret any)
func convT2E(typ *byte, elem *any) (ret any)
func convT2I(typ *byte, typ2 *byte, cache **byte, elem *any) (ret any)

// interface type assertions  x.(T)
func assertE2E(typ *byte, iface any) (ret any)
func assertE2E2(typ *byte, iface any) (ret any, ok bool)
func assertE2I(typ *byte, iface any) (ret any)
func assertE2I2(typ *byte, iface any) (ret any, ok bool)
func assertE2T(typ *byte, iface any) (ret any)
func assertE2T2(typ *byte, iface any) (ret any, ok bool)
func assertI2E(typ *byte, iface any) (ret any)
func assertI2E2(typ *byte, iface any) (ret any, ok bool)
func assertI2I(typ *byte, iface any) (ret any)
func assertI2I2(typ *byte, iface any) (ret any, ok bool)
func assertI2T(typ *byte, iface any) (ret any)
func assertI2T2(typ *byte, iface any) (ret any, ok bool)
func assertI2TOK(typ *byte, iface any) (ok bool)
func assertE2TOK(typ *byte, iface any) (ok bool)

func ifaceeq(i1 any, i2 any) (ret bool)
func efaceeq(i1 any, i2 any) (ret bool)
func ifacethash(i1 any) (ret uint32)
func efacethash(i1 any) (ret uint32)

// *byte is really *runtime.Type
func makemap(mapType *byte, hint int64) (hmap map[any]any)
func mapaccess1(mapType *byte, hmap map[any]any, key *any) (val *any)
func mapaccess1_fast32(mapType *byte, hmap map[any]any, key any) (val *any)
func mapaccess1_fast64(mapType *byte, hmap map[any]any, key any) (val *any)
func mapaccess1_faststr(mapType *byte, hmap map[any]any, key any) (val *any)
func mapaccess2(mapType *byte, hmap map[any]any, key *any) (val *any, pres bool)
func mapaccess2_fast32(mapType *byte, hmap map[any]any, key any) (val *any, pres bool)
func mapaccess2_fast64(mapType *byte, hmap map[any]any, key any) (val *any, pres bool)
func mapaccess2_faststr(mapType *byte, hmap map[any]any, key any) (val *any, pres bool)
func mapassign1(mapType *byte, hmap map[any]any, key *any, val *any)
func mapiterinit(mapType *byte, hmap map[any]any, hiter *any)
func mapdelete(mapType *byte, hmap map[any]any, key *any)
func mapiternext(hiter *any)

// *byte is really *runtime.Type
func makechan(chanType *byte, hint int64) (hchan chan any)
func chanrecv1(chanType *byte, hchan <-chan any, elem *any)
func chanrecv2(chanType *byte, hchan <-chan any, elem *any) bool
func chansend1(chanType *byte, hchan chan<- any, elem *any)
func closechan(hchan any)

// *byte is really *runtime.Type
func writebarrierptr(dst *any, src any)
func writebarrierstring(dst *any, src any)
func writebarrierslice(dst *any, src any)
func writebarrieriface(dst *any, src any)

// The unused *byte argument makes sure that src is 2-pointer-aligned,
// which is the maximum alignment on NaCl amd64p32
// (and possibly on 32-bit systems if we start 64-bit aligning uint64s).
func writebarrierfat2(dst *any, _ *byte, src any)
func writebarrierfat3(dst *any, _ *byte, src any)
func writebarrierfat4(dst *any, _ *byte, src any)
func writebarrierfat(typ *byte, dst *any, src *any)

func selectnbsend(chanType *byte, hchan chan<- any, elem *any) bool
func selectnbrecv(chanType *byte, elem *any, hchan <-chan any) bool
func selectnbrecv2(chanType *byte, elem *any, received *bool, hchan <-chan any) bool

func newselect(sel *byte, selsize int64, size int32)
func selectsend(sel *byte, hchan chan<- any, elem *any) (selected bool)
func selectrecv(sel *byte, hchan <-chan any, elem *any) (selected bool)
func selectrecv2(sel *byte, hchan <-chan any, elem *any, received *bool) (selected bool)
func selectdefault(sel *byte) (selected bool)
func selectgo(sel *byte)
func block()

func makeslice(typ *byte, nel int64, cap int64) (ary []any)
func growslice(typ *byte, old []any, n int64) (ary []any)
func memmove(to *any, frm *any, length uintptr)

func memequal(x, y *any, size uintptr) bool
func memequal8(x, y *any, size uintptr) bool
func memequal16(x, y *any, size uintptr) bool
func memequal32(x, y *any, size uintptr) bool
func memequal64(x, y *any, size uintptr) bool
func memequal128(x, y *any, size uintptr) bool

// only used on 32-bit
func int64div(int64, int64) int64
func uint64div(uint64, uint64) uint64
func int64mod(int64, int64) int64
func uint64mod(uint64, uint64) uint64
func float64toint64(float64) int64
func float64touint64(float64) uint64
func int64tofloat64(int64) float64
func uint64tofloat64(uint64) float64

func complex128div(num complex128, den complex128) (quo complex128)

// race detection
func racefuncenter(uintptr)
func racefuncexit()
func raceread(uintptr)
func racewrite(uintptr)
func racereadrange(addr, size uintptr)
func racewriterange(addr, size uintptr)
