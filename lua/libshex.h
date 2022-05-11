/* Code generated by cmd/cgo; DO NOT EDIT. */

/* package peon.top/shex/main */


#line 1 "cgo-builtin-export-prolog"

#include <stddef.h> /* for ptrdiff_t below */

#ifndef GO_CGO_EXPORT_PROLOGUE_H
#define GO_CGO_EXPORT_PROLOGUE_H

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#endif

#endif

/* Start of preamble from import "C" comments.  */




/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

#ifndef GO_CGO_GOSTRING_TYPEDEF
typedef _GoString_ GoString;
#endif
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif

extern GoInt sum(GoInt a, GoInt b);
extern void println(GoString str);
extern GoInt createExcelFile(GoString filename, GoString defaultSheetName);
extern void cell(GoInt resourceId, GoString tabIndex, GoString val);
extern void selectSheet(GoInt resourceId, GoString sheet);
extern void merge(GoInt resourceId, GoString si, GoString ei);
extern void save(GoInt resourceId);
extern GoInt registerStyle(GoInt resourceId, GoString style);
extern void setCellStyle(GoInt resourceId, GoString cellX, GoString cellY, GoInt styleId);
extern void setColWidth(GoInt resourceId, GoString startCol, GoString endCol, GoFloat64 width);
extern void setRowHeight(GoInt resourceId, GoInt row, GoFloat64 height);
extern void insertPageBreak(GoInt resourceId, GoString cell);
extern void setColStyle(GoInt resourceId, GoString columns, GoInt styleId);
extern GoInt getCellStyle(GoInt resourceId, GoString axis);

//实验功能
extern GoInt appendBoardStyle(GoInt resourceId, GoString board, GoString axis);
extern GoInt setPageMargins(GoInt resourceId, GoFloat64 top, GoFloat64 left, GoFloat64 right, GoFloat64 bottom, GoFloat64 header, GoFloat64 footer);

#ifdef __cplusplus
}
#endif
