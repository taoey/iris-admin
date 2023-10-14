#!/bin/bash

# 设置数据库文件路径和输出文件名
database_file="data/zhihu_hot.db"
output_file="data/output.csv"

# SQL 查询语句，选择要导出的数据
sql_query="SELECT id,question_id,title,url,hot_score,excerpt,date(create_time, 'unixepoch', 'localtime') create_time 
FROM zhihu_hot 
WHERE create_time >= strftime('%s', 'now', 'start of day','-1 day') AND create_time < strftime('%s', 'now', 'start of day', '+1 day')
group by question_id order by hot_score desc;"

# 导出数据到 CSV 文件
sqlite3 "$database_file" <<EOF
.headers on
.mode csv
.output "$output_file"
.separator "|||"
$sql_query
.quit
EOF


echo "数据已导出到 $output_file"
