/**
 * Created by bim on 28.02.2017.
 */
App.controller('createCtrl', function($scope, $http) {


    $scope.form = {"title": "", "latitude": 0, "longitude": 0};

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

    $scope.createUserTest = function() {
        if ($scope.questions.length >= 1 && $scope.questions[0].question.length > 0 && $scope.form.title.length > 1) {
            data = {"title": $scope.form.title, "email": $scope.email, "latitude": $scope.form.latitude, "longitude": $scope.form.longitude, "questions": $scope.questions};
            console.log(data);
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