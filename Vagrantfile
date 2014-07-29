# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box = "trusty64"
  config.vm.network "forwarded_port", guest: 11300, host: 11301
  config.vm.network "private_network", ip: "192.168.63.10"

  config.vm.provider "virtualbox" do |vb|  
    vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
  end

  config.vm.provision :shell, :inline => "apt-get update && apt-get upgrade -y && apt-get autoremove -y"
  config.vm.provision :shell, :inline => "apt-get install -y beanstalkd"
end
