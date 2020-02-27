brick_path=/mnt/gluster-data/vms
pvt_key="-i ~/.ssh/sesame_rsa"
for i in gluster-1 gluster-2 gluster-3
do
ssh $pvt_key $i "sudo setfattr -x trusted.glusterfs.volume-id $brick_path"
ssh $pvt_key $i "sudo setfattr -x trusted.gfid $brick_path"
ssh $pvt_key $i "sudo rm -rf $brick_path/.glusterfs"
ssh $pvt_key $i "sudo systemctl restart glusterfs-server.service"
done

