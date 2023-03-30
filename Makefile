blah: blah.o
	cc  blah.o -o blah #第三步
blah.o: blah.c 
	cc -c blah.c -o blah.o # 第二步
blah.c:
	echo "int main(){return 0;}" > blah.c # 第一步

all: one two three first second
	echo $@   
	echo $?
	echo $^

first second: # $@ 规则目标的文件名; 相当与一个集合，依次取出并执行
# $< 第一个先决条件的名称
# $? 比目标新的所有先决条件的名称,他们之间有空格
# $^ 所有先决条件的名称,他们之间有空格
	echo $@ 
one:
	touch one
two:
	touch two
three:
	touch three

.PHONY: clean
clean: # 伪目标 .PHONY
	rm -f one two three blah*

MY_VAR1= A text string # 使用命名时解析变量值 =被称为是递归变量（recursive）,因为其在使用命名时解析，所以不能进行递归定义
MY_VAR2 := A var	# 定义时立即解析变量值 := 被称为是简单变量（simply expanded）
one = hello
one ?= will not set one 
two ?= will not set two
two += append var # 追加
.PHONY: print
print:
	echo ${MY_VAR1}
	echo $(MY_VAR2)
	echo ${one}
	echo ${two}