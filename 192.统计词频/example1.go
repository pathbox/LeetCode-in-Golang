package LeetCode192

// cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -rnk1 | awk '{print $2,$1}'
-r : 降序排列

-n : 以数字排序，默认是按照字符排序的。

uniq用去取出连续的重复行

-c ：统计重复行的次数

cat words.txt | xargs -n 1 | sort | uniq -c | sort -nr | awk '{print $2" "$1}'