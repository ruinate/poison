#!/bin/bash

if [ -z "$1" ]; then
  echo
  echo usage: $0 [host]
  echo
  echo
  echo show snmp 1、2、3
  exit
fi
HOST="$1"
echo 'snmp-1'
snmpwalk -v 1 -c public $HOST 1.3.6.1.2.1.1.1
echo 'snmp-2'
snmpwalk -v 2c -c public $HOST 1.3.6.1.2.1.1.1
echo 'snmp-1 and 2'
snmpwalk -v 2c -c public $HOST 1.3.6.1.2.1.1.1
echo 'snmp-3'
snmpwalk -v 3 -u bolean -a SHA -A admin123 -x AES -X admin123 -l authPriv $HOST 1.3.6.1.2.1.1.1
