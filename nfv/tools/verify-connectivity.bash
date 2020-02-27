
for line in `cat hosts.file`
do
ping $line -c 3
done

