#setting up a host for ansible

KEY_FILE=~/.ssh/sesame_rsa
for i in `cat ./hosts.file`

do
	ssh $i -i $KEY_FILE  "sudo /usr/bin/apt-get -y install python2.7; sudo /bin/ln -s /usr/bin/python2.7 /usr/bin/python" 
	ansible $i -m command -a "whoami"
done
