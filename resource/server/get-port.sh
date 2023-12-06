#!/bin/sh

netstat -tlnp | grep xray | grep ::: | awk  '{print $4}' | sed s/://g
