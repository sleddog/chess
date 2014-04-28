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
Chess by Devin Gray
</title>
  <script src="//code.jquery.com/jquery-1.10.2.js"></script>
</head>
<body>
 
<b>Chess by Devin Gray</b><br />
open source code chess application (view on <a href="https://github.com/sleddog/chess">github.com/sleddog/chess</a>)<br />

<table><tr><td valign=top>
<?=draw_chess_board_from_config($initial_board_json);?>
</td>
<td valign=top>
History:<br />
<div id='move_history'>
</div>
</td>
</tr></table>

<br />

<div id='submit_move_div'>
<table><tr><td align="right" valign="center" width="300">
<input type='text' id='piece_from' style='width:40px; font-size: 35px' readonly="true" />
<span style="font-size: 25px">-</span>
<input type='text' id='piece_to' style='width:40px; font-size: 35px' readonly="true" />
</td>
<td valign="center">
<input id='submit_move_button' type='button' value='Submit Move' onclick='submit_move()' disabled='true' style='font-size: 26px' />
</td></tr></table>
</div>
 
<script>
var selectedSquare = 0;
var targetSquare = 0;
var legal_moves = [];
var board = null;
var num_to_letter = ['a','b','c','d','e','f','g','h'];
var pieces = JSON.parse('<?=json_encode($pieces); ?>');
//console.log(pieces);

function piece_to_unicode(piece) {
    return pieces[piece]['codepoint'];
}

function piece_to_html(piece) {
    return pieces[piece]['html'];
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
    //print_board_to_console(board);
}

