PROFILE=arunt
SYSID=4qkhkt
#OTODO: Programmitcally extract sysid for a host
maas $PROFILE machines allocate system_id=$SYSID
maas $PROFILE machine allocate system_id=$SYSID
