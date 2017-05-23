/**
 * Created by bim on 28.02.2017.
 */
App.controller('createCtrl', function($scope, $http) {


    $scope.form = {"title": "", "latitude": 0, "longitude": 0, "distance": 50};

    $scope.questions = [{"question": "", "type": "", "choices": [], "amount": 2}];


    $scope.init = function (email) {
        $scope.email = email
    }


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

    $scope.setFromMap = function (title, lat, lng) {
        $scope.form.title = title;
        $scope.form.latitude = lat;
        $scope.form.longitude = lng;
        console.log($scope.form.title);
    }

    $scope.createUserTest = function() {
        if ($scope.questions.length >= 1 && $scope.questions[0].question.length > 0 && $scope.form.title.length > 1 && $scope.form.distance > 20 && $scope.form.distance < 200) {
            data = {"title": $scope.form.title, "email": $scope.email, "latitude": $scope.form.latitude, "longitude": $scope.form.longitude, "questions": $scope.questions, "distance": $scope.form.distance};
            $http.post("/usertest/create", data);
            alert("You have just created your user test");
        } else {
            alert("Some required fields are empty");
        }
    }

    $scope.addChoices = function(index) {
        $scope.questions[index].choices = [];
        for(var i = 0; i < $scope.questions[index].amount; i++) {
            answer = {"answer": ""};
            $scope.questions[index].choices.push(answer);
        }
    }
});