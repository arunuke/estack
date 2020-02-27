#1: primary maas profile - maas_user
#2: primary dns - maas_dns
#3: proxy - maas_proxy
#4: subnet for maas (CND)
#5: default gateway
#6: ip dynamic range low
#7: ip dynamic range high
#8: name of rack controller, typically hostname

PROFILE=$1
MY_UPSTREAM_DNS=$2
MY_NAMESERVER=$2
EXTERNAL_PROXY=$3
SUBNET_CIDR=$4
MY_GATEWAY=$5
IP_DYNAMIC_RANGE_LOW=$6
IP_DYNAMIC_RANGE_HIGH=$7
PRIMARY_RACK_CONTROLLER=$8
MY_NTP=$9

#SSH_KEY="$9"

VLAN_TAG=0
#TODO: vlan tag in our case, but better to retrieve programmatically
API_KEY_FILE=api.key
API_SERVER=localhost
MAAS_URL=http://$API_SERVER/MAAS/api/2.0

echo " 1: $1 2:$2 3 $3 4:$4 5:$5 6: $6 7: $7 8: $8 9: $9" > bash_output.log
#create api key
sudo maas-region apikey --username=$PROFILE > $API_KEY_FILE

#controller login 
maas login $PROFILE $MAAS_URL - < $API_KEY_FILE

#set upstream DNS
maas $PROFILE maas set-config name=upstream_dns value=$MY_UPSTREAM_DNS
maas $PROFILE maas set-config name=ntp_servers value=$MY_NTP

#set default gateway
maas $PROFILE subnet update $SUBNET_CIDR gateway_ip=$MY_GATEWAY
maas $PROFILE subnet update $SUBNET_CIDR active_discovery=True

#copy key
#maas $PROFILE sshkeys create "key=$SSH_KEY"

#set proxies
maas $PROFILE maas set-config name=enable_http_proxy value=true
maas $PROFILE maas set-config name=http_proxy value=$EXTERNAL_PROXY

#set dnssec to no

maas $PROFILE maas set-config name=dnssec_validation value=no
sudo systemctl restart bind9.service

#import images
maas $PROFILE boot-resources import

#set subnet DNS
maas $PROFILE subnet update $SUBNET_CIDR dns_servers=$MY_GATEWAY

#Identify fabric ID for this subnet
FABRIC_ID=$(maas $PROFILE subnet read $SUBNET_CIDR \
    | grep fabric_id | cut -d ' ' -f 10 | cut -d ',' -f 1)

#set DHCP for this subnet/fabric
#current deployment uses untagged, so vlan_tag is 0

maas $PROFILE ipranges create type=dynamic \
    start_ip=$IP_DYNAMIC_RANGE_LOW end_ip=$IP_DYNAMIC_RANGE_HIGH \

#maas $PROFILE ipranges create type=reserved \
#    start_ip=$IP_STATIC_RANGE_LOW end_ip=$IP_STATIC_RANGE_HIGH \

maas $PROFILE vlan update $FABRIC_ID $VLAN_TAG dhcp_on=True \
    primary_rack=$PRIMARY_RACK_CONTROLLER
