#Add line to source
echo "deb https://download.virtualbox.org/virtualbox/debian bionic contrib" >> /etc/apt/sources.list

#Download virtualbox code
wget https://download.virtualbox.org/virtualbox/6.0.12/virtualbox-6.0_6.0.12-133076~Ubuntu~bionic_amd64.deb

#Add keys
sudo apt-key add oracle_vbox_2016.asc

#Install virtualbox
sudo apt-get update
sudo apt-get install virtualbox-6.0
