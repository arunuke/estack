

#Add line to source
echo 'deb https://download.virtualbox.org/virtualbox/debian bionic contrib' | sudo tee -a /etc/apt/sources.list

#Add keys
wget -q https://www.virtualbox.org/download/oracle_vbox_2016.asc -O- | sudo apt-key add -
wget -q https://www.virtualbox.org/download/oracle_vbox.asc -O- | sudo apt-key add -

#Install virtualbox
sudo apt-get update
sudo apt-get install virtualbox-6.1
