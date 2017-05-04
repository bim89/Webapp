/**
 * Created by pedjo on 04-May-17.
 */
google.charts.load('current', {'packages':['corechart']});
google.charts.setOnLoadCallback(drawChart);

function drawChart() {

    var data = google.visualization.arrayToDataTable([
        ['Response', 'Antall'],
        ['Fornøyd',     2],
        ['Litt fornøyd', 5],
        ['Nøytral',  2],
        ['Litt misfornøyd', 2],
        ['Misfornøyd',    5]
    ]);

    var options = {
        title: "",
        colors: ['#5CB73E', '#92B53D', '#E1C132', '#DC6137', '#E03434']
    };

    var chart = new google.visualization.PieChart(document.getElementById('piechart'));

    chart.draw(data, options);
}
