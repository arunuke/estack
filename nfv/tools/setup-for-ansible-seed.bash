#setting up a host for ansible


#for i in gluster-1 gluster-2 gluster-3
#for i in control network compute-1 compute-2 gluster-1 gluster-2 gluster-3
#for i in control network compute-1 compute-2 
#for i in jordan
for i in neo

do
	ssh $i "sudo /usr/bin/apt-get -y install python2.7; sudo /bin/ln -s /usr/bin/python2.7 /usr/bin/python" 
	ansible $i -m command -a "whoami"
done
