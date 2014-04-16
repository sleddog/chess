#include <stdio.h>
#include <stdlib.h>
#include <time.h>

/*
 *  Prototype version that will simply return a random move from a given board position
 *
 */

int main(void)
 {
   char randChar = ' ';
   int randNum =0;
   char * query = getenv("QUERY_STRING");

   //generate random letter between a-h
   srand(time(NULL));
   randNum = 8 * (rand() / (RAND_MAX + 1.0));
   randNum = randNum + 97;
   randChar = (char) randNum;

   printf("Content-type: text/json\n\n");
   printf("{");
   printf("\"query\":\"%s\",", query);
   //printf("board\":\"%s\",", board);
   //printf("move\":\"%s\",", move);

   //TODO testing just pawn moves for now. simply flip or mirror the move passed in
   printf("\"next-move\":\"%c7-%c5\"", randChar, randChar);
   printf("}\n");
   return 1;
 }

//TODO parse board into memory
//TODO parse move into memory
