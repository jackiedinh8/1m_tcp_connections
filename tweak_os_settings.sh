#!/bin/bash

echo 2621440 >  /proc/sys/net/netfilter/nf_conntrack_max

#cat /proc/sys/fs/file-max
sysctl -w fs.file-max=4000000
sysctl net.ipv4.ip_local_port_range="15000 61000"
sysctl net.ipv4.tcp_fin_timeout=30
#sysctl net.ipv4.tcp_tw_recycle=1
sysctl net.ipv4.tcp_tw_reuse=1
sysctl net.core.somaxconn=1024
sysctl net.core.netdev_max_backlog=2000
sysctl net.ipv4.tcp_max_syn_backlog=2048

