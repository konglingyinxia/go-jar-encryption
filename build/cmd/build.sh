#!/bin/bash

win() {
  echo "win 客户端编译"


  echo "win 客户端编译结束........."

}

linux() {
  echo "linux 客户端编译"


  echo "linux 客户端编译结束........"
}

usage() {
  echo "Usage: sh 执行脚本.sh [win|linux]"
  exit 1
}

case $1 in

"win")
  win
  ;;
"linux")
  linux
  ;;
*)
  usage
  ;;
esac
