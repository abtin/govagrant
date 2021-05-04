Last login: Tue May  4 12:48:11 on ttys001
➜  govagrant git:(main) ✗ export VAGRANT_VAGRANTFILE=$PWD/Vagrantfile.windows
➜  govagrant git:(main) ✗ vagrant box list
generic/centos8      (virtualbox, 3.2.18)
generic/ubuntu1804   (virtualbox, 3.2.18)
xnohat/windows10lite (virtualbox, 1.0.0)
➜  govagrant git:(main) ✗ 
➜  govagrant git:(main) ✗ vagrant up
Bringing machine 'windows' up with 'virtualbox' provider...
==> windows: Box 'xnohat/windows10lite' could not be found. Attempting to find and install...
    windows: Box Provider: virtualbox
    windows: Box Version: 1.0.0
==> windows: Loading metadata for box 'xnohat/windows10lite'
    windows: URL: https://vagrantcloud.com/xnohat/windows10lite
==> windows: Adding box 'xnohat/windows10lite' (v1.0.0) for provider: virtualbox
    windows: Downloading: https://vagrantcloud.com/xnohat/boxes/windows10lite/versions/1.0.0/providers/virtualbox.box
==> windows: Successfully added box 'xnohat/windows10lite' (v1.0.0) for 'virtualbox'!
==> windows: Importing base box 'xnohat/windows10lite'...
==> windows: Matching MAC address for NAT networking...
==> windows: Checking if box 'xnohat/windows10lite' version '1.0.0' is up to date...
==> windows: Setting the name of the VM: govagrant_windows_1620163097527_51017
==> windows: Clearing any previously set network interfaces...
==> windows: Preparing network interfaces based on configuration...
    windows: Adapter 1: nat
==> windows: Forwarding ports...
    windows: 3389 (guest) => 33899 (host) (adapter 1)
    windows: 5985 (guest) => 5985 (host) (adapter 1)
    windows: 5986 (guest) => 55986 (host) (adapter 1)
    windows: 22 (guest) => 2222 (host) (adapter 1)
==> windows: Booting VM...
==> windows: Waiting for machine to boot. This may take a few minutes...
    windows: WinRM address: 127.0.0.1:5985
    windows: WinRM username: vagrant
    windows: WinRM execution_time_limit: PT2H
    windows: WinRM transport: negotiate
==> windows: Machine booted and ready!
==> windows: Checking for guest additions in VM...
    windows: The guest additions on this VM do not match the installed version of
    windows: VirtualBox! In most cases this is fine, but in rare cases it can
    windows: prevent things such as shared folders from working properly. If you see
    windows: shared folder errors, please make sure the guest additions within the
    windows: virtual machine match the version of VirtualBox you have installed on
    windows: your host and reload your VM.
    windows: 
    windows: Guest Additions Version: 5.2.12
    windows: VirtualBox Version: 6.1
==> windows: Mounting shared folders...
    windows: /vagrant => /Users/abtin/1work/projs/govagrant
➜  govagrant git:(main) ✗ vagrant powershell windows -c "dir \vagrant"
    windows: 
    windows: 
    windows: 
    windows: 
    windows:     Directory: C:\vagrant
    windows: 
    windows: 
    windows: 
    windows: 
    windows: 
    windows: Mode                LastWriteTime         Length Name                                                                                                                                                                                                    
    windows: 
    windows: ----                -------------         ------ ----                                                                                                                                                                                                    
    windows: 
    windows: d-----         5/4/2021   8:58 AM                info                                                                                                                                                                                                    
    windows: 
    windows: ------         5/4/2021   9:32 AM             73 go.mod                                                                                                                                                                                                  
    windows: 
    windows: ------         5/4/2021  11:44 AM            262 Vagrantfile.ubuntu                                                                                                                                                                                      
    windows: 
    windows: ------         5/4/2021  11:11 AM           1062 LICENSE                                                                                                                                                                                                 
    windows: 
    windows: ------         5/4/2021   8:58 AM      129021323 go1.16.3.linux-amd64.tar.gz                                                                                                                                                                             
    windows: 
    windows: ------         5/4/2021   8:59 AM           2059 go.sum                                                                                                                                                                                                  
    windows: 
    windows: ------         5/4/2021  11:40 AM            172 Vagrantfile.windows                                                                                                                                                                                     
    windows: 
    windows: ------         5/4/2021  11:50 AM           2985 README.md                                                                                                                                                                                               
    windows: 
    windows: ------         5/4/2021  11:44 AM            259 Vagrantfile.centos                                                                                                                                                                                      
    windows: 
    windows: ------         5/4/2021   8:58 AM      124391424 go1.16.3.windows-amd64.msi                                                                                                                                                                              
    windows: 
    windows: ------         5/4/2021   9:32 AM            203 main.go                                                                                                                                                                                                 
    windows: 
    windows: 
    windows: 
    windows: 
    windows: 
