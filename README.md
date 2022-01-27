# go-weatherflow

```bash
$ go build
$ ./go-weatherflow
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245968,0.00,0]}
{"serial_number":"HB-00069665","type":"hub_status","firmware_revision":"177","uptime":728603,"rssi":-50,"timestamp":1643245971,"reset_flags":"BOR,PIN,POR","seq":72783,"radio_stats":[25,1,0,3,16637],"mqtt_stats":[10,108]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245974,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245977,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245980,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245983,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245986,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245989,0.00,0]}
{"serial_number":"HB-00069665","type":"hub_status","firmware_revision":"177","uptime":728623,"rssi":-49,"timestamp":1643245991,"reset_flags":"BOR,PIN,POR","seq":72785,"radio_stats":[25,1,0,3,16637],"mqtt_stats":[10,108]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245992,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245995,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245998,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246001,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246010,0.00,0]}
{"serial_number":"HB-00069665","type":"hub_status","firmware_revision":"177","uptime":728643,"rssi":-50,"timestamp":1643246011,"reset_flags":"BOR,PIN,POR","seq":72787,"radio_stats":[25,1,0,3,16637],"mqtt_stats":[10,108]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246013,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246015,0.00,0]}
{"serial_number":"ST-00055227","type":"device_status","hub_sn":"HB-00069665","timestamp":1643246016,"uptime":106625,"voltage":2.683,"firmware_revision":165,"rssi":-72,"hub_rssi":-66,"sensor_status":655364,"debug":0}
{"serial_number":"ST-00055227","type":"obs_st","hub_sn":"HB-00069665","obs":[[1643246016,0.00,0.00,0.00,0,3,1024.56,-12.82,47.84,0,0.00,0,0.000000,0,0,0,2.683,1]],"firmware_revision":165}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246025,0.39,308]}
```

OR

```bash
$ go get github.com/hrbrmstr/go-weatherflow/...
$ $GOPATH/bin/go-weatherflow
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245968,0.00,0]}
{"serial_number":"HB-00069665","type":"hub_status","firmware_revision":"177","uptime":728603,"rssi":-50,"timestamp":1643245971,"reset_flags":"BOR,PIN,POR","seq":72783,"radio_stats":[25,1,0,3,16637],"mqtt_stats":[10,108]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245974,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245977,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245980,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245983,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245986,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245989,0.00,0]}
{"serial_number":"HB-00069665","type":"hub_status","firmware_revision":"177","uptime":728623,"rssi":-49,"timestamp":1643245991,"reset_flags":"BOR,PIN,POR","seq":72785,"radio_stats":[25,1,0,3,16637],"mqtt_stats":[10,108]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245992,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245995,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643245998,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246001,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246010,0.00,0]}
{"serial_number":"HB-00069665","type":"hub_status","firmware_revision":"177","uptime":728643,"rssi":-50,"timestamp":1643246011,"reset_flags":"BOR,PIN,POR","seq":72787,"radio_stats":[25,1,0,3,16637],"mqtt_stats":[10,108]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246013,0.00,0]}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246015,0.00,0]}
{"serial_number":"ST-00055227","type":"device_status","hub_sn":"HB-00069665","timestamp":1643246016,"uptime":106625,"voltage":2.683,"firmware_revision":165,"rssi":-72,"hub_rssi":-66,"sensor_status":655364,"debug":0}
{"serial_number":"ST-00055227","type":"obs_st","hub_sn":"HB-00069665","obs":[[1643246016,0.00,0.00,0.00,0,3,1024.56,-12.82,47.84,0,0.00,0,0.000000,0,0,0,2.683,1]],"firmware_revision":165}
{"serial_number":"ST-00055227","type":"rapid_wind","hub_sn":"HB-00069665","ob":[1643246025,0.39,308]}
```
