<?php
/*
*  Author:  devingray@gmail.com
*  Very simple chess application written in PHP/HTML/Javascript on the frontend that 
*  invokes a CGI program written in C
*/

$num_to_letter = array('a','b','c','d','e','f','g','h');
$pieces = array(
    'wk'=>"&#9812;",
    'wq'=>"&#9813;",
    'wr'=>"&#9814;",
    'wb'=>"&#9815;",
    'wn'=>"&#9816;",
    'wp'=>"&#9817;",
    'bk'=>"&#9818;",
    'bq'=>"&#9819;",
    'br'=>"&#9820;",
    'bb'=>"&#9821;",
    'bn'=>"&#9822;",
    'bp'=>"&#9823;"
);
$initial_board_json = '{"a8":"br","b8":"bn","c8":"bb","d8":"bq","e8":"bk","f8":"bb","g8":"bn","h8":"br",';
$initial_board_json .= '"a7":"bp","b7":"bp","c7":"bp","d7":"bp","e7":"bp","f7":"bp","g7":"bp","h7":"bp",';
$initial_board_json .= '"a2":"wp","b2":"wp","c2":"wp","d2":"wp","e2":"wp","f2":"wp","g2":"wp","h2":"wp",';
$initial_board_json .= '"a1":"wr","b1":"wn","c1":"wb","d1":"wq","e1":"wk","f1":"wb","g1":"wn","h1":"wr"}';

function draw_chess_board_from_config($board_json) {
    global $pieces, $num_to_letter;

    //json representation of an initial chess board
    $config = json_decode($board_json);
    //echo "<pre>".print_r($board_config, true)."</pre>";

    $board = '<table><tbody>';
    for($i=8; $i>=1; $i--) {
        $board .= '<tr>';
        $board .= "<td>$i</td>";
        for($j=0; $j<8; $j++) {
            $l = $num_to_letter[$j];
            $square = "$l$i";
            $color = (($j+$i) % 2 != 0) ? 'silver' : 'white';
            $piece = property_exists($config, $square) ? $pieces[$config->$square] : '&nbsp;';
            $board .= "<td><button id='$square-$color' type='button' style='width:75px; height:75px; font-size: 60px; background-color: $color' onclick='square_select(this)'>$piece</button></td>";
        }
        $board .= '</tr>';
    }
    $board .= '</tr>';
    $board .= '<tr>';
    $board .= '<td></td>';
    for($j=0; $j<8; $j++) {
        $l = $num_to_letter[$j];
        $board .= "<td align='center'>$l</td>";
    }
    $board .= '</tr>';
    $board .= '</tbody></table>';
    return $board;
}


?>

<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
<title>
Devin Gray's chess experiment
</title>
  <style>
  img {
    height: 100px;
    float: left;
  }
  </style>
  <script src="//code.jquery.com/jquery-1.10.2.js"></script>
</head>
<body>
 
<h1>Devin Gray's chess experiment</h1>

<?=draw_chess_board_from_config($initial_board_json);?>
<div id="next-move"></div>
 
<script>
(function() {
  var chessAI = "http://www.devingray.com/cgi-bin/chess_ai.cgi";
  $.getJSON( chessAI, {
    board: '<?=$initial_board_json;?>',
    move: "e2-e4" // this will be a legal chess move
  })
    .done(function( data ) {
        console.log(data);
    });
})();

var selectedSquare = 0;
var targetSquare = 0;

function square_select(button) {
    console.log(button);
    console.log(button.id);
    console.log(button.style.backgroundColor);
    if(selectedSquare == 0) {
        button.style.backgroundColor="lightgreen";
        selectedSquare = button.id;
    }
    else {
        if(selectedSquare == button.id) {
            button.style.backgroundColor=selectedSquare.split('-')[1];
            selectedSquare = 0;
        }
        else { //set target
            button.style.backgroundColor="lightpink";
            targetSquare = button.id;
        }
    }
}
</script>

</body>
</html>