function print_board_to_console(board) {
    for (var i=7; i>=0; i--) {
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
    var coord = square_to_coord(button.id);
    var value = board[coord[0]][coord[1]];

    if(selectedSquare == 0) {
        //determine if there is even a piece on this square
        if(value == 0) {
            return; // there must be a piece here in order to select
        }
        if(value.substring(0,1) == 'b') {
            return; // only allow white pieces to be selected at first
        }
        highlight_initial_move(button);
    }
    else {
        //are they unselecting the first square?
        if(selectedSquare == button.id) {
            reset_initial_square();
            if(targetSquare != 0) {
                document.getElementById(targetSquare).style.backgroundColor = square_to_color(targetSquare);
            }
        }
        else { //set target
            if(targetSquare != 0) {
                //if square was a previous legal move use highlight color
                if(legal_moves.indexOf(targetSquare) > -1) {
                    resetColor = "lightyellow";
                }
                else {
                    resetColor = square_to_color(targetSquare);
                }
                document.getElementById(targetSquare).style.backgroundColor = resetColor;
            }
            //if they select another piece, restart selection
            if(value != 0) {
                if(value.substring(0,1) == 'w') {
                    reset_initial_square();
                    highlight_initial_move(button);
                    return;
                }
            }
            //make sure they are selecting legal moves
            if(legal_moves.indexOf(button.id) > -1) {
                //reset if they pick the same exact target
                if(targetSquare == button.id) {
                    reset_target_square();
                    return;
                }
                else {
                    button.style.backgroundColor = "lightpink";
                    targetSquare = button.id;
                    enable_submit_move();
                }
            }
            else {
                reset_target_square();
            }
        }
    }
}

function reset_initial_square() {
    clear_legal_moves();
    document.getElementById(selectedSquare).style.backgroundColor = square_to_color(selectedSquare);
    selectedSquare = 0;
    document.getElementById('piece_from').value = '';
}

function reset_target_square() {
    targetSquare = 0;
    document.getElementById('piece_to').value = '';
    document.getElementById('submit_move_button').disabled = true;
}

function enable_submit_move() {
    document.getElementById('submit_move_button').disabled = false;
    document.getElementById('piece_to').value = targetSquare;
}

function submit_move() {
    //validate one last time...
    if(selectedSquare == 0 || targetSquare == 0) {
        alert('error some how...');
        return;
    }
    var selectedMove = selectedSquare + '-' + targetSquare;

    //move the piece
    move_pieces(selectedSquare, targetSquare);

    //reset selections
    reset_initial_square();
    reset_target_square();

    //now call the AI to get the computer's move
    get_next_move(selectedMove);
}

function get_next_move(selectedMove) {
    //determine what mode we are in... human vs human, human vs AI
    //assuming for now to just be simple random AI
    get_move_from_server(selectedMove); 
}

function get_move_from_server(selectedMove) {
  var chessAI = "http://www.devingray.com/cgi-bin/chess_ai.cgi";
  $.getJSON( chessAI, {
    board: JSON.stringify(board), //'<?=$initial_board_json;?>',
    move: selectedMove //"e2-e4" this will be a legal chess move
  })
    .done(function( data ) {
        //console.log(data);
        if(data['next-move']) {
            make_move(data['next-move']);
        }
    });
}

function make_move(next_move) {
    var moves = next_move.split('-');
    if(moves.length != 2) {
        alert('error with move: ' + next_move);
        return;
    }
    //make sure this is a valid move
    var old_coord = square_to_coord(moves[0]);
    var piece = board[old_coord[0]][old_coord[1]];
    if(piece && piece.substring(0,1) == 'b') {
        move_pieces(moves[0], moves[1]);
    }
}

function move_pieces(from, to) {
    //move the piece
    var old_coord = square_to_coord(from);
    var new_coord = square_to_coord(to);
    var piece = board[old_coord[0]][old_coord[1]];
    board[old_coord[0]][old_coord[1]] = 0;
    board[new_coord[0]][new_coord[1]] = piece;
    //print_board_to_console(board);
    document.getElementById(from).innerHTML = '&nbsp;';
    document.getElementById(to).innerHTML = piece_to_html(piece);
}


function highlight_initial_move(button) {
    clear_legal_moves();
    selectedSquare = button.id;
    button.style.backgroundColor = "lightgreen";
    highlight_legal_moves(selectedSquare);
    document.getElementById('piece_from').value = selectedSquare;
}

function highlight_legal_moves(selectedSquare) {
    //inspect the board, and determine what the legal moves are
    //console.log(selectedSquare);
    var coord = square_to_coord(selectedSquare);
    //console.log(coord);
    //console.log(board);
    var piece = board[coord[0]][coord[1]];
    //console.log(piece);
    var color = piece.substring(0,1);
    var type = piece.substring(1,2);
    switch(type) {
        case 'p':
            return legal_pawn_moves(coord, color);
            break;
        default:
            console.log('default case');
            console.log(type);
            break;
    }
}

function coord_to_square(coord) {
    var letter = num_to_letter[coord[1]];
    return letter + (coord[0]+1);
}

//from the given coord, highlight legal pawn moves, for the respective color
function legal_pawn_moves(coord, color) {
    if(color == 'w') {
        var newCoord = [coord[0]+1, coord[1]];
        //if first square in front is blank 
        if(board[newCoord[0]][newCoord[1]] == 0) {
            add_legal_move(newCoord);

            //now check 2 moves in front if on the 2nd row
            if(coord[0] == 1) {
                var newCoord2 = [coord[0]+2, coord[1]];
                if(board[newCoord2[0]][newCoord2[1]] == 0) {
                    add_legal_move(newCoord2);
                }
            } 
        }

        //can you attack diagonally?
        diag_right = [coord[0]+1, coord[1]+1];
        if(diag_right[1] <= 7) {
            var piece = board[diag_right[0]][diag_right[1]];
            if(piece != 0) {
                //if a piece exists and is black
                if(piece.substring(0,1) == 'b') {
                    add_legal_move(diag_right);
                }
            }
        }
        diag_left = [coord[0]+1, coord[1]-1];
        if(diag_left[1] >= 0) {
            var piece = board[diag_left[0]][diag_left[1]];
            if(piece != 0) {
                //if a piece exists and is black
                if(piece.substring(0,1) == 'b') {
                    add_legal_move(diag_left);
                }
            }
        }

    }
    else { // color == 'b'
    }
}

//highlight the legal move on the board for this coord and track in global array
function add_legal_move(coord) {
    var square = coord_to_square(coord);
    document.getElementById(square).style.backgroundColor = 'lightyellow';
    legal_moves.push(square);
}

function clear_legal_moves() {
    for(var i=0; i<legal_moves.length; i++) {
        var square = legal_moves[i];
        document.getElementById(square).style.backgroundColor = square_to_color(square);
    }
    legal_moves = [];
}

//onload
(function() {
  init_board('<?=$initial_board_json; ?>');
})();

</script>

</body>
</html>
