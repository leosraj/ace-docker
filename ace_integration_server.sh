#!/bin/bash

# © Copyright IBM Corporation 2018.
#
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Eclipse Public License v2.0
# which accompanies this distribution, and is available at
# http://www.eclipse.org/legal/epl-v20.html

if [ -z "$MQSI_VERSION" ]; then
  source /opt/ibm/ace-11/server/bin/mqsiprofile
fi

if [ -s /home/aceuser/ace-server/odbc.ini ]; then
  export ODBCINI=/home/aceuser/ace-server/odbc.ini
fi

if ! [[ -z "${ACE_HTTP_ROUTE_NAME}" ]] || ! [[ -z "${ACE_HTTPS_ROUTE_NAME}" ]] ; then
  . /home/aceuser/portOverrides
elif ! [[ -z "${KUBERNETES_PORT}" ]] && ! [[ -z "${SERVICE_NAME}" ]] ; then
  . /home/aceuser/portOverrides
fi

exec IntegrationServer $*
