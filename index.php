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

function draw_chess_board_from_config($board_json) {
    global $pieces, $num_to_letter;

    //json representation of an initial chess board
    $config = json_decode($board_json);
    //echo "<pre>".print_r($board_config, true)."</pre>";

    $board = '<table class="table table-bordered" align="center" style="height:100%"><tbody>';
    for($i=8; $i>=1; $i--) {
        $board .= '<tr>';
        $board .= "<td>$i</td>";
        for($j=0; $j<8; $j++) {
            $l = $num_to_letter[$j];
            $square = "$l$i";
            $color = (($j+$i) % 2 != 0) ? 'silver' : 'white';
            $piece = property_exists($config, $square) ? $pieces[$config->$square]['html'] : '&nbsp;';
            $board .= "<td><button class='btn btn-default' id='$square' type='button' style='width:85px; height:85px; background-color: $color; font-size:50px' onclick='square_select(this)'>$piece</button></td>";
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


function get_initial_board()
{
    $board = '{"a8":"br","b8":"bn","c8":"bb","d8":"bq","e8":"bk","f8":"bb","g8":"bn","h8":"br",';
    $board .= '"a7":"bp","b7":"bp","c7":"bp","d7":"bp","e7":"bp","f7":"bp","g7":"bp","h7":"bp",';
    $board .= '"a2":"wp","b2":"wp","c2":"wp","d2":"wp","e2":"wp","f2":"wp","g2":"wp","h2":"wp",';
    $board .= '"a1":"wr","b1":"wn","c1":"wb","d1":"wq","e1":"wk","f1":"wb","g1":"wn","h1":"wr"}';
    return $board;
}


function get_random_board()
{
    global $num_to_letter, $pieces;
    $blank_chance = $_GET['blank_chance'] != null ? intval($_GET['blank_chance']) : rand(0,100);
    $board = '{';
    for($i=8; $i>=1; $i--) {
        for($j=0; $j<8; $j++) {
            $l = $num_to_letter[$j];
            $square = "$l$i";
            if(rand(1,100) > $blank_chance) {
                continue; //allow for blank pieces
            }
            $rand_piece = array_rand($pieces);
            $board .= '"'.$square.'":"'.$rand_piece.'",';
        }
    }
    $board = rtrim($board, ",");
    $board .= '}';
    return $board;
}


//define what board JSON to use for initial load
if($_GET['random']) {
    $board_to_draw = get_random_board(); 
}
else {
    $board_to_draw = get_initial_board();
}
?>

<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>Chess by Devin Gray</title>

    <!-- Bootstrap -->
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.1.1/css/bootstrap.min.css">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
      <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>
  <body>
<div class='container fill' style='height:100%'>
<table class='table' align="center" style='height:100%'><tr><td valign=top>
<div class="table-responsive" style='height:100%'>
<?=draw_chess_board_from_config($board_to_draw);?>
</div>
</td>
<td valign=top>
<b>Chess by Devin Gray</b><br />
open source chess application<br />
on <a href="https://github.com/sleddog/chess">github.com/sleddog/chess</a><br />
<hr />
<div id='submit_move_div'>
<input type='text' class='input-sm' style='width:45px' id='piece_from' readonly="true" />
<span>-</span>
<input type='text' class='input-sm' style='width:45px' id='piece_to' readonly="true" />
<input class='btn btn-default' id='submit_move_button' type='button' value='Submit Move' onclick='submit_move()' disabled='true'/>
</div>
<hr />
<table id='move_history_table' width='180'>
<tr><td>&nbsp;</td><td>White</td><td>Black</td></tr>
</table>
</td>
</tr></table>
</div>

<script>
var selectedSquare = 0;
var targetSquare = 0;
var legal_moves = [];
var board = null;
var num_to_letter = ['a','b','c','d','e','f','g','h'];
var pieces = JSON.parse('<?=json_encode($pieces); ?>');
var history = [];

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
    reset_target_square();
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

    //update the history
    update_history('white', selectedMove);

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
    board: JSON.stringify(board),
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
        //update the history
        update_history('black', next_move);
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
        case 'n':
            return legal_knight_moves(coord, color);
            break;
        case 'b':
            return legal_bishop_moves(coord, color);
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

function legal_knight_moves(coord, color) {
    if(color == 'w') {
        //create an array of coords, based on knight's movement (L shape)
        var moves = new Array(8); 
        moves[0] = [coord[0]+2, coord[1]-1];
        moves[1] = [coord[0]+2, coord[1]+1];
        moves[2] = [coord[0]-2, coord[1]-1];
        moves[3] = [coord[0]-2, coord[1]+1];
        moves[4] = [coord[0]+1, coord[1]-2];
        moves[5] = [coord[0]+1, coord[1]+2];
        moves[6] = [coord[0]-1, coord[1]-2];
        moves[7] = [coord[0]-1, coord[1]+2];
        //check each move
        for(var i=0; i<8; i++) {
            if(white_can_move(moves[i])) {
                add_legal_move(moves[i]);
            }
        }
    }
    else { // color == 'b'
    }
}

function white_can_move(coord) {
    //is off board?
    if(coord[0] < 0 || coord[0] > 7 || coord[1] < 0 || coord[1] > 7) {
        return false;
    }
    //is the square empty?
    var square = board[coord[0]][coord[1]];
    if(square == 0) {
        return true;
    }
    else {
        if(square.substring(0,1) == 'b') {
            //black piece exists in this square
            return true;
        }
        else {
            //white piece is already in place here
            return false;
        }
    }
}


function legal_bishop_moves(coord, color) {
    if(color == 'w') {
        //for each direction, travel until the edge of board or piece
        var directions = [[1,1],[1,-1],[-1,-1],[-1,1]];
        for(var i=0; i<4; i++) {
            var dir = directions[i];
            var step = 1;
            while(true) {
                var move = [coord[0]+(dir[0]*step), coord[1]+(dir[1]*step)];
                console.log(move);
                if(white_can_move(move)) {
                    add_legal_move(move);
                    step++;
                    //if move is a black piece, stop calculating on this line
                    var square = board[move[0]][move[1]];
                    if(square != 0 && square.substring(0,1) == 'b') {
                        break;
                    }
                }
                else {
                    break; //can't move here...
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

function update_history(player, move)
{
    var table = document.getElementById('move_history_table');
    if(!table) {
        return;
    }

    //based on the player, update the appropriate cell, or create new row
    var rowCount = table.rows.length
    if(player == 'white') {
        var row = table.insertRow(rowCount);
        var moveNumber = row.insertCell(0);
        moveNumber.innerHTML = rowCount;
        var white = row.insertCell(1);
        white.innerHTML = move;
    }
    else { // black
        var row = table.rows[rowCount-1];
        var black = row.insertCell(2);
        black.innerHTML = move;
    }

    //update the history array with this move
    history.push([player, move]);
    //console.log(history);
}

//onload
(function() {
  init_board('<?=$board_to_draw; ?>');
})();

</script>
<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-1105056-1");
pageTracker._trackPageview();
</script>

    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min.js"></script>
    <!-- Latest compiled and minified JavaScript -->
    <script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"></script>
  </body>
</html>
