for i in control network compute-1 compute-2
do
ssh $i '/bin/echo $HOSTNAME;/bin/ls -larth /mnt/shared'
done

