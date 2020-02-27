#Install vagrant on the host 

VAGRANT_REPO=https://releases.hashicorp.com/vagrant/2.2.7/vagrant_2.2.7_linux_amd64.zip
VAGRANT_FILE=vagrant.zip
wget $VAGRANT_REPO -O $VAGRANT_FILE
sudo apt-get install zip unzip
unzip $VAGRANT_FILE
sudo mv vagrant /usr/local/bin
rm $VAGRANT_FILE

