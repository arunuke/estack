IPADDR=10.22.0.15
for i in 1 2 3 4 5 6
do 
  curl https://$IPADDR/redfish/v1/Systems/1/EthernetInterfaces/$i/ --insecure -u admin:HP1nvent -L | json_pp > pci_eth_file_$i.json
done
