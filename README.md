# govagrant

This is a POC to demonstrate how OS specific testing can be done via Vagrant.  

It is important to remember these type of tests are considered as Integration Testing, since they cannot be done in isolation and on all environments as-is. To convey the message better, there are two test files under info/ folder. One for unit testing and the other one for integration testing. Note that there are hard-coded values in the integration testing code that works with the very specific version of Vagrant VMs at the time of writing this. 

Since the VMs do not come with golang installed and for the purpose of this POC we need to run `go test` command, we first need to prepare the VMs.

I'm doing this on a Mac to completely isolate the Host OS from the Guest ones.

Pre-requisites to run this POC:
- VirtualBox. I have version 6.1 
- Vagrant. I have version 2.2.16
- Golang. I have version 1.16.3

## 1) Prepare the VMs
### Install Golang on the Windows VM and halt it
```bash
$ vagrant up windows
$ # This vagrant VM shares the current folder under \vagrant
$ vagrant powershell -c "dir \vagrant"
$ vagrant powershell -c "start-process \vagrant\go1.16.3.windows-amd64.msi -Wait"
$ vagrant powershell -c '$env:Path += ";C:\Program Files\Go\bin"'
$ vagrant halt
```

### Install Golang on the CentOS VM and halt it
```bash
$ vagrant up centos
$ vagrant ssh centos -c "ls -lha ~/govagrant
$ vagrant ssh centos -c "cd ~/govagrant; sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz"
$ vagrant halt centos
```


### Install Golang on the Ubuntu VM and halt it
```bash
$ vagrant up ubuntu
$ vagrant ssh ubuntu -c "ls -lha ~/govagrant"
$ vagrant ssh ubuntu -c "cd ~/govagrant; sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz"
$ vagrant halt ubuntu
```

## 2) Run the unit test in Local environment (i.e. Mac) - Should PASS
```bash
$ go test -v ./... -tags unit
=== RUN   TestMd5sum
--- PASS: TestMd5sum (0.00s)
PASS
ok      govagrant/info  0.202s
```

## 3) Run integration test in Windows10 - Should PASS
```bash
$ vagrant up windows
$ vagrant powershell -c 'cd \Vagrant; dir; go test -v ./info -tags integration'
$ vagrant halt
```

## 4) Run integration test in centOS - Should PASS
```bash
$ vagrant ssh centos -c "export PATH=$PATH:/usr/local/go/bin; cd ~/govagrant; go test -v ./... -tags integration"
```

## 5) Run integration test in Ubuntu - Should FAIL
```bash
$ vagrant ssh centos -c "export PATH=$PATH:/usr/local/go/bin; cd ~/govagrant; go test -v ./... -tags integration"
```
