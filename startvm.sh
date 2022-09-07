#!/bin/bash

export VM_NAME="windows10"
export UUID=$(uuidgen)
export MAC=$(hexdump -n 6 -ve '1/1 "%.2x "' /dev/random | awk -v a="2,6,a,e" -v r="$RANDOM" 'BEGIN{srand(r);}NR==1{split(a,b,",");r=int(rand()*4+1);printf "%s%s:%s:%s:%s:%s:%s\n",substr($1,0,1),b[r],$2,$3,$4,$5,$6}')
export IMAGE="/home/rickylss/repo/win10.qcow2"

envsubst < vm.xml.tmp > vm.xml
virsh create vm.xml
virsh list --all
virsh domifaddr ${VM_NAME}

ssh Admin@