Vagrant.configure("2") do |config|

  config.vm.define "windows" do |windows|
    config.vm.box = "xnohat/windows10lite"
    config.vm.box_version = "1.0.0"
  end


  config.vm.define "centos" do |centos|
    config.vm.box = "generic/centos8"
    config.vm.box_version = "3.2.18"
    config.vm.synced_folder "/Users/abtin/1work/projs/govagrant", "/home/vagrant/govagrant"

  end

  config.vm.define "ubuntu" do |ubuntu|
    config.vm.box = "generic/ubuntu1804"
    config.vm.box_version = "3.2.18"
    config.vm.synced_folder "/Users/abtin/1work/projs/govagrant", "/home/vagrant/govagrant"

  end

end
