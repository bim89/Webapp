/**
 * Created by bim on 19.05.2017.
 */


App.config(
    (function (ChartJsProvider) {
        ChartJsProvider.setOptions({ colors : [ '#803690', '#00f', '#DCDCDC', '#46BFBD', '#FDB45C', '#949FB1', '#4D5360'] });
    })
).controller('indexCtrl', function($scope, $http) {

    $scope.init = function(data) {
        $scope.data = data;
        $scope.feedback = [];
        $scope.register = 0;
        $scope.unregister = 0;
        $scope.maleTotal = 0;
        $scope.femaleTotal = 0;
        $scope.twelve = 0;
        $scope.eighteen = 0;
        $scope.twentyfive = 0;
        $scope.thirtythree = 0;
        $scope.fortytwo = 0;
        $scope.sixty = 0;
        $scope.sixtyplus = 0;
        $scope.score = 0;
        $scope.scoreCount = 0;

        $scope.age = [];
        for (var i = 0; i < data.length; i++) {
            $scope.feedback.push(data[i].feedback)
            console.log(data[i].title + " index:" + i);
            for (var o = 0; o < data[i].feedback.length; o++) {
                if (data[i].feedback[o].user.username != "") {
                    gender = data[i].feedback[o].user.gender;
                    if (gender == "male" || gender == "Male") {
                        $scope.maleTotal++;
                        $scope.register++;
                        console.log("male:" + $scope.maleTotal);
                    } else if (gender == "women" || gender == "Women") {
                        $scope.femaleTotal++;
                        $scope.register++;
                        console.log("female:" + $scope.femaleTotal);
                    }
                    $scope.age.push(data[i].feedback[o].user.age);
                    if (data[i].feedback[o].user.age <= 12) {
                        $scope.twelve++;
                    } else if (data[i].feedback[o].user.age <= 18) {
                        $scope.eighteen++;
                    } else if (data[i].feedback[o].user.age <= 25) {
                        $scope.twentyfive++;
                    } else if (data[i].feedback[o].user.age <= 33) {
                        $scope.thirtythree++;
                    } else if (data[i].feedback[o].user.age <= 42) {
                        $scope.fortytwo++;
                    } else if (data[i].feedback[o].user.age <= 60) {
                        $scope.sixty++;
                    }  else if (data[i].feedback[o].user.age > 60) {
                        $scope.sixtyplus++;
                    }


                    for (var a = 0; a < data[i].feedback[o].answers.length; a++) {
                        if (data[i].questions[a] && data[i].questions[a].type == "stemning") {
                            console.log("score:" + data[i].feedback[o].answers[a].score);
                            $scope.score += data[i].feedback[o].answers[a].score;
                            $scope.scoreCount++;
                        }
                    }
                } else {
                    $scope.unregister++;
                }
            }


        }

        console.log($scope.maleTotal);
        console.log($scope.femaleTotal);

        $scope.avgScore = $scope.score / $scope.scoreCount + 1;
        console.log(Math.floor($scope.avgScore));
        $scope.labelsRegister = ["Registrerte brukere", "Uregistrerte brukere"];
        $scope.labelsGender = ["Menn", "Kvinner"];
        $scope.labelsAge = ["0-12år", "13-18år", "19-25år", "26-33år", "34-42år", "42-60år", "60+år"];
        $scope.dataRegister = [$scope.register, $scope.unregister];
        $scope.dataGender = [$scope.maleTotal, $scope.femaleTotal];
        $scope.dataAge = [$scope.twelve, $scope.eighteen, $scope.twentyfive, $scope.thirtythree, $scope.fortytwo, $scope.sixty, $scope.sixtyplus];
        $scope.colors = [ '#155965', '#72b978', '#DCDCDC', '#46BFBD', '#FDB45C', '#949FB1', '#4D5360'];
        $scope.options = [
            {
                size: {
                    height: 300,
                    width: 300
                }
            }
        ];
    }


    $scope.getImage = function() {

        if (Math.floor($scope.avgScore) == 1) {
            return "/assets/imgs/Frown.png";
        } else if (Math.floor($scope.avgScore) == 2) {
            return "/assets/imgs/sad.png";
        } else if (Math.floor($scope.avgScore) == 3) {
            return "/assets/imgs/neutral.png";
        } else if (Math.floor($scope.avgScore) == 4) {
            return "/assets/imgs/smile.png";
        } else if (Math.floor($scope.avgScore) == 5) {
            return "/assets/imgs/happy.png";
        }
    }

    $scope.getScore = function(usertest) {
        score = 0;
        if (usertest.feedback.length != 0) {
            score += usertest.feedback.length * 500;
            score += usertest.questions.length * 100;
            for (var i = 0; i < usertest.feedback.length; i++) {
                for (var o = 0; o < usertest.feedback[i].answers.length; o++) {
                    if (o < usertest.questions.length) {
                        if (usertest.questions[o].type = "stemning") {
                            score += usertest.feedback[i].answers[o].score * 100;
                        } else if (usertest.questions[o].type = "flervalg") {
                            score += 300;
                        } else if (usertest.questions[o].type = "text" && usertest.feedback[i].answers[o].answer != "") {
                            score += 500;
                        }
                    }
                }
            }
        }

        return score;
    }


});

