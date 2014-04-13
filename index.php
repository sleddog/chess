<?php
/*
*  Author:  devingray@gmail.com
*  Very simple chess application written in PHP/HTML/Javascript on the frontend that 
*  invokes a CGI program written in C
*/

$num_to_letter = array('a','b','c','d','e','f','g','h');
$pieces = array(
    'wk'=>array('html'=>"&#9812;",'codepoint'=>'\u2654'),
    'wq'=>array('html'=>"&#9813;",'codepoint'=>'\u2655'),
    'wr'=>array('html'=>"&#9814;",'codepoint'=>'\u2656'),
    'wb'=>array('html'=>"&#9815;",'codepoint'=>'\u2657'),
    'wn'=>array('html'=>"&#9816;",'codepoint'=>'\u2658'),
    'wp'=>array('html'=>"&#9817;",'codepoint'=>'\u2659'),
    'bk'=>array('html'=>"&#9818;",'codepoint'=>'\u265A'),
    'bq'=>array('html'=>"&#9819;",'codepoint'=>'\u265B'),
    'br'=>array('html'=>"&#9820;",'codepoint'=>'\u265C'),
    'bb'=>array('html'=>"&#9821;",'codepoint'=>'\u265D'),
    'bn'=>array('html'=>"&#9822;",'codepoint'=>'\u265E'),
    'bp'=>array('html'=>"&#9823;",'codepoint'=>'\u265F')
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
            $piece = property_exists($config, $square) ? $pieces[$config->$square]['html'] : '&nbsp;';
            $board .= "<td><button id='$square' type='button' style='width:75px; height:75px; font-size: 60px; background-color: $color' onclick='square_select(this)'>$piece</button></td>";
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
var selectedSquare = 0;
var targetSquare = 0;
var legal_moves = [];
var board = null;
var num_to_letter = ['a','b','c','d','e','f','g','h'];
var pieces = JSON.parse('<?=json_encode($pieces); ?>');
console.log(pieces);

function piece_to_unicode(piece) {
    return pieces[piece]['codepoint'];
}

function letter_to_num(letter) {
    var letters = "abcdefgh";
    return letters.indexOf(letter);
}

function init_board(board_json) {
    var board_config = JSON.parse(board_json);
    board = new Array(8); 
    for (var i=0; i<8; i++) {
        board[i] = new Array(8); 
        for (var j=0; j<8; j++) {
            var letter = num_to_letter[j]
            var square = letter + (i+1);
            if (board_config.hasOwnProperty(square)) {
                var value = board_config[square];
                board[i][j] = value;
            }
            else {
                board[i][j] = 0;
            }
        }
    }
    //console.log(board);
    print_board_to_console(board);
}

function print_board_to_console(board) {
    for (var i=0; i<8; i++) {
        var row = "";
        for (var j=0; j<8; j++) {
            row += "|" + ((board[i][j] == 0) ? "\u3000" : piece_to_unicode(board[i][j]));
            if (j==7) {
                row += "|";
            }
        }
        console.log(row);
    }
}


function get_move_from_server(selectedMove) {
  var chessAI = "http://www.devingray.com/cgi-bin/chess_ai.cgi";
  $.getJSON( chessAI, {
    board: '<?=$initial_board_json;?>',
    move: selectedMove //"e2-e4" this will be a legal chess move
  })
    .done(function( data ) {
        console.log(data);
    });
}

function square_to_color(square) {
  var letter = square.substring(0,1);
  var num = parseInt(square.substring(1,2));
  if("aceg".indexOf(letter) >= 0) {
    num++;
  }
  return  (num % 2 != 0) ? "white" : "silver";
}

function square_to_coord(square) {
    var letter = square.substring(0,1);
    var num = parseInt(square.substring(1,2));
    //[i,j]... a1 == [0,0]
    return [num-1, letter_to_num(letter)]
}

function square_select(button) {
    console.log(button);
    console.log(button.id);
    var coord = square_to_coord(button.id);
    console.log('coord=');
    console.log(coord);
    var value = board[coord[0]][coord[1]];
    console.log('value=');
    console.log(value);

    if(selectedSquare == 0) {
        //determine if there is even a piece on this square
        if(value == 0) {
            return; // there must be a piece here in order to select
        }
        if(value.substring(0,1) == 'b') {
            return; // only allow white pieces to be selected at first
        }
        selectedSquare = button.id;
        button.style.backgroundColor = "lightgreen";
        highlight_legal_moves(selectedSquare);
    }
    else {
        if(selectedSquare == button.id) {
            button.style.backgroundColor = square_to_color(selectedSquare);
            selectedSquare = 0;
            legal_moves = [];
            if(targetSquare != 0) {
                document.getElementById(targetSquare).style.backgroundColor = square_to_color(targetSquare);
            }
        }
        else { //set target
            if(targetSquare != 0) {
                document.getElementById(targetSquare).style.backgroundColor = square_to_color(targetSquare);
            }
            //if they select another piece, restart selection
            if(value != 0) {
                if(value.substring(0,1) == 'w') {
                    document.getElementById(selectedSquare).style.backgroundColor = square_to_color(selectedSquare);
                    selectedSquare = button.id;
                    button.style.backgroundColor = "lightgreen";
                    return;
                }
            }
            button.style.backgroundColor = "lightpink";
            targetSquare = button.id;
        }
    }
}

function highlight_legal_moves(selectedSquare) {
    //inspect the board, and determine what the legal moves are
}

//onload
(function() {
  init_board('<?=$initial_board_json; ?>');
})();

</script>

</body>
</html>
