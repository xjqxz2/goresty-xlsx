local Excel = require("xlsx")

-- 这是一个 DEMO 2
local table1 = Excel:new("表格1.xlsx","HELL")

for i=1,20 do
    table1:cell("A"..i,"H-"..i)
end

table1:insertPageBreak("A3")

table1:save()
