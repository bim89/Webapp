/**
 * Created by pedjo on 11-May-17.
 */
function initMap() {
    var map = new google.maps.Map(document.getElementById('map'), {
        zoom: 15,
        center: {lat: 59.916849, lng: 10.748032}
    });
    var geocoder = new google.maps.Geocoder();

    document.getElementById('submit').addEventListener('click', function() {
        geocodeAddress(geocoder, map);
    });
}

function geocodeAddress(geocoder, resultsMap) {
    var address = document.getElementById('address').value;
    geocoder.geocode({'address': address}, function(results, status) {
        if (status === 'OK') {
            resultsMap.setCenter(results[0].geometry.location);
            var marker = new google.maps.Marker({
                map: resultsMap,
                position: results[0].geometry.location

            });
            var lat = marker.getPosition().lat();
            var lng = marker.getPosition().lng();
            document.getElementById("title").value = document.getElementById("address").value;
            document.getElementById("latitude").value = lat;
            document.getElementById("longitude").value = lng;
        } else {
            alert("Geocode could not get an address" + status);
        }
    });
}