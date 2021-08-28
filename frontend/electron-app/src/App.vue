<template>
  <div id="app"></div>
</template>

<script>
import "leaflet/dist/leaflet.css";
import L from "leaflet";

delete L.Icon.Default.prototype._getIconUrl;

L.Icon.Default.mergeOptions({
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
  shadowUrl: require("leaflet/dist/images/marker-shadow.png"),
});

export default {
  data: function () {
    return {
      lat: 35.6825,
      lng: 139.752778,
      gpsMarker: null,
      redMarker: {
        icon: L.divIcon({ className: "red marker", iconSize: [16, 16] }),
      },
      map: null,
      message: "Hello Vue.js!",
    };
  },
  created: function () {
    var url = "http://localhost:3000/gps-viewer/v1/gps/00000";
    var http = require("http");
    var data = [];
    http.get(url, function (res) {
      res
        .on("data", function (chunk) {
          data.push(chunk);
        })
        .on("end", function () {
          var events = Buffer.concat(data);
          console.log(events);
        });
    });
    //var res = fetch(url, {
    //  method: "GET",
    // mode: "no-cors",
    //});
    //console.log(res);

    setInterval(() => {
      // this.lat += 0.001;
      // console.log(this.lat);
      // this.$forceUpdate();
      "location.reload()";
    }, 1000);
  },
  mounted: function () {
    this.map = L.map("app", {
      center: L.latLng(35.6825, 139.752778),
      zoom: 15,
    }).addLayer(L.tileLayer("http://{s}.tile.osm.org/{z}/{x}/{y}.png"));
    L.marker([this.lat, 139.752778], this.redMarker)
      .addTo(this.map)
      .bindPopup("Device Position")
      .openPopup();
  },
  beforeUpdated() {
    if (this.gpsMarker == null) {
      var popupContent = "Device Position";
      this.gpsMarker = L.marker([this.lat, this.lng], this.redMarker).addTo(
        this.map
      );
      this.gpsMarker.bindPopup(popupContent).openPopup();
    } else {
      this.gpsMarker.getPopup().setContent(popupContent);
      this.gpsMarker.setLatLng([this.lat, this.lng]);
    }
    L.marker([this.lat, 139.752778], this.redMarker)
      .addTo(this.map)
      .bindPopup("Device Position")
      .openPopup();
  },
  methods: {
    getUrl() {
      this.message = this.message.split("").reverse().join("");
    },
  },
};

//setInterval("location.reload()", 1000);
</script>

<style>
html,
body,
#app {
  height: 100%;
}
body {
  margin: 0;
}
</style>
