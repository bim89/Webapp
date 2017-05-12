/**
 * Created by pedjo on 11-May-17.
 */

var gmarkers = [];

function initMap() {
    var map = new google.maps.Map(document.getElementById('map'), {
        zoom: 15,
        center: {lat: 59.916849, lng: 10.748032}
    });
    var geocoder = new google.maps.Geocoder();

    document.getElementById('submit').addEventListener('click', function() {
        geocodeAddress(geocoder, map);
    });
    map.addListener('click', function(e) {
        placeMarker(e.latLng, map);
    });
}

function geocodeAddress(geocoder, resultsMap) {
    var address = document.getElementById('address').value;
    removeMarkers();
    geocoder.geocode({'address': address}, function(results, status) {
        if (status === 'OK') {
            resultsMap.setCenter(results[0].geometry.location);
            var marker = new google.maps.Marker({
                map: resultsMap,
                position: results[0].geometry.location

            });
            gmarkers.push(marker);
            var lat = marker.getPosition().lat();
            var lng = marker.getPosition().lng();
            document.getElementById("title").value = document.getElementById("address").value;
            document.getElementById("latitude").value = lat;
            document.getElementById("longitude").value = lng;
        } else {
            alert("Venligst skriv inn en mer spesifikk adresse");
        }
    });
}
function removeMarkers(){
    for(i=0; i<gmarkers.length; i++){
        gmarkers[i].setMap(null);
    }
}

function placeMarker(latLng, map) {
    removeMarkers();
    var marker = new google.maps.Marker({
        position: latLng,
        map: map
    });
    gmarkers.push(marker);
    map.panTo(latLng);
    var lat = marker.getPosition().lat();
    var lng = marker.getPosition().lng();
    document.getElementById("title").value = document.getElementById("address").value;
    document.getElementById("latitude").value = lat;
    document.getElementById("longitude").value = lng;
}
