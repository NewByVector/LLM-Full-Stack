## 全局操作
whoami       # current user
pwd          # current directory
su [user]    # switch user
[cmd] --help # get more inofrmation
man [cmd]    # open user manual
id           # user id
uname        # system info
cal          # calender

## 文件操作
ls    # list files
ls -a # list files include hidden files
ls -R # list files include sub files

cd    # move to home directory
cd ~  # move to home directory
cd /  # move to root directory
cd .. # move to parent directory

mkdir [dir]  # create new directory
rmdir [dir]  # remove a directory
touch [file] # create new file
rm [file]    # remove a file

cat [file]  # view file
more [file] # view large file
less [file] # view large file
head [file] # view first ten line of file
tail [file] # view last ten line of file

echo something          # print something
echo something >>[file] # input something into file

cp [file1] [file2] # copy file
mv [file1] [file2] # move file

find [dir] [file]   # search file in directory
locate [file]       # searh file by database
grep [word] [file]  # search word in file
cmp [file1] [file2] # compare file1 with file2

tar -zcvf [file.tar.gz] [file] # compress file
tar -zxvf [file.tar.gz]        # uncompress file
zip [file.zip] [file]          # compress file to .zip
unzip [file.zip]               # unzip file

alias k='cmd' # alias k replace cmd
unalias k     # cancel alias

sudo ufw status         # get firewall status
sudo ufw enable/disable # enable/disable firewall
sudo ufw allow http     # enable http port
sudo ufw deny ssh       # deny ssh port
sudo ufw reset          # reset firewall status

ps -e                      # view process
ps -e | grep 'processname' # view a process
top                        # view process
kill pid                   # kill process

systemctl start [service]   # start service
systemctl restart [service] # restart service
systemctl stop [service]    # stop service
