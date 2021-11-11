local ffi = require("ffi")

--  Load shex fii lib
local shex = ffi.load("shex")
ffi.cdef[[
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];
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
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;
typedef _GoString_ GoString;
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
]]

local Excel = {
    resourceId = 0,
    filename  = ""
}

-- Golang String 
function GoString(s) 
    local gstr = ffi.new("GoString")

    gstr.p = s
    gstr.n = #s

    return gstr
end



--  创建一个 Excel 文件
function Excel:new(filename,defaultSheetName) 
    o = o or {}
    setmetatable(o, self)
    self.__index = self
    self.filename = filename

    if #defaultSheetName <= 0 then 
        defaultSheetName = "Sheet1"
    end 

    self.resourceId = shex.createExcelFile(GoString(filename),GoString(defaultSheetName))
    return o
end 

--  选择一个工作表 默认为工作表名为 sheet1
function Excel:sel(sheet) 
    shex.selectSheet(self.resourceId ,GoString(sheet))
    return o
end

--  向工作区域中填充数据
function Excel:cell(index ,value) 
    shex.cell(
        self.resourceId ,
        GoString(index), 
        GoString(value)
    )

    return o
end

function Excel:merge(si ,ei) 
    shex.merge(self.resourceId,GoString(si),GoString(ei))
    return o
end

--  保存数据表
function Excel:save() 
    shex.save(self.resourceId)
end

--  注册Excel表样式
--  样式表请查阅这里 https://xuri.me/excelize/zh-hans/style.html 
function Excel:registerStyle(style) 
    return shex.registerStyle(
        self.resourceId,
        GoString(style)
    )
end 

--  设置单元格样式
function Excel:setCellStyle(cellIndexX ,cellIndexY,styleId)
    shex.setCellStyle(
        self.resourceId, 
        GoString(cellIndexX),
        GoString(cellIndexY),
        tonumber(styleId)
    )
end

--  设置列宽
function Excel:setColWidth(startCol ,endCol ,width) 
    --  设置 A 列 shex.setColWidth(self.resourceId,GoString("A"),GoString("A"),50.1)
    --  设置 A ~ C 列 shex.setColWidth(self.resourceId,GoString("A"),GoString("C"),50.1)
    shex.setColWidth(self.resourceId,GoString("A"),GoString("A"),50.1)
end

--  设置行高
function Excel:setRowHeight(row ,height) 
    shex.setRowHeight(self.resourceId,3,20.1)
end

-- 插入分页符
function Excel:insertPageBreak(cell) 
    shex.insertPageBreak(self.resourceId,cell)
end

-- 设置列样式
function Excel:setColStyle(columns ,styleId) 
    shex.setColStyle(self.resourceId,GoString(columns),tonumber(styleId))
end

-- 获取单元格样式
function Excel:getCellStyle(axis)
    shex.getCellStyle(self.resourceId,GoString(axis))
end

-- 追加单元格线条样式
function Excel:appendBoardStyle(board ,axis)
    return shex.appendBoardStyle(self.resourceId, GoString(board), GoString(axis))
end

return Excel