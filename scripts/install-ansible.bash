#Script to install ansible on 18.04
#Will be run as a "shell" provisioner from inside a vagrantfile

#Update the system
sudo apt update

#Configure common properties
sudo apt install software-properties-common

#Add the right ansible repos
sudo apt-add-repository ppa:ansible/ansible

#Update the system again
sudo apt update

#Install ansible
sudo apt-get -y install ansible 
