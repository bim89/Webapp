/**
 * Created by bim on 17.03.2017.
 */
App.config(
    (function (ChartJsProvider) {
        ChartJsProvider.setOptions({
            colors : [ '#803690', '#00f', '#DCDCDC', '#46BFBD', '#FDB45C', '#949FB1', '#4D5360'],
            scale : {
                ticks : {
                    beginAtZero : true
                }
            }
        });
    })
).controller('showCtrl', function($scope, $http) {
    $scope.moodLabels = ["Svært lite fornøyd", "Lite fornøyd", "Nøytral", "Fornøyd", "Svært fornøyd"];

    $scope.init = function(id) {


         $http.get("/usertest/show?t=" + id).then(function(response) {
             $scope.ut = response.data;
             $scope.moodData = [];
             $scope.answers = [];
             $scope.answerAmount = [];
             for (var i = 0; i < $scope.ut.questions.length; i++) {
                 $scope.moodData.push([0,0,0,0,0]);
                 $scope.answers.push($scope.ut.questions[i].choices.map(function(a) {return a.answer}));
                 $scope.answerAmount.push(Array.apply(null, {length: $scope.ut.questions[i].choices.length}).map(function() {return 0;}));
             }

             for (var i = 0; i < $scope.ut.feedback.length; i++) {
                 for (var b = 0;b < $scope.ut.feedback[i].answers.length; b++) {
                    if ($scope.ut.questions[b].type == "stemning") {
                        $scope.moodData[b][$scope.ut.feedback[i].answers[b].score]++;
                    } else if($scope.ut.questions[b].type == "flervalg") {
                        for(var c = 0; c < $scope.ut.questions[b].choices.length; c++) {
                            if ($scope.ut.feedback[i].answers[b].answer == $scope.ut.questions[b].choices[c].answer) {
                                $scope.answerAmount[b][c] += 1;
                            }
                        }
                    }

                }

             }



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

    $scope.getLabels = function() {
        return ["Svært lite fonøyd", "Lite fornøyd", "Nøytral", "Fornøyd", "Svært Fornøyd"];
    }

    /*
    $scope.getData = function(index) {
        data = [];
        $scope.$watch('ut.feedback', function (ut) {
            if (ut) {
                for (var i = 0; i < $scope.ut.feedback.length; i++) {
                    data.push($scope.ut.feedback[i].answers[index]);
                }

                return data
            } else {
                return null
            }
        });

    }
    */

    $scope.getImage = function(avg) {

        avg++;
        if (Math.floor(avg) == 1) {
            return "/assets/imgs/Frown.png";
        } else if (Math.floor(avg) == 2) {
            return "/assets/imgs/sad.png";
        } else if (Math.floor(avg) == 3) {
            return "/assets/imgs/neutral.png";
        } else if (Math.floor(avg) == 4) {
            return "/assets/imgs/smile.png";
        } else if (Math.floor(avg) == 5) {
            return "/assets/imgs/happy.png";
        }
    }

});