==> windows: Command: dir \vagrant executed successfully with output code 0.
➜  govagrant git:(main) ✗ vagrant powershell windows -c "start-process \vagrant\go1.16.3.windows-amd64.msi -Wait"
==> windows: Command: start-process \vagrant\go1.16.3.windows-amd64.msi -Wait executed successfully with output code 0.
➜  govagrant git:(main) ✗ vagrant powershell windows -c '$env:Path += ";C:\Program Files\Go\bin"'
==> windows: Command: $env:Path += ";C:\Program Files\Go\bin" executed successfully with output code 0.
➜  govagrant git:(main) ✗ vagrant suspend
==> windows: Saving VM state and suspending execution...
➜  govagrant git:(main) ✗ export VAGRANT_VAGRANTFILE=$PWD/Vagrantfile.centos
➜  govagrant git:(main) ✗ vagrant up
Bringing machine 'centos' up with 'virtualbox' provider...
==> centos: Box 'generic/centos8' could not be found. Attempting to find and install...
    centos: Box Provider: virtualbox
    centos: Box Version: 3.2.18
==> centos: Loading metadata for box 'generic/centos8'
    centos: URL: https://vagrantcloud.com/generic/centos8
==> centos: Adding box 'generic/centos8' (v3.2.18) for provider: virtualbox
    centos: Downloading: https://vagrantcloud.com/generic/boxes/centos8/versions/3.2.18/providers/virtualbox.box
    centos: Calculating and comparing box checksum...
==> centos: Successfully added box 'generic/centos8' (v3.2.18) for 'virtualbox'!
==> centos: Importing base box 'generic/centos8'...
==> centos: Matching MAC address for NAT networking...
==> centos: Checking if box 'generic/centos8' version '3.2.18' is up to date...
==> centos: Setting the name of the VM: govagrant_centos_1620163539674_71246
==> centos: Clearing any previously set network interfaces...
==> centos: Preparing network interfaces based on configuration...
    centos: Adapter 1: nat
==> centos: Forwarding ports...
    centos: 22 (guest) => 2222 (host) (adapter 1)
==> centos: Running 'pre-boot' VM customizations...
==> centos: Booting VM...
==> centos: Waiting for machine to boot. This may take a few minutes...
    centos: SSH address: 127.0.0.1:2222
    centos: SSH username: vagrant
    centos: SSH auth method: private key
    centos: 
    centos: Vagrant insecure key detected. Vagrant will automatically replace
    centos: this with a newly generated keypair for better security.
    centos: 
    centos: Inserting generated public key within guest...
    centos: Removing insecure key from the guest if it's present...
    centos: Key inserted! Disconnecting and reconnecting using new SSH key...
==> centos: Machine booted and ready!
==> centos: Checking for guest additions in VM...
    centos: The guest additions on this VM do not match the installed version of
    centos: VirtualBox! In most cases this is fine, but in rare cases it can
    centos: prevent things such as shared folders from working properly. If you see
    centos: shared folder errors, please make sure the guest additions within the
    centos: virtual machine match the version of VirtualBox you have installed on
    centos: your host and reload your VM.
    centos: 
    centos: Guest Additions Version: 5.2.44
    centos: VirtualBox Version: 6.1
==> centos: Mounting shared folders...
    centos: /home/vagrant/govagrant => /Users/abtin/1work/projs/govagrant
