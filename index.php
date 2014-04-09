<?php

$white_king = "&#9812;";
$white_queen  = "&#9813;";
$white_rook = "&#9814;";
$white_bishop = "&#9815;";
$white_knight = "&#9816;";
$white_pawn = "&#9817;";
$black_king = "&#9818;";
$black_queen  = "&#9819;";
$black_rook = "&#9820;";
$black_bishop = "&#9821;";
$black_knight = "&#9822;";
$black_pawn = "&#9823;";

$starting_positions = array(
    'white' => array(
        0 => $white_rook,
        1 => $white_knight,
        2 => $white_bishop,
        3 => $white_queen,
        4 => $white_king,
        5 => $white_bishop,
        6 => $white_knight,
        7 => $white_rook
    ),
    'black' => array(
        0 => $black_rook,
        1 => $black_knight,
        2 => $black_bishop,
        3 => $black_queen,
        4 => $black_king,
        5 => $black_bishop,
        6 => $black_knight,
        7 => $black_rook
    )
);

function start_pos($index, $color) {
    global $starting_positions;
    return $starting_positions[$color][$index];
}

function border_width($index) {
    return '1pt 1pt 1pt 1pt';
}

function bg_color($index) {
    if($index % 2 != 0) {
        return ' bgcolor="silver"';
    }
    return '';
}

function draw_chess_board() {
    global $white_pawn, $black_pawn;

    $board = '<table style="text-align:center;border-spacing:0pt; border-collapse:collapse; border-color: black; border-style: solid; border-width: 0pt 0pt 0pt 0pt">';
    $board .= '<tbody>';
    $board .= '<tr style="vertical-align:bottom;">';
    $board .= '<td style="vertical-align:middle;width:12pt">8</td>';
    for($i=0; $i<8; $i++) {
        $board .= '<td style="width:38pt; height:24pt; border-collapse:collapse; border-color: black; border-style: solid; border-width: ' . border_width($i) . '"' .  bg_color($i) . '><span style="font-size:250%;">' . start_pos($i, 'black') . '</span></td>';
    }
    $board .= '</tr>';
    $board .= '<tr style="vertical-align:bottom;">';
    $board .= '<td style="vertical-align:middle;width:12pt">7</td>';
    for($i=0; $i<8; $i++) {
        $board .= '<td style="width:38pt; height:24pt; border-collapse:collapse; border-color: black; border-style: solid; border-width: ' . border_width($i) . '"' .  bg_color($i+1) . '><span style="font-size:250%;">' . $black_pawn . '</span></td>';
    }
    $board .= '</tr>';
    for($j=0; $j<4; $j++) {
        $board .= '<tr style="vertical-align:bottom;">';
        $board .= '<td style="vertical-align:middle;width:12pt">' . (5-$j) . '</td>';
        for($i=0; $i<8; $i++) {
            $board .= '<td style="width:38pt; height:24pt; border-collapse:collapse; border-color: black; border-style: solid; border-width: ' . border_width($i) . '"' .  bg_color($i+$j) . '><span style="font-size:250%;">' . '&nbsp;' . '</span></td>';
        }
        $board .= '</tr>';
    }
    $board .= '<tr style="vertical-align:bottom;">';
    $board .= '<td style="vertical-align:middle;width:12pt">2</td>';
    for($i=0; $i<8; $i++) {
        $board .= '<td style="width:38pt; height:24pt; border-collapse:collapse; border-color: black; border-style: solid; border-width: ' . border_width($i) . '"' .  bg_color($i) . '><span style="font-size:250%;">' . $white_pawn . '</span></td>';
    }
    $board .= '</tr>';
    $board .= '<tr style="vertical-align:bottom;">';
    $board .= '<td style="vertical-align:middle;width:12pt">1</td>';
    for($i=0; $i<8; $i++) {
        $board .= '<td style="width:38pt; height:24pt; border-collapse:collapse; border-color: black; border-style: solid; border-width: ' . border_width($i) . '"' .  bg_color($i+1) . '><span style="font-size:250%;">' . start_pos($i, 'white') . '</span></td>';
    }
    $board .= '</tr>';
    $board .= '<tr>';
    $board .= '<td></td>';
    $board .= '<td>a</td>';
    $board .= '<td>b</td>';
    $board .= '<td>c</td>';
    $board .= '<td>d</td>';
    $board .= '<td>e</td>';
    $board .= '<td>f</td>';
    $board .= '<td>g</td>';
    $board .= '<td>h</td>';
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

<?=draw_chess_board();?>
<div id="images"></div>
 
<script>
(function() {
  var flickerAPI = "http://api.flickr.com/services/feeds/photos_public.gne?jsoncallback=?";
  $.getJSON( flickerAPI, {
    tags: "chess",
    tagmode: "any",
    format: "json"
  })
    .done(function( data ) {
      $.each( data.items, function( i, item ) {
        $( "<img>" ).attr( "src", item.media.m ).appendTo( "#images" );
        if ( i === 3 ) {
          return false;
        }
      });
    });
})();
</script>
 
</body>
</html>
