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

    $scope.dateLabels = ["Januar", "Februar", "Mars", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Desember"];
    $scope.dateData = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

    $scope.weekLabels = ["Mandag", "Tirsdag", "Onsdag", "Torsdag", "Fredag", "Lørdag", "Søndag"];
    $scope.weekData = [0, 0, 0, 0, 0, 0, 0];


    $scope.timeLabels = ["07:00-12:00", "12:00-15:00", "15:00-18:00", "18:00-21:00", "21:00-24:00", "24:00-07:00"];
    $scope.timeData = [0, 0, 0, 0, 0, 0];


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
                 var date =  new Date($scope.ut.feedback[i].date);
                 $scope.dateData[date.getMonth()] += 1;
                 $scope.weekData[date.getDay()] += 1;
                 if (date.getHours() >= 7 && date.getHours() < 12) {
                     $scope.timeData[0]++;
                 } else if (date.getHours() >= 12 && date.getHours() < 15) {
                     $scope.timeData[1]++;
                 } else if (date.getHours() >= 15 && date.getHours() < 18) {
                     $scope.timeData[2]++;
                 } else if (date.getHours() >= 18 && date.getHours() < 21) {
                     $scope.timeData[3]++;
                 } else if (date.getHours() >= 21 && date.getHours() <= 23) {
                     $scope.timeData[4]++;
                 } else {
                     $scope.timeData[5]++;
                 }

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

        return avg + 1
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