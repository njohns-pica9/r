# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure(2) do |config|
  config.vm.box = "trusty64"
  config.vm.network "forwarded_port", guest: 8080, host: 28000
  config.vm.network "forwarded_port", guest: 28015, host: 28015
  config.vm.network "private_network", ip: "192.168.33.10"
  config.vm.synced_folder "./data", "/vagrant_data", type: "nfs"
  config.vm.provision :shell, :inline => 'echo "deb http://download.rethinkdb.com/apt trusty main" | tee /etc/apt/sources.list.d/rethinkdb.list'
  config.vm.provision :shell, :inline => "wget -qO- http://download.rethinkdb.com/apt/pubkey.gpg | sudo apt-key add -"
  config.vm.provision :shell, :inline => "apt-get update"
  config.vm.provision :shell, :inline => "apt-get install -y build-essential rethinkdb"
end
