## Odin Chain
### Install Go
```bash
curl https://storage.googleapis.com/golang/go1.14.1.linux-amd64.tar.gz | tar -xzv -C /usr/local
#append the following line to ~/.profile:
export PATH=$PATH:/usr/local/go/bin
#refresh your profile
source ~/.profile
```
### Create Executable
```bash
root@ubuntu:~/go/src$ cd odinchain/cmd/godin
root@ubuntu:~/go/src/odinchain/cmd/godin$ go build
root@ubuntu:~/go/src/odinchain/cmd/godin$ cp godin /usr/local/bin
```