#setting up a host for ansible

	ssh $1 "sudo /usr/bin/apt-get -y install python2.7; sudo /bin/ln -s /usr/bin/python2.7 /usr/bin/python" 
	ansible $1 -m command -a "whoami"
