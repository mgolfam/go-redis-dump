#!/bin/bash

# cd /opt/redis2csv/
# make >> log 2>> log &
ps aux | grep redis2csv | awk '{print $2}' | xargs -i kill -9 {}
cd /opt/redis2csv/
make
rm -rf app
exit