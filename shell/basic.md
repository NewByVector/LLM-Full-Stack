## 变量
### 基础
Shell 变量分为系统变量和自定义变量。系统变量有$HOME、$PWD、$USER等。变量名可以由字母、数字、下划线组成，不能以数字开头。

- `定义变量：` 变量名=变量值，等号两侧不能有空格，变量名一般习惯用大写。
- `删除变量：` unset 变量名 。
- `声明静态变量：` readonly 变量名，静态变量不能unset。
- `使用变量：` $变量名
- `将命令返回值赋给变量：` A=\`ls\`  或者 A=$(ls) 
### 环境变量

- export 变量名=变量值，将 Shell 变量输出为环境变量。
- source 配置文件路径，让修改后的配置信息立即生效。
- echo $变量名，检查环境变量是否生效
### 位置参数变量

- `$n` ：$0 代表命令本身、`$1-$9` 代表第1到9个参数，10以上参数用花括号，如 ${10}。
- `$*` ：命令行中所有参数，且把所有参数看成一个整体。
- `$@` ：命令行中所有参数，且把每个参数区分对待。
- `$#` ：所有参数个数。
### 预定义变量

- `$$` ：当前进程的 PID 进程号。
- `$!` ：后台运行的最后一个进程的 PID 进程号。
- `$?` ：最后一次执行的命令的返回状态，0为执行正确，非0执行失败。
## 运算
```bash
# 第1种方式 $(()) 
echo $(((2+3)*4))   

# 第2种方式 $[]
echo $[(2+3)*4]  
```
## 语句
### 条件语句
注意condition前后要有空格。
```bash
#!/bin/bash
if [ $1 -ge 60 ]
then
    echo 及格
elif [ $1 -lt 60 ]
then
    echo "不及格" 
fi
```
也可以用 case 语法：
```bash
case $1 in
"1")
echo 周一
;;
"2")
echo 周二
;;
*)
echo 其它
;;
esac
```
### 循环语句
for in 循环
```bash
for j in "$@" 
do     
    echo "the arg is $j" 
done
```
for 循环
```bash
SUM=0  
for ((i=1;i<=100;i++)) 
do     
    SUM=$[$SUM+$i] 
done 

echo $SUM
```
while 循环
```bash
SUM=0
i=0

while [ $i -le $1 ]
do
    SUM=$[$SUM+$i]
    i=$[$i+1]
done       
echo $SUM
```
## 用户输入
```bash
read -p "请输入一个数num1=" NUM1
echo "你输入num1的值是：$NUM1"

read -t 10 -p "请在10秒内输入一个数num2=" NUM2
echo "你输入num2的值是：$NUM2"
```
## 函数
```bash
function getSum(){
    SUM=$[$n1+$n2]
    echo "sum=$SUM"
}   

read -p "请输入第一个参数n1：" n1
read -p "请输入第二个参数n2：" n2

# 调用 getSum 函数
getSum $n1 $n2
```
