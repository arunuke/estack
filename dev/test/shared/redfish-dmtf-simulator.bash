#Script to install and configure Redfish DMTF simulator
#

REDFISH_DMTF_EMULATOR_GIT_REPO=https://github.com/DMTF/Redfish-Interface-Emulator.git
BASE_DIR=
PIP_DIR=Redfish-Interface-Emulator/packageSets/
PIP_FILE=Env-Local-Python3.5.2_requirements.txt

#Refresh the current system
sudo apt-get update

#Upgrade all packages
sudo apt-get upgrade

#Install required packages
sudo apt-get install python-pip

#Do a git clone of the DMTF directory
git clone $REDFISH_DMTF_EMULATOR_GIT_REPO

#Install dependencies
pip install -r $PIP_DIR/$PIP_FILE

#Generate crt keys
openssl genrsa 2048 > host.key
chmod 400 host.key
openssl req -new -x509 -nodes -sha256 -days 365 -key host.key -out host.cert


