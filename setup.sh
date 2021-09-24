#docker run --network host -it tcpkali:latest
#!/bin/bash

NUM_INSTANCES=$1
NUM_CONNECTIONS=$2
GATEWAY_IP=$3 #172.21.1.1
MSG="aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

for (( c=0; c<${NUM_INSTANCES}; c++ ))
do
  docker run -l c10m --rm --network c10m-bridge -d tcpkali:latest /app/tcpkali -c $NUM_CONNECTIONS -w 8 -T 3600 $GATEWAY_IP:3540
done
