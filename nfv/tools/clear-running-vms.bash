#rm /home/arunt/.ssh/known_hosts
#for i in control network compute-1 compute-2 gluster-1 gluster-2 gluster-3
#for i in gluster-1 gluster-2 gluster-3
for i in control network compute-1 compute-2 
do
ssh kent "/usr/bin/virsh destroy $i; /usr/bin/virsh undefine $i"
done

