#setting up a host for ansible
#$1 is host, $2 is key file - cerebrus_rsa, $3 is ssh home directory /home/arunt/.ssh

	scp $3/$2* $1:$3/; ssh $1 "/bin/chmod 600 $3/$2*"; scp $3/config $1:$3/;
