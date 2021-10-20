local Excel = require("xlsx")

-- 这是一个 DEMO 2
local table1 = Excel:new("表格1.xlsx","HELL")

for i=1,200000 do
    table1:cell("A"..i,"H-"..i)
    table1:cell("A2","H-"..i)
end

table1:save()
