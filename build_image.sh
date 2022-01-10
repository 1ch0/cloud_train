#!/bin/bash
#
#***********************************************************************
#         Author:            ech0
#         Date:              2022-01-07
#         FileName:          build_image.sh
#         Description:       ech0 script
#         Copyright (C): 2022 All rights reserved
#***********************************************************************
# TAG=$1

# docker image build -t nginx:${TAG} ./

images=(
    '1ch0/test:httpserverv0.0.1'
)
for image in ${images[@]}; do
    docker build -t "$image" .
    docker push "$image"
done
