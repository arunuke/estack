sudo /sbin/dhclient ens3; sudo /sbin/dhclient ens8; 
sudo bash -c 'echo "Acquire::http::Proxy \"http://16.85.88.10:8080\";" > /etc/apt/apt.conf';
sudo apt-get install python2.7
sudo ln -s /usr/bin/python2.7 /usr/bin/python
