#!/bin/sh
while true; do
    echo `date` >> $SNAP_USER_DATA/heartbeats.log
    sleep 5m
done
