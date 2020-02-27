virt-install \
    --network=bridge:nest-br \
    --name=nest-1 \
    --disk path=/home/arunt/nest/nest-1.qcow2,format=qcow2,cache=none \
    --disk path=/home/arunt/nest/nest-1-ud.qcow2,format=qcow2,cache=none \
    --ram 512 \
    --vcpus=2 \
    --os-type linux \
    --autostart \
    --boot hd \
    --nographics 
 
