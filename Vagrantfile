# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  # Already have a vagrant box you want to use?  
  # Change it here
  config.vm.box = "trusty64"

  # Want a trusty64 basebox?  Uncomment the following line.
  # config.vm.box_url = "https://cloud-images.ubuntu.com/vagrant/trusty/current/trusty-server-cloudimg-amd64-vagrant-disk1.box"
  
  config.vm.network "forwarded_port", guest: 11300, host: 11301
  config.vm.network "private_network", ip: "192.168.63.10"

  config.vm.provider "virtualbox" do |vb|  
    vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
  end

  config.vm.provision :shell, :inline => "apt-get update && apt-get upgrade -y && apt-get autoremove -y"
  config.vm.provision :shell, :inline => "apt-get install -y beanstalkd"
end
