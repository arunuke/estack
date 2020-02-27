PROFILE=arunt
MARCH="amd64/generic"
MPOWER="ipmi"
MPOWERUSER="admin"
MPOWERPASS="HP1nvent"
MPOWERDRV="LAN_2_0"

maas arunt machines create \
  architecture=$MARCH \
  mac_addresses="3c:a8:2a:0e:58:40" \
  power_type=$MPOWER \
  power_parameters_power_address="10.11.0.3" \
  power_parameters_power_user=$MPOWERUSER \
  power_parameters_power_pass=$MPOWERPASS \
  power_parameters_power_driver=$MPOWERDRV
