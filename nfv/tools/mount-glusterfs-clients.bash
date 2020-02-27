for i in control network compute-1 compute-2 
do
ssh $i 'sudo /bin/mount -t glusterfs gluster-1:vm-vols /mnt/shared -o rw,user'
done

