## maps  

### 思路  

先将一个字符串变为用字符隔开的数组，像这样["h","e","l","l","o"],  
然后创建一个类型为`[string]int`的`map`，当出现相同的字符时，`value`值+1，  
遍历字符串数组，记录字符出现的次数。    
