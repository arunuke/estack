PROFILE=arunt
MY_UPSTREAM_DNS=10.255.0.10
EXTERNAL_PROXY=http://16.85.88.10:8080
SUBNET_CIDR=10.11.2.0/24
MY_GATEWAY=10.11.2.11
SSH_KEY=ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC1l5oo8AbP42eHWsp1vbeiQ0Gj6xtrssUhN37YiIzfcT3T19AQqGQuG83/xY4MDnpqEjhzDKJEAOU0Q1/FH/gWzxxdqPP5VPy6OrcgDiebgJ5YkiKzMD2cs8cf2BZyiMT76AjLXTQqmaYUjeJEo+yX59EP+zOI9ydMJ9jb3RiS/ytVeAmmmeX5h/0J8QYxG5Vo6R5SwS2bWobZydPFzghy9rf8QTmsvuADc60XpzJ5QVee8fIx2uR2hy836Hi4wBsW/XUiFHcU6PhncFg2uAPf6HYy0tl3kIR2dEHaniHJMCoexBzW+LI8mh5adCv59GF3UuLiEeQVqMVb9pQG+7uX arunt@potter

#set DNS
maas $PROFILE maas set-config name=upstream_dns value=$MY_UPSTREAM_DNS

#set default gateway
maas $PROFILE subnet update $SUBNET_CIDR gateway_ip=$MY_GATEWAY

#copy key
maas $PROFILE sshkeys create "key=$SSH_KEY"

#set proxies
maas $PROFILE maas set-config name=enable_http_proxy value=false
maas $PROFILE maas set-config name=http_proxy value=$EXTERNAL_PROXY
