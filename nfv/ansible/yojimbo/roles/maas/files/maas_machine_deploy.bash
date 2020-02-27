PROFILE=arunt
HOSTNAME=$1
SYSID=$(maas $PROFILE machines read hostname=$HOSTNAME | grep -m1 system_id | cut -d ':' -f2 | cut -d '"' -f2)
maas $PROFILE machines allocate system_id=$SYSID
sleep 5
maas $PROFILE machine deploy $SYSID

