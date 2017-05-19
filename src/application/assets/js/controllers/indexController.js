/**
 * Created by bim on 19.05.2017.
 */


App.config(
    (function (ChartJsProvider) {
        ChartJsProvider.setOptions({ colors : [ '#803690', '#00f', '#DCDCDC', '#46BFBD', '#FDB45C', '#949FB1', '#4D5360'] });
    })
).controller('indexCtrl', function($scope, $http) {
    $scope.labels = ["Registrerte brukere", "Uregistrerte brukere"];
    $scope.labelsGender = ["Menn", "Kvinner"];
    $scope.labelsAge = ["0-12år", "13-18år", "19-25år", "26-33år", "34-42år", "42-60år", "60+år"];
    $scope.data = [74, 26];
    $scope.dataGender = [62, 38];
    $scope.dataAge = [12, 25, 24, 17, 10, 7, 5];
    $scope.colors = [ '#155965', '#72b978', '#DCDCDC', '#46BFBD', '#FDB45C', '#949FB1', '#4D5360'];
    $scope.options = [
        {
            size: {
                height: 300,
                width: 300
            }
        }
        ];
});

