#include<stdio.h>

void main() {
	char str1[32];
	char str2[32] = "hansung";
	char str3[] = "hansung";

	char name1[][32] = { "kim", "lee", "choi", "park"};
	char* name2[] = { "kim", "lee", "choi", "park" };
	char* str = "korea";

	strcpy(str2, "PL programming langauge");
	printf("str2 = %s\n", str2);

	//strcpy(str3, "PL programming langauge");
	//printf("str2 = %s\n", str3);

	strcpy(name1[0], "PL programming langauge");
	printf("name1[0] = %s\n", name1[0]);

	//strcpy(name2[0], "PL programming langauge");
	//printf("name2[0] = %s\n", name2[0]);

	printf("LSH\n");
}