➜  govagrant git:(main) ✗ vagrant ssh -c "ls -lha ~/govagrant"
total 249M
drwxr-xr-x. 1 vagrant vagrant  544 May  4 18:47 .
drwx------. 4 vagrant vagrant   91 May  4 21:26 ..
drwxr-xr-x. 1 vagrant vagrant  416 May  4 18:15 .git
-rw-r-----. 1 vagrant vagrant  127 May  4 16:25 .gitattributes
-rw-r--r--. 1 vagrant vagrant   26 May  4 18:14 .gitignore
-rw-r--r--. 1 vagrant vagrant 124M May  4 15:58 go1.16.3.linux-amd64.tar.gz
-rw-r--r--. 1 vagrant vagrant 119M May  4 15:58 go1.16.3.windows-amd64.msi
-rw-r--r--. 1 vagrant vagrant   73 May  4 16:32 go.mod
-rw-r--r--. 1 vagrant vagrant 2.1K May  4 15:59 go.sum
drwxr-xr-x. 1 vagrant vagrant  160 May  4 15:58 info
-rw-r--r--. 1 vagrant vagrant 1.1K May  4 18:11 LICENSE
-rw-r--r--. 1 vagrant vagrant  203 May  4 16:32 main.go
-rw-r--r--. 1 vagrant vagrant 3.0K May  4 18:50 README.md
drwxr-xr-x. 1 vagrant vagrant  128 May  4 17:58 .vagrant
-rw-r--r--. 1 vagrant vagrant  259 May  4 18:44 Vagrantfile.centos
-rw-r--r--. 1 vagrant vagrant  262 May  4 18:44 Vagrantfile.ubuntu
-rw-r--r--. 1 vagrant vagrant  172 May  4 18:40 Vagrantfile.windows
Connection to 127.0.0.1 closed.
➜  govagrant git:(main) ✗ vagrant ssh  -c "cd ~/govagrant; sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz"
Connection to 127.0.0.1 closed.
➜  govagrant git:(main) ✗ vagrant suspend
==> centos: Saving VM state and suspending execution...
➜  govagrant git:(main) ✗ export VAGRANT_VAGRANTFILE=$PWD/Vagrantfile.ubuntu
➜  govagrant git:(main) ✗ vagrant up
Bringing machine 'ubuntu' up with 'virtualbox' provider...
==> ubuntu: Box 'generic/ubuntu1804' could not be found. Attempting to find and install...
    ubuntu: Box Provider: virtualbox
    ubuntu: Box Version: 3.2.18
==> ubuntu: Loading metadata for box 'generic/ubuntu1804'
    ubuntu: URL: https://vagrantcloud.com/generic/ubuntu1804
==> ubuntu: Adding box 'generic/ubuntu1804' (v3.2.18) for provider: virtualbox
    ubuntu: Downloading: https://vagrantcloud.com/generic/boxes/ubuntu1804/versions/3.2.18/providers/virtualbox.box
    ubuntu: Calculating and comparing box checksum...
==> ubuntu: Successfully added box 'generic/ubuntu1804' (v3.2.18) for 'virtualbox'!
==> ubuntu: Importing base box 'generic/ubuntu1804'...
==> ubuntu: Matching MAC address for NAT networking...
==> ubuntu: Checking if box 'generic/ubuntu1804' version '3.2.18' is up to date...
==> ubuntu: Setting the name of the VM: govagrant_ubuntu_1620163775535_77049
==> ubuntu: Clearing any previously set network interfaces...
==> ubuntu: Preparing network interfaces based on configuration...
    ubuntu: Adapter 1: nat
==> ubuntu: Forwarding ports...
    ubuntu: 22 (guest) => 2222 (host) (adapter 1)
==> ubuntu: Running 'pre-boot' VM customizations...
==> ubuntu: Booting VM...
==> ubuntu: Waiting for machine to boot. This may take a few minutes...
    ubuntu: SSH address: 127.0.0.1:2222
    ubuntu: SSH username: vagrant
    ubuntu: SSH auth method: private key
    ubuntu: 
    ubuntu: Vagrant insecure key detected. Vagrant will automatically replace
    ubuntu: this with a newly generated keypair for better security.
    ubuntu: 
    ubuntu: Inserting generated public key within guest...
    ubuntu: Removing insecure key from the guest if it's present...
    ubuntu: Key inserted! Disconnecting and reconnecting using new SSH key...
