#Variables
BF_BACKING_DIR=/home/arunt/backingstore
BF_BACKING_FILE=$BF_BACKING_DIR/glusterfile
BF_BLOCKSIZE=1k
BF_COUNT=4096
BF_FS=ext4


LO_LOOPBACK_DEVICE=/dev/loop1

DA_DATA_DIR=/data1/opensaf

GL_REPLICA=3
GL_NODE1=vsa1
GL_NODE2=vsa2
GL_NODE3=vsa3
GL_VOL_NAME=ctrlplane
GL_DATASTORE=/mnt/gluster_data

install_gluster()
{

apt-get install glusterfs-server
service glusterfs-server restart


}


setup_loopback_fs()
{
#setup loopback filesystem

mkdir -p $BF_BACKING_DIR
#create backing file
dd if=/dev/zero of=$BF_BACKING_FILE bs=$BF_BLOCKSIZE count=$BF_COUNT

#use backing file as loopback filesystem
losetup $LO_LOOPBACK_DEVICE $BF_BACKING_FILE

mkdir -p $GL_DATASTORE

mkfs -t $BF_FS $LO_LOOPBACK_DEVICE

mount -t $BF_FS $BF_BACKING_FILE $GL_DATASTORE

#add entry in fstab to keep it persistent across boot
}

configure_gluster()
{

gluster peer probe $GL_NODE2 $GL_NODE3
gluster volume create $GL_VOL_NAME replica $GL_REPLICA transport tcp $GL_NODE1:$GL_DATASTORE $GL_NODE2:$GL_DATASTORE $GL_NODE3:$GL_DATASTORE force


}

heal_gluster()
{

#remove the failed brick by reducing replica count
#restart service on the failed node
#add the same brick by increasing replica count
#perform heal action
#create a new file that performs a sync operation

echo "nothing"
}

setup_loopback_fs
install_gluster

#configure_gluster
