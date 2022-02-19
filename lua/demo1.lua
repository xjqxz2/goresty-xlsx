local Excel = require("xlsx")

-- 这是一个 DEMO 
local table1 = Excel:new("表格1.xlsx","HELL")

local redBoardStyle = "{\"type\": \"bottom\",\"color\":\"F1223F\",\"style\":1}"

--  预设单元格样式1
-- 蓝低无边框样式
local blueBackgroundStyleId = table1:registerStyle("{\"fill\":{\"type\":\"pattern\",\"color\":[\"#E0EBF5\"],\"pattern\":1}}")

-- 黑底无边框样式
local blackBackgroundStyleId = table1:registerStyle("{\"fill\":{\"type\":\"pattern\",\"color\":[\"#000000\"],\"pattern\":1}}")

--  将 A1-E10 区域设置为 Style1 的样式，即蓝底无线框样式
table1:setCellStyle("A1","E10",blueBackgroundStyleId)

--  将 A5-B6 区域设置为 Style1 的样式，即黑底无线框样式
table1:setCellStyle("A5","B6",blackBackgroundStyleId)
table1:cell("A3","H-1")
table1:sel("Sheet2")
table1:cell("B2","你好")
table1:sel("Sheet1")
table1:merge("A1","A3")
table1:cell("A3","H-2")
table1:sel("HELL")
table1:appendBoardStyle(redBoardStyle,"A3")
table1:insertPageBreak("C1")
table1:save()