==> ubuntu: Machine booted and ready!
==> ubuntu: Checking for guest additions in VM...
    ubuntu: The guest additions on this VM do not match the installed version of
    ubuntu: VirtualBox! In most cases this is fine, but in rare cases it can
    ubuntu: prevent things such as shared folders from working properly. If you see
    ubuntu: shared folder errors, please make sure the guest additions within the
    ubuntu: virtual machine match the version of VirtualBox you have installed on
    ubuntu: your host and reload your VM.
    ubuntu: 
    ubuntu: Guest Additions Version: 5.2.42
    ubuntu: VirtualBox Version: 6.1
==> ubuntu: Mounting shared folders...
    ubuntu: /home/vagrant/govagrant => /Users/abtin/1work/projs/govagrant
➜  govagrant git:(main) ✗ vagrant ssh -c "ls -lha ~/govagrant"
total 249M
drwxr-xr-x 1 vagrant vagrant  544 May  4 18:47 .
drwxr-xr-x 6 vagrant vagrant 4.0K May  4 21:29 ..
drwxr-xr-x 1 vagrant vagrant  416 May  4 18:15 .git
-rw-r----- 1 vagrant vagrant  127 May  4 16:25 .gitattributes
-rw-r--r-- 1 vagrant vagrant   26 May  4 18:14 .gitignore
-rw-r--r-- 1 vagrant vagrant 124M May  4 15:58 go1.16.3.linux-amd64.tar.gz
-rw-r--r-- 1 vagrant vagrant 119M May  4 15:58 go1.16.3.windows-amd64.msi
-rw-r--r-- 1 vagrant vagrant   73 May  4 16:32 go.mod
-rw-r--r-- 1 vagrant vagrant 2.1K May  4 15:59 go.sum
drwxr-xr-x 1 vagrant vagrant  160 May  4 15:58 info
-rw-r--r-- 1 vagrant vagrant 1.1K May  4 18:11 LICENSE
-rw-r--r-- 1 vagrant vagrant  203 May  4 16:32 main.go
-rw-r--r-- 1 vagrant vagrant 3.0K May  4 18:50 README.md
drwxr-xr-x 1 vagrant vagrant  128 May  4 17:58 .vagrant
-rw-r--r-- 1 vagrant vagrant  259 May  4 18:44 Vagrantfile.centos
-rw-r--r-- 1 vagrant vagrant  262 May  4 18:44 Vagrantfile.ubuntu
-rw-r--r-- 1 vagrant vagrant  172 May  4 18:40 Vagrantfile.windows
Connection to 127.0.0.1 closed.
➜  govagrant git:(main) ✗ vagrant ssh -c "cd ~/govagrant; sudo tar -C /usr/local -xzf go1.16.3.linux-amd64.tar.gz"
Connection to 127.0.0.1 closed.
➜  govagrant git:(main) ✗ vagrant suspend
==> ubuntu: Saving VM state and suspending execution...
➜  govagrant git:(main) ✗ go test -v ./... -tags unit
?   	govagrant	[no test files]
=== RUN   TestMd5sum
--- PASS: TestMd5sum (0.00s)
PASS
ok  	govagrant/info	(cached)
➜  govagrant git:(main) ✗ vagrant resume windows
The machine with the name 'windows' was not found configured for
this Vagrant environment.
➜  govagrant git:(main) ✗ export VAGRANT_VAGRANTFILE=$PWD/Vagrantfile.windows
➜  govagrant git:(main) ✗ vagrant resume
==> windows: Resuming suspended VM...
==> windows: Booting VM...
==> windows: Waiting for machine to boot. This may take a few minutes...
    windows: WinRM address: 127.0.0.1:5985
    windows: WinRM username: vagrant
    windows: WinRM execution_time_limit: PT2H
    windows: WinRM transport: negotiate
