#include <stdio.h>
#include <stdlib.h>

/*
 *  Prototype version that will simply return a random move from a given board position
 *
 */

int main(void)
 {
   char * query = getenv("QUERY_STRING");
   printf("Content-type: text/json\n\n");
   printf("{");
   printf("\"query\":\"%s\",", query);
   //printf("board\":\"%s\",", board);
   //printf("move\":\"%s\",", move);
   printf("\"next-move\":\"e7-e5\"");
   printf("}\n");
   return 1;
 }
