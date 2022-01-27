package main

import (
  "os"
  "fmt"
  "net"
  "strings"
  "encoding/json"
)

type obs_st struct {

  TempestSn string    `json:"serial_number"`
  RecordType string   `json:"type"`
  HubSn string        `json:"hub_sn"`
  Obs []([]float64)   `json:"obs"`
  FirmwareVersion int `json:"firmware_revision"`

}

type rapid_wind struct {

  TempestSn string  `json:"serial_number"`
  RecordType string `json:"type"`
  HubSn string      `json:"hub_sn"`
  Ob []float64      `json:"ob"`

}

type hub_status struct {

  HubSn string           `json:"serial_number"`
  RecordType string      `json:"type"`
  FirmwareVersion string `json:"firmware_revision"`
  Uptime int             `json:"uptime"`
  RSSI int               `json:"rssi"`
  Timestamp int          `json:"timestamp"`
  ResetFlags string      `json:"reset_flags"`
  Seq int                `json:"seq"`
  RadioStats []int       `json:"radio_stats"`
  MQTTStats []int        `json:"mqtt_stats"`

}

type device_status struct {

  DeviceSn string        `json:"serial_number"`
  RecordType string      `json:"type"`
  HubSn string           `json:"hub_sn"`
  Timestamp int          `json:"timestamp"`
  Uptime int             `json:"uptime"`
  Voltage float64        `json:"voltage"`
  FirmwareVersion int    `json:"firmware_revision"`
  RSSI int               `json:"rssi"`
  HubRSSI int            `json:"hub_rssi"`
  SensorStatus int       `json:"sensor_status"`
  Debug int              `json:"debug"`

}

func main() {

  // listen on port 50222 on all interfaces
  pc, err := net.ListenPacket("udp4", "0.0.0.0:50222")

  if err != nil {
    panic(err)
  }

  defer pc.Close()

  enc := json.NewEncoder(os.Stdout)

  buf := make([]byte, 1024) // this shld be plenty of space for Tempest msgs

  for { // forever

    n, _, err := pc.ReadFrom(buf) // read the broadcast messages

    if err != nil { // it is possible there is an error
      panic(err)
    }

    str1 := string(buf[:n])
    
    if (strings.Contains(str1, "obs_st")) {

      res_obs := obs_st{}
      err := json.Unmarshal(buf[:n], &res_obs)

      if (err != nil) {
        fmt.Println(err)
      }

      // fmt.Println(res_obs)
      enc.Encode(res_obs)

    } else if (strings.Contains(str1, "rapid_wind")) {

      res_wind := rapid_wind{}
      err := json.Unmarshal(buf[:n], &res_wind)

      if (err != nil) {
        fmt.Println(err)
      }

      // fmt.Println(res_wind)
      enc.Encode(res_wind)

    } else if (strings.Contains(str1, "hub_status")) {

      res_hub := hub_status{}
      err := json.Unmarshal(buf[:n], &res_hub)

      if (err != nil) {
        fmt.Println(err)
      }

      // fmt.Println(res_hub)
      enc.Encode(res_hub)


    } else if (strings.Contains(str1, "device_status")) {

      res_dev := device_status{}
      err := json.Unmarshal(buf[:n], &res_dev)

      if (err != nil) {
        fmt.Println(err)
      }

      // fmt.Println(res_dev)
      enc.Encode(res_dev)

    } else {

      fmt.Printf("%s\n", str1) // output to stdout

    }


  }

}
