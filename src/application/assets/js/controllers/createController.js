/**
 * Created by bim on 28.02.2017.
 */
App.controller('createCtrl', function($scope, $http) {
    $scope.questions = [{"question": "", "type": ""}];



    $scope.addQuestion = function() {
        if ($scope.questions[$scope.questions.length - 1].question) {
            $scope.questions.push({"question": "", "type": ""});
        } else {
            alert("Some of your fields have an empty question...");
        }
    }

    $scope.removeQuestion = function(index) {
        if ($scope.questions.length > 1) {
            $scope.questions.splice(index, 1);
        } else {
            alert("You only have one question");
        }
    }

    $scope.createUserTest = function() {
        if ($scope.questions.length >= 1 && $scope.questions[0].question.length > 0) {
            data = {"title": $scope.title, "latitude": $scope.latitude, "longitude": $scope.longitude, "questions": $scope.questions};
            $http.post("/usertest/create", data);
            alert("You have just created your user test");
        } else {
            alert("Your form is empty");
        }
    }

});