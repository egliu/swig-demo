/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.2
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ./compare_length/compare_length.i

package compare_length

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef long long swig_type_1;
typedef long long swig_type_2;
typedef long long swig_type_3;
typedef long long swig_type_4;
extern void _wrap_Swig_free_compare_length_f57c2e54145f5e02(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_compare_length_f57c2e54145f5e02(swig_intgo arg1);
extern uintptr_t _wrap_new_VecInt__SWIG_0_compare_length_f57c2e54145f5e02(void);
extern uintptr_t _wrap_new_VecInt__SWIG_1_compare_length_f57c2e54145f5e02(swig_type_1 arg1);
extern uintptr_t _wrap_new_VecInt__SWIG_2_compare_length_f57c2e54145f5e02(uintptr_t arg1);
extern swig_type_2 _wrap_VecInt_size_compare_length_f57c2e54145f5e02(uintptr_t arg1);
extern swig_type_3 _wrap_VecInt_capacity_compare_length_f57c2e54145f5e02(uintptr_t arg1);
extern void _wrap_VecInt_reserve_compare_length_f57c2e54145f5e02(uintptr_t arg1, swig_type_4 arg2);
extern _Bool _wrap_VecInt_isEmpty_compare_length_f57c2e54145f5e02(uintptr_t arg1);
extern void _wrap_VecInt_clear_compare_length_f57c2e54145f5e02(uintptr_t arg1);
extern void _wrap_VecInt_add_compare_length_f57c2e54145f5e02(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_VecInt_get_compare_length_f57c2e54145f5e02(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_VecInt_set_compare_length_f57c2e54145f5e02(uintptr_t arg1, swig_intgo arg2, swig_intgo arg3);
extern void _wrap_delete_VecInt_compare_length_f57c2e54145f5e02(uintptr_t arg1);
extern swig_intgo _wrap_compare_compare_length_f57c2e54145f5e02(uintptr_t arg1, uintptr_t arg2);
#undef intgo
#cgo LDFLAGS: -L../libs/ -ll2
*/
import "C"

import (
	_ "runtime/cgo"
	"sync"
	"unsafe"
)

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte

type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_compare_length_f57c2e54145f5e02(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrVecInt uintptr

func (p SwigcptrVecInt) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVecInt) SwigIsVecInt() {
}

func NewVecInt__SWIG_0() (_swig_ret VecInt) {
	var swig_r VecInt
	swig_r = (VecInt)(SwigcptrVecInt(C._wrap_new_VecInt__SWIG_0_compare_length_f57c2e54145f5e02()))
	return swig_r
}

func NewVecInt__SWIG_1(arg1 int64) (_swig_ret VecInt) {
	var swig_r VecInt
	_swig_i_0 := arg1
	swig_r = (VecInt)(SwigcptrVecInt(C._wrap_new_VecInt__SWIG_1_compare_length_f57c2e54145f5e02(C.swig_type_1(_swig_i_0))))
	return swig_r
}

func NewVecInt__SWIG_2(arg1 VecInt) (_swig_ret VecInt) {
	var swig_r VecInt
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (VecInt)(SwigcptrVecInt(C._wrap_new_VecInt__SWIG_2_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewVecInt(a ...interface{}) VecInt {
	argc := len(a)
	if argc == 0 {
		return NewVecInt__SWIG_0()
	}
	if argc == 1 {
		if _, ok := a[0].(int64); !ok {
			goto check_2
		}
		return NewVecInt__SWIG_1(a[0].(int64))
	}
check_2:
	if argc == 1 {
		return NewVecInt__SWIG_2(a[0].(VecInt))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrVecInt) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_VecInt_size_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVecInt) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_VecInt_capacity_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVecInt) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_VecInt_reserve_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0), C.swig_type_4(_swig_i_1))
}

func (arg1 SwigcptrVecInt) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_VecInt_isEmpty_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVecInt) Clear() {
	_swig_i_0 := arg1
	C._wrap_VecInt_clear_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVecInt) Add(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_VecInt_add_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrVecInt) Get(arg2 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int)(C._wrap_VecInt_get_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func (arg1 SwigcptrVecInt) Set(arg2 int, arg3 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	C._wrap_VecInt_set_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.swig_intgo(_swig_i_2))
}

func DeleteVecInt(arg1 VecInt) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_VecInt_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0))
}

type VecInt interface {
	Swigcptr() uintptr
	SwigIsVecInt()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 int)
	Get(arg2 int) (_swig_ret int)
	Set(arg2 int, arg3 int)
}

func Compare(arg1 VecInt, arg2 VecInt) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (int)(C._wrap_compare_compare_length_f57c2e54145f5e02(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1)))
	return swig_r
}