==> windows: Machine booted and ready!
==> windows: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> windows: flag to force provisioning. Provisioners marked to run always will still run.
➜  govagrant git:(main) ✗ vagrant powershell -c 'cd \Vagrant; go test -v ./info -tags integration' 
    windows: === RUN   TestNewOsInfo_windows10
    windows: 
    windows: --- PASS: TestNewOsInfo_windows10 (0.00s)
    windows: 
    windows: === RUN   TestNewOsInfo_linux_ubuntu
    windows: 
    windows:     osinfo_integration_test.go:36: Skipping TestNewOsInfo_linux_ubuntu on a non linux host
    windows: 
    windows: --- PASS: TestNewOsInfo_linux_ubuntu (0.00s)
    windows: 
    windows: PASS
    windows: 
    windows: ok  	govagrant/info	(cached)
    windows: 
==> windows: Command: cd \Vagrant; go test -v ./info -tags integration executed successfully with output code 0.
➜  govagrant git:(main) ✗ vagrant suspend
==> windows: Saving VM state and suspending execution...
➜  govagrant git:(main) ✗ export VAGRANT_VAGRANTFILE=$PWD/Vagrantfile.centos
➜  govagrant git:(main) ✗ vagrant resume
==> centos: Resuming suspended VM...
==> centos: Booting VM...
==> centos: Waiting for machine to boot. This may take a few minutes...
    centos: SSH address: 127.0.0.1:2222
    centos: SSH username: vagrant
    centos: SSH auth method: private key
==> centos: Machine booted and ready!
==> centos: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> centos: flag to force provisioning. Provisioners marked to run always will still run.
➜  govagrant git:(main) ✗ vagrant ssh centos -c "export PATH=$PATH:/usr/local/go/bin; cd ~/govagrant; go test -v ./... -tags integration"
go: downloading github.com/shirou/gopsutil/v3 v3.21.4
go: downloading golang.org/x/sys v0.0.0-20210217105451-b926d437f341
?   	govagrant	[no test files]
=== RUN   TestNewOsInfo_windows10
    osinfo_integration_test.go:14: Skipping TestNewOsInfo_windows10 on a non windows host
--- PASS: TestNewOsInfo_windows10 (0.00s)
=== RUN   TestNewOsInfo_linux_ubuntu
--- PASS: TestNewOsInfo_linux_ubuntu (0.00s)
PASS
ok  	govagrant/info	0.003s
Connection to 127.0.0.1 closed.
➜  govagrant git:(main) ✗ vagrant suspend                                   
==> centos: Saving VM state and suspending execution...
➜  govagrant git:(main) ✗ export VAGRANT_VAGRANTFILE=$PWD/Vagrantfile.ubuntu
➜  govagrant git:(main) ✗ vagrant resume                                    
==> ubuntu: Resuming suspended VM...
==> ubuntu: Booting VM...
==> ubuntu: Waiting for machine to boot. This may take a few minutes...
    ubuntu: SSH address: 127.0.0.1:2222
    ubuntu: SSH username: vagrant
    ubuntu: SSH auth method: private key
==> ubuntu: Machine booted and ready!
==> ubuntu: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> ubuntu: flag to force provisioning. Provisioners marked to run always will still run.
➜  govagrant git:(main) ✗ vagrant ssh centos -c "export PATH=$PATH:/usr/local/go/bin; cd ~/govagrant; go test -v ./... -tags integration"
The machine with the name 'centos' was not found configured for
this Vagrant environment.
➜  govagrant git:(main) ✗ vagrant ssh  -c "export PATH=$PATH:/usr/local/go/bin; cd ~/govagrant; go test -v ./... -tags integration" 
go: downloading github.com/shirou/gopsutil/v3 v3.21.4
go: downloading golang.org/x/sys v0.0.0-20210217105451-b926d437f341
?   	govagrant	[no test files]
=== RUN   TestNewOsInfo_windows10
    osinfo_integration_test.go:14: Skipping TestNewOsInfo_windows10 on a non windows host
--- PASS: TestNewOsInfo_windows10 (0.00s)
=== RUN   TestNewOsInfo_linux_ubuntu
    osinfo_integration_test.go:51: 
        wanted:	Platform: centos, Family: rhel, Version: 8.3.2011
        got:	Platform: ubuntu, Family: debian, Version: 18.04
--- FAIL: TestNewOsInfo_linux_ubuntu (0.00s)
FAIL
FAIL	govagrant/info	0.003s
FAIL
Connection to 127.0.0.1 closed.
➜  govagrant git:(main) ✗ vagrant suspend
==> ubuntu: Saving VM state and suspending execution...
