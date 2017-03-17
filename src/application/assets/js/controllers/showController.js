/**
 * Created by bim on 17.03.2017.
 */
App.controller('showCtrl', function($scope, $http) {

    $scope.init = function(id) {
         $http.get("/usertest/show?t=" + id).then(function(response) {
             $scope.ut = response.data;
             console.log($scope.ut);
        });

    }

    $scope.getAverage = function(index) {
        var total = 0;
        for(var i = 0; i < $scope.ut.feedback.length; i++ ) {
            total += $scope.ut.feedback[i].answers[index].score;
        }

        var avg = total / $scope.ut.feedback.length;

        return avg
    }

    $scope.getScore = function() {
        var total = 0;
        for(var i = 0; i < $scope.ut.questions.length; i++ ) {
            total += $scope.getAverage(i);
        }

        return total;
    }

    $scope.hasFeedback = function() {
        return $scope.ut.feedback.length > 0
    }

});